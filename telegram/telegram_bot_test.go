// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/telegram"
)

var _ = Describe("Bot", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	Context("Validate", func() {
		It("accepts the empty bot as the default bot", func() {
			Expect(telegram.Bot("").Validate(ctx)).To(BeNil())
		})
		It("accepts a named bot", func() {
			Expect(telegram.Bot("noise").Validate(ctx)).To(BeNil())
		})
	})
	Context("String", func() {
		It("returns the value", func() {
			Expect(telegram.Bot("noise").String()).To(Equal("noise"))
		})
	})
	Context("Ptr", func() {
		It("returns a pointer to the value", func() {
			Expect(*telegram.Bot("noise").Ptr()).To(Equal(telegram.Bot("noise")))
		})
	})
})
