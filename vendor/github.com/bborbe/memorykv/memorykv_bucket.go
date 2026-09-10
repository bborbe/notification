// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	"context"
	"sort"
	"sync"

	libkv "github.com/bborbe/kv"
)

type ValueGetter interface {
	Get(ctx context.Context, key []byte) (libkv.Item, error)
}

func NewBucket() libkv.Bucket {
	return &bucket{
		data: make(map[string][]byte),
	}
}

type bucket struct {
	data map[string][]byte
	mux  sync.Mutex
}

func (b *bucket) Iterator() libkv.Iterator {
	keys := b.keys()
	sort.Strings(keys)
	return NewIterator(keys, b)
}

func (b *bucket) IteratorReverse() libkv.Iterator {
	keys := b.keys()
	sort.Strings(keys)
	return NewIteratorReverse(keys, b)
}

func (b *bucket) Get(ctx context.Context, key []byte) (libkv.Item, error) {
	b.mux.Lock()
	defer b.mux.Unlock()
	value := b.data[string(key)]
	return libkv.NewByteItem(key, value), nil
}

func (b *bucket) Put(ctx context.Context, key []byte, value []byte) error {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.data[string(key)] = value
	return nil
}

func (b *bucket) Delete(ctx context.Context, key []byte) error {
	b.mux.Lock()
	defer b.mux.Unlock()
	delete(b.data, string(key))
	return nil
}

func (b *bucket) keys() []string {
	b.mux.Lock()
	defer b.mux.Unlock()
	result := make([]string, 0, len(b.data))
	for k := range b.data {
		result = append(result, k)
	}
	return result
}
