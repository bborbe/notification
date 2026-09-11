// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"
	"strings"

	"github.com/bborbe/collection"
	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
)

const (
	AccountHitLossLimitNotificationType   NotificationType = "account-loss-limit"
	AccountHitProfitLimitNotificationType NotificationType = "account-profit-limit"
	BacktestCompletedNotificationType     NotificationType = "backtest-completed"
	BacktestFailedNotificationType        NotificationType = "backtest-failed"
	BacktestStartedNotificationType       NotificationType = "backtest-started"
	GchatRelevantNotificationType         NotificationType = "gchat-relevant"
	GoReleaseNotificationType             NotificationType = "go-release"
	MantraNotificationType                NotificationType = "mantra"
	PendingApprovalNotificationType       NotificationType = "pending-approval"
	SignalNotificationType                NotificationType = "signal"
	TestNotificationType                  NotificationType = "test"
)

var AvailableNotificationTypes = NotificationTypes{
	AccountHitLossLimitNotificationType,
	AccountHitProfitLimitNotificationType,
	BacktestCompletedNotificationType,
	BacktestFailedNotificationType,
	BacktestStartedNotificationType,
	GchatRelevantNotificationType,
	GoReleaseNotificationType,
	MantraNotificationType,
	PendingApprovalNotificationType,
	SignalNotificationType,
	TestNotificationType,
}

type NotificationTypes []NotificationType

func (n NotificationTypes) Contains(value NotificationType) bool {
	return collection.Contains(n, value)
}

func (n NotificationTypes) Strings() []string {
	result := make([]string, len(n))
	for i, bb := range n {
		result[i] = bb.String()
	}
	return result
}

func (n NotificationTypes) Interfaces() []interface{} {
	result := make([]interface{}, len(n))
	for i, ss := range n {
		result[i] = ss
	}
	return result
}

func ParseNotificationTypesFromString(value string) NotificationTypes {
	return ParseNotificationTypes(strings.FieldsFunc(value, func(r rune) bool {
		return r == ','
	}))
}

func ParseNotificationTypes(values []string) NotificationTypes {
	result := make(NotificationTypes, len(values))
	for i, value := range values {
		result[i] = NotificationType(value)
	}
	return result
}

type NotificationType string

func (n NotificationType) String() string {
	return string(n)
}

func (n NotificationType) Validate(ctx context.Context) error {
	if !AvailableNotificationTypes.Contains(n) {
		return errors.Wrapf(ctx, validation.Error, "notificationType(%s) is invalid", n)
	}
	return nil
}
