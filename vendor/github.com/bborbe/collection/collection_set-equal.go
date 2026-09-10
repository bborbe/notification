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

// HasEqual represents types that can compare themselves for equality with another value.
type HasEqual[V any] interface {
	Equal(value V) bool
}

// SetEqual represents a thread-safe set for types that implement HasEqual.
// Elements are uniquely identified by their Equal method.
//
// Performance: This implementation uses a slice-based approach. Operations have
// O(n) complexity where n is the number of elements. For large sets or performance-critical
// code, consider using SetHashCode which provides O(1) average-case operations.
type SetEqual[T HasEqual[T]] interface {
	// Add inserts elements into the set, using the Equal method for uniqueness checking.
	// Duplicate elements are automatically ignored.
	// Multiple elements can be added in a single call with only one mutex lock.
	Add(elements ...T)
	// Remove deletes elements from the set using its Equal method for matching.
	// Multiple elements can be removed in a single call with only one mutex lock.
	Remove(elements ...T)
	// Contains reports whether an element matching the given value is present in the set.
	Contains(element T) bool
	// ContainsAll reports whether all given elements are present in the set using the Equal method.
	ContainsAll(elements ...T) bool
	// ContainsAny reports whether at least one of the given elements is present in the set using the Equal method.
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
	// Elements are iterated in insertion order (FIFO).
	Each(ctx context.Context, fn func(ctx context.Context, value T) error) error
	// Clone returns a new SetEqual containing all elements from the current set.
	// The returned set is a shallow copy - modifications to it won't affect the original.
	Clone() SetEqual[T]
	// Without returns a new SetEqual containing all elements from the current set
	// except those specified in the elements parameter.
	// The original set is not modified.
	Without(elements ...T) SetEqual[T]
	// UnmarshalJSON deserializes a JSON array into set elements.
	// It implements json.Unmarshaler for automatic JSON parsing.
	UnmarshalJSON(data []byte) error
	// MarshalJSON serializes set elements to a JSON array.
	// It implements json.Marshaler for automatic JSON serialization.
	MarshalJSON() ([]byte, error)
}

// NewSetEqual creates a new thread-safe set for types that implement HasEqual.
// It accepts optional initial elements to populate the set.
// Duplicate elements are automatically handled using the Equal method.
//
// Performance: This implementation uses a slice-based approach with O(n) operations.
// Initialization with k elements has O(k²) complexity due to uniqueness checks.
// For better performance with large sets, use NewSetHashCode instead.
//
// Example:
//
//	type User struct { ID int; Name string }
//	func (u User) Equal(other User) bool { return u.ID == other.ID }
//	set := collection.NewSetEqual(User{1, "Alice"}, User{2, "Bob"})
func NewSetEqual[T HasEqual[T]](elements ...T) SetEqual[T] {
	s := &setEqual[T]{
		data: make([]T, 0),
	}
	s.Add(elements...)
	return s
}

type setEqual[T HasEqual[T]] struct {
	mux  sync.Mutex
	data []T
}

func (s *setEqual[T]) Add(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if s.contains(element) {
			continue
		}
		s.data = append(s.data, element)
	}
}

func (s *setEqual[T]) Remove(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]T, 0, len(s.data))
	for _, e := range s.data {
		shouldRemove := false
		for _, element := range elements {
			if e.Equal(element) {
				shouldRemove = true
				break
			}
		}
		if shouldRemove {
			continue
		}
		result = append(result, e)
	}
	s.data = result
}

func (s *setEqual[T]) Contains(element T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.contains(element)
}

func (s *setEqual[T]) ContainsAll(elements ...T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if !s.contains(element) {
			return false
		}
	}
	return true
}

func (s *setEqual[T]) ContainsAny(elements ...T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if s.contains(element) {
			return true
		}
	}
	return false
}

func (s *setEqual[T]) contains(element T) bool {
	for _, e := range s.data {
		if e.Equal(element) {
			return true
		}
	}
	return false
}

func (s *setEqual[T]) Slice() []T {
	s.mux.Lock()
	defer s.mux.Unlock()
	return Copy(s.data)
}

func (s *setEqual[T]) Length() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return len(s.data)
}

// Strings returns all elements as their string representations in sorted order.
// This provides deterministic output suitable for debugging and logging.
func (s *setEqual[T]) Strings() []string {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]string, 0, len(s.data))
	for _, e := range s.data {
		result = append(result, elementToString(e))
	}

	sort.Strings(result)
	return result
}

// String returns a human-readable string representation of the set.
// Format: "SetEqual[element1, element2, ...]" for non-empty sets, "SetEqual[]" for empty sets.
// Elements are sorted by their string representation for deterministic output.
func (s *setEqual[T]) String() string {
	return formatSetString("SetEqual[", s.Strings())
}

// Each calls fn for each element in the set. Iteration stops on first error.
// Elements are iterated in insertion order (FIFO).
func (s *setEqual[T]) Each(ctx context.Context, fn func(ctx context.Context, value T) error) error {
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

// Clone returns a new SetEqual containing all elements from the current set.
// The returned set is a shallow copy - modifications to it won't affect the original.
func (s *setEqual[T]) Clone() SetEqual[T] {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := &setEqual[T]{
		data: make([]T, len(s.data)),
	}
	copy(result.data, s.data)

	return result
}

// Without returns a new SetEqual containing all elements from the current set
// except those specified in the elements parameter.
// The original set is not modified.
func (s *setEqual[T]) Without(elements ...T) SetEqual[T] {
	result := s.Clone()
	result.Remove(elements...)
	return result
}

// MarshalJSON implements json.Marshaler for SetEqual.
// It serializes the set as a JSON array of elements, preserving insertion order.
func (s *setEqual[T]) MarshalJSON() ([]byte, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	return json.Marshal(s.data)
}

// UnmarshalJSON implements json.Unmarshaler for SetEqual.
// It deserializes a JSON array into set elements using the Equal method for uniqueness.
func (s *setEqual[T]) UnmarshalJSON(data []byte) error {
	var elements []T
	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	s.data = make([]T, 0, len(elements))

	for _, element := range elements {
		if !s.contains(element) {
			s.data = append(s.data, element)
		}
	}

	return nil
}
