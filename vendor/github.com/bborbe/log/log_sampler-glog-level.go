// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import "github.com/golang/glog"

// NewSamplerGlogLevel creates a sampler that samples based on the current glog verbosity level.
// It returns true when the current glog verbosity is at or above the specified level.
// This allows for dynamic sampling based on the runtime log level configuration.
//
// Example:
//
//	// Sample when glog verbosity is 3 or higher
//	sampler := log.NewSamplerGlogLevel(3)
//	if sampler.IsSample() {
//	    glog.V(2).Infof("This logs when -v=3 or higher is set")
//	}
//
//	// Combine with other samplers for conditional high-verbosity logging
//	conditionalSampler := log.SamplerList{
//	    log.NewSamplerGlogLevel(4),           // Always sample at high verbosity
//	    log.NewSampleTime(30 * time.Second),  // Otherwise, sample every 30 seconds
//	}
//
// Parameters:
//   - level: The minimum glog verbosity level required for sampling
//
// This sampler has no internal state and queries glog's current verbosity setting
// on each call, making it inherently thread-safe and responsive to runtime changes.
func NewSamplerGlogLevel(level glog.Level) Sampler {
	return SamplerFunc(func() bool {
		return bool(glog.V(level))
	})
}
