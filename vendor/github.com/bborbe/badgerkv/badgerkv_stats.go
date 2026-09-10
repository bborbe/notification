// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package badgerkv

import (
	"context"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

// Stats returns a fast overview: total LSM + value-log size and bucket inventory
// (names only). Per-bucket KeyCount is left at zero — call StatsDetailed for
// counts, but note that Badger has no native per-prefix counter so detailed
// counting requires scanning every key in each bucket.
func (b *badgerdb) Stats(ctx context.Context) (*libkv.Stats, error) {
	return b.statsImpl(ctx, false)
}

// StatsDetailed returns Stats plus per-bucket KeyCount. Cost: O(total keys)
// — Badger scans every key with each bucket's prefix. Do not poll hot.
func (b *badgerdb) StatsDetailed(ctx context.Context) (*libkv.Stats, error) {
	return b.statsImpl(ctx, true)
}

func (b *badgerdb) statsImpl(ctx context.Context, detailed bool) (*libkv.Stats, error) {
	lsm, vlog := b.db.Size()
	s := &libkv.Stats{
		Backend:  "badger",
		SizeB:    lsm + vlog,
		Detailed: detailed,
	}
	err := b.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
		names, err := tx.ListBucketNames(ctx)
		if err != nil {
			return errors.Wrapf(ctx, err, "list bucket names failed")
		}
		for _, name := range names {
			select {
			case <-ctx.Done():
				return errors.Wrap(ctx, ctx.Err(), "context cancelled")
			default:
			}

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
