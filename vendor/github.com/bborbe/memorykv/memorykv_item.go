// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	"context"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

func NewItem(
	key []byte,
	valueGetter ValueGetter,
) libkv.Item {
	return &item{
		key:         key,
		valueGetter: valueGetter,
	}
}

type item struct {
	valueGetter ValueGetter
	key         []byte
}

func (i *item) Exists() bool {
	return true
}

func (i *item) Key() []byte {
	return i.key
}

func (i *item) Value(fn func(val []byte) error) error {
	// TODO: add context to Value
	ctx := context.Background()

	item, err := i.valueGetter.Get(ctx, i.key)
	if err != nil {
		return errors.Wrapf(ctx, err, "get value failed")
	}
	return item.Value(fn)
}
