// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"time"

	libkv "github.com/bborbe/kv"
	"github.com/bborbe/log"
	libmemorykv "github.com/bborbe/memorykv"
	libtime "github.com/bborbe/time"
	libtimetest "github.com/bborbe/time/test"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/pkg"
	"github.com/bborbe/notification/cmd/frontend/pkg/handler"
	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
	libindex "github.com/bborbe/notification/index"
)

var _ = Describe("CleanHandler", func() {
	var (
		ctx             context.Context
		cancel          context.CancelFunc
		db              libkv.DB
		currentDateTime libtime.CurrentDateTime
		defaultMaxAge   libtime.Duration
		defaultLimit    int
		cleaner         pkg.NotificationCleaner
		h               http.Handler
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		var err error
		db, err = libmemorykv.OpenMemory(ctx)
		Expect(err).NotTo(HaveOccurred())

		bleveIndex, err := libindex.OpenTemp(ctx, index.NewIndexMapping())
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(bleveIndex.Close)

		currentDateTime = libtime.NewCurrentDateTime()
		currentDateTime.SetNow(libtimetest.ParseDateTime("2026-02-28T00:00:00Z"))
		defaultMaxAge = libtime.Duration(30 * 24 * libtime.Hour)
		defaultLimit = 10000

		cleaner = pkg.NewNotificationCleaner(
			currentDateTime,
			index.NewNotificationIndexer(libindex.NewIndex(bleveIndex)),
			core.NewNotificationStoreTx(),
			log.DefaultSamplerFactory,
		)

		h = handler.NewCleanHandler(ctx, db, cleaner, defaultMaxAge, defaultLimit)
	})

	AfterEach(func() {
		cancel()
		time.Sleep(50 * time.Millisecond) // let background goroutine complete before closing DB
		db.Close()
	})

	It("returns 200 with no query params", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("returns 200 with valid maxAge", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?maxAge=168h", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("returns 200 with valid cleanLimit", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?cleanLimit=5000", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("returns 200 with both params", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?maxAge=72h&cleanLimit=1000", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("returns 200 when cleanLimit is zero (error logged in background)", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?cleanLimit=0", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("returns 200 when cleanLimit is negative (error logged in background)", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?cleanLimit=-1", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("uses default maxAge when param is invalid", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?maxAge=notaduration", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("returns 200 when maxAge is zero (error logged in background)", func() {
		req := httptest.NewRequest(http.MethodGet, "/clean?maxAge=0s", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})
})
