// Copyright (c) 2024 Benjamin Borbe All rights reserved.
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

const NotificationPublishCommandOperation base.CommandOperation = "notification-publish"

//counterfeiter:generate -o ../../mocks/command-notification-publish-command-sender.go --fake-name CommandNotificationPublishCommandSender . NotificationPublishCommandSender
type NotificationPublishCommandSender interface {
	SendPublishNotificationCommand(
		ctx context.Context,
		command NotificationPublishCommand,
	) error
}

type NotificationPublishCommandSenderFunc func(ctx context.Context, command NotificationPublishCommand) error

func (n NotificationPublishCommandSenderFunc) SendPublishNotificationCommand(
	ctx context.Context,
	command NotificationPublishCommand,
) error {
	return n(ctx, command)
}

func NewNotificationPublishCommandSender(
	commandCreator base.CommandCreator,
	commandObjectSender cdb.CommandObjectSender,
	initiator cqrsiam.Initiator,
) NotificationPublishCommandSender {
	return NotificationPublishCommandSenderFunc(
		func(ctx context.Context, command NotificationPublishCommand) error {
			event, err := base.ParseEvent(ctx, command)
			if err != nil {
				return errors.Wrapf(ctx, err, "parse command failed")
			}
			commandObject := cdb.CommandObject{
				Command: commandCreator.NewCommand(
					NotificationPublishCommandOperation,
					initiator,
					"",
					event,
				),
				SchemaID: core.NotificationV1SchemaID,
			}
			if err := commandObjectSender.SendCommandObject(ctx, commandObject); err != nil {
				return errors.Wrapf(ctx, err, "send notify command failed")
			}
			return nil
		},
	)
}
