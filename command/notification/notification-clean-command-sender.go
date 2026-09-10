// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/cqrs/cdb"
	cqrsiam "github.com/bborbe/cqrs/iam"
	"github.com/bborbe/errors"

	"github.com/bborbe/notification"
)

const NotificationCleanCommandOperation base.CommandOperation = "clean-notifications"

//counterfeiter:generate -o ../../mocks/command-notification-clean-command-sender.go --fake-name CommandNotificationCleanCommandSender . NotificationCleanCommandSender
type NotificationCleanCommandSender interface {
	SendCleanNotificationCommand(
		ctx context.Context,
		command NotificationCleanCommand,
	) error
}

type NotificationCleanCommandSenderFunc func(ctx context.Context, command NotificationCleanCommand) error

func (n NotificationCleanCommandSenderFunc) SendCleanNotificationCommand(
	ctx context.Context,
	command NotificationCleanCommand,
) error {
	return n(ctx, command)
}

func NewNotificationCleanCommandSender(
	commandCreator base.CommandCreator,
	commandObjectSender cdb.CommandObjectSender,
	initiator cqrsiam.Initiator,
) NotificationCleanCommandSender {
	return NotificationCleanCommandSenderFunc(
		func(ctx context.Context, command NotificationCleanCommand) error {
			event, err := base.ParseEvent(ctx, command)
			if err != nil {
				return errors.Wrapf(ctx, err, "parse command failed")
			}
			commandObject := cdb.CommandObject{
				Command: commandCreator.NewCommand(
					NotificationCleanCommandOperation,
					initiator,
					"",
					event,
				),
				SchemaID: core.NotificationV1SchemaID,
			}
			if err := commandObjectSender.SendCommandObject(ctx, commandObject); err != nil {
				return errors.Wrapf(ctx, err, "send clean command failed")
			}
			return nil
		},
	)
}
