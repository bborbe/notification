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
	// Bot names the bot that delivers this message. Empty means the default
	// bot, so commands produced before this field existed keep reaching the
	// bot that has always handled them. Consumers filter on it, which is what
	// lets several bots share one chat id — see telegram.Bot.
	Bot telegram.Bot `json:"bot,omitempty"`
}

func (s SendCommand) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ChatID", s.ChatID),
		validation.Name("Message", s.Message),
		validation.Name("Bot", s.Bot),
	}.Validate(ctx)
}
