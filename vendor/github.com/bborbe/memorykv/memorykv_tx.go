// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	"context"
	"sync"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

func NewTx() libkv.Tx {
	return &tx{
		data: make(map[string]libkv.Bucket),
	}
}

type tx struct {
	mux  sync.Mutex
	data map[string]libkv.Bucket
}

func (t *tx) ListBucketNames(ctx context.Context) (libkv.BucketNames, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	result := libkv.BucketNames{}
	for bucketName := range t.data {
		result = append(result, libkv.BucketName(bucketName))
	}
	return result, nil
}

func (t *tx) Bucket(ctx context.Context, name libkv.BucketName) (libkv.Bucket, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	bucket, ok := t.data[name.String()]
	if !ok {
		return nil, errors.Wrapf(ctx, libkv.BucketNotFoundError, "bucket %s not found", name)
	}
	return bucket, nil
}

func (t *tx) CreateBucket(ctx context.Context, name libkv.BucketName) (libkv.Bucket, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	_, ok := t.data[name.String()]
	if ok {
		return nil, errors.Wrapf(
			ctx,
			libkv.BucketAlreadyExistsError,
			"bucket %s already exists",
			name,
		)
	}
	bucket := NewBucket()
	t.data[name.String()] = bucket
	return bucket, nil
}

func (t *tx) CreateBucketIfNotExists(
	ctx context.Context,
	name libkv.BucketName,
) (libkv.Bucket, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	bucket, ok := t.data[name.String()]
	if ok {
		return bucket, nil
	}
	bucket = NewBucket()
	t.data[name.String()] = bucket
	return bucket, nil
}

func (t *tx) DeleteBucket(ctx context.Context, name libkv.BucketName) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	_, ok := t.data[name.String()]
	if !ok {
		return errors.Wrapf(ctx, libkv.BucketNotFoundError, "bucket %s not found", name)
	}
	delete(t.data, name.String())
	return nil
}
