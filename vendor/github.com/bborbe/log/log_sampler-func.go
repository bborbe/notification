// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

// SamplerFunc is a function type that implements the Sampler interface.
// It allows regular functions to be used as samplers, enabling custom sampling logic.
//
// Example:
//
//	// Custom sampler that samples based on random chance
//	randomSampler := log.SamplerFunc(func() bool {
//	    return rand.Float64() < 0.1 // 10% chance
//	})
//
//	// Custom sampler with external state
//	counter := 0
//	customSampler := log.SamplerFunc(func() bool {
//	    counter++
//	    return counter%3 == 0 // Every 3rd call
//	})
type SamplerFunc func() bool

// IsSample implements the Sampler interface by calling the underlying function.
func (s SamplerFunc) IsSample() bool {
	return s()
}
