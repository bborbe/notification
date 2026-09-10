// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/bborbe/errors"
)

func handleCustomTypeValidation(
	ctx context.Context,
	tf reflect.StructField,
	ef reflect.Value,
	createError func() error,
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
			// For custom string types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		case reflect.Bool:
			// Bool types are never considered "empty" for required validation
			return true, nil
		case reflect.Int, reflect.Int32, reflect.Int64:
			// For custom int types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		case reflect.Uint, reflect.Uint64:
			// For custom uint types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		case reflect.Float64:
			// For custom float64 types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		}
	}
	return false, nil
}

// ValidateRequired fields are set and returns an error if not.
func ValidateRequired(ctx context.Context, data interface{}) error {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		ef := e.Field(i)
		argName, ok := tf.Tag.Lookup("required")
		if !ok || argName != "true" {
			continue
		}
		if err := validateRequiredField(ctx, tf, ef); err != nil {
			return err
		}
	}
	return nil
}

// validateRequiredField checks if a single required field is set.
func validateRequiredField(ctx context.Context, tf reflect.StructField, ef reflect.Value) error {
	createError := func() error {
		buf := bytes.NewBufferString("Required field empty, ")
		argName, argOk := tf.Tag.Lookup("arg")
		if argOk {
			fmt.Fprintf(buf, "define parameter %s", argName)
		}
		envName, envOk := tf.Tag.Lookup("env")
		if envOk {
			if argOk {
				fmt.Fprintf(buf, " or ")
			}
			fmt.Fprintf(buf, "define env %s", envName)
		}
		return errors.New(ctx, buf.String())
	}

	switch ef.Interface().(type) {
	case string:
		var empty string
		if empty == ef.Interface() {
			return createError()
		}
	case bool:
	case int:
		var empty int
		if empty == ef.Interface() {
			return createError()
		}
	case int64:
		var empty int64
		if empty == ef.Interface() {
			return createError()
		}
	case uint:
		var empty uint
		if empty == ef.Interface() {
			return createError()
		}
	case uint64:
		var empty uint64
		if empty == ef.Interface() {
			return createError()
		}
	case int32:
		var empty int32
		if empty == ef.Interface() {
			return createError()
		}
	case float64:
		var empty float64
		if empty == ef.Interface() {
			return createError()
		}
	case *float64:
		var empty *float64
		if empty == ef.Interface() {
			return createError()
		}
	case time.Duration:
		var empty time.Duration
		if empty == ef.Interface() {
			return createError()
		}
	default:
		// Handle pointers
		if ef.Kind() == reflect.Pointer {
			if ef.IsNil() {
				return createError()
			}
		} else if ef.Kind() == reflect.Slice {
			// Handle slices
			if ef.Len() == 0 {
				return createError()
			}
		} else {
			// Check if it's a custom type with underlying primitive type
			if handled, err := handleCustomTypeValidation(ctx, tf, ef, createError); handled {
				if err != nil {
					return err
				}
			} else {
				return errors.Errorf(ctx, "field %s with type %T is unsupported", tf.Name, ef.Interface())
			}
		}
	}
	return nil
}

