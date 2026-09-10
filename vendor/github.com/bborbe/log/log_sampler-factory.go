// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import "time"

// DefaultSamplerFactory provides a default sampler configuration that combines
// time-based sampling (every 10 seconds) with glog level sampling (level 4+).
var DefaultSamplerFactory SamplerFactory = SamplerFactoryFunc(func() Sampler {
	return SamplerList{
		NewSampleTime(10 * time.Second),
		NewSamplerGlogLevel(4),
	}
})

//counterfeiter:generate -o mocks/log-sampler-factory.go --fake-name LogSamplerFactory . SamplerFactory

// SamplerFactory provides a factory pattern for creating Sampler instances.
// This interface enables dependency injection and makes testing easier.
//
// For testing, use log.DefaultSamplerFactory instead of mocking this interface.
type SamplerFactory interface {
	// Sampler creates and returns a new Sampler instance.
	Sampler() Sampler
}

// SamplerFactoryFunc is a function type that implements the SamplerFactory interface.
// It allows regular functions to be used as sampler factories.
type SamplerFactoryFunc func() Sampler

// Sampler implements the SamplerFactory interface by calling the underlying function.
func (s SamplerFactoryFunc) Sampler() Sampler {
	return s()
}
