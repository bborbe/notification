// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"context"
	"encoding"
	"flag"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
)

// ParseArgs parses command-line arguments into the given struct using arg struct tags.
// See Parse() documentation for supported types and struct tag options.
//
// Parameters:
//   - ctx: Context for error handling
//   - data: Pointer to struct with arg tags
//   - args: Command-line arguments (typically os.Args[1:])
//
// Returns error if parsing fails or if default values are malformed.
func ParseArgs(ctx context.Context, data interface{}, args []string) error {
	values, err := argsToValues(ctx, data, args)
	if err != nil {
		return errors.Wrap(ctx, err, "args to values failed")
	}
	if err := Fill(ctx, data, values); err != nil {
		return errors.Wrap(ctx, err, "fill failed")
	}
	return nil
}

func handleCustomType(
	ctx context.Context,
	values map[string]interface{},
	tf reflect.StructField,
	ef reflect.Value,
	argName, defaultString string,
	found bool,
	usage string,
) (bool, error) {
	// Get the underlying type
	underlyingType := ef.Type()
	for underlyingType.Kind() == reflect.Pointer {
		underlyingType = underlyingType.Elem()
	}

	// Check if it's a named type (custom type) with an underlying primitive type
	if underlyingType.PkgPath() != "" && underlyingType.Kind() != reflect.Struct {
		switch underlyingType.Kind() {
		case reflect.String:
			values[tf.Name] = flag.CommandLine.String(argName, defaultString, usage)
			return true, nil
		case reflect.Bool:
			defaultValue, _ := strconv.ParseBool(defaultString)
			values[tf.Name] = flag.CommandLine.Bool(argName, defaultValue, usage)
			return true, nil
		case reflect.Int:
			defaultValue, _ := strconv.Atoi(defaultString)
			values[tf.Name] = flag.CommandLine.Int(argName, defaultValue, usage)
			return true, nil
		case reflect.Int64:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int64(argName, defaultValue, usage)
			return true, nil
		case reflect.Uint:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint(argName, uint(defaultValue), usage)
			return true, nil
		case reflect.Uint64:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint64(argName, defaultValue, usage)
			return true, nil
		case reflect.Int32:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int(argName, int(defaultValue), usage)
			return true, nil
		case reflect.Float64:
			defaultValue, _ := strconv.ParseFloat(defaultString, 64)
			values[tf.Name] = flag.CommandLine.Float64(argName, defaultValue, usage)
			return true, nil
		}
	}
	return false, nil
}

// parseSliceFromString splits a string by separator, trims whitespace from each element,
// and converts to the appropriate slice type based on the element type.
func parseSliceFromString(
	ctx context.Context,
	value string,
	separator string,
	elemType reflect.Type,
) (interface{}, error) {
	if value == "" {
		// Return empty slice of the appropriate type
		return reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0).Interface(), nil
	}

	parts := strings.Split(value, separator)
	// Trim whitespace from each part
	trimmed := make([]string, 0, len(parts))
	for _, p := range parts {
		t := strings.TrimSpace(p)
		if t != "" { // Skip empty parts after trimming
			trimmed = append(trimmed, t)
		}
	}

	// If all parts were empty after trimming, return empty slice
	if len(trimmed) == 0 {
		return reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0).Interface(), nil
	}

	// Check if element type implements encoding.TextUnmarshaler
	ptrType := reflect.PointerTo(elemType)
	if ptrType.Implements(reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()) {
		result := reflect.MakeSlice(reflect.SliceOf(elemType), len(trimmed), len(trimmed))
		for i, p := range trimmed {
			elem := result.Index(i).Addr().Interface().(encoding.TextUnmarshaler)
			if err := elem.UnmarshalText([]byte(p)); err != nil {
				return nil, errors.Wrapf(ctx, err, "unmarshal text %q failed", p)
			}
		}
		return result.Interface(), nil
	}

	// Convert based on element type
	switch elemType.Kind() {
	case reflect.String:
		return trimmed, nil
	case reflect.Int:
		result := make([]int, len(trimmed))
		for i, p := range trimmed {
			v, err := strconv.Atoi(p)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "parse int %q failed", p)
			}
			result[i] = v
		}
		return result, nil
	case reflect.Int64:
		result := make([]int64, len(trimmed))
		for i, p := range trimmed {
			v, err := strconv.ParseInt(p, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "parse int64 %q failed", p)
			}
			result[i] = v
		}
		return result, nil
	case reflect.Uint:
		result := make([]uint, len(trimmed))
		for i, p := range trimmed {
			v, err := strconv.ParseUint(p, 10, 0)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "parse uint %q failed", p)
			}
			result[i] = uint(v)
		}
		return result, nil
	case reflect.Uint64:
		result := make([]uint64, len(trimmed))
		for i, p := range trimmed {
			v, err := strconv.ParseUint(p, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "parse uint64 %q failed", p)
			}
			result[i] = v
		}
		return result, nil
	case reflect.Float64:
		result := make([]float64, len(trimmed))
		for i, p := range trimmed {
			v, err := strconv.ParseFloat(p, 64)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "parse float64 %q failed", p)
			}
			result[i] = v
		}
		return result, nil
	case reflect.Bool:
		result := make([]bool, len(trimmed))
		for i, p := range trimmed {
			v, err := strconv.ParseBool(p)
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "parse bool %q failed", p)
			}
			result[i] = v
		}
		return result, nil
	default:
		// Check for custom string types like Username
		if elemType.PkgPath() != "" && elemType.Kind() == reflect.String {
			// Return as []string, Fill() will convert via JSON to custom type
			return trimmed, nil
		}
		return nil, errors.Errorf(ctx, "unsupported slice element type: %v", elemType)
	}
}

