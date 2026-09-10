// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import (
	"bytes"
	"encoding/json"
	"strings"
)

// BucketNames represents a collection of bucket names with utility methods.
type BucketNames []BucketName

// Contains checks if the collection contains the specified bucket name.
func (t BucketNames) Contains(value BucketName) bool {
	for _, tt := range t {
		if tt.Equal(value) {
			return true
		}
	}
	return false
}

// BucketFromStrings creates a bucket name by joining multiple strings with underscores.
func BucketFromStrings(values ...string) BucketName {

	return NewBucketName(strings.Join(values, "_"))
}

// NewBucketName creates a new bucket name from a string.
func NewBucketName(name string) BucketName {
	return BucketName(name)
}

// BucketName represents a bucket identifier as a byte slice.
type BucketName []byte

// String returns the bucket name as a string.
func (b BucketName) String() string {
	return string(b)
}

// Bytes returns the bucket name as a byte slice.
func (b BucketName) Bytes() []byte {
	return b
}

// Equal compares two bucket names for equality.
func (b BucketName) Equal(value BucketName) bool {
	return bytes.Equal(b, value)
}

// MarshalJSON encodes the bucket name as a plain JSON string rather than the
// default base64 representation Go uses for []byte values.
func (b BucketName) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(b))
}

// UnmarshalJSON decodes a JSON string back into a bucket name.
func (b *BucketName) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*b = BucketName(s)
	return nil
}
