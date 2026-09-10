// Copyright (c) 2024-2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentry

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/run"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
)

// NewSkipErrorAndReport creates a run.Func that executes the given action and reports
// any errors to Sentry without propagating the error. The returned function always
// returns nil, making it suitable for fire-and-forget operations where errors should
// be logged and reported but not halt execution.
func NewSkipErrorAndReport(sentryClient Client, action run.Runnable) run.Func {
	return func(ctx context.Context) error {
		if err := action.Run(ctx); err != nil {
			data := errors.DataFromError(err)
			sentryClient.CaptureException(
				err,
				&sentry.EventHint{
					Context: ctx,
					Data:    data,
				},
				sentry.NewScope(),
			)
			glog.Warningf("run action failed: %v %+v", err, data)
		}
		return nil
	}
}
