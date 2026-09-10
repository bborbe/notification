// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	libkv "github.com/bborbe/kv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/db"
)

var _ = Describe("Stats Handler", func() {
	var (
		ctx     context.Context
		libDB   libkv.DB
		handler http.Handler
	)

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		libDB, err = db.OpenMemoryDB(ctx)
		Expect(err).To(BeNil())
		handler = db.NewStatsHandler(libDB)
	})

	AfterEach(func() {
		_ = libDB.Close()
	})

	It("returns a non-nil handler", func() {
		Expect(handler).NotTo(BeNil())
	})

	It("responds 200 with application/json", func() {
		req := httptest.NewRequest(http.MethodGet, "/dbstats", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Header().Get("Content-Type")).To(ContainSubstring("application/json"))
	})

	It("returns backend name in JSON body", func() {
		req := httptest.NewRequest(http.MethodGet, "/dbstats", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		var got libkv.Stats
		Expect(json.Unmarshal(rr.Body.Bytes(), &got)).To(Succeed())
		Expect(got.Backend).To(Equal("memory"))
	})

	It("returns empty buckets for fresh db", func() {
		req := httptest.NewRequest(http.MethodGet, "/dbstats", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		var got libkv.Stats
		Expect(json.Unmarshal(rr.Body.Bytes(), &got)).To(Succeed())
		Expect(got.Buckets).To(BeEmpty())
	})

	It("returns bucket names but leaves KeyCount at zero for fast Stats", func() {
		bucketName := libkv.NewBucketName("test-bucket")
		err := libDB.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(ctx, bucketName)
			Expect(err).To(BeNil())
			Expect(bucket.Put(ctx, []byte("k1"), []byte("v1"))).To(Succeed())
			Expect(bucket.Put(ctx, []byte("k2"), []byte("v2"))).To(Succeed())
			return nil
		})
		Expect(err).To(BeNil())

		req := httptest.NewRequest(http.MethodGet, "/dbstats", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		var got libkv.Stats
		Expect(json.Unmarshal(rr.Body.Bytes(), &got)).To(Succeed())
		Expect(got.Detailed).To(BeFalse())
		Expect(got.Buckets).To(HaveLen(1))
		Expect(got.Buckets[0].Name).To(Equal(bucketName))
		Expect(got.Buckets[0].KeyCount).To(Equal(int64(0)))
	})

	It("reports bucket key counts when details=true", func() {
		bucketName := libkv.NewBucketName("test-bucket")
		err := libDB.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(ctx, bucketName)
			Expect(err).To(BeNil())
			Expect(bucket.Put(ctx, []byte("k1"), []byte("v1"))).To(Succeed())
			Expect(bucket.Put(ctx, []byte("k2"), []byte("v2"))).To(Succeed())
			return nil
		})
		Expect(err).To(BeNil())

		req := httptest.NewRequest(http.MethodGet, "/dbstats?details=true", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		var got libkv.Stats
		Expect(json.Unmarshal(rr.Body.Bytes(), &got)).To(Succeed())
		Expect(got.Detailed).To(BeTrue())
		Expect(got.Buckets).To(HaveLen(1))
		Expect(got.Buckets[0].Name).To(Equal(bucketName))
		Expect(got.Buckets[0].KeyCount).To(Equal(int64(2)))
	})
})
