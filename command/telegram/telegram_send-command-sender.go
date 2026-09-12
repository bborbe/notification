// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegram

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/cqrs/cdb"
	cqrsiam "github.com/bborbe/cqrs/iam"
	"github.com/bborbe/errors"
	"github.com/golang/glog"

	"github.com/bborbe/notification"
)

//counterfeiter:generate -o ../../mocks/command-telegram-send-command-object-sender.go --fake-name CommandTelegramSendCommandObjectSender . SendCommandObjectSender
type SendCommandObjectSender interface {
	SendCommand(ctx context.Context, sendCommand SendCommand) error
}

func NewSendCommandObjectSender(
	commandCreator base.CommandCreator,
	commandSender cdb.CommandObjectSender,
	initiator cqrsiam.Initiator,
) SendCommandObjectSender {
	return &sendCommandObjectSender{
		initiator:      initiator,
		commandSender:  commandSender,
		commandCreator: commandCreator,
	}
}

type sendCommandObjectSender struct {
	commandSender  cdb.CommandObjectSender
	commandCreator base.CommandCreator
	initiator      cqrsiam.Initiator
}

func (c *sendCommandObjectSender) SendCommand(
	ctx context.Context,
	sendCommand SendCommand,
) error {
	glog.V(4).Infof("send telegram command to cdb started")

	if err := sendCommand.Validate(ctx); err != nil {
		return errors.Wrap(ctx, err, "validate send command failed")
	}

	event, err := base.ParseEvent(ctx, sendCommand)
	if err != nil {
		return errors.Wrapf(ctx, err, "parse event failed")
	}

	commandObject := cdb.CommandObject{
		SchemaID: core.TelegramV1SchemaID,
		Command: c.commandCreator.NewCommand(
			SendCommandOperation,
			c.initiator,
			"",
			event,
		),
	}

	if err := c.commandSender.SendCommandObject(ctx, commandObject); err != nil {
		return errors.Wrapf(ctx, err, "send command failed")
	}

	glog.V(4).Infof("send telegram command to cdb completed")
	return nil
}
