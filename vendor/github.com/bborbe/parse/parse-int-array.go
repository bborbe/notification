// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"context"

	"github.com/bborbe/errors"
)

// ParseIntArray converts an interface{} value to an int slice.
// Supported types: []int, []interface{}, []int32, []int64, []float32, []float64, []string.
// Each element is converted using ParseInt.
// Returns an error if the value cannot be converted to []int.
func ParseIntArray(ctx context.Context, value interface{}) ([]int, error) {
	switch v := value.(type) {
	case []int:
		return v, nil
	case []interface{}:
		return ParseIntArrayFromInterfaces(ctx, v)
	case []int32:
		return ParseIntArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []int64:
		return ParseIntArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []float32:
		return ParseIntArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []float64:
		return ParseIntArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []string:
		return ParseIntArrayFromInterfaces(ctx, ToInterfaceList(v))
	default:
		return nil, errors.Errorf(ctx, "invalid type %T", v)
	}
}

// ParseIntArrayDefault converts an interface{} value to an int slice, returning defaultValue on error.
// This is a convenience wrapper around ParseIntArray that never returns an error.
func ParseIntArrayDefault(ctx context.Context, value interface{}, defaultValue []int) []int {
	result, err := ParseIntArray(ctx, value)
	if err != nil {
		return defaultValue
	}
	return result
}

// ParseIntArrayFromInterfaces converts a slice of interface{} values to an int slice.
// Each element is converted using ParseInt.
// Returns an error if any element cannot be converted to int.
func ParseIntArrayFromInterfaces(ctx context.Context, values []interface{}) ([]int, error) {
	result := make([]int, len(values))
	for i, vv := range values {
		pi, err := ParseInt(ctx, vv)
		if err != nil {
			return nil, errors.Wrapf(ctx, err, "parse int failed")
		}
		result[i] = pi
	}
	return result, nil
}

// ToInterfaceList converts a typed slice to a slice of interface{}.
// This is a generic helper function used internally for type conversion.
func ToInterfaceList[T any](values []T) []interface{} {
	result := make([]interface{}, len(values))
	for i, value := range values {
		result[i] = value
	}
	return result
}
