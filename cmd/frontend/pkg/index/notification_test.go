// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index_test

import (
	"context"

	"github.com/bborbe/collection"
	"github.com/bborbe/cqrs/base"
	libtime "github.com/bborbe/time"
	libtimetest "github.com/bborbe/time/test"
	"github.com/blevesearch/bleve/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
	libindex "github.com/bborbe/notification/index"
)

var _ = Describe("Notification Index", func() {
	var (
		ctx                  context.Context
		bleveIndex           bleve.Index
		notificationIndexer  index.NotificationIndexer
		notificationSearcher index.NotificationIndexSearcher
	)

	BeforeEach(func() {
		ctx = context.Background()

		var err error
		bleveIndex, err = libindex.OpenTemp(ctx, index.NewIndexMapping())
		Expect(err).To(BeNil())

		notificationIndexer = index.NewNotificationIndexer(libindex.NewIndex(bleveIndex))
		notificationSearcher = index.NewNotificationIndexSearcher(bleveIndex)
	})

	AfterEach(func() {
		if bleveIndex != nil {
			bleveIndex.Close()
		}
	})

	Context("Add and Search", func() {
		It("adds notification to index and finds it", func() {
			notification := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-1"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
				},
				Type:     core.NotificationType("test"),
				Message:  core.NotificationMessage("Test notification message"),
				Target:   collection.Ptr(core.NotificationTarget("test-channel")),
				Metadata: map[string]string{"brokerIdentifier": "test-broker", "accountIdentifier": "test-account", "stage": "demo"},
			}

			err := notificationIndexer.Add(ctx, notification)
			Expect(err).To(BeNil())

			ids, err := notificationSearcher.Search(ctx, nil, nil, nil, nil, nil, nil, nil, "", 10)
			Expect(err).To(BeNil())
			Expect(ids).To(HaveLen(1))
			Expect(ids[0]).To(Equal(base.Identifier("test-notification-1")))
		})

		It("searches by type", func() {
			notification1 := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-1"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
				},
				Type:    core.NotificationType("test"),
				Message: core.NotificationMessage("Test notification message"),
			}

			notification2 := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-2"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T11:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T11:00:00Z")),
				},
				Type:    core.NotificationType("signal"),
				Message: core.NotificationMessage("Signal notification message"),
			}

			err := notificationIndexer.Add(ctx, notification1)
			Expect(err).To(BeNil())

			err = notificationIndexer.Add(ctx, notification2)
			Expect(err).To(BeNil())

			ids, err := notificationSearcher.Search(
				ctx,
				nil,
				nil,
				core.NotificationTypes{core.NotificationType("test")},
				nil,
				nil,
				nil,
				nil,
				"",
				10,
			)
			Expect(err).To(BeNil())
			Expect(ids).To(HaveLen(1))
			Expect(ids[0]).To(Equal(base.Identifier("test-notification-1")))
		})

		It("searches by message content", func() {
			notification := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-1"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
				},
				Type:    core.NotificationType("test"),
				Message: core.NotificationMessage("Buy signal for EURUSD"),
			}

			err := notificationIndexer.Add(ctx, notification)
			Expect(err).To(BeNil())

			ids, err := notificationSearcher.Search(
				ctx,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				"EURUSD",
				10,
			)
			Expect(err).To(BeNil())
			Expect(ids).To(HaveLen(1))
			Expect(ids[0]).To(Equal(base.Identifier("test-notification-1")))
		})

		It("searches by date range", func() {
			notification1 := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-1"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
				},
				Type:    core.NotificationType("test"),
				Message: core.NotificationMessage("Old notification"),
			}

			notification2 := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-2"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-02T10:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-02T10:00:00Z")),
				},
				Type:    core.NotificationType("test"),
				Message: core.NotificationMessage("New notification"),
			}

			err := notificationIndexer.Add(ctx, notification1)
			Expect(err).To(BeNil())

			err = notificationIndexer.Add(ctx, notification2)
			Expect(err).To(BeNil())

			from := collection.Ptr(
				libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T12:00:00Z")),
			)
			ids, err := notificationSearcher.Search(ctx, from, nil, nil, nil, nil, nil, nil, "", 10)
			Expect(err).To(BeNil())
			Expect(ids).To(HaveLen(1))
			Expect(ids[0]).To(Equal(base.Identifier("test-notification-2")))
		})

		It("removes notification from index", func() {
			notification := core.Notification{
				Object: base.Object[base.Identifier]{
					Identifier: base.Identifier("test-notification-1"),
					Created:    libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
					Modified:   libtime.DateTime(libtimetest.ParseDateTime("2024-01-01T10:00:00Z")),
				},
				Type:    core.NotificationType("test"),
				Message: core.NotificationMessage("Test notification message"),
			}

			err := notificationIndexer.Add(ctx, notification)
			Expect(err).To(BeNil())

			ids, err := notificationSearcher.Search(ctx, nil, nil, nil, nil, nil, nil, nil, "", 10)
			Expect(err).To(BeNil())
			Expect(ids).To(HaveLen(1))

			err = notificationIndexer.Remove(ctx, base.Identifier("test-notification-1"))
			Expect(err).To(BeNil())

			ids, err = notificationSearcher.Search(ctx, nil, nil, nil, nil, nil, nil, nil, "", 10)
			Expect(err).To(BeNil())
			Expect(ids).To(HaveLen(0))
		})
	})
})
