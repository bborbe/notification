// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"
	"encoding/json"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	"github.com/golang/glog"
)

type NotificationMapperTx interface {
	Map(
		ctx context.Context,
		tx libkv.Tx,
		fn func(ctx context.Context, notification Notification) error,
	) error
}

type NotificationAdderTx interface {

	// Add notification to store
	Add(ctx context.Context, tx libkv.Tx, notification Notification) error
}

type NotificationRemoverTx interface {

	// Remove notification to store
	Remove(ctx context.Context, tx libkv.Tx, identifier base.Identifier) error
}

type NotificationGetterTx interface {

	// Get notification from store
	Get(ctx context.Context, tx libkv.Tx, identifier base.Identifier) (*Notification, error)
}

type NotificationStreamerTx interface {

	// Stream all notifications from store
	Stream(ctx context.Context, tx libkv.Tx, ch chan<- Notification) error
}

type NotificationStreamerTxFunc func(ctx context.Context, tx libkv.Tx, ch chan<- Notification) error

func (s NotificationStreamerTxFunc) Stream(
	ctx context.Context,
	tx libkv.Tx,
	ch chan<- Notification,
) error {
	return s(ctx, tx, ch)
}

type NotificationFinderTx interface {
	// Find matching trades
	Find(
		ctx context.Context,
		tx libkv.Tx,
		match func(notification Notification) bool,
		ch chan<- Notification,
	) error
}

//counterfeiter:generate -o ../mocks/core-notification-store-tx.go --fake-name CoreNotificationStoreTx . NotificationStoreTx
type NotificationStoreTx interface {
	NotificationAdderTx
	NotificationRemoverTx
	NotificationGetterTx
	NotificationStreamerTx
	NotificationFinderTx
	NotificationMapperTx
}

func NewNotificationStoreTx() NotificationStoreTx {
	return NewNotificationStoreWithBucketTx(
		libkv.NewBucketName("notification-store"),
	)
}

func NewNotificationStoreWithBucketTx(
	bucketName libkv.BucketName,
) NotificationStoreTx {
	return &notificationStoreTx{
		bucket: bucketName,
	}
}

type notificationStoreTx struct {
	bucket libkv.BucketName
}

func (c *notificationStoreTx) Remove(
	ctx context.Context,
	tx libkv.Tx,
	identifier base.Identifier,
) error {
	bucket, err := tx.CreateBucketIfNotExists(ctx, c.bucket)
	if err != nil {
		return errors.Wrapf(ctx, err, "get bucket failed")
	}
	if err := bucket.Delete(ctx, identifier.Bytes()); err != nil {
		return errors.Wrapf(ctx, err, "remove %s failed", identifier)
	}
	return nil
}

func (c *notificationStoreTx) Get(
	ctx context.Context,
	tx libkv.Tx,
	identifier base.Identifier,
) (*Notification, error) {
	var notification Notification
	bucket, err := tx.Bucket(ctx, c.bucket)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "get bucket failed")
	}
	item, err := bucket.Get(ctx, identifier.Bytes())
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "get %s failed", identifier)
	}
	err = item.Value(func(val []byte) error {
		if len(val) == 0 {
			return errors.Wrapf(
				ctx,
				libkv.KeyNotFoundError,
				"notification(%s) not found",
				identifier,
			)
		}
		return json.Unmarshal(val, &notification)
	})
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "handel value failed")
	}
	return &notification, nil
}

func (c *notificationStoreTx) Add(
	ctx context.Context,
	tx libkv.Tx,
	notification Notification,
) error {
	bucket, err := tx.CreateBucketIfNotExists(ctx, c.bucket)
	if err != nil {
		return errors.Wrapf(ctx, err, "get bucket failed")
	}
	value, err := json.Marshal(notification)
	if err != nil {
		return errors.Wrapf(ctx, err, "marshal json failed")
	}
	if err = bucket.Put(ctx, notification.Identifier.Bytes(), value); err != nil {
		return errors.Wrapf(ctx, err, "set failed")
	}
	return nil
}

func (c *notificationStoreTx) Map(
	ctx context.Context,
	tx libkv.Tx,
	fn func(ctx context.Context, notification Notification) error,
) error {
	var counter int
	bucket, err := tx.Bucket(ctx, c.bucket)
	if err != nil {
		if errors.Is(err, libkv.BucketNotFoundError) {
			glog.Warningf("bucket %s not found", c.bucket)
			return nil
		}
		return errors.Wrapf(ctx, err, "get bucket failed")
	}
	it := bucket.Iterator()
	defer it.Close()
	for it.Rewind(); it.Valid(); it.Next() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			item := it.Item()
			err := item.Value(func(v []byte) error {
				var notification Notification
				if err := json.Unmarshal(v, &notification); err != nil {
					return errors.Wrapf(ctx, err, "unmarshal notification failed")
				}
				if err := fn(ctx, notification); err != nil {
					return errors.Wrapf(ctx, err, "call fn failed")
				}
				counter++
				return nil
			})
			if err != nil {
				return errors.Wrapf(ctx, err, "handle value failed")
			}
		}
	}
	glog.V(4).Infof("found %d notifications", counter)
	return nil
}

func (c *notificationStoreTx) Stream(
	ctx context.Context,
	tx libkv.Tx,
	ch chan<- Notification,
) error {
	return c.Map(ctx, tx, func(ctx context.Context, notification Notification) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- notification:
			return nil
		}
	})
}

func (c *notificationStoreTx) Find(
	ctx context.Context,
	tx libkv.Tx,
	match func(notification Notification) bool,
	ch chan<- Notification,
) error {
	return c.Map(ctx, tx, func(ctx context.Context, notification Notification) error {
		if match(notification) {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- notification:
			}
		}
		return nil
	})
}
