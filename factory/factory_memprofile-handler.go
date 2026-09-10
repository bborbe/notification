// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"net/http"

	libhttp "github.com/bborbe/http"
)

func CreateMemoryProfileHandler() http.Handler {
	return libhttp.NewErrorHandler(libhttp.NewMemoryProfileDownloadHandler())
}
