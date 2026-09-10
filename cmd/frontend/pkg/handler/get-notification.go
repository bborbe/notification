// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"net/http"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkv "github.com/bborbe/kv"
	"github.com/gorilla/mux"

	"github.com/bborbe/notification"
)

func NewGetNotificationHandler(
	notificationStoreTx core.NotificationStoreTx,
) libhttp.JSONHandlerTx {
	return libhttp.JSONHandlerTxFunc(
		func(ctx context.Context, tx libkv.Tx, req *http.Request) (interface{}, error) {
			vars := mux.Vars(req)
			identifier := base.Identifier(vars["id"])
			result, err := notificationStoreTx.Get(ctx, tx, identifier)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "get notification failed")
			}
			return result, nil
		},
	)
}
