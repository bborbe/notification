// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"
	stderrors "errors"

	"github.com/bborbe/collection"
	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/log"
	libtime "github.com/bborbe/time"
	"github.com/golang/glog"

	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
)

// ErrLimitReached is returned when the cleanup limit is reached during find operation
var ErrLimitReached = stderrors.New("limit reached")

//counterfeiter:generate -o ../mocks/notification-cleaner.go --fake-name NotificationCleaner . NotificationCleaner
type NotificationCleaner interface {
	Clean(ctx context.Context, tx libkv.Tx, maxAge libtime.Duration, limit int) error
}

func NewNotificationCleaner(
	currentDateTime libtime.CurrentDateTime,
	notificationIndexer index.NotificationIndexer,
	notificationStoreTx core.NotificationStoreTx,
	logSamplerFactory log.SamplerFactory,
) NotificationCleaner {
	return &notificationCleaner{
		currentDateTime:     currentDateTime,
		notificationIndexer: notificationIndexer,
		notificationStoreTx: notificationStoreTx,
		logSampler:          logSamplerFactory.Sampler(),
	}
}

type notificationCleaner struct {
	currentDateTime     libtime.CurrentDateTime
	notificationIndexer index.NotificationIndexer
	notificationStoreTx core.NotificationStoreTx
	logSampler          log.Sampler
}

func (c *notificationCleaner) Clean(
	ctx context.Context,
	tx libkv.Tx,
	maxAge libtime.Duration,
	limit int,
) error {
	identifiers, err := c.find(ctx, tx, maxAge, limit)
	if err != nil {
		return errors.Wrapf(ctx, err, "find failed")
	}
	removed := len(identifiers)
	glog.Infof("found %d notifications to remove (max age: %s)", removed, maxAge)
	for i, id := range identifiers {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx, ctx.Err(), "context cancelled during cleanup")
		default:
		}

		if err := c.notificationStoreTx.Remove(ctx, tx, id); err != nil {
			return errors.Wrapf(ctx, err, "remove from db failed")
		}
		if err := c.notificationIndexer.Remove(ctx, id); err != nil {
			return errors.Wrap(ctx, err, "remove from index failed")
		}
		// Log progress using sampler (outputs every ~10s)
		if c.logSampler.IsSample() {
			glog.V(2).Infof("removed %d/%d notifications (sample)", i+1, removed)
		}
	}
	glog.Infof("removed %d notifications completed", removed)
	return nil
}

func (c *notificationCleaner) find(
	ctx context.Context,
	tx libkv.Tx,
	maxAge libtime.Duration,
	limit int,
) (base.Identifiers, error) {
	identifiers := collection.NewSet[base.Identifier]()
	cutoffTime := c.currentDateTime.Now().Time().Add(-maxAge.Duration())
	count := 0

	err := c.notificationStoreTx.Map(
		ctx,
		tx,
		func(ctx context.Context, notification core.Notification) error {
			if notification.Created.Time().Before(cutoffTime) {
				identifiers.Add(notification.Identifier)
				count++
				// Stop early if we've reached limit
				if limit > 0 && count >= limit {
					return ErrLimitReached
				}
			}
			return nil
		},
	)
	if err != nil && !stderrors.Is(err, ErrLimitReached) {
		return nil, errors.Wrapf(ctx, err, "map failed")
	}
	return identifiers.Slice(), nil
}