//nolint:gocyclo // TODO: Refactor to reduce complexity (currently 101, limit is 30)
func argsToValues(
	ctx context.Context,
	data interface{},
	args []string,
) (map[string]interface{}, error) {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	values := make(map[string]interface{})
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		ef := e.Field(i)
		argName, ok := tf.Tag.Lookup("arg")
		if !ok {
			continue
		}
		defaultString, found := tf.Tag.Lookup("default")
		usage := tf.Tag.Get("usage")
		switch ef.Interface().(type) {
		case string:
			values[tf.Name] = flag.CommandLine.String(argName, defaultString, usage)
		case bool:
			defaultValue, _ := strconv.ParseBool(defaultString)
			values[tf.Name] = flag.CommandLine.Bool(argName, defaultValue, usage)
		case int:
			defaultValue, _ := strconv.Atoi(defaultString)
			values[tf.Name] = flag.CommandLine.Int(argName, defaultValue, usage)
		case int64:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int64(argName, defaultValue, usage)
		case uint:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint(argName, uint(defaultValue), usage)
		case uint64:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint64(argName, defaultValue, usage)
		case int32:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int(argName, int(defaultValue), usage)
		case float64:
			defaultValue, _ := strconv.ParseFloat(defaultString, 64)
			values[tf.Name] = flag.CommandLine.Float64(argName, defaultValue, usage)
		case *float64:
			if found {
				defaultValue, _ := strconv.ParseFloat(defaultString, 64)
				values[tf.Name] = defaultValue
			}
			flag.CommandLine.Func(argName, usage, func(s string) error {
				if s == "" {
					return nil
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return errors.Wrap(ctx, err, "parse float failed")
				}
				values[tf.Name] = v
				return nil
			})
		case time.Duration:
			if found {
				defaultValue, _ := libtime.ParseDuration(ctx, defaultString)
				if defaultValue != nil {
					values[tf.Name] = defaultValue.Duration()
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse duration failed")
				}
				values[tf.Name] = duration.Duration()
				return nil
			})
		case time.Time:
			if found {
				defaultValue, err := libtime.ParseTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				t, err := libtime.ParseTime(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse time failed")
				}
				values[tf.Name] = *t
				return nil
			})
		case *time.Time:
			if found {
				defaultValue, err := libtime.ParseTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				t, err := libtime.ParseTime(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse time failed")
				}
				values[tf.Name] = *t
				return nil
			})
		case *time.Duration:
			if found {
				defaultValue, err := libtime.ParseDuration(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = defaultValue.Duration()
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse duration failed")
				}
				values[tf.Name] = duration.Duration()
				return nil
			})
		case libtime.Duration:
			if found {
				defaultValue, err := libtime.ParseDuration(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse duration failed")
				}
				values[tf.Name] = *duration
				return nil
			})
		case *libtime.Duration:
			if found {
				defaultValue, err := libtime.ParseDuration(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse duration failed")
				}
				values[tf.Name] = *duration
				return nil
			})
		case libtime.DateTime:
			if found {
				defaultValue, err := libtime.ParseDateTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				dateTime, err := libtime.ParseDateTime(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse datetime failed")
				}
				values[tf.Name] = *dateTime
				return nil
			})
		case *libtime.DateTime:
			if found {
				defaultValue, err := libtime.ParseDateTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				dateTime, err := libtime.ParseDateTime(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse datetime failed")
				}
				values[tf.Name] = *dateTime
				return nil
			})
		case libtime.Date:
			if found {
				defaultValue, err := libtime.ParseDate(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				date, err := libtime.ParseDate(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse date failed")
				}
				values[tf.Name] = *date
				return nil
			})
		case *libtime.Date:
			if found {
				defaultValue, err := libtime.ParseDate(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				date, err := libtime.ParseDate(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse date failed")
				}
				values[tf.Name] = *date
				return nil
			})
		case libtime.UnixTime:
			if found {
				defaultValue, err := libtime.ParseUnixTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				unixTime, err := libtime.ParseUnixTime(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse unixtime failed")
				}
				values[tf.Name] = *unixTime
				return nil
			})
		case *libtime.UnixTime:
			if found {
				defaultValue, err := libtime.ParseUnixTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				unixTime, err := libtime.ParseUnixTime(ctx, value)
				if err != nil {
					return errors.Wrap(ctx, err, "parse unixtime failed")
				}
				values[tf.Name] = *unixTime
				return nil
			})
		default:
			// Check if type implements encoding.TextUnmarshaler BEFORE checking for slice
			// This allows slice types like kafka.Brokers to implement TextUnmarshaler on the slice itself
			ptrType := reflect.PointerTo(ef.Type())
			if ptrType.Implements(reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()) {
				// Handle default value
				if found && defaultString != "" {
					unmarshaler := reflect.New(ef.Type()).Interface().(encoding.TextUnmarshaler)
					if err := unmarshaler.UnmarshalText([]byte(defaultString)); err != nil {
						return nil, errors.Wrapf(ctx, err, "unmarshal default value %q for field %s failed", defaultString, tf.Name)
					}
					values[tf.Name] = reflect.ValueOf(unmarshaler).Elem().Interface()
				}

				flag.CommandLine.Func(argName, usage, func(value string) error {
					if value == "" {
						return nil
					}
					unmarshaler := reflect.New(ef.Type()).Interface().(encoding.TextUnmarshaler)
					if err := unmarshaler.UnmarshalText([]byte(value)); err != nil {
						return errors.Wrap(ctx, err, "unmarshal text failed")
					}
					values[tf.Name] = reflect.ValueOf(unmarshaler).Elem().Interface()
					return nil
				})
				continue
			}

			// Check if it's a slice type (for slices that don't implement TextUnmarshaler)
			if ef.Type().Kind() == reflect.Slice {
				separator := tf.Tag.Get("separator")
				if separator == "" {
					separator = ","
				}
				elemType := ef.Type().Elem()

				// Handle default value for slices
				if found && defaultString != "" {
					parsed, err := parseSliceFromString(ctx, defaultString, separator, elemType)
					if err != nil {
						return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
					}
					values[tf.Name] = parsed
				}

				flag.CommandLine.Func(argName, usage, func(value string) error {
					parsed, err := parseSliceFromString(ctx, value, separator, elemType)
					if err != nil {
						return err
					}
					values[tf.Name] = parsed
					return nil
				})
				continue
			}

			// Check if it's a custom type with underlying primitive type
			if handled, err := handleCustomType(ctx, values, tf, ef, argName, defaultString, found, usage); handled {
				if err != nil {
					return nil, err
				}
			} else {
				return nil, errors.Errorf(ctx, "field %s with type %T is unsupported", tf.Name, ef.Interface())
			}
		}
	}
	if err := flag.CommandLine.Parse(args); err != nil {
		return nil, errors.Wrap(ctx, err, "parse commandline failed")
	}
	return values, nil
}

// argsToValuesExplicit returns only values that were explicitly set via command-line arguments.
// Unlike argsToValues, this does not include default values for unset flags.
// This is used internally to ensure proper precedence: args > env > defaults.
func argsToValuesExplicit(
	ctx context.Context,
	data interface{},
	args []string,
) (map[string]interface{}, error) {
	// First get all values (including defaults)
	allValues, err := argsToValues(ctx, data, args)
	if err != nil {
		return nil, err
	}

	// Then filter to only explicitly-set flags
	actuallySet := make(map[string]interface{})
	visitedFlags := make(map[string]bool)
	flag.CommandLine.Visit(func(f *flag.Flag) {
		visitedFlags[f.Name] = true
	})

	// Map flag names back to struct field names
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		argName, ok := tf.Tag.Lookup("arg")
		if !ok {
			continue
		}
		if visitedFlags[argName] {
			if val, exists := allValues[tf.Name]; exists {
				actuallySet[tf.Name] = val
			}
		}
	}

	return actuallySet, nil
}
