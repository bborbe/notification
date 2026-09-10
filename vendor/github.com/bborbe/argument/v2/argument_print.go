// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Print all configured arguments. Set display:"hidden" to hide or display:"length" to only print the arguments length.
func Print(ctx context.Context, data interface{}) error {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		// Skip unexported fields: reflect.Value.Interface() panics on them, and
		// they are never argument targets (Parse only fills tagged exported fields).
		if !t.Field(i).IsExported() {
			continue
		}
		ef := e.Field(i)
		argName := t.Field(i).Tag.Get("display")
		if argName == "hidden" {
			continue
		}
		if argName == "length" {
			log.Printf(
				"Argument: %s length %d",
				t.Field(i).Name,
				len(fmt.Sprintf("%v", ef.Interface())),
			)
			continue
		}
		if ef.Kind() == reflect.Slice {
			// Format slices as comma-separated values with count
			length := ef.Len()
			if length == 0 {
				log.Printf("Argument: %s []", t.Field(i).Name)
			} else {
				values := make([]string, length)
				for j := 0; j < length; j++ {
					values[j] = fmt.Sprintf("%v", ef.Index(j).Interface())
				}
				log.Printf("Argument: %s [%d]: %s", t.Field(i).Name, length, strings.Join(values, ", "))
			}
		} else if ef.Kind() == reflect.Pointer || ef.Kind() == reflect.Interface {
			if ef.IsZero() {
				log.Printf("Argument: %s <nil>", t.Field(i).Name)
			} else {
				log.Printf("Argument: %s '%v'", t.Field(i).Name, ef.Elem())
			}
		} else {
			log.Printf("Argument: %s '%v'", t.Field(i).Name, ef.Interface())
		}
	}
	return nil
}
