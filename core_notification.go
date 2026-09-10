// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"github.com/bborbe/cqrs/base"
	"github.com/bborbe/validation"
)

type Notification struct {
	base.Object[base.Identifier]

	// Type of the notification
	Type NotificationType `json:"type"`

	// Target of the notification (discord channelName, telegram room ...)
	Target *NotificationTarget `json:"target,omitempty"`

	// Message of the notification
	Message NotificationMessage `json:"message"`

	// Metadata carries source-specific fields as opaque key/value pairs
	// (e.g. trading producers pack brokerIdentifier/stage/accountIdentifier here).
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (n Notification) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("Object", n.Object),
		validation.Name("Type", n.Type),
		validation.Name("Target", validation.NilOrValid(n.Target)),
		validation.Name("Message", n.Message),
	}.Validate(ctx)
}
