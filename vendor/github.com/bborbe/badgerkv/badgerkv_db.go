// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package badgerkv provides a standardized key-value store interface built on top of BadgerDB.
// It implements the github.com/bborbe/kv interface, offering a clean and consistent API for
// database operations including transactions, buckets, and key-value operations.
//
// Basic usage:
//
//	// Open a file-based database
//	db, err := badgerkv.OpenPath(ctx, "/path/to/db")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	// Perform operations within a transaction
//	err = db.Update(ctx, func(ctx context.Context, tx badgerkv.Tx) error {
//		bucket, err := tx.CreateBucketIfNotExists([]byte("users"))
//		if err != nil {
//			return err
//		}
//		return bucket.Put([]byte("user:1"), []byte(`{"name": "John"}`))
//	})
//
// The package supports both file-based and in-memory databases, with full transaction support
// and bucket-based data organization.
package badgerkv

import (
	"context"
	"os"

	"github.com/bborbe/collection"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	"github.com/dgraph-io/badger/v4"
	"github.com/golang/glog"
)

type contextKey string

const stateCtxKey contextKey = "state"

type DB interface {
	libkv.DB
	DB() *badger.DB
}

type ChangeOptions func(opts *badger.Options)

// MinMemoryUsageOptions configures BadgerDB for minimal memory usage.
// This is useful for resource-constrained environments or when running multiple instances.
func MinMemoryUsageOptions(opts *badger.Options) {
	opts.MemTableSize = 16 << 20
	opts.NumMemtables = 3
	opts.NumLevelZeroTables = 3
	opts.NumLevelZeroTablesStall = 8
}

// OpenPath opens a file-based BadgerDB database at the specified path.
// Optional ChangeOptions functions can be provided to customize BadgerDB options.
//
// Example:
//
//	db, err := badgerkv.OpenPath(ctx, "/tmp/mydb")
//	db, err := badgerkv.OpenPath(ctx, "/tmp/mydb", badgerkv.MinMemoryUsageOptions)
func OpenPath(ctx context.Context, path string, fn ...ChangeOptions) (DB, error) {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil
	for _, f := range fn {
		f(&opts)
	}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "open badger db failed")
	}
	return NewDB(db), nil
}

// OpenMemory opens an in-memory BadgerDB database.
// This is useful for testing or temporary data storage.
// Optional ChangeOptions functions can be provided to customize BadgerDB options.
//
// Example:
//
//	db, err := badgerkv.OpenMemory(ctx)
func OpenMemory(ctx context.Context, fn ...ChangeOptions) (DB, error) {
	opts := badger.DefaultOptions("").WithInMemory(true)
	opts.Logger = nil
	for _, f := range fn {
		f(&opts)
	}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "open badger db failed")
	}
	return NewDB(db), nil
}

func NewDB(db *badger.DB) DB {
	return &badgerdb{
		db: db,
	}
}

type badgerdb struct {
	db *badger.DB
}

func (b *badgerdb) Remove() error {
	opts := b.db.Opts()
	paths := collection.Unique([]string{opts.Dir, opts.ValueDir})
	for _, path := range paths {
		if path == "" {
			continue
		}
		if _, err := os.Stat(path); err != nil {
			continue
		}
		if err := os.RemoveAll(path); err != nil {
			return err
		}
		glog.V(4).Infof("remove files from %s", path)
	}
	return nil
}

func (b *badgerdb) Sync() error {
	return b.db.Sync()
}

func (b *badgerdb) DB() *badger.DB {
	return b.db
}

func (b *badgerdb) Close() error {
	return b.db.Close()
}

func (b *badgerdb) Update(
	ctx context.Context,
	fn func(ctx context.Context, tx libkv.Tx) error,
) error {
	return b.runTx(ctx, "update", b.db.Update, fn)
}

func (b *badgerdb) View(
	ctx context.Context,
	fn func(ctx context.Context, tx libkv.Tx) error,
) error {
	return b.runTx(ctx, "view", b.db.View, fn)
}

func (b *badgerdb) runTx(
	ctx context.Context,
	op string,
	badgerFn func(func(*badger.Txn) error) error,
	fn func(ctx context.Context, tx libkv.Tx) error,
) error {
	glog.V(4).Infof("db %s started", op)
	if IsTransactionOpen(ctx) {
		return errors.Wrapf(ctx, libkv.TransactionAlreadyOpenError, "transaction already open")
	}
	err := badgerFn(func(tx *badger.Txn) error {
		glog.V(4).Infof("db %s started", op)
		ctx = SetOpenState(ctx)
		if err := fn(ctx, NewTx(tx)); err != nil {
			return errors.Wrapf(ctx, err, "db %s failed", op)
		}
		glog.V(4).Infof("db %s completed", op)
		return nil
	})
	if err != nil {
		return errors.Wrapf(ctx, err, "db %s failed", op)
	}
	glog.V(4).Infof("db %s completed", op)
	return nil
}

// IsTransactionOpen checks if a transaction is currently active in the given context.
// BadgerKV prevents nested transactions, so this function can be used to verify
// transaction state before attempting database operations.
func IsTransactionOpen(ctx context.Context) bool {
	return ctx.Value(stateCtxKey) != nil
}

// SetOpenState marks the context as having an active transaction.
// This is used internally to prevent nested transactions.
func SetOpenState(ctx context.Context) context.Context {
	return context.WithValue(ctx, stateCtxKey, "open")
}
