// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification_test

import (
	"context"
	"time"

	libtime "github.com/bborbe/time"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/command/notification"
)

var _ = Describe("NotificationCleanCommand", func() {
	var ctx context.Context
	var err error
	var notificationCleanCommand notification.NotificationCleanCommand
	BeforeEach(func() {
		ctx = context.Background()
		notificationCleanCommand = notification.NotificationCleanCommand{
			MaxAge: libtime.Duration(720 * time.Hour), // 30 days
			Limit:  10000,
		}
	})
	Context("Validate", func() {
		JustBeforeEach(func() {
			err = notificationCleanCommand.Validate(ctx)
		})
		Context("valid with positive duration and limit", func() {
			BeforeEach(func() {
				notificationCleanCommand.MaxAge = libtime.Duration(720 * time.Hour)
				notificationCleanCommand.Limit = 10000
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("valid with 1 hour and 1 limit", func() {
			BeforeEach(func() {
				notificationCleanCommand.MaxAge = libtime.Duration(1 * time.Hour)
				notificationCleanCommand.Limit = 1
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("invalid with zero duration", func() {
			BeforeEach(func() {
				notificationCleanCommand.MaxAge = libtime.Duration(0)
				notificationCleanCommand.Limit = 10000
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("MaxAge"))
				Expect(err.Error()).To(ContainSubstring("must be positive"))
			})
		})
		Context("invalid with negative duration", func() {
			BeforeEach(func() {
				notificationCleanCommand.MaxAge = libtime.Duration(-24 * time.Hour)
				notificationCleanCommand.Limit = 10000
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("MaxAge"))
				Expect(err.Error()).To(ContainSubstring("must be positive"))
			})
		})
		Context("invalid with zero limit", func() {
			BeforeEach(func() {
				notificationCleanCommand.MaxAge = libtime.Duration(720 * time.Hour)
				notificationCleanCommand.Limit = 0
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("Limit"))
				Expect(err.Error()).To(ContainSubstring("must be positive"))
			})
		})
		Context("invalid with negative limit", func() {
			BeforeEach(func() {
				notificationCleanCommand.MaxAge = libtime.Duration(720 * time.Hour)
				notificationCleanCommand.Limit = -100
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("Limit"))
				Expect(err.Error()).To(ContainSubstring("must be positive"))
			})
		})
	})
	Context("Ptr", func() {
		It("returns pointer to command", func() {
			ptr := notificationCleanCommand.Ptr()
			Expect(ptr).NotTo(BeNil())
			Expect(*ptr).To(Equal(notificationCleanCommand))
		})
	})
})
