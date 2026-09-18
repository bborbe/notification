// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram

import (
	"context"
)

// Bot names which Telegram bot delivers a message. A chat id does not identify
// a bot — for a private chat it is the recipient's own user id, so every bot
// messaging the same person uses the same chat id. Bot is therefore the only
// discriminator available when more than one bot serves one recipient.
//
// The empty Bot is the default bot, which keeps commands written before this
// field existed routing to the bot that has always handled them. Consumers
// filter on this value, so each running instance delivers only the commands
// naming it and skips the rest.
type Bot string

func (b Bot) String() string {
	return string(b)
}

func (b Bot) Ptr() *Bot {
	return &b
}

// Validate accepts every value, including the empty one. Bot is optional by
// design: an absent field means the default bot, and requiring a value here
// would reject every command produced before the field was introduced.
func (b Bot) Validate(ctx context.Context) error {
	return nil
}
