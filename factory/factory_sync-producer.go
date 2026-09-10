// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"context"

	"github.com/bborbe/errors"
	libkafka "github.com/bborbe/kafka"
)

// NewSyncProducerWithName creates a sync producer that adds a 'name' header to all messages.
func NewSyncProducerWithName(
	ctx context.Context,
	brokers libkafka.Brokers,
	name string,
	opts ...libkafka.SaramaConfigOptions,
) (libkafka.SyncProducer, error) {
	syncProducer, err := libkafka.NewSyncProducerWithName(
		ctx,
		brokers,
		name,
		opts...,
	)
	if err != nil {
		return nil, errors.Wrap(ctx, err, "create syncProducer failed")
	}
	return libkafka.NewSyncProducerMetrics(syncProducer), nil
}
