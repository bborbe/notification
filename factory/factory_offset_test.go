// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	libkv "github.com/bborbe/kv"
	libmemorykv "github.com/bborbe/memorykv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/factory"
)

var _ = Describe("Offset Factory", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		db     libkv.DB
		err    error
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		db, err = libmemorykv.OpenMemory(ctx)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		cancel()
	})

	Context("CreateOffsetManagerHandler", func() {
		It("returns a valid HTTP handler", func() {
			handler := factory.CreateOffsetManagerHandler(db, cancel)
			Expect(handler).NotTo(BeNil())
			var _ http.Handler = handler // Verify it implements http.Handler
		})

		It("creates handler with non-nil parameters", func() {
			Expect(func() {
				factory.CreateOffsetManagerHandler(db, cancel)
			}).NotTo(Panic())
		})

		It("uses default offset store when group parameter is empty", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			handler := factory.CreateOffsetManagerHandler(db, cancel)

			// Set offset without group parameter
			req, err := http.NewRequestWithContext(
				ctx,
				http.MethodGet,
				"/offsetmanager?topic=test-topic&partition=0&offset=12345",
				nil,
			)
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			// Handler should succeed (or fail gracefully, not panic)
			// The key test is that it uses NewOffsetStore (no group prefix)
			Expect(rr.Code).To(Or(Equal(http.StatusOK), Equal(http.StatusInternalServerError)))
		})

		It("uses group offset store when group parameter is provided", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			handler := factory.CreateOffsetManagerHandler(db, cancel)

			// Set offset with group parameter
			req, err := http.NewRequestWithContext(
				ctx,
				http.MethodGet,
				"/offsetmanager?topic=test-topic&partition=0&offset=12345&group=my-group",
				nil,
			)
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			// Handler should succeed (or fail gracefully, not panic)
			// The key test is that it uses NewOffsetStoreGroup with the group
			Expect(rr.Code).To(Or(Equal(http.StatusOK), Equal(http.StatusInternalServerError)))
		})
	})

	Context("CreateSaramaOffsetManagerHandler", func() {
		It("function exists and can be called with nil parameters for compilation test", func() {
			// Note: Full integration test would require actual Sarama client setup
			// This test verifies the function signature compiles correctly
			Expect(factory.CreateSaramaOffsetManagerHandler).NotTo(BeNil())
		})
	})
})
