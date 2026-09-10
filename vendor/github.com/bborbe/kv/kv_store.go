// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import (
	"context"

	"github.com/bborbe/errors"
)

// StoreMapper provides mapping functionality over all key-value pairs in a store.
type StoreMapper[KEY ~[]byte | ~string, OBJECT any] interface {
	Map(ctx context.Context, fn func(ctx context.Context, key KEY, object OBJECT) error) error
}

// StoreAdder provides functionality to add objects to a store.
type StoreAdder[KEY ~[]byte | ~string, OBJECT any] interface {
	Add(ctx context.Context, key KEY, object OBJECT) error
}

// StoreRemover provides functionality to remove objects from a store.
type StoreRemover[KEY ~[]byte | ~string] interface {
	Remove(ctx context.Context, key KEY) error
}

// StoreGetter provides functionality to retrieve objects from a store.
type StoreGetter[KEY ~[]byte | ~string, OBJECT any] interface {
	Get(ctx context.Context, key KEY) (*OBJECT, error)
}

// StoreExists provides functionality to check object existence in a store.
type StoreExists[KEY ~[]byte | ~string, OBJECT any] interface {
	Exists(ctx context.Context, key KEY) (bool, error)
}

// StoreStreamer provides functionality to stream all objects from a store.
type StoreStreamer[KEY ~[]byte | ~string, OBJECT any] interface {
	Stream(ctx context.Context, ch chan<- OBJECT) error
}

// StoreStream is deprecated: use StoreStreamer instead.
type StoreStream[KEY ~[]byte | ~string, OBJECT any] = StoreStreamer[KEY, OBJECT]

// StoreLister provides functionality to list all objects from a store.
type StoreLister[KEY ~[]byte | ~string, OBJECT any] interface {
	List(ctx context.Context) ([]OBJECT, error)
}

// StoreList is deprecated: use StoreLister instead.
type StoreList[KEY ~[]byte | ~string, OBJECT any] = StoreLister[KEY, OBJECT]

// Store provides a complete type-safe key-value store interface combining all store operations.
type Store[KEY ~[]byte | ~string, OBJECT any] interface {
	StoreAdder[KEY, OBJECT]
	StoreRemover[KEY]
	StoreGetter[KEY, OBJECT]
	StoreMapper[KEY, OBJECT]
	StoreExists[KEY, OBJECT]
	StoreStreamer[KEY, OBJECT]
	StoreLister[KEY, OBJECT]
}

// NewStore returns a Store
func NewStore[KEY ~[]byte | ~string, OBJECT any](db DB, bucketName BucketName) Store[KEY, OBJECT] {
	return NewStoreFromTx(
		db,
		NewStoreTx[KEY, OBJECT](bucketName),
	)
}

// NewStoreFromTx returns a Store from a existing StoreTx
func NewStoreFromTx[KEY ~[]byte | ~string, OBJECT any](
	db DB,
	storeTx StoreTx[KEY, OBJECT],
) Store[KEY, OBJECT] {
	return &store[KEY, OBJECT]{
		db:    db,
		store: storeTx,
	}
}

type store[KEY ~[]byte | ~string, OBJECT any] struct {
	db    DB
	store StoreTx[KEY, OBJECT]
}

func (s store[KEY, OBJECT]) Add(ctx context.Context, key KEY, object OBJECT) error {
	return s.db.Update(ctx, func(ctx context.Context, tx Tx) error {
		return s.store.Add(ctx, tx, key, object)
	})
}

func (s store[KEY, OBJECT]) Remove(ctx context.Context, key KEY) error {
	return s.db.Update(ctx, func(ctx context.Context, tx Tx) error {
		return s.store.Remove(ctx, tx, key)
	})
}

func (s store[KEY, OBJECT]) Get(ctx context.Context, key KEY) (*OBJECT, error) {
	var object *OBJECT
	err := s.db.View(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		object, err = s.store.Get(ctx, tx, key)
		return err
	})
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "view failed")
	}
	return object, nil
}

func (s store[KEY, OBJECT]) Exists(ctx context.Context, key KEY) (bool, error) {
	var object bool
	err := s.db.View(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		object, err = s.store.Exists(ctx, tx, key)
		return err
	})
	if err != nil {
		return false, errors.Wrapf(ctx, err, "view failed")
	}
	return object, nil
}

func (s store[KEY, OBJECT]) Map(
	ctx context.Context,
	fn func(ctx context.Context, key KEY, object OBJECT) error,
) error {
	return s.db.View(ctx, func(ctx context.Context, tx Tx) error {
		return s.store.Map(ctx, tx, fn)
	})
}

func (s store[KEY, OBJECT]) Stream(ctx context.Context, ch chan<- OBJECT) error {
	return s.Map(ctx, func(ctx context.Context, key KEY, object OBJECT) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- object:
			return nil
		}
	})
}

func (s store[KEY, OBJECT]) List(ctx context.Context) ([]OBJECT, error) {
	objects := make([]OBJECT, 0)
	err := s.Map(ctx, func(ctx context.Context, key KEY, object OBJECT) error {
		objects = append(objects, object)
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "map failed")
	}
	return objects, nil
}
