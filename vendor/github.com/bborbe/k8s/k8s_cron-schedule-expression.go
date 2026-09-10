// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	"github.com/robfig/cron/v3"
)

type CronScheduleExpression string

func (c CronScheduleExpression) String() string {
	return string(c)
}

func (c CronScheduleExpression) Validate(ctx context.Context) error {
	if len(c) == 0 {
		return errors.Wrapf(ctx, validation.Error, "CronScheduleExpression empty")
	}
	if _, err := cron.ParseStandard(c.String()); err != nil {
		return errors.Wrapf(ctx, validation.Error, "parse CronScheduleExpression failed")
	}
	return nil
}
