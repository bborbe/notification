// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cron

import (
	"context"

	"github.com/bborbe/cron"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
)

func NewExpressionCronTx(
	sentryClient libsentry.Client,
	db libkv.DB,
	action libkv.FuncTx,
	expression cron.Expression,
) run.Func {
	return NewExpressionCron(
		sentryClient,
		func(ctx context.Context) error {
			return db.Update(ctx, action)
		},
		expression,
	)
}
