// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram

import (
	"context"

	"github.com/bborbe/validation"
)

// ChatID identifies the Telegram chat a message is delivered to. Telegram ids
// are numeric and negative for groups and channels, so it is carried as a
// string — the Bot API accepts either form for chat_id.
type ChatID string

func (c ChatID) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ChatID", validation.NotEmptyString(c)),
	}.Validate(ctx)
}

func (c ChatID) String() string {
	return string(c)
}

func (c ChatID) Ptr() *ChatID {
	return &c
}
