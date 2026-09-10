// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"context"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	libtime "github.com/bborbe/time"
	"github.com/golang/glog"

	"github.com/bborbe/notification"
)

type NotificationFinder interface {
	Find(
		ctx context.Context,
		tx libkv.Tx,
		from *libtime.DateTime,
		until *libtime.DateTime,
		types core.NotificationTypes,
		targets []core.NotificationTarget,
		brokerIdentifiers []string,
		accountIdentifiers []string,
		stages []string,
		query string,
		limit int,
	) ([]core.Notification, error)
}

func NewNotificationFinder(
	notificationGetterTx core.NotificationGetterTx,
	notificationIndex NotificationIndexSearcher,
) NotificationFinder {
	return &notificationFinder{
		notificationIndex:    notificationIndex,
		notificationGetterTx: notificationGetterTx,
	}
}

type notificationFinder struct {
	notificationGetterTx core.NotificationGetterTx
	notificationIndex    NotificationIndexSearcher
}

func (a *notificationFinder) Find(
	ctx context.Context,
	tx libkv.Tx,
	from *libtime.DateTime,
	until *libtime.DateTime,
	types core.NotificationTypes,
	targets []core.NotificationTarget,
	brokerIdentifiers []string,
	accountIdentifiers []string,
	stages []string,
	query string,
	limit int,
) ([]core.Notification, error) {
	notificationIdentifiers, err := a.notificationIndex.Search(
		ctx,
		from,
		until,
		types,
		targets,
		brokerIdentifiers,
		accountIdentifiers,
		stages,
		query,
		limit,
	)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "search failed")
	}
	glog.V(3).Infof("found %d notificationIdentifiers in bleve index", len(notificationIdentifiers))

	result := []core.Notification{}
	for _, notificationIdentifier := range notificationIdentifiers {
		notification, err := a.notificationGetterTx.Get(ctx, tx, notificationIdentifier)
		if err != nil {
			glog.V(2).Infof("get notification(%s) failed: %v", notificationIdentifier, err)
			continue
		}
		result = append(result, *notification)
	}
	glog.V(3).Infof("found %d notifications in db", len(result))
	return result, nil
}