// ValidateHasValidation validates data using the HasValidation interface.
// It first checks if the top-level struct implements HasValidation.
// Then it iterates through struct fields:
//   - For slices: validates the slice type first, then falls back to validating each element
//   - For other types: validates if they implement HasValidation
//
// Important: ValidateHasValidation runs on ALL fields that implement HasValidation,
// regardless of whether they have the `required:"true"` tag. This is by design:
//   - The `required` tag checks if a field is present (non-zero)
//   - The Validate() method checks if a field's value is valid
//
// These are separate concerns. If you have an optional field with a default value,
// that default value should still be validated. If you want zero values to be valid
// for optional fields, your Validate() implementation should explicitly handle that:
//
//	func (p Port) Validate(ctx context.Context) error {
//	    if p == 0 {
//	        return nil // zero value is valid for optional ports
//	    }
//	    if p < 1024 {
//	        return fmt.Errorf("port must be >= 1024")
//	    }
//	    return nil
//	}
//
// Returns the first validation error encountered.
//
// Example usage (automatic validation via Parse):
//
//	type Port int
//	func (p Port) Validate(ctx context.Context) error {
//	    if p < 1 || p > 65535 {
//	        return fmt.Errorf("port must be between 1 and 65535, got %d", p)
//	    }
//	    return nil
//	}
//
//	type Config struct {
//	    Port Port `arg:"port" default:"8080"`
//	}
//
//	var config Config
//	if err := argument.Parse(ctx, &config); err != nil {
//	    // Parse automatically calls ValidateHasValidation
//	    // Validation error will be returned if port is out of range
//	}
//
// Example usage (manual validation workflow):
//
//	var config Config
//	if err := argument.ParseOnly(ctx, &config); err != nil {
//	    return err
//	}
//	// Custom logic here (e.g., override certain fields, apply business rules)
//	if config.Port == 0 {
//	    config.Port = 8080
//	}
//	// Then run standard validation
//	if err := argument.ValidateRequired(ctx, &config); err != nil {
//	    return err
//	}
//	if err := argument.ValidateHasValidation(ctx, &config); err != nil {
//	    return err
//	}
//
// For slice validation, the slice type is checked first:
//
//	type Brokers []Broker
//	func (b Brokers) Validate(ctx context.Context) error {
//	    if len(b) == 0 {
//	        return fmt.Errorf("at least one broker required")
//	    }
//	    // Will automatically validate each Broker if it implements HasValidation
//	    return nil
//	}
func ValidateHasValidation(ctx context.Context, data interface{}) error {
	// First, check if the top-level struct implements HasValidation
	if validator, ok := data.(HasValidation); ok {
		if err := validator.Validate(ctx); err != nil {
			return errors.Wrap(ctx, err, "validation failed")
		}
	}

	// Now validate fields
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		ef := e.Field(i)
		tf := t.Field(i)

		// Skip unexported fields
		if !ef.CanInterface() {
			continue
		}

		if err := validateField(ctx, tf.Name, ef); err != nil {
			return err
		}
	}

	return nil
}

// validateField validates a single field that may implement HasValidation.
func validateField(ctx context.Context, fieldName string, fieldValue reflect.Value) error {
	// Handle slices specially
	if fieldValue.Kind() == reflect.Slice {
		return validateSlice(ctx, fieldName, fieldValue)
	}

	// For non-slice types, check if they implement HasValidation
	if fieldValue.CanInterface() {
		// Skip nil pointers to avoid panic when calling methods
		if fieldValue.Kind() == reflect.Pointer && fieldValue.IsNil() {
			return nil
		}

		if validator, ok := fieldValue.Interface().(HasValidation); ok {
			if err := validator.Validate(ctx); err != nil {
				return errors.Wrapf(
					ctx,
					err,
					"field %s (type %s) validation failed",
					fieldName,
					fieldValue.Type(),
				)
			}
		}
	}

	return nil
}

// validateSlice validates a slice field.
// First checks if the slice type itself implements HasValidation.
// Falls back to validating each element if they implement HasValidation.
func validateSlice(ctx context.Context, fieldName string, sliceValue reflect.Value) error {
	// First, check if the slice itself implements HasValidation
	if sliceValue.CanInterface() {
		if validator, ok := sliceValue.Interface().(HasValidation); ok {
			if err := validator.Validate(ctx); err != nil {
				return errors.Wrapf(
					ctx,
					err,
					"field %s (type %s) validation failed",
					fieldName,
					sliceValue.Type(),
				)
			}
			return nil
		}
	}

	// Fallback: validate each element
	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i)
		if elem.CanInterface() {
			if validator, ok := elem.Interface().(HasValidation); ok {
				if err := validator.Validate(ctx); err != nil {
					return errors.Wrapf(
						ctx,
						err,
						"field %s[%d] (type %s) validation failed",
						fieldName,
						i,
						elem.Type(),
					)
				}
			}
		}
	}

	return nil
}
