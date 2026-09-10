// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// Copy returns a new slice that is a shallow copy of the input slice.
func Copy[T any](values []T) []T {
	result := make([]T, 0, len(values))
	result = append(result, values...)
	return result
}
