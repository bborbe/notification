// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"github.com/bborbe/cqrs/base"
	libkafka "github.com/bborbe/kafka"
	libkv "github.com/bborbe/kv"
)

func NewNotificationStoreMessageHandlerTx(
	notificationStoreTx NotificationStoreTx,
) libkafka.MessageHandlerTx {
	return NewNotificationMessageHandlerTx(
		NotificationHandlerTxFunc(
			func(ctx context.Context, tx libkv.Tx, notification Notification) error {
				return notificationStoreTx.Add(ctx, tx, notification)
			},
			func(ctx context.Context, tx libkv.Tx, identifier base.Identifier) error {
				return notificationStoreTx.Remove(ctx, tx, identifier)
			},
		),
	)
}
