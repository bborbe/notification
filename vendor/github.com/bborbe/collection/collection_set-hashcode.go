// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"context"
	"encoding/json"
	"sort"
	"sync"
)

// HasHashCode represents types that can provide a string hash code for themselves.
//
// Security: HashCode() MUST return unique values for distinct elements.
// Hash collisions will cause silent element overwrites in SetHashCode.
// Recommended: Use fmt.Sprintf("%#v", obj) for comprehensive hashing.
type HasHashCode interface {
	HashCode() string
}

// SetHashCode represents a thread-safe set for types that implement HasHashCode.
// Elements are uniquely identified by their hash code.
//
// Performance: This implementation uses a map-based approach with O(1) average-case
// operations for Add, Remove, and Contains. This provides better performance than
// SetEqual for large sets or performance-critical code.
type SetHashCode[T HasHashCode] interface {
	// Add inserts elements into the set, using their hash codes for uniqueness.
	// Duplicate elements (same hash code) are automatically ignored.
	// Multiple elements can be added in a single call with only one mutex lock.
	Add(elements ...T)
	// Remove deletes elements from the set by their hash codes.
	// Multiple elements can be removed in a single call with only one mutex lock.
	Remove(elements ...T)
	// Contains reports whether an element with the given hash code is present in the set.
	Contains(element T) bool
	// ContainsAll reports whether all given elements are present in the set by their hash codes.
	ContainsAll(elements ...T) bool
	// ContainsAny reports whether at least one of the given elements is present in the set by their hash codes.
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
	// Clone returns a new SetHashCode containing all elements from the current set.
	// The returned set is a shallow copy - modifications to it won't affect the original.
	Clone() SetHashCode[T]
	// Without returns a new SetHashCode containing all elements from the current set
	// except those specified in the elements parameter.
	// The original set is not modified.
	Without(elements ...T) SetHashCode[T]
	// UnmarshalJSON deserializes a JSON array into set elements.
	// It implements json.Unmarshaler for automatic JSON parsing.
	UnmarshalJSON(data []byte) error
	// MarshalJSON serializes set elements to a JSON array.
	// It implements json.Marshaler for automatic JSON serialization.
	MarshalJSON() ([]byte, error)
}

// NewSetHashCode creates a new thread-safe set for types that implement HasHashCode.
// It accepts optional initial elements to populate the set.
// Duplicate elements (same hash code) are automatically handled.
//
// Performance: This implementation uses a map-based approach with O(1) average-case
// operations. Initialization is O(n) for n elements, making it suitable for large sets.
//
// Example:
//
//	type User struct { ID int; Name string }
//	func (u User) HashCode() string { return fmt.Sprintf("user-%d", u.ID) }
//	set := collection.NewSetHashCode(User{1, "Alice"}, User{2, "Bob"})
func NewSetHashCode[T HasHashCode](elements ...T) SetHashCode[T] {
	s := &setHashCode[T]{
		data: make(map[string]T),
	}
	s.Add(elements...)
	return s
}

type setHashCode[T HasHashCode] struct {
	mux  sync.Mutex
	data map[string]T
}

func (s *setHashCode[T]) Add(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		s.data[element.HashCode()] = element
	}
}

func (s *setHashCode[T]) Remove(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		delete(s.data, element.HashCode())
	}
}

func (s *setHashCode[T]) Contains(element T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, found := s.data[element.HashCode()]
	return found
}

func (s *setHashCode[T]) ContainsAll(elements ...T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if _, found := s.data[element.HashCode()]; !found {
			return false
		}
	}
	return true
}

func (s *setHashCode[T]) ContainsAny(elements ...T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if _, found := s.data[element.HashCode()]; found {
			return true
		}
	}
	return false
}

func (s *setHashCode[T]) Slice() []T {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]T, 0, len(s.data))
	for _, v := range s.data {
		result = append(result, v)
	}
	return result
}

func (s *setHashCode[T]) Length() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return len(s.data)
}

// Strings returns all elements as their string representations in sorted order.
// This provides deterministic output suitable for debugging and logging.
func (s *setHashCode[T]) Strings() []string {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]string, 0, len(s.data))
	for _, v := range s.data {
		result = append(result, elementToString(v))
	}

	sort.Strings(result)
	return result
}

// String returns a human-readable string representation of the set.
// Format: "SetHashCode[element1, element2, ...]" for non-empty sets, "SetHashCode[]" for empty sets.
// Elements are sorted by their string representation for deterministic output.
func (s *setHashCode[T]) String() string {
	return formatSetString("SetHashCode[", s.Strings())
}

// Each calls fn for each element in the set. Iteration stops on first error.
// The order of iteration is arbitrary and not guaranteed to be consistent.
func (s *setHashCode[T]) Each(
	ctx context.Context,
	fn func(ctx context.Context, value T) error,
) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range s.data {
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

// Clone returns a new SetHashCode containing all elements from the current set.
// The returned set is a shallow copy - modifications to it won't affect the original.
func (s *setHashCode[T]) Clone() SetHashCode[T] {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := &setHashCode[T]{
		data: make(map[string]T, len(s.data)),
	}

	for k, v := range s.data {
		result.data[k] = v
	}

	return result
}

// Without returns a new SetHashCode containing all elements from the current set
// except those specified in the elements parameter.
// The original set is not modified.
func (s *setHashCode[T]) Without(elements ...T) SetHashCode[T] {
	result := s.Clone()
	result.Remove(elements...)
	return result
}

// MarshalJSON implements json.Marshaler for SetHashCode.
// It serializes the set as a JSON array of elements in arbitrary order.
func (s *setHashCode[T]) MarshalJSON() ([]byte, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	elements := make([]T, 0, len(s.data))
	for _, v := range s.data {
		elements = append(elements, v)
	}
	return json.Marshal(elements)
}

// UnmarshalJSON implements json.Unmarshaler for SetHashCode.
// It deserializes a JSON array into set elements using hash codes for uniqueness.
func (s *setHashCode[T]) UnmarshalJSON(data []byte) error {
	var elements []T
	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	s.data = make(map[string]T)

	for _, element := range elements {
		s.data[element.HashCode()] = element
	}

	return nil
}
