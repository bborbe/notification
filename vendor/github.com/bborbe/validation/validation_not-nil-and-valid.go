// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// NotNilAndValid validates that the value is not nil AND passes the provided validation.
// It ensures the validator is not nil before executing the validation logic.
// This is useful for required fields that must be present and valid.
func NotNilAndValid(value HasValidation) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if Nil(value).Validate(ctx) == nil {
			return errors.Wrapf(ctx, Error, "should be not nil")
		}
		return value.Validate(ctx)
	})
}
