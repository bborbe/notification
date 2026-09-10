// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"net/http"

	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/parse"
	libtime "github.com/bborbe/time"

	"github.com/bborbe/notification/cmd/frontend/pkg"
)

func NewCleanHandler(
	ctx context.Context,
	db libkv.DB,
	cleaner pkg.NotificationCleaner,
	defaultMaxAge libtime.Duration,
	defaultCleanLimit int,
) http.Handler {
	return libhttp.NewBackgroundRunRequestHandler(
		ctx,
		func(ctx context.Context, req *http.Request) error {
			maxAge := libtime.ParseDurationDefault(ctx, req.FormValue("maxAge"), defaultMaxAge)
			if maxAge <= 0 {
				return errors.Errorf(ctx, "invalid maxAge: %s", maxAge)
			}
			cleanLimit := parse.ParseIntDefault(ctx, req.FormValue("cleanLimit"), defaultCleanLimit)
			if cleanLimit <= 0 {
				return errors.Errorf(ctx, "invalid cleanLimit: %d", cleanLimit)
			}
			return db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				return cleaner.Clean(ctx, tx, maxAge, cleanLimit)
			})
		},
	)
}
