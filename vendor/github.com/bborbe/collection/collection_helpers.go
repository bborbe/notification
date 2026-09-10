// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"fmt"
	"strings"
)

// elementToString converts an element to its string representation.
// It checks for fmt.Stringer interface first, then handles string type directly,
// and falls back to fmt.Sprintf for all other types.
func elementToString[T any](element T) string {
	switch v := any(element).(type) {
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatSetString creates a formatted string representation of a set.
// It takes a prefix (e.g., "Set[", "SetEqual[") and a slice of string elements.
// Returns "prefix]" for empty slices, or "prefix + comma-separated elements]" otherwise.
func formatSetString(prefix string, elements []string) string {
	if len(elements) == 0 {
		return prefix + "]"
	}

	var b strings.Builder
	b.WriteString(prefix)
	for i, str := range elements {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(str)
	}
	b.WriteString("]")
	return b.String()
}
