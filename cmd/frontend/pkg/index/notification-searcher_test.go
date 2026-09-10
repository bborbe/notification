// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index_test

import (
	"context"

	"github.com/bborbe/cqrs/base"
	libtime "github.com/bborbe/time"
	"github.com/blevesearch/bleve/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
	libindex "github.com/bborbe/notification/index"
	libtest "github.com/bborbe/notification/test"
)

var _ = Describe("NotificationIndex", func() {
	var ctx context.Context
	var err error
	var from *libtime.DateTime
	var until *libtime.DateTime
	var types core.NotificationTypes
	var targets []core.NotificationTarget
	var brokerIdentifiers []string
	var accountIdentifiers []string
	var stages []string
	var query string
	var limit int
	var result base.Identifiers
	var notifications []core.Notification
	var bleveIndex bleve.Index
	var notificationIndexer index.NotificationIndexer
	var notificationIndexSearcher index.NotificationIndexSearcher
	BeforeEach(func() {
		ctx = context.Background()

		notifications = []core.Notification{}
		from = nil
		until = nil
		types = nil
		targets = nil
		brokerIdentifiers = nil
		accountIdentifiers = nil
		stages = nil
		query = ""
		limit = 100

		bleveIndex, err = libindex.OpenTemp(ctx, index.NewIndexMapping())
		Expect(err).To(BeNil())

		notificationIndexer = index.NewNotificationIndexer(bleveIndex)
		notificationIndexSearcher = index.NewNotificationIndexSearcher(bleveIndex)
	})
	AfterEach(func() {
		_ = bleveIndex.Close()
	})
	Context("Search", func() {
		JustBeforeEach(func() {
			for _, notification := range notifications {
				Expect(notificationIndexer.Add(ctx, notification))
			}
			result, err = notificationIndexSearcher.Search(
				ctx,
				from,
				until,
				types,
				targets,
				brokerIdentifiers,
				accountIdentifiers,
				stages,
				query,
				limit,
			)
		})
		Context("empty bleveIndex", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns no results", func() {
				Expect(result).To(HaveLen(0))
			})
		})
		Context("with data", func() {
			BeforeEach(func() {
				notifications = []core.Notification{
					{
						Object: base.Object[base.Identifier]{
							Identifier: "1",
							Created:    libtime.DateTime(libtest.ParseTime("2023-07-19T20:05:00Z")),
						},
						Type:    core.TestNotificationType,
						Message: "Info message",
						Target:  core.NotificationTarget("slack").Ptr(),
						Metadata: map[string]string{
							"brokerIdentifier":  "capitalcom",
							"accountIdentifier": "account-1",
							"stage":             "live",
						},
					},
					{
						Object: base.Object[base.Identifier]{
							Identifier: "2",
							Created:    libtime.DateTime(libtest.ParseTime("2023-07-21T20:05:00Z")),
						},
						Type:    core.SignalNotificationType,
						Message: "Error message",
						Target:  core.NotificationTarget("slack").Ptr(),
						Metadata: map[string]string{
							"brokerIdentifier":  "capitalcom",
							"accountIdentifier": "account-2",
							"stage":             "demo",
						},
					},
					{
						Object: base.Object[base.Identifier]{
							Identifier: "3",
							Created:    libtime.DateTime(libtest.ParseTime("2023-07-10T20:05:00Z")),
						},
						Type:    core.BacktestCompletedNotificationType,
						Message: "Warning message",
						Target:  core.NotificationTarget("discord").Ptr(),
						Metadata: map[string]string{
							"brokerIdentifier":  "darwinex",
							"accountIdentifier": "account-3",
							"stage":             "live",
						},
					},
				}
			})
			Context("without filter", func() {
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns all results", func() {
					Expect(result).To(HaveLen(3))
					Expect(result[0]).To(Equal(base.Identifier("2")))
					Expect(result[1]).To(Equal(base.Identifier("1")))
					Expect(result[2]).To(Equal(base.Identifier("3")))
				})
			})
			Context("limit 2", func() {
				BeforeEach(func() {
					limit = 2
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns limited results", func() {
					Expect(result).To(HaveLen(2))
				})
			})
			Context("type filter", func() {
				BeforeEach(func() {
					types = core.NotificationTypes{core.SignalNotificationType}
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns filtered results", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("2")))
				})
			})
			Context("target filter", func() {
				BeforeEach(func() {
					targets = []core.NotificationTarget{core.NotificationTarget("discord")}
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns filtered results", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("3")))
				})
			})
			Context("broker filter", func() {
				BeforeEach(func() {
					brokerIdentifiers = []string{"darwinex"}
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns filtered results", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("3")))
				})
			})
			Context("account filter", func() {
				BeforeEach(func() {
					accountIdentifiers = []string{"account-2"}
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns filtered results", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("2")))
				})
			})
			Context("stage filter", func() {
				BeforeEach(func() {
					stages = []string{"demo"}
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns filtered results", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("2")))
				})
			})
			Context("message query", func() {
				BeforeEach(func() {
					query = "Error"
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns matching results", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("2")))
				})
			})
			Context("from", func() {
				BeforeEach(func() {
					from = libtime.DateTime(libtest.ParseTime("2023-07-21T20:05:00Z")).Ptr()
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns notifications from the specified date forward", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("2")))
				})
			})
			Context("from (regression test for date range bug)", func() {
				BeforeEach(func() {
					from = libtime.DateTime(libtest.ParseTime("2023-07-15T00:00:00Z")).Ptr()
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns notifications after the from date when until is nil", func() {
					// This validates the fix for the bug where until=nil was set to zero time
					Expect(result).To(HaveLen(2))
					Expect(result).To(ContainElements(
						base.Identifier("1"), // 2023-07-19
						base.Identifier("2"), // 2023-07-21
					))
				})
				It("does not return notifications before the from date", func() {
					Expect(result).ToNot(ContainElement(base.Identifier("3"))) // 2023-07-10
				})
			})
			Context("until", func() {
				BeforeEach(func() {
					until = libtime.DateTime(libtest.ParseTime("2023-07-21T20:05:00Z")).Ptr()
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns notifications up to the specified date", func() {
					Expect(result).To(HaveLen(2))
					Expect(result).To(ContainElements(
						base.Identifier("1"), // 2023-07-19
						base.Identifier("3"), // 2023-07-10
					))
				})
			})
			Context("from and until together", func() {
				BeforeEach(func() {
					from = libtime.DateTime(libtest.ParseTime("2023-07-15T00:00:00Z")).Ptr()
					until = libtime.DateTime(libtest.ParseTime("2023-07-20T23:59:59Z")).Ptr()
				})
				It("returns no error", func() {
					Expect(err).To(BeNil())
				})
				It("returns notifications within the date range", func() {
					Expect(result).To(HaveLen(1))
					Expect(result[0]).To(Equal(base.Identifier("1"))) // 2023-07-19
				})
			})
		})
	})
})
