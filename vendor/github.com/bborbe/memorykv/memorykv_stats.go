// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	"context"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

// Stats returns a fast overview: backend name and bucket inventory (names only).
// Per-bucket KeyCount is left at zero — call StatsDetailed for that.
func (b *memorydb) Stats(ctx context.Context) (*libkv.Stats, error) {
	return b.statsImpl(ctx, false)
}

// StatsDetailed returns Stats plus per-bucket KeyCount. Cheap on the in-memory
// backend, but kept consistent with the libkv contract.
func (b *memorydb) StatsDetailed(ctx context.Context) (*libkv.Stats, error) {
	return b.statsImpl(ctx, true)
}

func (b *memorydb) statsImpl(ctx context.Context, detailed bool) (*libkv.Stats, error) {
	s := &libkv.Stats{
		Backend:  "memory",
		Detailed: detailed,
	}
	err := b.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
		names, err := tx.ListBucketNames(ctx)
		if err != nil {
			return errors.Wrapf(ctx, err, "list bucket names failed")
		}
		for _, name := range names {
			bs := libkv.BucketStats{Name: name}
			if detailed {
				bucket, err := tx.Bucket(ctx, name)
				if err != nil {
					return errors.Wrapf(ctx, err, "get bucket %s failed", name)
				}
				count, err := libkv.Count(ctx, bucket)
				if err != nil {
					return errors.Wrapf(ctx, err, "count bucket %s failed", name)
				}
				bs.KeyCount = count
			}
			s.Buckets = append(s.Buckets, bs)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "stats failed")
	}
	return s, nil
}
