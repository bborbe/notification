// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"github.com/bborbe/errors"
	libparse "github.com/bborbe/parse"
	"github.com/bborbe/validation"
)

func ParseNotificationTargets(ctx context.Context, values any) (NotificationTargets, error) {
	strings, err := libparse.ParseStrings(ctx, values)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "parse strings failed")
	}
	var targets NotificationTargets
	for _, value := range strings {
		if value != "" {
			targets = append(targets, NotificationTarget(value))
		}
	}
	return targets, nil
}

type NotificationTargets []NotificationTarget

type NotificationTarget string

func (n NotificationTarget) String() string {
	return string(n)
}

func (n NotificationTarget) Validate(ctx context.Context) error {
	if len(n) == 0 {
		return errors.Wrapf(ctx, validation.Error, "NotificationTarget empty")
	}
	return nil
}

func (n NotificationTarget) Ptr() *NotificationTarget {
	return &n
}
