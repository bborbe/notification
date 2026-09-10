// Copyright (c) 2025 Benjamin Borbe All rights reserved.
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
	"github.com/bborbe/notification/cmd/frontend/pkg/index"
	"github.com/bborbe/notification"
)

func NewListNotificationsHandler(
	searcher index.NotificationIndexSearcher,
	notificationLoader pkg.NotificationLoader,
) libhttp.JSONHandlerTx {
	return libhttp.JSONHandlerTxFunc(
		func(ctx context.Context, tx libkv.Tx, req *http.Request) (interface{}, error) {
			if err := req.ParseForm(); err != nil {
				return nil, errors.Wrapf(ctx, err, "parse form failed")
			}
			from, _ := libtime.ParseDateTime(ctx, req.FormValue("from"))
			until, _ := libtime.ParseDateTime(ctx, req.FormValue("until"))
			types := core.ParseNotificationTypes(req.Form["type"])
			stages := req.Form["stage"]
			brokers := req.Form["broker"]
			accountIdentifiers := req.Form["accountIdentifier"]
			targets, _ := core.ParseNotificationTargets(ctx, req.Form["target"])
			query := req.FormValue("query")
			limit := parse.ParseIntDefault(ctx, req.FormValue("limit"), 10000)

			notificationIdentifiers, err := searcher.Search(
				ctx,
				from,
				until,
				types,
				targets,
				brokers,
				accountIdentifiers,
				stages,
				query,
				limit,
			)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "search failed")
			}
			notifications, err := notificationLoader.GetAll(ctx, tx, notificationIdentifiers)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "get all notifications failed")
			}
			return notifications, nil
		},
	)
}
