// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

const DefaultSearchLimit = 10000

//counterfeiter:generate -o ../mocks/index.go --fake-name Index . Index
type Index interface {
	Index(id string, data interface{}) error
	Delete(id string) error
}
