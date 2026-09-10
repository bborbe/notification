// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// Filter returns a new slice containing only the elements from the input slice
// that satisfy the given predicate function.
func Filter[T any](list []T, match func(value T) bool) []T {
	result := make([]T, 0)
	for _, e := range list {
		if match(e) {
			result = append(result, e)
		}
	}
	return result
}
