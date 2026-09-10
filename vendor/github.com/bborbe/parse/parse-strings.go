// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"context"
	"reflect"

	"github.com/bborbe/errors"
)

// HasStrings interface is implemented by types that can provide a string slice representation.
type HasStrings interface {
	Strings() []string
}

// HasString interface is implemented by types that can provide a string representation.
// This is similar to fmt.Stringer but used specifically for ParseStrings conversion.
type HasString interface {
	String() string
}

// ParseStrings converts an interface{} value to a string slice.
// Supported types: []string, []interface{}, []float64, []bool, []int, []int32, []int64, string,
// HasStrings interface, HasString interface, slices of string subtypes (e.g., []Direction where type Direction string),
// and slices of types implementing String() string method.
// A single string value is returned as a slice with one element.
// Returns nil for nil input.
// Returns an error if the value cannot be converted to []string.
func ParseStrings(ctx context.Context, value interface{}) ([]string, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case []string:
		return v, nil
	case []interface{}:
		return toStringList(ctx, v)
	case []float64:
		return toStringList(ctx, v)
	case []bool:
		return toStringList(ctx, v)
	case []int:
		return toStringList(ctx, v)
	case []int32:
		return toStringList(ctx, v)
	case []int64:
		return toStringList(ctx, v)
	case string:
		str, err := ParseString(ctx, v)
		if err != nil {
			return nil, err
		}
		return []string{str}, nil
	case HasStrings:
		return v.Strings(), nil
	case HasString:
		return []string{v.String()}, nil
	default:
		if isSliceOfStringSubtype(value) {
			return convertSliceToStrings(ctx, value)
		}
		if isSliceOfHasString(value) {
			return convertSliceToStrings(ctx, value)
		}
		return nil, errors.Errorf(ctx, "unsupported type %T", value)
	}
}

func toStringList[T any](ctx context.Context, input []T) ([]string, error) {
	result := make([]string, len(input))
	for i, a := range input {
		str, err := ParseString(ctx, a)
		if err != nil {
			return nil, err
		}
		result[i] = str
	}
	return result, nil
}

// ParseStringsDefault converts an interface{} value to a string slice, returning defaultValue on error.
// This is a convenience wrapper around ParseStrings that never returns an error.
func ParseStringsDefault(ctx context.Context, value interface{}, defaultValue []string) []string {
	result, err := ParseStrings(ctx, value)
	if err != nil {
		return defaultValue
	}
	return result
}

// isSliceOfStringSubtype checks if the value is a slice whose elements have an underlying string type.
// For example, []Direction where type Direction string would return true.
func isSliceOfStringSubtype(value interface{}) bool {
	t := reflect.TypeOf(value)
	if t == nil || t.Kind() != reflect.Slice {
		return false
	}
	elemType := t.Elem()
	return elemType.Kind() == reflect.String
}

// isSliceOfHasString checks if the value is a slice whose elements implement the String() string method.
// This handles []T where T implements fmt.Stringer or HasString interface, as well as []HasString (interface type).
func isSliceOfHasString(value interface{}) bool {
	t := reflect.TypeOf(value)
	if t == nil || t.Kind() != reflect.Slice {
		return false
	}
	elemType := t.Elem()

	// First check if element type is assignable to HasString interface
	hasStringInterfaceType := reflect.TypeOf((*HasString)(nil)).Elem()
	if elemType.Implements(hasStringInterfaceType) {
		return true
	}

	// Also check for pointer types implementing HasString
	if elemType.Kind() == reflect.Pointer {
		if elemType.Implements(hasStringInterfaceType) {
			return true
		}
	}

	// Fallback: Check if element type has a String() string method
	// This handles concrete types that implement String() but aren't explicitly typed as interface
	stringMethod, hasMethod := elemType.MethodByName("String")
	if !hasMethod {
		return false
	}

	// Verify the method signature: func() string
	methodType := stringMethod.Type
	// Method type includes receiver, so: func(T) string has 1 in, 1 out
	if methodType.NumIn() != 1 || methodType.NumOut() != 1 {
		return false
	}

	// Check return type is string
	stringType := reflect.TypeOf("")
	return methodType.Out(0) == stringType
}

// convertSliceToStrings converts a slice of any type to []string using reflection.
// Each element is converted using ParseString which handles string subtypes and fmt.Stringer.
func convertSliceToStrings(ctx context.Context, value interface{}) ([]string, error) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice {
		return nil, errors.Errorf(ctx, "value is not a slice: %T", value)
	}

	length := v.Len()
	result := make([]string, length)

	for i := 0; i < length; i++ {
		elem := v.Index(i).Interface()
		str, err := ParseString(ctx, elem)
		if err != nil {
			return nil, errors.Wrapf(ctx, err, "failed to convert element %d", i)
		}
		result[i] = str
	}

	return result, nil
}
