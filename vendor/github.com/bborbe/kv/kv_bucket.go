// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import (
	"context"
	"errors"
)

// ErrKeyNotFound is currently not returned, but can be used as common error.
var ErrKeyNotFound = errors.New("key not found")

// KeyNotFoundError is deprecated: use ErrKeyNotFound instead.
var KeyNotFoundError = ErrKeyNotFound

//counterfeiter:generate -o mocks/bucket.go --fake-name Bucket . Bucket

// Bucket represents a key-value bucket within a transaction that supports CRUD operations and iteration.
type Bucket interface {
	Put(ctx context.Context, key []byte, value []byte) error
	Get(ctx context.Context, bytes []byte) (Item, error)
	Delete(ctx context.Context, bytes []byte) error
	Iterator() Iterator
	IteratorReverse() Iterator
}
