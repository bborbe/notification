// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discord

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/validation"

	"github.com/bborbe/notification/discord"
)

const SendCommandOperation base.CommandOperation = "send"

type SendCommand struct {
	ChannelName discord.ChannelName `json:"channelName"`
	Message     discord.Message     `json:"message"`
}

func (s SendCommand) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ChannelName", s.ChannelName),
		validation.Name("Message", s.Message),
	}.Validate(ctx)
}
