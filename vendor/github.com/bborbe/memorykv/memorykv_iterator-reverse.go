// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	libkv "github.com/bborbe/kv"
	"golang.org/x/exp/slices"
)

func NewIteratorReverse(
	keys []string,
	valueGetter ValueGetter,
) libkv.Iterator {
	return &iteratorReverse{
		keys:        keys,
		pos:         0,
		valueGetter: valueGetter,
	}
}

type iteratorReverse struct {
	pos         int
	keys        []string
	valueGetter ValueGetter
}

func (i *iteratorReverse) Close() {
}

func (i *iteratorReverse) Item() libkv.Item {
	return NewItem([]byte(i.keys[i.pos]), i.valueGetter)
}

func (i *iteratorReverse) Next() {
	i.pos--
}

func (i *iteratorReverse) Valid() bool {
	return i.pos >= 0 && i.pos < len(i.keys)
}

func (i *iteratorReverse) Rewind() {
	i.pos = len(i.keys) - 1
}

func (i *iteratorReverse) Seek(key []byte) {
	pos, found := slices.BinarySearch(i.keys, string(key))
	if found {
		i.pos = pos
	} else {
		i.pos = pos - 1
	}
}
