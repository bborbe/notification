// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/v2/mapping"
)

func NewNotificationDocumentMapping() *mapping.DocumentMapping {
	documentMapping := bleve.NewDocumentStaticMapping()

	typeMapping := bleve.NewTextFieldMapping()
	typeMapping.Analyzer = keyword.Name
	documentMapping.AddFieldMappingsAt("NotificationType", typeMapping)

	messageMapping := bleve.NewTextFieldMapping()
	documentMapping.AddFieldMappingsAt("Message", messageMapping)

	targetMapping := bleve.NewTextFieldMapping()
	targetMapping.Analyzer = keyword.Name
	documentMapping.AddFieldMappingsAt("Target", targetMapping)

	brokerIdentifierMapping := bleve.NewTextFieldMapping()
	brokerIdentifierMapping.Analyzer = keyword.Name
	documentMapping.AddFieldMappingsAt("BrokerIdentifier", brokerIdentifierMapping)

	accountIdentifierMapping := bleve.NewTextFieldMapping()
	accountIdentifierMapping.Analyzer = keyword.Name
	documentMapping.AddFieldMappingsAt("AccountIdentifier", accountIdentifierMapping)

	stageMapping := bleve.NewTextFieldMapping()
	stageMapping.Analyzer = keyword.Name
	documentMapping.AddFieldMappingsAt("Stage", stageMapping)

	createdMapping := bleve.NewDateTimeFieldMapping()
	documentMapping.AddFieldMappingsAt("Created", createdMapping)

	modifiedMapping := bleve.NewDateTimeFieldMapping()
	documentMapping.AddFieldMappingsAt("Modified", modifiedMapping)

	return documentMapping
}
