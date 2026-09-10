// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"github.com/blevesearch/bleve/v2"
	blevequery "github.com/blevesearch/bleve/v2/search/query"
)

func NewMatchAny(field string, values []string) blevequery.Query {
	switch len(values) {
	case 1:
		query := blevequery.NewMatchQuery(values[0])
		query.SetField(field)
		return query
	default:
		var queries []blevequery.Query
		for _, value := range values {
			query := blevequery.NewMatchQuery(value)
			query.SetField(field)
			queries = append(queries, query)
		}
		return bleve.NewDisjunctionQuery(queries...)
	}
}
