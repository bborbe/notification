// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db_test

import (
	"context"
	"os"

	libkv "github.com/bborbe/kv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/db"
)

var _ = Describe("DB Open", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("OpenMemoryDB", func() {
		It("returns a working libkv.DB", func() {
			d, err := db.OpenMemoryDB(ctx)
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			defer func() { _ = d.Close() }()

			err = d.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				_, err := tx.CreateBucketIfNotExists(ctx, libkv.NewBucketName("b"))
				return err
			})
			Expect(err).To(BeNil())
		})

		It("returns memory-backend Stats", func() {
			d, err := db.OpenMemoryDB(ctx)
			Expect(err).To(BeNil())
			defer func() { _ = d.Close() }()

			stats, err := d.Stats(ctx)
			Expect(err).To(BeNil())
			Expect(stats.Backend).To(Equal("memory"))
		})
	})

	Context("OpenBoltDB", func() {
		var dataDir string

		BeforeEach(func() {
			var err error
			dataDir, err = os.MkdirTemp("", "db_open_test_bolt_")
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			_ = os.RemoveAll(dataDir)
		})

		It("returns a working libkv.DB", func() {
			d, err := db.OpenBoltDB(ctx, dataDir, true)
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			defer func() { _ = d.Close() }()

			err = d.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				_, err := tx.CreateBucketIfNotExists(ctx, libkv.NewBucketName("b"))
				return err
			})
			Expect(err).To(BeNil())
		})

		It("returns bolt-backend Stats", func() {
			d, err := db.OpenBoltDB(ctx, dataDir, true)
			Expect(err).To(BeNil())
			defer func() { _ = d.Close() }()

			stats, err := d.Stats(ctx)
			Expect(err).To(BeNil())
			Expect(stats.Backend).To(Equal("bolt"))
			Expect(stats.SizeB).To(BeNumerically(">", int64(0)))
		})
	})

	Context("OpenBadgerDB", func() {
		var dataDir string

		BeforeEach(func() {
			var err error
			dataDir, err = os.MkdirTemp("", "db_open_test_badger_")
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			_ = os.RemoveAll(dataDir)
		})

		It("returns a working libkv.DB", func() {
			d, err := db.OpenBadgerDB(ctx, dataDir)
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			defer func() { _ = d.Close() }()

			err = d.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				_, err := tx.CreateBucketIfNotExists(ctx, libkv.NewBucketName("b"))
				return err
			})
			Expect(err).To(BeNil())
		})

		It("returns badger-backend Stats", func() {
			d, err := db.OpenBadgerDB(ctx, dataDir)
			Expect(err).To(BeNil())
			defer func() { _ = d.Close() }()

			stats, err := d.Stats(ctx)
			Expect(err).To(BeNil())
			Expect(stats.Backend).To(Equal("badger"))
		})
	})
})
