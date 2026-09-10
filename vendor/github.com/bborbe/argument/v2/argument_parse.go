// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"context"
	"os"

	"github.com/bborbe/errors"
)

// Parse combines all functionality. It parses command-line arguments and environment variables
// into a struct using struct tags, then validates required fields are set.
//
// Supported Types:
//   - Basic types: string, bool, int, int32, int64, uint, uint64, float64
//   - Pointer types: *float64 (optional values, nil if not provided)
//   - Slice types: []string, []int, []int64, []uint, []uint64, []float64, []bool
//   - Custom type slices: []Username where type Username string
//   - Custom types implementing encoding.TextUnmarshaler: For complex parsing logic
//   - Standard library time types:
//   - time.Time and *time.Time: RFC3339 format (e.g., "2006-01-02T15:04:05Z")
//   - time.Duration and *time.Duration: Extended format supporting days (e.g., "1d2h30m", "7d")
//   - github.com/bborbe/time types:
//   - libtime.Duration and *libtime.Duration: Extended duration with weeks (e.g., "2w", "1w3d")
//   - libtime.DateTime and *libtime.DateTime: Timestamp with timezone
//   - libtime.Date and *libtime.Date: Date only (e.g., "2006-01-02")
//   - libtime.UnixTime and *libtime.UnixTime: Unix timestamp (seconds since epoch)
//
// Pointer types (*Type) are optional and will be nil if not provided or if provided as empty string.
// Non-pointer types will use zero values if not provided.
//
// Slice types support comma-separated values by default (e.g., "alice,bob,charlie").
// Whitespace around each element is automatically trimmed.
// Use the separator tag to customize the delimiter (e.g., separator:":").
//
// Struct Tags:
//   - arg: Command-line argument name (required to parse field)
//   - env: Environment variable name (optional)
//   - default: Default value if not provided (optional)
//   - separator: Separator for slice values (default: ",", optional)
//   - required: Mark field as required (optional)
//   - display: Control how value is displayed - "length" shows only length for sensitive data (optional)
//   - usage: Help text for the argument (optional)
//
// Example:
//
//	type Config struct {
//	    Host     string        `arg:"host" env:"HOST" default:"localhost" usage:"Server hostname"`
//	    Port     int           `arg:"port" env:"PORT" default:"8080" required:"true"`
//	    Timeout  time.Duration `arg:"timeout" default:"30s" usage:"Request timeout"`
//	    StartAt  *time.Time    `arg:"start" usage:"Optional start time"`
//	    Password string        `arg:"password" env:"PASSWORD" display:"length" usage:"API password"`
//	    Names    []string      `arg:"names" env:"NAMES" default:"alice,bob" usage:"User names"`
//	    Ports    []int         `arg:"ports" env:"PORTS" separator:":" usage:"Port numbers"`
//	}
//
// Custom types can implement encoding.TextUnmarshaler for specialized parsing:
//
//	type Broker string
//
//	func (b *Broker) UnmarshalText(text []byte) error {
//	    value := string(text)
//	    if !strings.Contains(value, "://") {
//	        value = "plain://" + value  // Add default schema
//	    }
//	    *b = Broker(value)
//	    return nil
//	}
//
//	type Config struct {
//	    Broker  Broker   `arg:"broker" default:"localhost:9092"`
//	    Brokers []Broker `arg:"brokers" env:"BROKERS" usage:"Kafka brokers"`
//	}
//
// Precedence: Command-line arguments override environment variables, which override defaults.
func Parse(ctx context.Context, data interface{}) error {
	if err := ParseOnly(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "parse failed")
	}
	if err := ValidateRequired(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "validate required failed")
	}
	if err := ValidateHasValidation(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "validate failed")
	}
	return nil
}

// ParseAndPrint parses command-line arguments and environment variables into a struct,
// prints the parsed configuration to stdout, then validates required fields.
// It combines Parse() functionality with Print() output, useful for debugging and
// confirming configuration during application startup.
//
// See Parse() documentation for supported types and struct tag options.
func ParseAndPrint(ctx context.Context, data interface{}) error {
	if err := ParseOnly(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "parse failed")
	}
	if err := Print(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "print failed")
	}
	if err := ValidateRequired(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "validate required failed")
	}
	if err := ValidateHasValidation(ctx, data); err != nil {
		return errors.Wrap(ctx, err, "validate failed")
	}
	return nil
}

// ParseOnly parses command-line arguments and environment variables into a struct
// WITHOUT any validation. This is useful when you want to handle validation separately
// or implement custom validation logic.
//
// ParseOnly combines arguments, environment variables, and defaults into the struct
// but skips both ValidateRequired and ValidateHasValidation checks.
//
// Example custom validation workflow:
//
//	if err := argument.ParseOnly(ctx, &config); err != nil {
//	    return err
//	}
//	// Custom validation
//	if config.Port < 1024 {
//	    return errors.New(ctx, "port must be >= 1024")
//	}
//	// Then run standard validation if needed
//	if err := argument.ValidateRequired(ctx, &config); err != nil {
//	    return err
//	}
//	if err := argument.ValidateHasValidation(ctx, &config); err != nil {
//	    return err
//	}
//
// See Parse() documentation for supported types and struct tag options.
func ParseOnly(ctx context.Context, data interface{}) error {
	argsValues, err := argsToValuesExplicit(ctx, data, os.Args[1:])
	if err != nil {
		return errors.Wrap(ctx, err, "arg to values failed")
	}
	envValues, err := envToValues(ctx, data, os.Environ())
	if err != nil {
		return errors.Wrap(ctx, err, "env to values failed")
	}
	defaultValues, err := DefaultValues(ctx, data)
	if err != nil {
		return errors.Wrap(ctx, err, "default values failed")
	}
	if err := Fill(ctx, data, mergeValues(defaultValues, envValues, argsValues)); err != nil {
		return errors.Wrap(ctx, err, "fill failed")
	}
	return nil
}

func mergeValues(list ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, values := range list {
		for k, v := range values {
			result[k] = v
		}
	}
	return result
}
