// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"context"

	"github.com/bborbe/errors"
)

// ParseInt64Array converts an interface{} value to an int64 slice.
// Supported types: []int64, []interface{}, []int, []int32, []float32, []float64, []string.
// Each element is converted using ParseInt64.
// Returns an error if the value cannot be converted to []int64.
func ParseInt64Array(ctx context.Context, value interface{}) ([]int64, error) {
	switch v := value.(type) {
	case []int64:
		return v, nil
	case []interface{}:
		return ParseInt64ArrayFromInterfaces(ctx, v)
	case []int:
		return ParseInt64ArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []int32:
		return ParseInt64ArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []float32:
		return ParseInt64ArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []float64:
		return ParseInt64ArrayFromInterfaces(ctx, ToInterfaceList(v))
	case []string:
		return ParseInt64ArrayFromInterfaces(ctx, ToInterfaceList(v))
	default:
		return nil, errors.Errorf(ctx, "invalid type %T", v)
	}
}

// ParseInt64ArrayDefault converts an interface{} value to an int64 slice, returning defaultValue on error.
// This is a convenience wrapper around ParseInt64Array that never returns an error.
func ParseInt64ArrayDefault(ctx context.Context, value interface{}, defaultValue []int64) []int64 {
	result, err := ParseInt64Array(ctx, value)
	if err != nil {
		return defaultValue
	}
	return result
}

// ParseInt64ArrayFromInterfaces converts a slice of interface{} values to an int64 slice.
// Each element is converted using ParseInt64.
// Returns an error if any element cannot be converted to int64.
func ParseInt64ArrayFromInterfaces(ctx context.Context, values []interface{}) ([]int64, error) {
	result := make([]int64, len(values))
	for i, vv := range values {
		pi, err := ParseInt64(ctx, vv)
		if err != nil {
			return nil, errors.Wrapf(ctx, err, "parse int64 failed")
		}
		result[i] = pi
	}
	return result, nil
}
