// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"os"
	"path"

	"github.com/bborbe/errors"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/golang/glog"
)

func OpenTemp(ctx context.Context, indexMapping mapping.IndexMapping) (bleve.Index, error) {
	if err := indexMapping.Validate(); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate indexMapping failed")
	}
	return bleve.NewMemOnly(indexMapping)
}

func OpenDir(
	ctx context.Context,
	indexMapping mapping.IndexMapping,
	indexPath string,
) (bleve.Index, error) {
	return OpenPath(ctx, indexMapping, BuildIndexPath(indexPath))
}

func BuildIndexPath(indexPath string) string {
	return path.Join(indexPath, "index.bleve")
}

func OpenPath(
	ctx context.Context,
	indexMapping mapping.IndexMapping,
	indexPath string,
) (bleve.Index, error) {
	if err := indexMapping.Validate(); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate indexMapping failed")
	}
	glog.V(2).Infof("open index with path %s", indexPath)
	index, err := bleve.Open(indexPath)
	if err != nil {
		index, err = bleve.New(indexPath, indexMapping)
		if err != nil {
			return nil, errors.Wrap(ctx, err, "create index failed")
		}
		glog.V(2).Infof("new index created")
	}
	return index, nil
}

func RemoveIndex(ctx context.Context, indexPath string) error {
	if err := os.Remove(BuildIndexPath(indexPath)); err != nil {
		return errors.Wrapf(ctx, err, "remove index failed")
	}
	glog.V(2).Infof("remove index %s completed", indexPath)
	return nil
}
