// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/bborbe/cqrs/base"
	libcron "github.com/bborbe/cron"
	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkafka "github.com/bborbe/kafka"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/log"
	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
	libtime "github.com/bborbe/time"
	"github.com/blevesearch/bleve/v2"
	"github.com/golang/glog"

	"github.com/bborbe/notification/cmd/frontend/pkg"
	"github.com/bborbe/notification/cmd/frontend/pkg/handler"
	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
	"github.com/bborbe/notification/cron"
	libindex "github.com/bborbe/notification/index"
)

func CreateNotificationConsumer(
	saramaClientProvider libkafka.SaramaClientProvider,
	bleveIndex bleve.Index,
	db libkv.DB,
	branch base.Branch,
	batchSize libkafka.BatchSize,
) run.Func {
	return func(ctx context.Context) error {
		notificationIndexer := index.NewNotificationIndexer(libindex.NewIndex(bleveIndex))
		notificationStoreTx := core.NewNotificationStoreTx()
		return libkafka.NewOffsetConsumerHighwaterMarksBatchWithProvider(
			saramaClientProvider,
			core.NotificationV1SchemaID.EventTopic(
				base.TopicPrefixFromBranch(branch),
			),
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
							libkafka.MessageHandlerTxFunc(
								func(ctx context.Context, tx libkv.Tx, msg *sarama.ConsumerMessage) error {
									var notification core.Notification
									if err := json.Unmarshal(msg.Value, &notification); err != nil {
										return errors.Wrapf(ctx, err, "unmarshal failed")
									}
									if err := notificationStoreTx.Add(ctx, tx, notification); err != nil {
										return errors.Wrapf(ctx, err, "add to store failed")
									}
									if err := notificationIndexer.Add(ctx, notification); err != nil {
										return errors.Wrapf(ctx, err, "add to index failed")
									}
									return nil
								},
							),
							libkafka.NewMetrics(),
						),
						log.DefaultSamplerFactory,
					),
				),
			),
			batchSize,
			run.NewTrigger(),
			log.DefaultSamplerFactory,
		).Consume(ctx)
	}
}

func CreateGetNotificationHandler(db libkv.DB) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.NewJsonHandlerViewTx(
			db,
			handler.NewGetNotificationHandler(
				core.NewNotificationStoreTx(),
			),
		),
	)
}

func CreateListNotificationsHandler(bleveIndex bleve.Index, db libkv.DB) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.NewJsonHandlerViewTx(
			db,
			handler.NewListNotificationsHandler(
				index.NewNotificationIndexSearcher(bleveIndex),
				pkg.NewNotificationLoader(
					core.NewNotificationStoreTx(),
					log.DefaultSamplerFactory,
				),
			),
		),
	)
}

func CreateCleanHandler(
	ctx context.Context,
	db libkv.DB,
	bleveIndex bleve.Index,
	currentDateTime libtime.CurrentDateTime,
	defaultMaxAge libtime.Duration,
	defaultCleanLimit int,
) http.Handler {
	return handler.NewCleanHandler(
		ctx,
		db,
		NewNotificationCleaner(bleveIndex, currentDateTime),
		defaultMaxAge,
		defaultCleanLimit,
	)
}

func CreateCleanerCron(
	sentryClient libsentry.Client,
	db libkv.DB,
	bleveIndex bleve.Index,
	currentDateTime libtime.CurrentDateTime,
	maxAge libtime.Duration,
	cleanLimit int,
	cronExpression libcron.Expression,
) run.Func {
	notificationCleaner := NewNotificationCleaner(bleveIndex, currentDateTime)
	return cron.NewExpressionCronTx(
		sentryClient,
		db,
		func(ctx context.Context, tx libkv.Tx) error {
			glog.V(2).Infof("cleanup cron started")
			if err := notificationCleaner.Clean(ctx, tx, maxAge, cleanLimit); err != nil {
				return errors.Wrapf(ctx, err, "clean failed")
			}
			glog.V(2).Infof("cleanup cron completed")
			return nil
		},
		cronExpression,
	)
}

func CreateStatusHandler(bleveIndex bleve.Index, db libkv.DB) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.NewJsonHandlerViewTx(
			db,
			handler.NewStatusHandler(
				bleveIndex,
				libkv.NewBucketName("notification-store"),
			),
		),
	)
}

func NewNotificationCleaner(
	bleveIndex bleve.Index,
	currentDateTime libtime.CurrentDateTime,
) pkg.NotificationCleaner {
	return pkg.NewNotificationCleaner(
		currentDateTime,
		index.NewNotificationIndexer(libindex.NewIndex(bleveIndex)),
		core.NewNotificationStoreTx(),
		log.DefaultSamplerFactory,
	)
}
