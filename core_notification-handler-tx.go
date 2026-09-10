// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

type NotificationHandlerTx interface {
	UpdateNotification(ctx context.Context, tx libkv.Tx, notification Notification) error
	DeleteNotification(ctx context.Context, tx libkv.Tx, identifier base.Identifier) error
}

func NotificationHandlerTxFunc(
	update func(ctx context.Context, tx libkv.Tx, notification Notification) error,
	delete func(ctx context.Context, tx libkv.Tx, identifier base.Identifier) error,
) NotificationHandlerTx {
	return &notificationHandlerTxFunc{
		update: update,
		delete: delete,
	}
}

type notificationHandlerTxFunc struct {
	update func(ctx context.Context, tx libkv.Tx, notification Notification) error
	delete func(ctx context.Context, tx libkv.Tx, identifier base.Identifier) error
}

func (e *notificationHandlerTxFunc) UpdateNotification(
	ctx context.Context,
	tx libkv.Tx,
	notification Notification,
) error {
	if e.update == nil {
		return nil
	}
	if err := e.update(ctx, tx, notification); err != nil {
		return errors.Wrapf(ctx, err, "update failed")
	}
	return nil
}

func (e *notificationHandlerTxFunc) DeleteNotification(
	ctx context.Context,
	tx libkv.Tx,
	identifier base.Identifier,
) error {
	if e.delete == nil {
		return nil
	}
	if err := e.delete(ctx, tx, identifier); err != nil {
		return errors.Wrapf(ctx, err, "delete failed")
	}
	return nil
}

type NotificationHandlerTxList []NotificationHandlerTx

func (c NotificationHandlerTxList) UpdateNotification(
	ctx context.Context,
	tx libkv.Tx,
	notification Notification,
) error {
	for _, mm := range c {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := mm.UpdateNotification(ctx, tx, notification); err != nil {
				return errors.Wrapf(ctx, err, "consume message failed")
			}
		}
	}
	return nil
}

func (c NotificationHandlerTxList) DeleteNotification(
	ctx context.Context,
	tx libkv.Tx,
	identifier base.Identifier,
) error {
	for _, mm := range c {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := mm.DeleteNotification(ctx, tx, identifier); err != nil {
				return errors.Wrapf(ctx, err, "consume message failed")
			}
		}
	}
	return nil
}
