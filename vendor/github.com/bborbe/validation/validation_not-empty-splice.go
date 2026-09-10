// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// NotEmptySlice validates that the slice is not empty.
// It accepts slices of any type and checks that the length is greater than zero.
// It returns an error if the slice is empty, otherwise returns nil.
func NotEmptySlice[T any](values []T) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(values) == 0 {
			return errors.Wrapf(ctx, Error, "empty slice")
		}
		return nil
	})
}
