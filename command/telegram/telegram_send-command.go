// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/validation"

	"github.com/bborbe/notification/telegram"
)

const SendCommandOperation base.CommandOperation = "send"

type SendCommand struct {
	ChatID  telegram.ChatID  `json:"chatId"`
	Message telegram.Message `json:"message"`
}

func (s SendCommand) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ChatID", s.ChatID),
		validation.Name("Message", s.Message),
	}.Validate(ctx)
}
