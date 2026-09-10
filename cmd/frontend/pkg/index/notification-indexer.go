// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

//counterfeiter:generate -o ../../mocks/notification-indexer.go --fake-name NotificationIndexer . NotificationIndexer

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	"github.com/golang/glog"

	"github.com/bborbe/notification"
	"github.com/bborbe/notification/index"
)

type NotificationIndexer interface {
	Add(ctx context.Context, notification core.Notification) error
	Remove(ctx context.Context, notificationIdentifier base.Identifier) error
}

func NewNotificationIndexer(index index.Index) NotificationIndexer {
	return &notificationIndexer{
		index: index,
	}
}

type notificationIndexer struct {
	index index.Index
}

func (a *notificationIndexer) Add(ctx context.Context, notification core.Notification) error {
	if err := notification.Identifier.Validate(ctx); err != nil {
		glog.V(3).Infof("notificationIdentifier is not valid => skip: %v", err)
		return nil
	}

	indexEntry := NotificationEntry{
		Created:           notification.Created.Time(),
		Modified:          notification.Modified.Time(),
		NotificationType:  notification.Type,
		Message:           notification.Message,
		Target:            notification.Target,
		BrokerIdentifier:  notification.Metadata["brokerIdentifier"],
		AccountIdentifier: notification.Metadata["accountIdentifier"],
		Stage:             notification.Metadata["stage"],
	}

	if err := a.index.Index(notification.Identifier.String(), indexEntry); err != nil {
		return errors.Wrapf(ctx, err, "add to index failed")
	}
	return nil
}

func (a *notificationIndexer) Remove(
	ctx context.Context,
	notificationIdentifier base.Identifier,
) error {
	if err := notificationIdentifier.Validate(ctx); err != nil {
		glog.V(3).Infof("notificationIdentifier is not valid => skip: %v", err)
		return nil
	}
	if err := a.index.Delete(notificationIdentifier.String()); err != nil {
		return errors.Wrapf(ctx, err, "remove from index failed")
	}
	return nil
}
