// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// LengthEqual validates that the slice length equals the expected length.
// It returns an error if the lengths don't match, otherwise returns nil.
func LengthEqual[T any](list []T, expectedLength int) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(list) == expectedLength {
			return nil
		}
		return errors.Wrapf(ctx, Error, "length is not %d", expectedLength)
	})
}

// LengthGt validates that the slice length is greater than the expected length.
// It returns an error if the slice is not longer, otherwise returns nil.
func LengthGt[T any](list []T, expectedLength int) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(list) > expectedLength {
			return nil
		}
		return errors.Wrapf(ctx, Error, "length is not > %d", expectedLength)
	})
}

// LengthGe validates that the slice length is greater than or equal to the expected length.
// It returns an error if the slice is shorter, otherwise returns nil.
func LengthGe[T any](list []T, expectedLength int) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(list) >= expectedLength {
			return nil
		}
		return errors.Wrapf(ctx, Error, "length is not >= %d", expectedLength)
	})
}

// LengthLt validates that the slice length is less than the expected length.
// It returns an error if the slice is not shorter, otherwise returns nil.
func LengthLt[T any](list []T, expectedLength int) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(list) < expectedLength {
			return nil
		}
		return errors.Wrapf(ctx, Error, "length is not < %d", expectedLength)
	})
}

// LengthLe validates that the slice length is less than or equal to the expected length.
// It returns an error if the slice is longer, otherwise returns nil.
func LengthLe[T any](list []T, expectedLength int) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if len(list) <= expectedLength {
			return nil
		}
		return errors.Wrapf(ctx, Error, "length is not <= %d", expectedLength)
	})
}
