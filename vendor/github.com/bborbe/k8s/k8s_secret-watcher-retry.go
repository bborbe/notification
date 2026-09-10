// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
	"github.com/golang/glog"
)

// NewSecretWatcherRetry wraps a SecretWatcher with automatic retry logic.
// If the underlying watcher returns an error that is not context cancellation
// or deadline exceeded, it will wait for the specified duration and retry.
// Context cancellation or deadline exceeded errors are returned immediately.
func NewSecretWatcherRetry(
	secretWatcher SecretWatcher,
	waiter libtime.WaiterDuration,
	duration libtime.Duration,
) SecretWatcher {
	return &secretWatcherRetry{
		secretWatcher: secretWatcher,
		waiter:        waiter,
		duration:      duration,
	}
}

type secretWatcherRetry struct {
	secretWatcher SecretWatcher
	waiter        libtime.WaiterDuration
	duration      libtime.Duration
}

func (s *secretWatcherRetry) Watch(ctx context.Context) error {
	glog.V(2).Infof("secret watcher retry started")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := s.secretWatcher.Watch(ctx); err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					return errors.Wrapf(ctx, err, "secret watcher watch failed")
				}
				glog.V(2).Infof("watch failed, retrying: %v", err)
			} else {
				glog.V(2).Infof("watch completed, restarting")
			}

			// Small delay before reconnecting
			if err := s.waiter.Wait(ctx, s.duration); err != nil {
				return errors.Wrapf(ctx, err, "wait before retry failed")
			}
		}
	}
}
