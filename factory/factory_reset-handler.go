// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"context"
	"net/http"

	libhttp "github.com/bborbe/http"
	libkv "github.com/bborbe/kv"
)

func CreateResetHandler(db libkv.DB, cancel context.CancelFunc) http.Handler {
	return libhttp.NewDangerousHandlerWrapper(libkv.NewResetHandler(db, cancel))
}

func CreateResetBucketHandler(db libkv.DB, cancel context.CancelFunc) http.Handler {
	return libhttp.NewDangerousHandlerWrapper(libkv.NewResetBucketHandler(db, cancel))
}
