// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification_test

import (
	"context"
	"time"

	"github.com/bborbe/cqrs/base"
	cqrsiam "github.com/bborbe/cqrs/iam"
	cqrsmocks "github.com/bborbe/cqrs/mocks"
	libtime "github.com/bborbe/time"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/command/notification"
)

var _ = Describe("NotificationCleanCommandSender", func() {
	var ctx context.Context
	var err error
	var notificationCleanCommandSender notification.NotificationCleanCommandSender
	var commandObjectSender *cqrsmocks.CDBCommandObjectSender
	var cleanCommand notification.NotificationCleanCommand
	BeforeEach(func() {
		ctx = context.Background()
		cleanCommand = notification.NotificationCleanCommand{
			MaxAge: libtime.Duration(720 * time.Hour), // 30 days
		}

		commandObjectSender = &cqrsmocks.CDBCommandObjectSender{}
		notificationCleanCommandSender = notification.NewNotificationCleanCommandSender(
			base.NewCommandCreator(
				base.RequestIDChannel(ctx),
			),
			commandObjectSender,
			cqrsiam.Initiator("test"),
		)
	})
	JustBeforeEach(func() {
		err = notificationCleanCommandSender.SendCleanNotificationCommand(ctx, cleanCommand)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("calls send", func() {
		Expect(commandObjectSender.SendCommandObjectCallCount()).To(Equal(1))
	})
})
