// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// All represents a logical AND operation for multiple validators.
// All validators must pass for the validation to succeed.
type All []HasValidation

// Validate executes all validators and returns an error if any validator fails.
// It stops at the first validation failure and returns that error wrapped with context.
func (l All) Validate(ctx context.Context) error {
	for _, ll := range l {
		if err := ll.Validate(ctx); err != nil {
			return errors.Wrapf(ctx, err, "validate failed")
		}
	}
	return nil
}
