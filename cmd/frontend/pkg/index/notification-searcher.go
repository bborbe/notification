// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

//counterfeiter:generate -o ../../mocks/notification-index-searcher.go --fake-name NotificationIndexSearcher . NotificationIndexSearcher

import (
	"context"
	"strings"
	"time"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
	"github.com/blevesearch/bleve/v2"
	blevequery "github.com/blevesearch/bleve/v2/search/query"
	"github.com/golang/glog"

	"github.com/bborbe/notification"
	"github.com/bborbe/notification/index"
)

type NotificationIndexSearcher interface {
	Search(
		ctx context.Context,
		from *libtime.DateTime,
		until *libtime.DateTime,
		types core.NotificationTypes,
		targets []core.NotificationTarget,
		brokerIdentifiers []string,
		accountIdentifiers []string,
		stages []string,
		query string,
		limit int,
	) (base.Identifiers, error)
}

func NewNotificationIndexSearcher(
	index bleve.Index,
) NotificationIndexSearcher {
	return &notificationIndexSearcher{
		index: index,
	}
}

type notificationIndexSearcher struct {
	index bleve.Index
}

func (a *notificationIndexSearcher) Search(
	ctx context.Context,
	from *libtime.DateTime,
	until *libtime.DateTime,
	types core.NotificationTypes,
	targets []core.NotificationTarget,
	brokerIdentifiers []string,
	accountIdentifiers []string,
	stages []string,
	query string,
	limit int,
) (base.Identifiers, error) {
	glog.V(3).Info("search started")
	var queries []blevequery.Query

	if from != nil || until != nil {
		var startTime time.Time
		var endTime time.Time
		if from != nil {
			startTime = from.Time()
		}
		if until != nil {
			endTime = until.Time()
		}
		query := blevequery.NewDateRangeQuery(startTime, endTime)
		query.SetField("Created")
		queries = append(queries, query)
	}

	switch len(types) {
	case 0:
		glog.V(4).Infof("no types filter")
	case 1:
		matchQuery := blevequery.NewMatchQuery(types[0].String())
		matchQuery.SetField("NotificationType")
		queries = append(queries, matchQuery)
	default:
		typeStrings := make([]string, len(types))
		for i, t := range types {
			typeStrings[i] = t.String()
		}
		queries = append(queries, index.NewMatchAny("NotificationType", typeStrings))
	}

	switch len(targets) {
	case 0:
		glog.V(4).Infof("no targets filter")
	case 1:
		matchQuery := blevequery.NewMatchQuery(targets[0].String())
		matchQuery.SetField("Target")
		queries = append(queries, matchQuery)
	default:
		targetStrings := make([]string, len(targets))
		for i, target := range targets {
			targetStrings[i] = target.String()
		}
		queries = append(queries, index.NewMatchAny("Target", targetStrings))
	}

	switch len(brokerIdentifiers) {
	case 0:
		glog.V(4).Infof("no brokerIdentifiers filter")
	case 1:
		matchQuery := blevequery.NewMatchQuery(brokerIdentifiers[0])
		matchQuery.SetField("BrokerIdentifier")
		queries = append(queries, matchQuery)
	default:
		queries = append(
			queries,
			index.NewMatchAny("BrokerIdentifier", brokerIdentifiers),
		)
	}

	switch len(accountIdentifiers) {
	case 0:
		glog.V(4).Infof("no accountIdentifiers filter")
	case 1:
		matchQuery := blevequery.NewMatchQuery(accountIdentifiers[0])
		matchQuery.SetField("AccountIdentifier")
		queries = append(queries, matchQuery)
	default:
		queries = append(
			queries,
			index.NewMatchAny("AccountIdentifier", accountIdentifiers),
		)
	}

	switch len(stages) {
	case 0:
		glog.V(4).Infof("no stages filter")
	case 1:
		matchQuery := blevequery.NewMatchQuery(stages[0])
		matchQuery.SetField("Stage")
		queries = append(queries, matchQuery)
	default:
		queries = append(queries, index.NewMatchAny("Stage", stages))
	}

	if strings.TrimSpace(query) != "" {
		messageQuery := blevequery.NewMatchQuery(query)
		messageQuery.SetField("Message")
		queries = append(queries, messageQuery)
	}

	request := bleve.NewSearchRequestOptions(
		index.And(queries...),
		limit,
		0,
		false,
	)
	request.SortBy([]string{"-Created"})
	request.Fields = []string{}

	result, err := a.index.SearchInContext(ctx, request)
	if err != nil {
		return nil, errors.Wrap(ctx, err, "search failed")
	}

	var ids base.Identifiers
	for _, hit := range result.Hits {
		ids = append(ids, base.Identifier(hit.ID))
	}
	return ids, nil
}
