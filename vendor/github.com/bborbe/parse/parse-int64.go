// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bborbe/math"
)

// ParseInt64 converts an interface{} value to an int64.
// Supported types: int64, int32, int, float32, float64, string.
// Float values are rounded to the nearest integer.
// String values are parsed using strconv.ParseInt.
// Returns an error if the value cannot be converted to int64.
func ParseInt64(ctx context.Context, value interface{}) (int64, error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case int:
		return int64(v), nil
	case float32:
		return int64(math.Round(float64(v))), nil
	case float64:
		return int64(math.Round(v)), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return ParseInt64(ctx, fmt.Sprintf("%v", value))
	}
}

// ParseInt64Default converts an interface{} value to an int64, returning defaultValue on error.
// This is a convenience wrapper around ParseInt64 that never returns an error.
func ParseInt64Default(ctx context.Context, value interface{}, defaultValue int64) int64 {
	result, err := ParseInt64(ctx, value)
	if err != nil {
		return defaultValue
	}
	return result
}
