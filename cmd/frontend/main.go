// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/bborbe/cqrs/base"
	libcron "github.com/bborbe/cron"
	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkafka "github.com/bborbe/kafka"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
	"github.com/bborbe/service"
	libtime "github.com/bborbe/time"
	"github.com/blevesearch/bleve/v2"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/bborbe/notification/cmd/frontend/pkg/factory"
	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification/db"
	libfactory "github.com/bborbe/notification/factory"
	libindex "github.com/bborbe/notification/index"
	libmetrics "github.com/bborbe/notification/metrics"
)

const maxAge = libtime.Duration(720 * time.Hour) // 30 days

func main() {
	app := &application{}
	os.Exit(service.Main(context.Background(), app, &app.SentryDSN, &app.SentryProxy))
}

type application struct {
	SentryDSN           string             `required:"true"  arg:"sentry-dsn"            env:"SENTRY_DSN"            usage:"SentryDSN"                                   display:"length"`
	SentryProxy         string             `required:"false" arg:"sentry-proxy"          env:"SENTRY_PROXY"          usage:"Sentry Proxy"`
	Listen              string             `required:"true"  arg:"listen"                env:"LISTEN"                usage:"address to listen to"`
	DataDir             string             `required:"true"  arg:"datadir"               env:"DATADIR"               usage:"data directory"`
	KafkaBrokers        libkafka.Brokers   `required:"true"  arg:"kafka-brokers"         env:"KAFKA_BROKERS"         usage:"Comma separated list of Kafka brokers"`
	BatchSize           libkafka.BatchSize `required:"true"  arg:"batch-size"            env:"BATCH_SIZE"            usage:"batch consume size"                                           default:"1"`
	NoSync              bool               `required:"true"  arg:"no-sync"               env:"NO_SYNC"               usage:"no sync"                                                      default:"false"`
	Branch              base.Branch        `required:"true"  arg:"branch"                env:"BRANCH"                usage:"branch"`
	CleanLimit          int                `required:"true"  arg:"clean-limit"           env:"CLEAN_LIMIT"           usage:"max notifications to delete per cleanup run"                  default:"10000"`
	CleanCronExpression libcron.Expression `required:"true"  arg:"clean-cron-expression" env:"CLEAN_CRON_EXPRESSION" usage:"cron expression for notification cleanup"                     default:"0 3 * * * ?"`
	BuildGitCommit      string             `required:"false" arg:"build-git-commit"      env:"BUILD_GIT_COMMIT"      usage:"Build Git commit hash"                                        default:"none"`
	BuildDate           *libtime.DateTime  `required:"false" arg:"build-date"            env:"BUILD_DATE"            usage:"Build timestamp (RFC3339)"`
}

func (a *application) Run(ctx context.Context, sentryClient libsentry.Client) error {
	libmetrics.NewBuildInfoMetrics().SetBuildInfo(a.BuildDate)

	saramaClientProvider, err := libkafka.NewSaramaClientProviderByType(
		ctx,
		libkafka.SaramaClientProviderTypeReused,
		a.KafkaBrokers,
		func(config *sarama.Config) {
			config.Consumer.MaxWaitTime = 1000 * time.Millisecond
		},
	)
	if err != nil {
		return errors.Wrapf(ctx, err, "create sarama client failed")
	}
	defer saramaClientProvider.Close()

	db, err := db.OpenBoltDB(ctx, a.DataDir, a.NoSync)
	if err != nil {
		return errors.Wrapf(ctx, err, "open db failed")
	}
	defer db.Close()

	bleveIndex, err := libindex.OpenDir(ctx, index.NewIndexMapping(), a.DataDir)
	if err != nil {
		return errors.Wrapf(ctx, err, "open bleve index failed")
	}
	defer bleveIndex.Close()

	currentDateTime := libtime.NewCurrentDateTime()

	return service.Run(
		ctx,
		a.createCleanerCron(sentryClient, db, bleveIndex, currentDateTime),
		a.createNotificationConsumer(saramaClientProvider, bleveIndex, db),
		a.createHTTPServer(bleveIndex, db, currentDateTime),
	)
}

func (a *application) createNotificationConsumer(
	saramaClientProvider libkafka.SaramaClientProvider,
	bleveIndex bleve.Index,
	db libkv.DB,
) run.Func {
	return factory.CreateNotificationConsumer(
		saramaClientProvider,
		bleveIndex,
		db,
		a.Branch,
		a.BatchSize,
	)
}

func (a *application) createCleanerCron(
	sentryClient libsentry.Client,
	db libkv.DB,
	bleveIndex bleve.Index,
	currentDateTime libtime.CurrentDateTime,
) run.Func {
	return factory.CreateCleanerCron(
		sentryClient,
		db,
		bleveIndex,
		currentDateTime,
		maxAge,
		a.CleanLimit,
		a.CleanCronExpression,
	)
}

func (a *application) createHTTPServer(
	bleveIndex bleve.Index,
	db libkv.DB,
	currentDateTime libtime.CurrentDateTime,
) run.Func {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		router := mux.NewRouter()
		router.Path("/healthz").Handler(libhttp.NewPrintHandler("OK"))
		router.Path("/readiness").Handler(libhttp.NewPrintHandler("OK"))
		router.Path("/metrics").Handler(promhttp.Handler())
		router.Path("/resetdb").Handler(libkv.NewResetHandler(db, cancel))
		router.Path("/resetbucket/{BucketName}").Handler(libkv.NewResetBucketHandler(db, cancel))
		router.Path("/setloglevel/{level}").Handler(libfactory.CreateSetLoglevelHandler(ctx))
		router.Path("/offsetmanager").Handler(libfactory.CreateOffsetManagerHandler(db, cancel))
		router.Path("/clean").
			Handler(factory.CreateCleanHandler(ctx, db, bleveIndex, currentDateTime, maxAge, a.CleanLimit))
		router.Path("/status").
			Handler(factory.CreateStatusHandler(bleveIndex, db))

		router.Path("/api/1.0/notification/{id:.+}").Handler(
			factory.CreateGetNotificationHandler(db),
		)

		router.Path("/api/1.0/notification").Handler(
			factory.CreateListNotificationsHandler(bleveIndex, db),
		)

		glog.V(2).Infof("starting http server listen on %s", a.Listen)
		return libhttp.NewServer(
			a.Listen,
			router,
		).Run(ctx)
	}
}
