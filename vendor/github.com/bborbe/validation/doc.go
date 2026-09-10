// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package validation provides functional composition for validation rules using the HasValidation interface pattern.
//
// The core principle is that all validators implement the HasValidation interface with a single method:
//
//	Validate(ctx context.Context) error
//
// This allows for consistent composition of validation rules using logical operators like All and Any.
//
// Basic usage:
//
//	validator := validation.All{
//		validation.NotNil(user),
//		validation.NotEmptyString(user.Name),
//		validation.LengthGe(user.Email, 1),
//	}
//
//	if err := validator.Validate(ctx); err != nil {
//		// Handle validation error
//	}
//
// The package provides validators for common scenarios:
//   - Basic validators: NotNil, Nil, True, False, Equal
//   - String/slice validators: NotEmptyString, NotEmptySlice
//   - Length validators: LengthEqual, LengthGt, LengthGe, LengthLt, LengthLe
//   - Logical operators: All, Any, Either, Not
//   - Conditional validators: NilOrValid, NotNilAndValid
//
// All validation errors are wrapped using github.com/bborbe/errors for consistent error handling.
package validation
