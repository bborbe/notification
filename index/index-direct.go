// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import "github.com/blevesearch/bleve/v2"

func NewIndex(bleveIndex bleve.Index) Index {
	return &indexDirect{index: bleveIndex}
}

type indexDirect struct {
	index bleve.Index
}

func (i *indexDirect) Index(id string, data interface{}) error {
	return i.index.Index(id, data)
}

func (i *indexDirect) Delete(id string) error {
	return i.index.Delete(id)
}
