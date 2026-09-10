// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// Not inverts the result of a validator.
// It returns nil if the wrapped validator fails, and an error if it succeeds.
// Non-validation errors are passed through unchanged.
func Not(hasValidation HasValidation) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if err := hasValidation.Validate(ctx); err != nil {
			if errors.Is(err, Error) {
				return nil
			}
			return err
		}
		return errors.Wrapf(ctx, Error, "expected not valid")
	})
}
