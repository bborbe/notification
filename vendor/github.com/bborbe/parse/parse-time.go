// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"context"
	"time"

	"github.com/bborbe/errors"
)

// ParseTime converts an interface{} value to a time.Time using the specified format.
// The value is first converted to a string using ParseString, then parsed using time.Parse.
// Format should follow Go's time format layout (e.g., "2006-01-02", "2006-01-02T15:04:05Z07:00").
// Returns an error if the value cannot be converted to time.Time.
func ParseTime(ctx context.Context, value interface{}, format string) (time.Time, error) {
	str, err := ParseString(ctx, value)
	if err != nil {
		return time.Time{}, errors.Wrapf(ctx, err, "parse %v as string failed", value)
	}
	t, err := time.Parse(format, str)
	if err != nil {
		return time.Time{}, errors.Wrapf(
			ctx,
			err,
			"parse '%s' with format '%s' failed",
			value,
			format,
		)
	}
	return t, nil
}

// ParseTimeDefault converts an interface{} value to a time.Time using the specified format,
// returning defaultValue on error.
// This is a convenience wrapper around ParseTime that never returns an error.
func ParseTimeDefault(
	ctx context.Context,
	value interface{},
	format string,
	defaultValue time.Time,
) time.Time {
	result, err := ParseTime(ctx, value, format)
	if err != nil {
		return defaultValue
	}
	return result
}
