// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libkafka "github.com/bborbe/kafka"
	libkv "github.com/bborbe/kv"
)

func NewNotificationMessageHandlerTx(
	notificationHandler NotificationHandlerTx,
) libkafka.MessageHandlerTx {
	return libkafka.MessageHandlerTxFunc(
		func(ctx context.Context, tx libkv.Tx, msg *sarama.ConsumerMessage) error {
			if len(msg.Value) == 0 {
				return notificationHandler.DeleteNotification(ctx, tx, base.Identifier(msg.Key))
			}
			var notification Notification
			if err := json.Unmarshal(msg.Value, &notification); err != nil {
				return errors.Wrapf(ctx, err, "unmarshal notification failed")
			}
			return notificationHandler.UpdateNotification(ctx, tx, notification)
		},
	)
}
