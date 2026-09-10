// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"

	"github.com/bborbe/cqrs/base"
	libkv "github.com/bborbe/kv"
	"github.com/bborbe/log"
	"github.com/golang/glog"

	"github.com/bborbe/notification"
)

type NotificationLoader interface {
	GetAll(
		ctx context.Context,
		tx libkv.Tx,
		baseIdentifiers base.Identifiers,
	) ([]core.Notification, error)
}

func NewNotificationLoader(
	notificationGetterTx core.NotificationGetterTx,
	logSamplerFactory log.SamplerFactory,
) NotificationLoader {
	return &notificationLoader{
		notificationGetterTx: notificationGetterTx,
		logSampler:           logSamplerFactory.Sampler(),
	}
}

type notificationLoader struct {
	notificationGetterTx core.NotificationGetterTx
	logSampler           log.Sampler
}

func (a *notificationLoader) GetAll(
	ctx context.Context,
	tx libkv.Tx,
	baseIdentifiers base.Identifiers,
) ([]core.Notification, error) {
	result := []core.Notification{}
	for _, baseIdentifier := range baseIdentifiers {
		notification, err := a.notificationGetterTx.Get(ctx, tx, baseIdentifier)
		if err != nil {
			if a.logSampler.IsSample() {
				glog.V(2).Infof("get notification(%s) failed: %v (sample)", baseIdentifier, err)
			}
			continue
		}
		result = append(result, *notification)
	}
	glog.V(3).Infof("found %d notifications in db", len(result))
	return result, nil
}
