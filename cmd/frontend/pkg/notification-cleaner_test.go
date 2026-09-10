// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg_test

import (
	"context"
	"time"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	libkvmocks "github.com/bborbe/kv/mocks"
	"github.com/bborbe/log"
	libtime "github.com/bborbe/time"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/mocks"
	"github.com/bborbe/notification/cmd/frontend/pkg"
	"github.com/bborbe/notification"
	coremocks "github.com/bborbe/notification/mocks"
)

var _ = Describe("NotificationCleaner", func() {
	var (
		ctx                 context.Context
		tx                  *libkvmocks.Tx
		currentDateTime     libtime.CurrentDateTime
		notificationIndexer *mocks.NotificationIndexer
		notificationStoreTx *coremocks.CoreNotificationStoreTx
		cleaner             pkg.NotificationCleaner
		maxAge              libtime.Duration
		now                 time.Time
		cutoffTime          time.Time
	)

	BeforeEach(func() {
		ctx = context.Background()
		tx = &libkvmocks.Tx{}
		now = time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)
		currentDateTime = libtime.NewCurrentDateTime()
		currentDateTime.SetNow(libtime.DateTime(now))
		notificationIndexer = &mocks.NotificationIndexer{}
		notificationStoreTx = &coremocks.CoreNotificationStoreTx{}

		cleaner = pkg.NewNotificationCleaner(
			currentDateTime,
			notificationIndexer,
			notificationStoreTx,
			log.DefaultSamplerFactory,
		)

		maxAge = libtime.Duration(720 * time.Hour) // 30 days
		cutoffTime = now.Add(-maxAge.Duration())
	})

	Context("when cleaning notifications", func() {
		var oldNotifications []core.Notification
		var recentNotifications []core.Notification

		BeforeEach(func() {
			// Old notifications (should be cleaned)
			oldNotifications = []core.Notification{
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old1"),
						Created:    libtime.DateTime(cutoffTime.Add(-24 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old2"),
						Created:    libtime.DateTime(cutoffTime.Add(-48 * time.Hour)),
					},
				},
			}

			// Recent notifications (should be kept)
			recentNotifications = []core.Notification{
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("recent1"),
						Created:    libtime.DateTime(cutoffTime.Add(24 * time.Hour)),
					},
				},
			}

			allNotifications := append(oldNotifications, recentNotifications...)

			// Mock Map to iterate over all notifications
			notificationStoreTx.MapCalls(
				func(ctx context.Context, tx libkv.Tx, fn func(context.Context, core.Notification) error) error {
					for _, n := range allNotifications {
						if err := fn(ctx, n); err != nil {
							return err
						}
					}
					return nil
				},
			)

			// Mock Remove to succeed
			notificationStoreTx.RemoveReturns(nil)
			notificationIndexer.RemoveReturns(nil)
		})

		It("should clean old notifications and keep recent ones", func() {
			err := cleaner.Clean(ctx, tx, maxAge, 10000)

			Expect(err).NotTo(HaveOccurred())

			// Should remove 2 old notifications from store
			Expect(notificationStoreTx.RemoveCallCount()).To(Equal(2))
			_, _, id1 := notificationStoreTx.RemoveArgsForCall(0)
			_, _, id2 := notificationStoreTx.RemoveArgsForCall(1)
			Expect([]base.Identifier{id1, id2}).To(ConsistOf(
				base.Identifier("old1"),
				base.Identifier("old2"),
			))

			// Should remove 2 old notifications from index
			Expect(notificationIndexer.RemoveCallCount()).To(Equal(2))
			_, idxId1 := notificationIndexer.RemoveArgsForCall(0)
			_, idxId2 := notificationIndexer.RemoveArgsForCall(1)
			Expect([]base.Identifier{idxId1, idxId2}).To(ConsistOf(
				base.Identifier("old1"),
				base.Identifier("old2"),
			))
		})
	})

	Context("when no notifications need cleaning", func() {
		BeforeEach(func() {
			// All notifications are recent
			recentNotifications := []core.Notification{
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("recent1"),
						Created:    libtime.DateTime(cutoffTime.Add(24 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("recent2"),
						Created:    libtime.DateTime(cutoffTime.Add(48 * time.Hour)),
					},
				},
			}

			notificationStoreTx.MapCalls(
				func(ctx context.Context, tx libkv.Tx, fn func(context.Context, core.Notification) error) error {
					for _, n := range recentNotifications {
						if err := fn(ctx, n); err != nil {
							return err
						}
					}
					return nil
				},
			)
		})

		It("should not remove any notifications", func() {
			err := cleaner.Clean(ctx, tx, maxAge, 10000)

			Expect(err).NotTo(HaveOccurred())
			Expect(notificationStoreTx.RemoveCallCount()).To(Equal(0))
			Expect(notificationIndexer.RemoveCallCount()).To(Equal(0))
		})
	})

	Context("when store remove fails", func() {
		BeforeEach(func() {
			oldNotification := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("old1"),
					Created:    libtime.DateTime(cutoffTime.Add(-24 * time.Hour)),
				},
			}

			notificationStoreTx.MapCalls(
				func(ctx context.Context, tx libkv.Tx, fn func(context.Context, core.Notification) error) error {
					return fn(ctx, oldNotification)
				},
			)

			notificationStoreTx.RemoveReturns(errors.New(ctx, "store remove failed"))
		})

		It("should return error", func() {
			err := cleaner.Clean(ctx, tx, maxAge, 10000)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("store remove failed"))
		})
	})

	Context("when index remove fails", func() {
		BeforeEach(func() {
			oldNotification := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("old1"),
					Created:    libtime.DateTime(cutoffTime.Add(-24 * time.Hour)),
				},
			}

			notificationStoreTx.MapCalls(
				func(ctx context.Context, tx libkv.Tx, fn func(context.Context, core.Notification) error) error {
					return fn(ctx, oldNotification)
				},
			)

			notificationStoreTx.RemoveReturns(nil)
			notificationIndexer.RemoveReturns(errors.New(ctx, "index remove failed"))
		})

		It("should return error", func() {
			err := cleaner.Clean(ctx, tx, maxAge, 10000)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("index remove failed"))
		})
	})

	Context("when Map fails", func() {
		BeforeEach(func() {
			notificationStoreTx.MapReturns(errors.New(ctx, "map failed"))
		})

		It("should return error", func() {
			err := cleaner.Clean(ctx, tx, maxAge, 10000)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("map failed"))
		})
	})

	Context("when limit is reached", func() {
		BeforeEach(func() {
			// Create 5 old notifications but limit to 3
			oldNotifications := []core.Notification{
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old1"),
						Created:    libtime.DateTime(cutoffTime.Add(-24 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old2"),
						Created:    libtime.DateTime(cutoffTime.Add(-48 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old3"),
						Created:    libtime.DateTime(cutoffTime.Add(-72 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old4"),
						Created:    libtime.DateTime(cutoffTime.Add(-96 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old5"),
						Created:    libtime.DateTime(cutoffTime.Add(-120 * time.Hour)),
					},
				},
			}

			notificationStoreTx.MapCalls(
				func(ctx context.Context, tx libkv.Tx, fn func(context.Context, core.Notification) error) error {
					for _, n := range oldNotifications {
						if err := fn(ctx, n); err != nil {
							return err
						}
					}
					return nil
				},
			)

			notificationStoreTx.RemoveReturns(nil)
			notificationIndexer.RemoveReturns(nil)
		})

		It("should only remove limit number of notifications", func() {
			err := cleaner.Clean(ctx, tx, maxAge, 3)

			Expect(err).NotTo(HaveOccurred())
			Expect(notificationStoreTx.RemoveCallCount()).To(Equal(3))
			Expect(notificationIndexer.RemoveCallCount()).To(Equal(3))
		})
	})

	Context("when context is cancelled", func() {
		var cancelCtx context.Context
		var cancel context.CancelFunc

		BeforeEach(func() {
			cancelCtx, cancel = context.WithCancel(context.Background())

			// Create multiple old notifications
			oldNotifications := []core.Notification{
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old1"),
						Created:    libtime.DateTime(cutoffTime.Add(-24 * time.Hour)),
					},
				},
				{
					Object: base.Object[base.Identifier]{
						Identifier: base.Identifier("old2"),
						Created:    libtime.DateTime(cutoffTime.Add(-48 * time.Hour)),
					},
				},
			}

			notificationStoreTx.MapCalls(
				func(ctx context.Context, tx libkv.Tx, fn func(context.Context, core.Notification) error) error {
					for _, n := range oldNotifications {
						if err := fn(ctx, n); err != nil {
							return err
						}
					}
					return nil
				},
			)

			// Cancel context after first removal
			callCount := 0
			notificationStoreTx.RemoveStub = func(ctx context.Context, tx libkv.Tx, id base.Identifier) error {
				callCount++
				if callCount == 1 {
					cancel() // Cancel after first removal
				}
				return nil
			}

			notificationIndexer.RemoveReturns(nil)
		})

		It("should stop cleanup and return context error", func() {
			err := cleaner.Clean(cancelCtx, tx, maxAge, 10000)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context cancelled during cleanup"))
			// Should have removed only 1 notification before cancellation
			Expect(notificationStoreTx.RemoveCallCount()).To(Equal(1))
		})
	})
})
