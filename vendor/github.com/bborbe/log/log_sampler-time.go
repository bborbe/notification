// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"sync"
	stdtime "time"

	libtime "github.com/bborbe/time"
)

// NewSampleTime creates a time-based sampler that limits log sampling to a maximum frequency.
// It ensures that IsSample() returns true at most once per the specified duration.
//
// Example:
//
//	sampler := log.NewSampleTime(5 * time.Second) // Sample at most once every 5 seconds
//	for {
//	    if sampler.IsSample() {
//	        glog.V(2).Infof("This message appears at most once every 5 seconds")
//	    }
//	    time.Sleep(100 * time.Millisecond)
//	}
//
// Parameters:
//   - duration: The minimum time interval between samples. Must be > 0.
//
// The sampler is thread-safe and can be used concurrently from multiple goroutines.
// It uses github.com/bborbe/time for consistent time handling across the library.
func NewSampleTime(duration stdtime.Duration) Sampler {
	var mux sync.Mutex
	var lastlog stdtime.Time
	return SamplerFunc(func() bool {
		mux.Lock()
		defer mux.Unlock()
		if libtime.Now().Sub(lastlog) <= duration {
			return false
		}
		lastlog = libtime.Now()
		return true
	})
}
