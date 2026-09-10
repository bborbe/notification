// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"context"
	"fmt"
	"strings"

	"github.com/bborbe/errors"
)

// ParseBool converts an interface{} value to a bool.
// Supported types: bool, string (case-insensitive "true"/"false"), fmt.Stringer.
// String values are converted to lowercase and compared against "true" and "false".
// Returns an error if the value cannot be converted to bool.
func ParseBool(ctx context.Context, value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		switch strings.ToLower(v) {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
		return false, errors.Errorf(ctx, "invalid type")
	case fmt.Stringer:
		return ParseBool(ctx, v.String())
	default:
		return ParseBool(ctx, fmt.Sprintf("%v", value))
	}
}

// ParseBoolDefault converts an interface{} value to a bool, returning defaultValue on error.
// This is a convenience wrapper around ParseBool that never returns an error.
func ParseBoolDefault(ctx context.Context, value interface{}, defaultValue bool) bool {
	result, err := ParseBool(ctx, value)
	if err != nil {
		return defaultValue
	}
	return result
}
