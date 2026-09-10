// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"runtime"

	"github.com/golang/glog"
)

// MemoryStats is a helper that reads memory stats and logs them with the given prefix
func MemoryStats(prefix string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	glog.Infof(
		"MEMORY USAGE - %s - Alloc: %.2f MB, Sys: %.2f MB, HeapInUse: %.2f MB, HeapObjects: %d, NumGC: %d",
		prefix,
		float64(memStats.Alloc)/1024/1024,
		float64(memStats.Sys)/1024/1024,
		float64(memStats.HeapInuse)/1024/1024,
		memStats.HeapObjects,
		memStats.NumGC,
	)
}
