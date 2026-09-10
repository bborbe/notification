// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

// False validates that the boolean value is false.
// It returns an error if the value is true, otherwise returns nil.
func False(value bool) HasValidation {
	return Equal(value, false)
}
