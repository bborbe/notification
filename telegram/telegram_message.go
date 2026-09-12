// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram

import (
	"context"

	"github.com/bborbe/validation"
)

type Message string

func (m Message) String() string {
	return string(m)
}

func (m Message) Ptr() *Message {
	return &m
}

func (m Message) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("Message", validation.NotEmptyString(m)),
	}.Validate(ctx)
}
