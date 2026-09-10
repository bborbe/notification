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

// ParseInt converts an interface{} value to an int.
// Supported types: int, int32, int64, float32, float64, string, fmt.Stringer.
// Float values are rounded to the nearest integer.
// String values are parsed using strconv.Atoi.
// Returns an error if the value cannot be converted to int.
func ParseInt(ctx context.Context, value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float32:
		return int(math.Round(float64(v))), nil
	case float64:
		return int(math.Round(v)), nil
	case string:
		return strconv.Atoi(v)
	case fmt.Stringer:
		return strconv.Atoi(v.String())
	default:
		return ParseInt(ctx, fmt.Sprintf("%v", value))
	}
}

// ParseIntDefault converts an interface{} value to an int, returning defaultValue on error.
// This is a convenience wrapper around ParseInt that never returns an error.
func ParseIntDefault(ctx context.Context, value interface{}, defaultValue int) int {
	result, err := ParseInt(ctx, value)
	if err != nil {
		return defaultValue
	}
	return result
}
