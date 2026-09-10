// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"time"

	libtimetest "github.com/bborbe/time/test"
)

func ParseTime(value interface{}) time.Time {
	return libtimetest.ParseTime(value)
}
