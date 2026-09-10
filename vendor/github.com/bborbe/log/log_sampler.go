// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

//counterfeiter:generate -o mocks/log-sampler.go --fake-name LogSampler . Sampler

// Sampler defines an interface for log sampling decisions.
// It allows reducing log volume by selectively determining which log entries should be emitted.
//
// Example usage:
//
//	sampler := log.NewSampleMod(10)
//	if sampler.IsSample() {
//	    glog.V(2).Infof("This message is sampled")
//	}
type Sampler interface {
	// IsSample returns true if the current log entry should be emitted.
	// Implementations may use various strategies such as counters, time intervals,
	// or log levels to determine sampling behavior.
	IsSample() bool
}
