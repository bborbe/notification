// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

// SamplerList combines multiple samplers using OR logic.
// It returns true if ANY of the contained samplers returns true.
// This allows for complex sampling strategies by combining different sampling methods.
//
// Example:
//
//	// Sample if either time-based OR mod-based condition is met
//	sampler := log.SamplerList{
//	    log.NewSampleTime(30 * time.Second),  // At most once per 30 seconds
//	    log.NewSampleMod(100),                // Every 100th call
//	}
//
//	// Sample if high log level OR time interval OR random chance
//	complexSampler := log.SamplerList{
//	    log.NewSamplerGlogLevel(4),           // When glog level >= 4
//	    log.NewSampleTime(10 * time.Second),  // At most once per 10 seconds
//	    log.SamplerFunc(func() bool {         // 5% random chance
//	        return rand.Float64() < 0.05
//	    }),
//	}
//
// The samplers are evaluated in order and the first one to return true
// causes the entire list to return true (short-circuit evaluation).
type SamplerList []Sampler

// IsSample implements the Sampler interface using OR logic across all contained samplers.
// It returns true if any sampler in the list returns true.
func (s SamplerList) IsSample() bool {
	for _, sampler := range s {
		if sampler.IsSample() {
			return true
		}
	}
	return false
}
