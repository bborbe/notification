// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification_test

import (
	"context"

	"github.com/bborbe/cqrs/base"
	cqrsiam "github.com/bborbe/cqrs/iam"
	cqrsmocks "github.com/bborbe/cqrs/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/command/notification"
	"github.com/bborbe/notification"
)

var _ = Describe("NotificationPublishCommandSender", func() {
	var ctx context.Context
	var err error
	var notificationPublishCommandSender notification.NotificationPublishCommandSender
	var commandObjectSender *cqrsmocks.CDBCommandObjectSender
	var publishCommand notification.NotificationPublishCommand
	BeforeEach(func() {
		ctx = context.Background()
		publishCommand = notification.NotificationPublishCommand{
			Type:    core.SignalNotificationType,
			Message: "test notification",
		}

		commandObjectSender = &cqrsmocks.CDBCommandObjectSender{}
		notificationPublishCommandSender = notification.NewNotificationPublishCommandSender(
			base.NewCommandCreator(
				base.RequestIDChannel(ctx),
			),
			commandObjectSender,
			cqrsiam.Initiator("test"),
		)
	})
	JustBeforeEach(func() {
		err = notificationPublishCommandSender.SendPublishNotificationCommand(ctx, publishCommand)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("calls send", func() {
		Expect(commandObjectSender.SendCommandObjectCallCount()).To(Equal(1))
	})
})
