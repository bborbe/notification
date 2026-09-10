// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification

import (
	"context"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
	"github.com/bborbe/validation"
)

type NotificationCleanCommand struct {
	// MaxAge is the maximum age of notifications to keep (e.g., "720h" for 30 days)
	MaxAge libtime.Duration `json:"maxAge"`
	// Limit is the maximum number of notifications to delete in one run (default: 10000)
	Limit int `json:"limit"`
}

func (n NotificationCleanCommand) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("MaxAge", validation.HasValidationFunc(func(ctx context.Context) error {
			if n.MaxAge.Duration() <= 0 {
				return errors.Errorf(ctx, "must be positive")
			}
			return nil
		})),
		validation.Name("Limit", validation.HasValidationFunc(func(ctx context.Context) error {
			if n.Limit <= 0 {
				return errors.Errorf(ctx, "must be positive")
			}
			return nil
		})),
	}.Validate(ctx)
}

func (n NotificationCleanCommand) Ptr() *NotificationCleanCommand {
	return &n
}
