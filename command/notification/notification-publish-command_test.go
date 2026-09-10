// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/command/notification"
	"github.com/bborbe/notification"
)

var _ = Describe("NotificationPublishCommand", func() {
	var ctx context.Context
	var err error
	var notificationPublishCommand notification.NotificationPublishCommand
	BeforeEach(func() {
		ctx = context.Background()
		notificationPublishCommand = notification.NotificationPublishCommand{
			Type:     core.SignalNotificationType,
			Message:  "hello world",
			Target:   nil,
			Metadata: map[string]string{},
		}
	})
	Context("Validate", func() {
		BeforeEach(func() {
			err = notificationPublishCommand.Validate(ctx)
		})
		Context("valid without target", func() {
			BeforeEach(func() {
				notificationPublishCommand.Target = nil
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("valid with target", func() {
			BeforeEach(func() {
				notificationPublishCommand.Target = core.NotificationTarget("discord").Ptr()
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
