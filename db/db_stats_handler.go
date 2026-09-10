// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"context"
	"net/http"

	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkv "github.com/bborbe/kv"
	libparse "github.com/bborbe/parse"
)

// NewStatsHandler returns an HTTP handler that exposes libkv.DB.Stats() as JSON.
// Backend-agnostic: works against any libkv.DB implementation
// (boltkv, badgerkv, memorykv).
//
// Default: fast Stats — bucket inventory and total file size, no per-bucket key counts.
// With ?details=true: slow StatsDetailed — full per-bucket KeyCount and SizeB.
//
// Example usage:
//
//	router.Path("/dbstats").Handler(db.NewStatsHandler(libDB))
//	curl ".../dbstats"
//	curl ".../dbstats?details=true"
func NewStatsHandler(libDB libkv.DB) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.NewJSONHandler(
			libhttp.JSONHandlerFunc(
				func(ctx context.Context, req *http.Request) (interface{}, error) {
					details := libparse.ParseBoolDefault(ctx, req.FormValue("details"), false)
					var (
						stats *libkv.Stats
						err   error
					)
					if details {
						stats, err = libDB.StatsDetailed(ctx)
					} else {
						stats, err = libDB.Stats(ctx)
					}
					if err != nil {
						return nil, errors.Wrapf(ctx, err, "db stats failed")
					}
					return stats, nil
				},
			),
		),
	)
}
