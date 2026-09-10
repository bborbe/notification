// Copyright (c) 2024-2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentry

// ExcludeErrors is a collection of ExcludeError functions that can be used to filter
// out specific errors from being sent to Sentry.
type ExcludeErrors []ExcludeError

// IsExcluded checks if the given error matches any of the exclude conditions.
// It returns true if any ExcludeError function in the collection returns true for the error.
func (e ExcludeErrors) IsExcluded(err error) bool {
	for _, ee := range e {
		if ee(err) {
			return true
		}
	}
	return false
}

// ExcludeError is a function type that determines whether an error should be excluded
// from Sentry reporting. It returns true if the error should be excluded.
type ExcludeError func(err error) bool
