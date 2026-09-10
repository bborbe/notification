// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import "sync"

// NewSampleMod creates a counter-based sampler that samples every Nth log entry.
// It maintains an internal counter that increments on each IsSample() call,
// returning true when the counter is divisible by the modulus value.
//
// Example:
//
//	sampler := log.NewSampleMod(10) // Sample every 10th log entry
//	for i := 0; i < 100; i++ {
//	    if sampler.IsSample() {
//	        glog.V(2).Infof("Log entry %d", i) // This will be called 10 times
//	    }
//	}
//
// Parameters:
//   - mod: The modulus value for sampling (must be > 0). Every mod-th call will return true.
//
// The sampler is thread-safe and can be used concurrently from multiple goroutines.
func NewSampleMod(mod uint64) Sampler {
	var counter uint64
	var mux sync.Mutex
	return SamplerFunc(func() bool {
		mux.Lock()
		defer mux.Unlock()
		counter++
		return counter%mod == 0
	})
}
