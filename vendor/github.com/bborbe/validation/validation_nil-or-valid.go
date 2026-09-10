// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

// NilOrValid validates that the value is either nil or passes the provided validation.
// It succeeds if the validator itself is nil OR if the validator passes its validation.
// This is useful for optional fields that should be validated only when present.
func NilOrValid(validation HasValidation) HasValidation {
	return Any{
		Nil(validation),
		validation,
	}
}
