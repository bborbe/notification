// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package log provides logging utilities focused on log sampling and dynamic log level management.
//
// This library extends Google's glog with sampling mechanisms to reduce log volume in production
// while maintaining debuggability through runtime log level changes.
//
// # Core Features
//
// - Log Sampling: Reduce log volume using time-based, counter-based, or level-based sampling
// - Dynamic Log Levels: Change log verbosity at runtime via HTTP endpoints
// - Multiple Sampler Strategies: ModSampler, TimeSampler, GlogLevelSampler, TrueSampler
// - Sampler Composition: Combine multiple samplers using OR logic with SamplerList
// - Memory Monitoring: Track memory usage with periodic logging
//
// # Quick Start
//
// Sample every 10th log entry:
//
//	sampler := log.NewSampleMod(10)
//	if sampler.IsSample() {
//	    glog.V(2).Infof("Sampled message")
//	}
//
// Combine samplers to log if high verbosity OR every 10 seconds:
//
//	sampler := log.SamplerList{
//	    log.NewSamplerGlogLevel(4),
//	    log.NewSampleTime(10 * time.Second),
//	}
//	if sampler.IsSample() {
//	    glog.V(2).Infof("This logs if verbosity >= 4 OR at most once per 10 seconds")
//	}
//
// # Dynamic Log Level Management
//
// Set up an HTTP endpoint to change log levels at runtime:
//
//	logLevelSetter := log.NewLogLevelSetter(glog.Level(1), 5*time.Minute)
//	router.Handle("/debug/loglevel/{level}", log.NewSetLoglevelHandler(ctx, logLevelSetter))
//
// Change log level via HTTP:
//
//	curl http://localhost:8080/debug/loglevel/4
//
// The log level will automatically reset after 5 minutes.
//
// # Sampler Types
//
// ModSampler - Sample every Nth occurrence:
//
//	sampler := log.NewSampleMod(100)  // Sample every 100th call
//
// TimeSampler - Sample at most once per time interval:
//
//	sampler := log.NewSampleTime(30 * time.Second)  // At most once per 30 seconds
//
// GlogLevelSampler - Sample based on glog verbosity:
//
//	sampler := log.NewSamplerGlogLevel(3)  // Sample if -v=3 or higher
//
// TrueSampler - Always sample (useful for debugging):
//
//	sampler := log.NewSamplerTrue()  // Always returns true
//
// SamplerList - Combine samplers with OR logic:
//
//	sampler := log.SamplerList{
//	    log.NewSampleMod(1000),
//	    log.NewSampleTime(5 * time.Minute),
//	}  // Sample every 1000th call OR at most once per 5 minutes
//
// # Memory Monitoring
//
// Monitor memory usage periodically:
//
//	monitor := log.NewMemoryMonitor()
//	monitor.LogMemoryUsage("startup")
//	// ... do work ...
//	monitor.LogMemoryUsage("after-processing")
//
// Or use the convenience function:
//
//	log.MemoryStats("checkpoint-1")
//
// # Thread Safety
//
// All sampler implementations are thread-safe and can be used concurrently
// from multiple goroutines without additional synchronization.
//
// # Testing
//
// Use TrueSampler for testing to ensure consistent behavior:
//
//	func TestYourCode(t *testing.T) {
//	    sampler := log.NewSamplerTrue()
//	    // Test code that always samples
//	}
//
// Or use SamplerFunc for controlled testing:
//
//	shouldSample := true
//	sampler := log.SamplerFunc(func() bool {
//	    return shouldSample
//	})
package log
