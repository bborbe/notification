// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"context"
	"encoding/json"
	"sort"
	"strings"
	"sync"
	"unsafe"
)

// Set represents a thread-safe collection of unique elements.
// This implementation uses a map for O(1) average-case lookups, additions, and deletions.
type Set[T comparable] interface {
	// Add inserts elements into the set. Duplicate elements are automatically ignored.
	// Multiple elements can be added in a single call with only one mutex lock.
	Add(elements ...T)
	// Remove deletes elements from the set.
	// Multiple elements can be removed in a single call with only one mutex lock.
	Remove(elements ...T)
	// Contains reports whether an element is present in the set.
	Contains(element T) bool
	// ContainsAll reports whether all given elements are present in the set.
	ContainsAll(elements ...T) bool
	// ContainsAny reports whether at least one of the given elements is present in the set.
	ContainsAny(elements ...T) bool
	// Slice returns all elements as a slice in arbitrary order.
	Slice() []T
	// Length returns the number of elements in the set.
	Length() int
	// Strings returns all elements as their string representations in sorted order.
	// This provides deterministic output suitable for debugging and logging.
	Strings() []string
	// String returns a human-readable string representation of the set.
	String() string
	// Each calls fn for each element in the set. Iteration stops on first error.
	// The order of iteration is arbitrary and not guaranteed to be consistent.
	Each(ctx context.Context, fn func(ctx context.Context, value T) error) error
	// Clone returns a new Set containing all elements from the current set.
	// The returned set is a shallow copy - modifications to it won't affect the original.
	Clone() Set[T]
	// Without returns a new Set containing all elements from the current set
	// except those specified in the elements parameter.
	// The original set is not modified.
	Without(elements ...T) Set[T]
	// UnmarshalText parses comma-separated text into set elements.
	// It implements encoding.TextUnmarshaler for automatic parsing with argument packages.
	UnmarshalText(text []byte) error
	// MarshalText converts set elements to comma-separated text.
	// It implements encoding.TextMarshaler for automatic serialization.
	MarshalText() ([]byte, error)
	// UnmarshalJSON deserializes a JSON array into set elements.
	// It implements json.Unmarshaler for automatic JSON parsing.
	UnmarshalJSON(data []byte) error
	// MarshalJSON serializes set elements to a JSON array.
	// It implements json.Marshaler for automatic JSON serialization.
	MarshalJSON() ([]byte, error)
}

// NewSet creates a new thread-safe set for comparable types.
// It accepts optional initial elements to populate the set.
// Duplicate elements are automatically handled.
//
// Performance: This implementation uses a map-based approach with O(1) average-case
// operations for Add, Remove, and Contains. Initialization is O(n) for n elements.
//
// Example:
//
//	set := collection.NewSet(1, 2, 3)
//	empty := collection.NewSet[int]()
func NewSet[T comparable](elements ...T) Set[T] {
	s := &set[T]{
		data: make(map[T]struct{}),
	}
	s.Add(elements...)
	return s
}

type set[T comparable] struct {
	mux  sync.Mutex
	data map[T]struct{}
}

func (s *set[T]) Add(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, element := range elements {
		s.data[element] = struct{}{}
	}
}

func (s *set[T]) Remove(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		delete(s.data, element)
	}
}

func (s *set[T]) Contains(element T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, found := s.data[element]
	return found
}

func (s *set[T]) ContainsAll(elements ...T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if _, found := s.data[element]; !found {
			return false
		}
	}
	return true
}

func (s *set[T]) ContainsAny(elements ...T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if _, found := s.data[element]; found {
			return true
		}
	}
	return false
}

func (s *set[T]) Slice() []T {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]T, 0, len(s.data))
	for k := range s.data {
		result = append(result, k)
	}
	return result
}

func (s *set[T]) Length() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return len(s.data)
}

