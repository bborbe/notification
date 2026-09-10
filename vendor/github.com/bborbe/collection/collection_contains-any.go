// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// ContainsAny returns true if at least one element from slice b is present in slice a.
// It checks whether a and b have any intersection (a ∩ b ≠ ∅).
// Empty slice b always returns false (no elements to check).
//
// Performance: O(len(a) + len(b)) time, O(len(a)) space.
// For very large slices (> 10M elements), consider validating input sizes.
func ContainsAny[T comparable](a []T, b []T) bool {
	if len(b) == 0 {
		return false
	}
	mapA := make(map[T]struct{})
	for _, aa := range a {
		mapA[aa] = struct{}{}
	}
	for _, bb := range b {
		if _, found := mapA[bb]; found {
			return true
		}
	}
	return false
}
