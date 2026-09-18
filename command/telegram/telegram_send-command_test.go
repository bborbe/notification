// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram_test

import (
	"context"
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/command/telegram"
	telegram2 "github.com/bborbe/notification/telegram"
)

var _ = Describe("SendCommand", func() {
	var ctx context.Context
	var err error
	var content []byte
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
		Context("empty bot", func() {
			BeforeEach(func() {
				sendCommand.Bot = ""
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("bot set", func() {
			BeforeEach(func() {
				sendCommand.Bot = "noise"
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
	Context("Unmarshal", func() {
		var unmarshalled telegram.SendCommand
		JustBeforeEach(func() {
			err = json.Unmarshal(content, &unmarshalled)
		})
		Context("without bot field", func() {
			BeforeEach(func() {
				content = []byte(`{"chatId":"112230768","message":"hello world"}`)
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("defaults bot to empty", func() {
				Expect(unmarshalled.Bot).To(Equal(telegram2.Bot("")))
			})
			It("keeps the other fields", func() {
				Expect(unmarshalled.ChatID).To(Equal(telegram2.ChatID("112230768")))
				Expect(unmarshalled.Message).To(Equal(telegram2.Message("hello world")))
			})
		})
		Context("with bot field", func() {
			BeforeEach(func() {
				content = []byte(`{"chatId":"112230768","message":"hello world","bot":"noise"}`)
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("reads the bot", func() {
				Expect(unmarshalled.Bot).To(Equal(telegram2.Bot("noise")))
			})
		})
	})
	Context("Marshal", func() {
		JustBeforeEach(func() {
			content, err = json.Marshal(sendCommand)
		})
		Context("without bot", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("omits the bot field", func() {
				Expect(string(content)).NotTo(ContainSubstring("bot"))
			})
		})
		Context("with bot", func() {
			BeforeEach(func() {
				sendCommand.Bot = "noise"
			})
			It("writes the bot field", func() {
				Expect(string(content)).To(ContainSubstring(`"bot":"noise"`))
			})
		})
	})
})