// Strings returns all elements as their string representations in sorted order.
// This provides deterministic output suitable for debugging and logging.
func (s *set[T]) Strings() []string {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]string, 0, len(s.data))
	for k := range s.data {
		result = append(result, elementToString(k))
	}

	sort.Strings(result)
	return result
}

// String returns a human-readable string representation of the set.
// Format: "Set[element1, element2, ...]" for non-empty sets, "Set[]" for empty sets.
// Elements are sorted by their string representation for deterministic output.
func (s *set[T]) String() string {
	return formatSetString("Set[", s.Strings())
}

// Each calls fn for each element in the set. Iteration stops on first error.
// The order of iteration is arbitrary and not guaranteed to be consistent.
func (s *set[T]) Each(ctx context.Context, fn func(ctx context.Context, value T) error) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	for element := range s.data {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := fn(ctx, element); err != nil {
				return err
			}
		}
	}
	return nil
}

// Clone returns a new Set containing all elements from the current set.
// The returned set is a shallow copy - modifications to it won't affect the original.
func (s *set[T]) Clone() Set[T] {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := &set[T]{
		data: make(map[T]struct{}, len(s.data)),
	}

	for element := range s.data {
		result.data[element] = struct{}{}
	}

	return result
}

// Without returns a new Set containing all elements from the current set
// except those specified in the elements parameter.
// The original set is not modified.
func (s *set[T]) Without(elements ...T) Set[T] {
	result := s.Clone()
	result.Remove(elements...)
	return result
}

// ParseSetFromStrings converts a slice of strings into a Set with string-based type.
// T must be string or a type based on string (using ~string constraint).
func ParseSetFromStrings[T ~string](values []string) Set[T] {
	result := make([]T, len(values))
	for i, v := range values {
		result[i] = T(v)
	}
	return NewSet(result...)
}

// ParseSetFromString parses a comma-separated string into a Set with string-based type.
// T must be string or a type based on string (using ~string constraint).
func ParseSetFromString[T ~string](value string) Set[T] {
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == ','
	})
	// Trim whitespace from each part
	trimmed := make([]string, 0, len(parts))
	for _, part := range parts {
		if t := strings.TrimSpace(part); t != "" {
			trimmed = append(trimmed, t)
		}
	}
	return ParseSetFromStrings[T](trimmed)
}

// MarshalText implements encoding.TextMarshaler for Set.
func (s *set[S]) MarshalText() ([]byte, error) {
	return []byte(strings.Join(s.Strings(), ",")), nil
}

// UnmarshalText implements encoding.TextUnmarshaler for set with string element type.
// This allows Set[string] to be automatically parsed from comma-separated strings
// when used with github.com/bborbe/argument.
func (s *set[S]) UnmarshalText(text []byte) error {
	value := string(text)
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == ','
	})

	// Clear existing data
	s.mux.Lock()
	defer s.mux.Unlock()
	s.data = make(map[S]struct{})

	// Add trimmed parts
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			// Convert string to S type using unsafe pointer conversion
			// This works for any type S where the underlying type is string
			// The conversion is safe because both string and ~string types have identical memory layout
			element := *(*S)(unsafe.Pointer(&trimmed)) //#nosec G103 -- Safe conversion between string-based types
			s.data[element] = struct{}{}
		}
	}

	return nil
}

// MarshalJSON implements json.Marshaler for Set.
// It serializes the set as a JSON array of elements, supporting primitives,
// complex types, maps, and objects.
func (s *set[T]) MarshalJSON() ([]byte, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	elements := make([]T, 0, len(s.data))
	for k := range s.data {
		elements = append(elements, k)
	}
	return json.Marshal(elements)
}

// UnmarshalJSON implements json.Unmarshaler for Set.
// It deserializes a JSON array into set elements, supporting primitives,
// complex types, maps, and objects.
func (s *set[T]) UnmarshalJSON(data []byte) error {
	var elements []T
	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	s.data = make(map[T]struct{})

	for _, element := range elements {
		s.data[element] = struct{}{}
	}

	return nil
}
