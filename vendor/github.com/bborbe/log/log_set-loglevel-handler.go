// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// NewSetLoglevelHandler creates an HTTP handler for dynamically changing log levels via REST API.
// The handler expects a URL path variable named "level" containing the desired log level.
//
// Usage with gorilla/mux:
//
//	router := mux.NewRouter()
//	logLevelSetter := log.NewLogLevelSetter(glog.Level(1), 5*time.Minute)
//	router.Handle("/debug/loglevel/{level}", log.NewSetLoglevelHandler(ctx, logLevelSetter))
//
// Example HTTP requests:
//
//	GET  /debug/loglevel/4  - Set log level to 4
//	POST /debug/loglevel/2  - Set log level to 2
//
// Parameters:
//   - ctx: Context for the log level setter operations
//   - logLevelSetter: The LogLevelSetter implementation to use for changing levels
//
// Returns an http.Handler that can be registered with any HTTP router.
func NewSetLoglevelHandler(ctx context.Context, logLevelSetter LogLevelSetter) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		level, err := strconv.ParseInt(vars["level"], 10, 32)
		if err != nil {
			fmt.Fprintf(resp, "parse loglevel failed: %v\n", err)
			return
		}
		if err := logLevelSetter.Set(ctx, glog.Level(int32(level))); err != nil {
			fmt.Fprintf(resp, "set loglevel failed: %v\n", err)
			return
		}
		fmt.Fprintf(
			resp,
			"set loglevel to %d completed\n",
			level,
		) // #nosec G705 - level is an integer, not user-controlled string
	})
}
