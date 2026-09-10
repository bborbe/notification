// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"context"
	"encoding"
	"reflect"
	"strconv"
	"time"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
)

// ParseEnv parses environment variables into the given struct using env struct tags.
// See Parse() documentation for supported types and struct tag options.
//
// Parameters:
//   - ctx: Context for error handling
//   - data: Pointer to struct with env tags
//   - environ: Environment variables (typically os.Environ())
//
// Returns error if parsing fails.
func ParseEnv(ctx context.Context, data interface{}, environ []string) error {
	values, err := envToValues(ctx, data, environ)
	if err != nil {
		return errors.Wrap(ctx, err, "env to values failed")
	}
	if err := Fill(ctx, data, values); err != nil {
		return errors.Wrap(ctx, err, "fill failed")
	}
	return nil
}

//nolint:dupl // TODO: Extract shared logic with handleCustomTypeDefault to eliminate duplication
func handleCustomTypeEnv(
	ctx context.Context,
	values map[string]interface{},
	tf reflect.StructField,
	ef reflect.Value,
	value string,
) (bool, error) {
	// Get the underlying type
	underlyingType := ef.Type()
	for underlyingType.Kind() == reflect.Pointer {
		underlyingType = underlyingType.Elem()
	}

	// Check if it's a named type (custom type) with an underlying primitive type
	if underlyingType.PkgPath() != "" && underlyingType.Kind() != reflect.Struct {
		var err error
		switch underlyingType.Kind() {
		case reflect.String:
			values[tf.Name] = value
			return true, nil
		case reflect.Bool:
			values[tf.Name], err = strconv.ParseBool(value)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			return true, nil
		case reflect.Int:
			values[tf.Name], err = strconv.Atoi(value)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			return true, nil
		case reflect.Int64:
			values[tf.Name], err = strconv.ParseInt(value, 10, 0)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			return true, nil
		case reflect.Uint:
			values[tf.Name], err = strconv.ParseUint(value, 10, 0)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			return true, nil
		case reflect.Uint64:
			values[tf.Name], err = strconv.ParseUint(value, 10, 0)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			return true, nil
		case reflect.Int32:
			v, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			values[tf.Name] = int32(v)
			return true, nil
		case reflect.Float64:
			values[tf.Name], err = strconv.ParseFloat(value, 64)
			if err != nil {
				return true, errors.Errorf(
					ctx,
					"parse field %s as %T failed: %v",
					tf.Name,
					ef.Interface(),
					err,
				)
			}
			return true, nil
		}
	}
	return false, nil
}

//nolint:gocyclo // TODO: Refactor to reduce complexity (currently 53, limit is 30)
func envToValues(
	ctx context.Context,
	data interface{},
	environ []string,
) (map[string]interface{}, error) {
	var err error
	envValues := make(map[string]string)
	for _, env := range environ {
		for i := 0; i < len(env); i++ {
			if env[i] == '=' {
				envValues[env[:i]] = env[i+1:]
			}
		}
	}
	values := make(map[string]interface{})
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		ef := e.Field(i)
		argName, ok := tf.Tag.Lookup("env")
		if !ok {
			continue
		}
		value, ok := envValues[argName]
		if !ok {
			continue
		}
		switch ef.Interface().(type) {
		case string:
			values[tf.Name] = value
		case bool:
			values[tf.Name], err = strconv.ParseBool(value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
		case int:
			values[tf.Name], err = strconv.Atoi(value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
		case int64:
			values[tf.Name], err = strconv.ParseInt(value, 10, 0)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
		case uint:
			values[tf.Name], err = strconv.ParseUint(value, 10, 0)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
		case uint64:
			values[tf.Name], err = strconv.ParseUint(value, 10, 0)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
		case int32:
			v, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = int32(v)
		case float64:
			values[tf.Name], err = strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
		case time.Duration:
			duration, err := libtime.ParseDuration(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = duration.Duration()
		case time.Time:
			t, err := libtime.ParseTime(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *t
		case *time.Time:
			t, err := libtime.ParseTime(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *t
		case *time.Duration:
			duration, err := libtime.ParseDuration(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = duration.Duration()
		case libtime.Duration:
			duration, err := libtime.ParseDuration(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *duration
		case *libtime.Duration:
			duration, err := libtime.ParseDuration(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *duration
		case libtime.DateTime:
			dateTime, err := libtime.ParseDateTime(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *dateTime
		case *libtime.DateTime:
			dateTime, err := libtime.ParseDateTime(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *dateTime
		case libtime.Date:
			date, err := libtime.ParseDate(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *date
		case *libtime.Date:
			date, err := libtime.ParseDate(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *date
		case libtime.UnixTime:
			unixTime, err := libtime.ParseUnixTime(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *unixTime
		case *libtime.UnixTime:
			unixTime, err := libtime.ParseUnixTime(ctx, value)
			if err != nil {
				return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
			}
			values[tf.Name] = *unixTime
		//nolint:dupl // TODO: Extract shared type handling logic with defaultToValues switch statement
		default:
			// Check if type implements encoding.TextUnmarshaler BEFORE checking for slice
			// This allows slice types like kafka.Brokers to implement TextUnmarshaler on the slice itself
			ptrType := reflect.PointerTo(ef.Type())
			if ptrType.Implements(reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()) {
				unmarshaler := reflect.New(ef.Type()).Interface().(encoding.TextUnmarshaler)
				if err := unmarshaler.UnmarshalText([]byte(value)); err != nil {
					return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
				}
				values[tf.Name] = reflect.ValueOf(unmarshaler).Elem().Interface()
				continue
			}

			// Check if it's a slice type (for slices that don't implement TextUnmarshaler)
			if ef.Type().Kind() == reflect.Slice {
				separator := tf.Tag.Get("separator")
				if separator == "" {
					separator = ","
				}
				elemType := ef.Type().Elem()

				parsed, err := parseSliceFromString(ctx, value, separator, elemType)
				if err != nil {
					return nil, errors.Errorf(ctx, "parse field %s as %T failed: %v", tf.Name, ef.Interface(), err)
				}
				values[tf.Name] = parsed
				continue
			}

			// Check if it's a custom type with underlying primitive type
			if handled, err := handleCustomTypeEnv(ctx, values, tf, ef, value); handled {
				if err != nil {
					return nil, err
				}
			} else {
				return nil, errors.Errorf(ctx, "field %s with type %T is unsupported", tf.Name, ef.Interface())
			}
		}
	}
	return values, nil
}
