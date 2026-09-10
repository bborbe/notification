// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cron

import (
	"context"
	"time"

	libkv "github.com/bborbe/kv"
	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
)

func NewIntervalCronTx(
	sentryClient libsentry.Client,
	db libkv.DB,
	action libkv.FuncTx,
	interval time.Duration,
) run.Func {
	return NewIntervalCron(
		sentryClient,
		func(ctx context.Context) error {
			return db.Update(ctx, action)
		},
		interval,
	)
}
