// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"context"
	"net/http"

	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkafka "github.com/bborbe/kafka"
	libkv "github.com/bborbe/kv"
)

// CreateOffsetManagerHandler creates a standardized HTTP handler for DB-based offset management.
// This handler provides a RESTful interface to manage Kafka consumer offsets stored in a database.
//
// Parameters:
//   - db: Database connection for storing offset information
//   - cancel: Context cancellation function for graceful shutdown
//
// Returns an HTTP handler that can be mounted at /offsetmanager endpoint.
//
// Example usage:
//
//	router.Path("/offsetmanager").Handler(factory.CreateOffsetManagerHandler(db, cancel))
func CreateOffsetManagerHandler(db libkv.DB, cancel context.CancelFunc) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.WithErrorFunc(
			func(ctx context.Context, resp http.ResponseWriter, req *http.Request) error {
				req.Body = http.MaxBytesReader(resp, req.Body, 1<<20)
				offsetStore := libkafka.NewOffsetStore(db)
				if group := libkafka.Group(req.FormValue("group")); group != "" {
					offsetStore = libkafka.NewOffsetStoreGroup(db, group)
				}

				// Create offset manager for this group
				offsetManager := libkafka.NewStoreOffsetManager(
					offsetStore,
					libkafka.OffsetOldest,
					libkafka.OffsetNewest,
				)

				// Delegate to the existing handler logic
				return libkafka.NewOffsetManagerHandler(offsetManager, cancel).
					ServeHTTP(ctx, resp, req)
			},
		),
	)
}

// CreateSaramaOffsetManagerHandler creates a standardized HTTP handler for Sarama-based offset management.
// This handler provides a RESTful interface to manage Kafka consumer offsets using Sarama client.
//
// Parameters:
//   - saramaClient: Sarama Kafka client for offset operations
//   - group: Kafka consumer group for offset management
//   - cancel: Context cancellation function for graceful shutdown
//
// Returns an HTTP handler that can be mounted at /offsetmanager endpoint.
//
// Example usage:
//
//	router.Path("/offsetmanager").Handler(factory.CreateSaramaOffsetManagerHandler(saramaClient, group, cancel))
func CreateSaramaOffsetManagerHandler(
	saramaClientProvider libkafka.SaramaClientProvider,
	group libkafka.Group,
	cancel context.CancelFunc,
) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.WithErrorFunc(
			func(ctx context.Context, resp http.ResponseWriter, req *http.Request) error {
				saramaClient, err := saramaClientProvider.Client(ctx)
				if err != nil {
					return errors.Wrapf(ctx, err, "get client failed")
				}
				return libkafka.NewOffsetManagerHandler(
					libkafka.NewSaramaOffsetManager(
						saramaClient,
						group,
						libkafka.OffsetOldest,
						libkafka.OffsetNewest,
					),
					cancel,
				).ServeHTTP(ctx, resp, req)
			},
		),
	)
}
