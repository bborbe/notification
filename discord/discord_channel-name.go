// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discord

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	"github.com/bwmarrin/discordgo"
)

type ChannelNames []ChannelName

func (c ChannelNames) Contains(channelName ChannelName) bool {
	for _, cc := range c {
		if cc == channelName {
			return true
		}
	}
	return false
}

type ChannelName string

func (c ChannelName) Validate(ctx context.Context) error {
	if len(c) < 2 {
		return errors.Wrapf(ctx, validation.Error, "channel name to short")
	}
	if len(c) > 100 {
		return errors.Wrapf(ctx, validation.Error, "channel name to long")
	}
	return nil
}

func (c ChannelName) String() string {
	return string(c)
}

func (c ChannelName) Bytes() []byte {
	return []byte(c)
}

func ChannelNamesFromChannel(channel *discordgo.Channel) ChannelName {
	return ChannelName(channel.Name)
}

func ChannelNamesFromChannels(channels []*discordgo.Channel) ChannelNames {
	result := ChannelNames{}
	for _, channel := range channels {
		result = append(result, ChannelNamesFromChannel(channel))
	}
	return result
}
