// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"context"

	"github.com/bborbe/collection"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
	"github.com/blevesearch/bleve/v2"
	"github.com/golang/glog"
)

type ReindexStreamerTx[T any] interface {
	Stream(ctx context.Context, tx libkv.Tx, ch chan<- T) error
}

type IndexerTx[T any] interface {
	Add(ctx context.Context, tx libkv.Tx, object T) error
}

type ReindexBatchIndexProviderTx[T any] interface {
	CreateIndexer(batch *bleve.Batch) IndexerTx[T]
}

type ReindexBatchIndexProviderFuncTx[T any] func(batch *bleve.Batch) IndexerTx[T]

func (r ReindexBatchIndexProviderFuncTx[T]) CreateIndexer(batch *bleve.Batch) IndexerTx[T] {
	return r(batch)
}

// NewReindexer creates a Tx-based reindexer that operates within a transaction.
// The caller is responsible for wrapping this in db.View() to provide the transaction.
func NewReindexer[T any](
	reindexBatchIndexProvider ReindexBatchIndexProviderTx[T],
	streamer ReindexStreamerTx[T],
	bleveIndex bleve.Index,
	batchSize uint64,
) libkv.RunnableTx {
	return libkv.FuncTx(func(ctx context.Context, tx libkv.Tx) error {
		batch := bleveIndex.NewBatch()
		indexer := reindexBatchIndexProvider.CreateIndexer(batch)

		var counter uint64
		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- T) error {
				return streamer.Stream(ctx, tx, ch)
			},
			func(ctx context.Context, actualTrade T) error {
				if err := indexer.Add(ctx, tx, actualTrade); err != nil {
					return errors.Wrapf(ctx, err, "add actualTrade to index failed")
				}
				counter++
				if counter%batchSize == 0 {
					if err := bleveIndex.Batch(batch); err != nil {
						return errors.Wrapf(ctx, err, "execute batch failed")
					}
					glog.V(2).Infof("reindex %d completed", counter)
					batch = bleveIndex.NewBatch()
					indexer = reindexBatchIndexProvider.CreateIndexer(batch)
				}
				return nil
			},
		)
		if err != nil {
			return errors.Wrapf(ctx, err, "index failed")
		}

		if counter%batchSize != 0 {
			glog.V(2).Infof("index remaining trades")
			if err := bleveIndex.Batch(batch); err != nil {
				return errors.Wrapf(ctx, err, "execute batch failed")
			}
			glog.V(2).Infof("reindex %d completed", counter)
		}
		glog.V(2).Infof("reindex all completed")
		return nil
	})
}
