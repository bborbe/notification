// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	libkv "github.com/bborbe/kv"
	"golang.org/x/exp/slices"
)

func NewIterator(
	keys []string,
	valueGetter ValueGetter,
) libkv.Iterator {
	return &iterator{
		keys:        keys,
		pos:         0,
		valueGetter: valueGetter,
	}
}

type iterator struct {
	pos         int
	keys        []string
	valueGetter ValueGetter
}

func (i *iterator) Close() {
}

func (i *iterator) Item() libkv.Item {
	return NewItem([]byte(i.keys[i.pos]), i.valueGetter)
}

func (i *iterator) Next() {
	i.pos++
}

func (i *iterator) Valid() bool {
	return i.pos >= 0 && i.pos < len(i.keys)
}

func (i *iterator) Rewind() {
	i.pos = 0
}

func (i *iterator) Seek(key []byte) {
	pos, _ := slices.BinarySearch(i.keys, string(key))
	i.pos = pos
}
