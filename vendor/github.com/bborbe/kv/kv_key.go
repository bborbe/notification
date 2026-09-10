// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

// Key represents a key in the key-value store as a byte slice.
type Key []byte

// String returns the key as a string.
func (f Key) String() string {
	return string(f)
}

// Bytes returns the key as a byte slice.
func (f Key) Bytes() []byte {
	return f
}
