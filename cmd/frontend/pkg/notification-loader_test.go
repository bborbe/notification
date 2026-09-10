// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg_test

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	libkvmocks "github.com/bborbe/kv/mocks"
	"github.com/bborbe/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/pkg"
	"github.com/bborbe/notification"
	coremocks "github.com/bborbe/notification/mocks"
)

var _ = Describe("NotificationLoader", func() {
	var (
		ctx                 context.Context
		tx                  *libkvmocks.Tx
		notificationStoreTx *coremocks.CoreNotificationStoreTx
		loader              pkg.NotificationLoader
		identifiers         base.Identifiers
	)

	BeforeEach(func() {
		ctx = context.Background()
		tx = &libkvmocks.Tx{}
		notificationStoreTx = &coremocks.CoreNotificationStoreTx{}

		loader = pkg.NewNotificationLoader(
			notificationStoreTx,
			log.DefaultSamplerFactory,
		)

		identifiers = base.Identifiers{
			base.Identifier("id1"),
			base.Identifier("id2"),
			base.Identifier("id3"),
		}
	})

	Context("when all notifications are found", func() {
		var notifications []*core.Notification

		BeforeEach(func() {
			notifications = []*core.Notification{
				{
					Object:  base.Object[base.Identifier]{Identifier: base.Identifier("id1")},
					Message: "msg1",
				},
				{
					Object:  base.Object[base.Identifier]{Identifier: base.Identifier("id2")},
					Message: "msg2",
				},
				{
					Object:  base.Object[base.Identifier]{Identifier: base.Identifier("id3")},
					Message: "msg3",
				},
			}

			notificationStoreTx.GetCalls(
				func(ctx context.Context, tx libkv.Tx, id base.Identifier) (*core.Notification, error) {
					for _, n := range notifications {
						if n.Identifier == id {
							return n, nil
						}
					}
					return nil, errors.New(ctx, "not found")
				},
			)
		})

		It("should return all notifications", func() {
			result, err := loader.GetAll(ctx, tx, identifiers)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(HaveLen(3))
			Expect(result[0].Identifier).To(Equal(base.Identifier("id1")))
			Expect(result[1].Identifier).To(Equal(base.Identifier("id2")))
			Expect(result[2].Identifier).To(Equal(base.Identifier("id3")))
		})
	})

	Context("when some notifications are not found", func() {
		BeforeEach(func() {
			notificationStoreTx.GetCalls(
				func(ctx context.Context, tx libkv.Tx, id base.Identifier) (*core.Notification, error) {
					if id == base.Identifier("id1") {
						return &core.Notification{
							Object:  base.Object[base.Identifier]{Identifier: id},
							Message: "msg1",
						}, nil
					}
					if id == base.Identifier("id3") {
						return &core.Notification{
							Object:  base.Object[base.Identifier]{Identifier: id},
							Message: "msg3",
						}, nil
					}
					return nil, errors.New(ctx, "not found")
				},
			)
		})

		It("should return only found notifications", func() {
			result, err := loader.GetAll(ctx, tx, identifiers)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(HaveLen(2))
			Expect(result[0].Identifier).To(Equal(base.Identifier("id1")))
			Expect(result[1].Identifier).To(Equal(base.Identifier("id3")))
		})
	})

	Context("when no notifications are found", func() {
		BeforeEach(func() {
			notificationStoreTx.GetReturns(nil, errors.New(ctx, "not found"))
		})

		It("should return empty slice", func() {
			result, err := loader.GetAll(ctx, tx, identifiers)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEmpty())
		})
	})

	Context("when identifiers list is empty", func() {
		It("should return empty slice", func() {
			result, err := loader.GetAll(ctx, tx, base.Identifiers{})

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEmpty())
		})
	})
})
