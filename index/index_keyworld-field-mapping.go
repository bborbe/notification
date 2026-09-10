// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/v2/mapping"
)

func NewKeyworldFieldMapping() *mapping.FieldMapping {
	strategyIdentifierMapping := bleve.NewTextFieldMapping()
	strategyIdentifierMapping.Analyzer = keyword.Name
	return strategyIdentifierMapping
}
