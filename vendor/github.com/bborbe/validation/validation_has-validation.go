// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import "context"

//counterfeiter:generate -o mocks/validation-has-validation.go --fake-name ValidationHasValidation . HasValidation

// HasValidation defines the interface that all validators must implement.
// It provides a single method for validating data with context support.
type HasValidation interface {
	// Validate performs validation logic and returns an error if validation fails.
	// The context can be used for cancellation and timeout handling.
	Validate(ctx context.Context) error
}

// HasValidationFunc is a function type that implements the HasValidation interface.
// It allows functions to be used as validators without creating custom types.
type HasValidationFunc func(ctx context.Context) error

// Validate implements the HasValidation interface by calling the function itself.
func (v HasValidationFunc) Validate(ctx context.Context) error {
	return v(ctx)
}
