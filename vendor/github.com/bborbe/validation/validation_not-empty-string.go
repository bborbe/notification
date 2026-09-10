// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// NotEmptyString validates that the string value is not empty.
// It accepts any type that has string as its underlying type.
// It returns an error if the string is empty, otherwise returns nil.
func NotEmptyString[T ~string](value T) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(value) == 0 {
			return errors.Wrapf(ctx, Error, "empty string")
		}
		return nil
	})
}
