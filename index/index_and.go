// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"github.com/blevesearch/bleve/v2"
	blevequery "github.com/blevesearch/bleve/v2/search/query"
)

func And(queries ...blevequery.Query) blevequery.Query {
	switch len(queries) {
	case 0:
		return blevequery.NewMatchAllQuery()
	case 1:
		return queries[0]
	default:
		return bleve.NewConjunctionQuery(queries...)
	}
}
