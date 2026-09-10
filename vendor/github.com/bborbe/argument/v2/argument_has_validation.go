// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import "context"

//counterfeiter:generate -o mocks/has_validation.go --fake-name HasValidation . HasValidation

// HasValidation defines the interface that validators must implement.
// It provides a single method for validating data with context support.
// Types implementing this interface can define custom validation logic
// beyond the built-in required field validation.
//
// Example:
//
//	type Port int
//
//	func (p Port) Validate(ctx context.Context) error {
//		if p < 1024 {
//			return errors.New(ctx, "port must be >= 1024")
//		}
//		return nil
//	}
//
// For more complex validation scenarios (e.g., field comparisons, conditional validation,
// validation helpers), see github.com/bborbe/validation which provides additional
// validation utilities and patterns that work well with this interface.
type HasValidation interface {
	// Validate performs validation logic and returns an error if validation fails.
	// The context can be used for cancellation and timeout handling.
	Validate(ctx context.Context) error
}
