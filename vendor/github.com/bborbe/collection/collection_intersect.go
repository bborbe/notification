// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// Intersect returns a new slice containing only the elements that appear in both input slices.
// It returns the intersection (a ∩ b) of the two slices.
// The order of elements in the result matches their order of first occurrence in slice a.
// Duplicate elements in the input slices are automatically handled.
//
// Performance: O(len(a) + len(b)) time, O(len(b)) space.
func Intersect[T comparable](a []T, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return []T{}
	}

	// Build a map of elements in b for O(1) lookup
	mapB := make(map[T]struct{}, len(b))
	for _, bb := range b {
		mapB[bb] = struct{}{}
	}

	// Track already added elements to avoid duplicates in result
	seen := make(map[T]struct{})
	result := make([]T, 0)

	// Add elements from a that exist in b
	for _, aa := range a {
		if _, inB := mapB[aa]; inB {
			if _, alreadyAdded := seen[aa]; !alreadyAdded {
				result = append(result, aa)
				seen[aa] = struct{}{}
			}
		}
	}

	return result
}
