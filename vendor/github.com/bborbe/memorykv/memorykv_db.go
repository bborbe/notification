// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	"context"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	"github.com/golang/glog"
)

type contextKey string

const stateCtxKey contextKey = "state"

type DB interface {
	libkv.DB
}

func OpenMemory(ctx context.Context) (libkv.DB, error) {
	return NewDB(), nil
}

func NewDB() DB {
	return &memorydb{
		tx: NewTx(),
	}
}

type memorydb struct {
	tx libkv.Tx
}

func (b *memorydb) Remove() error {
	return nil
}

func (b *memorydb) Sync() error {
	return nil
}

func (b *memorydb) Close() error {
	b.tx = nil
	return nil
}

func (b *memorydb) Update(
	ctx context.Context,
	fn func(ctx context.Context, tx libkv.Tx) error,
) error {
	glog.V(4).Infof("db update started")
	if IsTransactionOpen(ctx) {
		return errors.Wrapf(ctx, libkv.TransactionAlreadyOpenError, "transaction already open")
	}
	glog.V(4).Infof("db update started")
	ctx = SetOpenState(ctx)
	if err := fn(ctx, b.tx); err != nil {
		return errors.Wrapf(ctx, err, "db update failed")
	}
	glog.V(4).Infof("db update completed")
	return nil
}

func (b *memorydb) View(
	ctx context.Context,
	fn func(ctx context.Context, tx libkv.Tx) error,
) error {
	glog.V(4).Infof("db view started")
	if IsTransactionOpen(ctx) {
		return errors.Wrapf(ctx, libkv.TransactionAlreadyOpenError, "transaction already open")
	}
	glog.V(4).Infof("db view started")
	ctx = SetOpenState(ctx)
	if err := fn(ctx, b.tx); err != nil {
		return errors.Wrapf(ctx, err, "db view failed")
	}
	glog.V(4).Infof("db view completed")
	return nil
}

func IsTransactionOpen(ctx context.Context) bool {
	return ctx.Value(stateCtxKey) != nil
}

func SetOpenState(ctx context.Context) context.Context {
	return context.WithValue(ctx, stateCtxKey, "open")
}
