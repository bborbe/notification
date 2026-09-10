// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import "strings"

// Compare compares two string-like values and returns an integer comparing a and b.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func Compare[T ~string](a, b T) int {
	return strings.Compare(string(a), string(b))
}
