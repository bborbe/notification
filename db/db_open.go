// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"context"

	libbadgerkv "github.com/bborbe/badgerkv"
	libboltkv "github.com/bborbe/boltkv"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	libmemorykv "github.com/bborbe/memorykv"
	bolt "go.etcd.io/bbolt"
)

func OpenBoltDB(ctx context.Context, dataDir string, noSync bool) (libkv.DB, error) {
	db, err := libboltkv.OpenDir(ctx, dataDir, func(opts *bolt.Options) {
		opts.NoSync = noSync
	})
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "open bolt db failed")
	}
	return libkv.NewDBWithMetrics(
		db,
		libkv.NewMetrics(),
	), nil
}

func OpenBadgerDB(ctx context.Context, dataDir string) (libkv.DB, error) {
	db, err := libbadgerkv.OpenPath(ctx, dataDir)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "open badger db failed")
	}
	return libkv.NewDBWithMetrics(
		db,
		libkv.NewMetrics(),
	), nil
}

func OpenMemoryDB(ctx context.Context) (libkv.DB, error) {
	db, err := libmemorykv.OpenMemory(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "open memory db failed")
	}
	return libkv.NewDBWithMetrics(
		db,
		libkv.NewMetrics(),
	), nil
}
