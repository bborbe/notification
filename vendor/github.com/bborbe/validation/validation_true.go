// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

// True validates that the boolean value is true.
// It returns an error if the value is false, otherwise returns nil.
func True(value bool) HasValidation {
	return Equal(value, true)
}
