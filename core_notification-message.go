// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"
	"fmt"

	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
)

func NotificationMessagef(format string, a ...any) NotificationMessage {
	return NotificationMessage(fmt.Sprintf(format, a...))
}

type NotificationMessage string

func (n NotificationMessage) String() string {
	return string(n)
}

func (n NotificationMessage) Validate(ctx context.Context) error {
	if len(n) == 0 {
		return errors.Wrapf(ctx, validation.Error, "NotificationMessage empty")
	}
	return nil
}
