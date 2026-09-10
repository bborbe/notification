// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"net/http"

	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkv "github.com/bborbe/kv"
	"github.com/blevesearch/bleve/v2"
)

func NewStatusHandler(
	bleveIndex bleve.Index,
	bucketName libkv.BucketName,
) libhttp.JSONHandlerTx {
	return libhttp.JSONHandlerTxFunc(
		func(ctx context.Context, tx libkv.Tx, req *http.Request) (interface{}, error) {
			indexCount, err := bleveIndex.DocCount()
			if err != nil {
				return nil, errors.Wrapf(ctx, err, "get index doc count failed")
			}
			var storeCount int64
			bucket, err := tx.Bucket(ctx, bucketName)
			if err != nil {
				if !errors.Is(err, libkv.BucketNotFoundError) {
					return nil, errors.Wrapf(ctx, err, "get bucket failed")
				}
			} else {
				storeCount, err = libkv.Count(ctx, bucket)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "count bucket entries failed")
				}
			}
			return map[string]interface{}{
				"index": indexCount,
				"store": storeCount,
			}, nil
		},
	)
}
