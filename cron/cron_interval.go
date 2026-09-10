// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cron

import (
	"context"
	"time"

	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
)

func NewIntervalCron(
	sentryClient libsentry.Client,
	action run.Func,
	interval time.Duration,
) run.Func {
	return func(ctx context.Context) error {
		for {
			glog.V(3).Infof("cron started")
			if err := action(ctx); err != nil {
				sentryClient.CaptureException(
					err,
					&sentry.EventHint{
						Context:           ctx,
						OriginalException: err,
					},
					sentry.NewScope(),
				)
			}
			glog.V(3).Infof("cron completed")

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.NewTimer(interval).C:
				glog.V(3).Infof("wait for %v completed", interval)
			}
		}
	}
}
