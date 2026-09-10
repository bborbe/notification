// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import "strings"

// ParseStatefulSetsFromString parses a comma-separated string into StatefulSets.
func ParseStatefulSetsFromString(value string) StatefulSets {
	return ParseStatefulSets(strings.FieldsFunc(value, func(r rune) bool {
		return r == ','
	}))
}

// ParseStatefulSets converts a slice of strings into StatefulSets.
func ParseStatefulSets(values []string) StatefulSets {
	result := make(StatefulSets, len(values))
	for i, value := range values {
		result[i] = StatefulSet(value)
	}
	return result
}

// StatefulSets is a collection of StatefulSet names.
type StatefulSets []StatefulSet

// Contains returns true if the collection contains the specified statefulSet.
func (c StatefulSets) Contains(statefulSet StatefulSet) bool {
	for _, o := range c {
		if o == statefulSet {
			return true
		}
	}
	return false
}

// StatefulSet represents a Kubernetes StatefulSet name.
type StatefulSet string

// String returns the string representation of the StatefulSet name.
func (s StatefulSet) String() string {
	return string(s)
}
