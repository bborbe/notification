// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notification

import (
	"context"

	"github.com/bborbe/validation"

	"github.com/bborbe/notification"
)

type NotificationPublishCommand struct {
	// Type of the notification
	Type core.NotificationType `json:"type"`

	// Target of the notification (discord channelName, telegram room ...)
	Target *core.NotificationTarget `json:"target,omitempty"`

	// Message of the notification
	Message core.NotificationMessage `json:"message"`

	// Metadata carries source-specific fields as opaque key/value pairs
	// (e.g. trading producers pack brokerIdentifier/stage/accountIdentifier here).
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (n NotificationPublishCommand) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("Type", n.Type),
		validation.Name("Target", validation.NilOrValid(n.Target)),
		validation.Name("Message", n.Message),
	}.Validate(ctx)
}

func (n NotificationPublishCommand) Ptr() *NotificationPublishCommand {
	return &n
}
