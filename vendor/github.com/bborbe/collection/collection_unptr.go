// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// UnPtr dereferences a pointer and returns the value.
// If the pointer is nil, it returns the zero value of type T.
func UnPtr[T any](ptr *T) T {
	var t T
	if ptr != nil {
		t = *ptr
	}
	return t
}
