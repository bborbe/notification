// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler_test

import (
	"context"
	"encoding/json"
	"net/http/httptest"

	"github.com/bborbe/cqrs/base"
	libhttp "github.com/bborbe/http"
	libkv "github.com/bborbe/kv"
	libmemorykv "github.com/bborbe/memorykv"
	"github.com/blevesearch/bleve/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/pkg/handler"
	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
	libindex "github.com/bborbe/notification/index"
)

var _ = Describe("StatusHandler", func() {
	var (
		ctx        context.Context
		db         libkv.DB
		bleveIndex bleve.Index
		bucketName libkv.BucketName
		h          libhttp.JSONHandlerTx
	)

	BeforeEach(func() {
		ctx = context.Background()

		var err error
		db, err = libmemorykv.OpenMemory(ctx)
		Expect(err).NotTo(HaveOccurred())

		bleveIndex, err = libindex.OpenTemp(ctx, index.NewIndexMapping())
		Expect(err).NotTo(HaveOccurred())

		bucketName = libkv.NewBucketName("notification-store")

		h = handler.NewStatusHandler(bleveIndex, bucketName)
	})

	AfterEach(func() {
		if bleveIndex != nil {
			bleveIndex.Close()
		}
		db.Close()
	})

	It("returns zero counts for empty database and index", func() {
		req := httptest.NewRequest("GET", "/status", nil)

		var result interface{}
		var err error
		err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
			result, err = h.ServeHTTP(ctx, tx, req)
			return err
		})
		Expect(err).NotTo(HaveOccurred())

		bytes, err := json.Marshal(result)
		Expect(err).NotTo(HaveOccurred())

		var status map[string]interface{}
		err = json.Unmarshal(bytes, &status)
		Expect(err).NotTo(HaveOccurred())

		Expect(status["index"]).To(BeEquivalentTo(0))
		Expect(status["store"]).To(BeEquivalentTo(0))
	})

	It("returns correct store count after adding notifications", func() {
		notificationStoreTx := core.NewNotificationStoreTx()

		err := db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
			return notificationStoreTx.Add(ctx, tx, core.Notification{
				Object:  base.Object[base.Identifier]{Identifier: "test-1"},
				Type:    "test",
				Message: "test notification",
			})
		})
		Expect(err).NotTo(HaveOccurred())

		req := httptest.NewRequest("GET", "/status", nil)

		var result interface{}
		err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
			result, err = h.ServeHTTP(ctx, tx, req)
			return err
		})
		Expect(err).NotTo(HaveOccurred())

		bytes, err := json.Marshal(result)
		Expect(err).NotTo(HaveOccurred())

		var status map[string]interface{}
		err = json.Unmarshal(bytes, &status)
		Expect(err).NotTo(HaveOccurred())

		Expect(status["store"]).To(BeEquivalentTo(1))
	})

	It("returns correct index count after indexing notifications", func() {
		notificationIndexer := index.NewNotificationIndexer(libindex.NewIndex(bleveIndex))

		err := notificationIndexer.Add(ctx, core.Notification{
			Object:  base.Object[base.Identifier]{Identifier: "test-1"},
			Type:    "test",
			Message: "test notification",
		})
		Expect(err).NotTo(HaveOccurred())

		req := httptest.NewRequest("GET", "/status", nil)

		var result interface{}
		err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
			result, err = h.ServeHTTP(ctx, tx, req)
			return err
		})
		Expect(err).NotTo(HaveOccurred())

		bytes, err := json.Marshal(result)
		Expect(err).NotTo(HaveOccurred())

		var status map[string]interface{}
		err = json.Unmarshal(bytes, &status)
		Expect(err).NotTo(HaveOccurred())

		Expect(status["index"]).To(BeEquivalentTo(1))
	})
})
