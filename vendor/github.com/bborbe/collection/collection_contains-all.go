// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// ContainsAll returns true if all elements from slice b are present in slice a.
// It checks whether a is a superset of b (a ⊇ b).
// Empty slice b always returns true. Duplicate elements in b are treated as single elements.
//
// Performance: O(len(a) + len(b)) time, O(len(a)) space.
// For very large slices (> 10M elements), consider validating input sizes.
func ContainsAll[T comparable](a []T, b []T) bool {
	mapA := make(map[T]struct{})
	for _, aa := range a {
		mapA[aa] = struct{}{}
	}
	for _, bb := range b {
		if _, found := mapA[bb]; !found {
			return false
		}
	}
	return true
}
