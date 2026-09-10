// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"net/http"
	"os"

	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	"github.com/blevesearch/bleve/v2"
	"github.com/golang/glog"
)

func NewResetIndexHandler(index bleve.Index, cancel context.CancelFunc) http.Handler {
	return libhttp.NewErrorHandler(
		libhttp.WithErrorFunc(
			func(ctx context.Context, resp http.ResponseWriter, req *http.Request) error {
				indexPath := index.Name()
				if err := os.RemoveAll(indexPath); err != nil {
					return errors.Wrapf(ctx, err, "remove index %s failed", indexPath)
				}
				glog.V(2).Infof("remove %s completed", indexPath)
				cancel()
				resp.WriteHeader(http.StatusOK)
				_, _ = libhttp.WriteAndGlog(resp, "removed bleveIndex %s successful", indexPath)
				return nil
			},
		),
	)
}
