// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

// Stats reports the size and bucket composition of a DB.
// When Detailed is false, only Backend, SizeB and Buckets[].Name are guaranteed
// to be populated; per-bucket KeyCount and SizeB are filled only by StatsDetailed.
type Stats struct {
	Backend  string        `json:"backend"`
	SizeB    int64         `json:"size_bytes,omitempty"`
	Buckets  []BucketStats `json:"buckets"`
	Detailed bool          `json:"detailed"`
}

// BucketStats reports per-bucket metrics. KeyCount and SizeB are populated only
// by DB.StatsDetailed; DB.Stats leaves them at zero (omitted from JSON).
type BucketStats struct {
	Name     BucketName `json:"name"`
	KeyCount int64      `json:"keys,omitempty"`
	SizeB    int64      `json:"size_bytes,omitempty"`
}
