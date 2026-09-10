// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"github.com/bborbe/cqrs/base"
	libkafka "github.com/bborbe/kafka"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/log"
	"github.com/bborbe/run"
)

func RunNotificationStoreConsumer(
	saramaClient libkafka.SaramaClient,
	db libkv.DB,
	branch base.Branch,
	batchSize libkafka.BatchSize,
	trigger run.Fire,
) run.Func {
	return func(ctx context.Context) error {
		logSamplerFactory := log.DefaultSamplerFactory
		notificationStoreTx := NewNotificationStoreTx()
		return libkafka.NewOffsetConsumerHighwaterMarksBatch(
			saramaClient,
			NotificationV1SchemaID.EventTopic(base.TopicPrefixFromBranch(branch)),
			libkafka.NewStoreOffsetManager(
				libkafka.NewOffsetStore(db),
				libkafka.OffsetOldest,
				libkafka.OffsetNewest,
			),
			libkafka.NewMessageHandlerBatchTxUpdate(
				db,
				libkafka.NewMessageHandlerBatchTx(
					libkafka.NewMessageHandlerTxSkipErrors(
						libkafka.NewMessageHandlerTxMetrics(
							NewNotificationMessageHandlerTx(
								NotificationHandlerTxFunc(
									func(ctx context.Context, tx libkv.Tx, notification Notification) error {
										return notificationStoreTx.Add(ctx, tx, notification)
									},
									func(ctx context.Context, tx libkv.Tx, identifier base.Identifier) error {
										return notificationStoreTx.Remove(ctx, tx, identifier)
									},
								),
							),
							libkafka.NewMetrics(),
						),
						logSamplerFactory,
					),
				),
			),
			batchSize,
			trigger,
			logSamplerFactory,
		).Consume(ctx)
	}
}
