// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	libtime "github.com/bborbe/time"
	"github.com/golang/glog"
)

//counterfeiter:generate -o mocks/memory-monitor.go --fake-name MemoryMonitor . MemoryMonitor

// MemoryMonitor provides memory usage monitoring functionality
type MemoryMonitor interface {
	LogMemoryUsage(name string)
	LogMemoryUsagef(format string, args ...interface{})
	LogMemoryUsageOnStart()
	LogMemoryUsageOnEnd()
}

// NewMemoryMonitor creates a new memory monitor that logs memory usage at specified intervals
func NewMemoryMonitor(logInterval time.Duration) MemoryMonitor {
	return &memoryMonitor{
		logInterval: logInterval,
		lastLogTime: time.Time{}, // zero time initially
	}
}

type memoryMonitor struct {
	logInterval time.Duration

	mutex       sync.Mutex
	lastLogTime time.Time
}

// LogMemoryUsage logs current memory usage if enough time has passed since last log (thread-safe)
func (m *memoryMonitor) LogMemoryUsage(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	now := libtime.Now()

	// Check if enough time has passed since last log
	if m.lastLogTime.IsZero() || now.Sub(m.lastLogTime) >= m.logInterval {
		// Update the last log time first
		m.lastLogTime = now

		// Then log the memory usage
		MemoryStats(name)
	}
}

// LogMemoryUsagef logs current memory usage with formatted message if enough time has passed since last log (thread-safe)
func (m *memoryMonitor) LogMemoryUsagef(format string, args ...interface{}) {
	name := fmt.Sprintf(format, args...)
	m.LogMemoryUsage(name)
}

// LogMemoryUsageOnStart logs memory usage at the beginning
func (m *memoryMonitor) LogMemoryUsageOnStart() {
	glog.Infof("MEMORY MONITOR - Started")
	MemoryStats("START")
}

// LogMemoryUsageOnEnd logs memory usage at the end
func (m *memoryMonitor) LogMemoryUsageOnEnd() {
	glog.Infof("MEMORY MONITOR - Completed")
	MemoryStats("END")

	// Force garbage collection and log again to see the difference
	runtime.GC()
	time.Sleep(100 * time.Millisecond) // Give GC a moment to complete
	MemoryStats("after GC")
}
