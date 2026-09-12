// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/command/telegram"
)

var _ = Describe("SendCommand", func() {
	var ctx context.Context
	var err error
	var sendCommand telegram.SendCommand
	BeforeEach(func() {
		ctx = context.Background()
		sendCommand = telegram.SendCommand{
			ChatID:  "112230768",
			Message: "hello world",
		}
	})
	Context("Validate", func() {
		JustBeforeEach(func() {
			err = sendCommand.Validate(ctx)
		})
		Context("success", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("invalid chat id", func() {
			BeforeEach(func() {
				sendCommand.ChatID = ""
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
		Context("invalid message", func() {
			BeforeEach(func() {
				sendCommand.Message = ""
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
