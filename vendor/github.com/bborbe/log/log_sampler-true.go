// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

// NewSamplerTrue creates a sampler that always returns true.
// This is useful for testing, debugging, or situations where you want to
// temporarily disable sampling and log everything.
//
// Example:
//
//	// For debugging, log everything regardless of other sampling rules
//	debugSampler := log.NewSamplerTrue()
//	if debugSampler.IsSample() {
//	    glog.V(4).Infof("Debug message - always logged")
//	}
//
//	// Use in tests to ensure all log paths are exercised
//	testSampler := log.NewSamplerTrue()
//
// This sampler has no internal state and is inherently thread-safe.
func NewSamplerTrue() Sampler {
	return SamplerFunc(func() bool {
		return true
	})
}
