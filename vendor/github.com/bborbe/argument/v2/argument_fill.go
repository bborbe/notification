// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"bytes"
	"context"
	"encoding"
	"encoding/json"
	"reflect"

	"github.com/bborbe/errors"
)

// Fill populates the given struct with values from the provided map using JSON marshaling.
// It encodes the map to JSON and then decodes it back into the struct, allowing for
// flexible type conversion and nested structure population.
//
// Types implementing encoding.TextMarshaler are converted to strings before JSON encoding
// to ensure compatibility with types that implement encoding.TextUnmarshaler.
// For slices of TextMarshaler, each element is converted to a string separately to maintain
// array structure in JSON, allowing proper unmarshaling into slice types.
//
// Parameters:
//   - ctx: Context for error handling
//   - data: Pointer to struct to populate
//   - values: Map of field names to values
//
// Returns error if JSON encoding or decoding fails.
func Fill(ctx context.Context, data interface{}, values map[string]interface{}) error {
	// Convert TextMarshaler types to strings for JSON compatibility
	jsonValues := make(map[string]interface{}, len(values))
	for k, v := range values {
		if v == nil {
			jsonValues[k] = v
			continue
		}

		// If the value implements json.Marshaler, let JSON encoding handle it directly.
		// This is important for types like UnixTime where MarshalJSON and MarshalText
		// produce different formats (e.g., integer vs ISO string).
		if _, ok := v.(json.Marshaler); ok {
			jsonValues[k] = v
			continue
		}

		// Check if the value itself implements encoding.TextMarshaler (including slices)
		// This must be checked BEFORE checking element types, because types like
		// kafka.Brokers implement TextUnmarshaler on the slice type itself
		if marshaler, ok := v.(encoding.TextMarshaler); ok {
			text, err := marshaler.MarshalText()
			if err != nil {
				return errors.Wrapf(ctx, err, "marshal text for field %s failed", k)
			}
			jsonValues[k] = string(text)
			continue
		}

		rv := reflect.ValueOf(v)

		// Check if value is a slice where elements implement TextMarshaler
		// but the slice type itself does not
		if rv.Kind() == reflect.Slice {
			elemType := rv.Type().Elem()
			// Check if element type implements TextMarshaler
			if elemType.Implements(reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()) {
				// Convert []TextMarshaler to []string for JSON compatibility
				// This allows JSON to unmarshal into slice types where each element
				// implements TextUnmarshaler
				strSlice := make([]string, rv.Len())
				for i := 0; i < rv.Len(); i++ {
					elem := rv.Index(i).Interface().(encoding.TextMarshaler)
					text, err := elem.MarshalText()
					if err != nil {
						return errors.Wrapf(ctx, err, "marshal text for field %s[%d] failed", k, i)
					}
					strSlice[i] = string(text)
				}
				jsonValues[k] = strSlice
				continue
			}
		}

		// Pass through as-is for standard types
		jsonValues[k] = v
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(jsonValues); err != nil {
		return errors.Wrap(ctx, err, "encode json failed")
	}
	if err := json.NewDecoder(buf).Decode(data); err != nil {
		return errors.Wrap(ctx, err, "decode json failed")
	}
	return nil
}
