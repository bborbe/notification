// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package collection provides generic, type-safe collection utilities and data structures for Go.
//
// This library leverages Go generics to offer a comprehensive set of operations for working with
// slices, sets, channels, and other collection types. It includes functional-style operations,
// concurrent-safe data structures, and utility functions for common collection manipulations.
//
// Key Features:
//
// - Slice Operations: Filter, Map, Find, Contains, Unique, Reverse, Copy, Exclude, Join
// - Set Data Structures: Thread-safe generic sets with Equal and HashCode support
// - Channel Processing: Concurrent operations with context support
// - Pointer Utilities: Generic Ptr/Unptr functions for value-pointer conversion
// - Comparison Utilities: Generic equality and string comparison helpers
//
// The package follows functional programming patterns with higher-order functions while
// maintaining thread safety for concurrent operations through proper mutex synchronization.
package collection
