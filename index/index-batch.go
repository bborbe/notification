// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import "github.com/blevesearch/bleve/v2"

func NewBatchIndex(batch *bleve.Batch) Index {
	return &indexBatch{batch: batch}
}

type indexBatch struct {
	batch *bleve.Batch
}

func (i *indexBatch) Index(id string, data interface{}) error {
	return i.batch.Index(id, data)
}

func (i *indexBatch) Delete(id string) error {
	i.batch.Delete(id)
	return nil
}
