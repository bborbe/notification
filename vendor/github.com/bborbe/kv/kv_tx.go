// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import (
	"context"
	"errors"
)

// ErrBucketNotFound is returned when attempting to access a bucket that does not exist.
var ErrBucketNotFound = errors.New("bucket not found")

// BucketNotFoundError is deprecated: use ErrBucketNotFound instead.
var BucketNotFoundError = ErrBucketNotFound

// ErrBucketAlreadyExists is returned when attempting to create a bucket that already exists.
var ErrBucketAlreadyExists = errors.New("bucket already exists")

// BucketAlreadyExistsError is deprecated: use ErrBucketAlreadyExists instead.
var BucketAlreadyExistsError = ErrBucketAlreadyExists

//counterfeiter:generate -o mocks/tx.go --fake-name Tx . Tx

// Tx represents a database transaction that provides bucket management operations.
type Tx interface {
	Bucket(ctx context.Context, name BucketName) (Bucket, error)
	CreateBucket(ctx context.Context, name BucketName) (Bucket, error)
	CreateBucketIfNotExists(ctx context.Context, name BucketName) (Bucket, error)
	DeleteBucket(ctx context.Context, name BucketName) error
	ListBucketNames(ctx context.Context) (BucketNames, error)
}
