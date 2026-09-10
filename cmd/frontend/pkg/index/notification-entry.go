// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"time"

	"github.com/bborbe/notification"
)

const NotificationType = "notification"

type NotificationEntry struct {
	Created           time.Time
	Modified          time.Time
	NotificationType  core.NotificationType
	Message           core.NotificationMessage
	Target            *core.NotificationTarget
	BrokerIdentifier  string
	AccountIdentifier string
	Stage             string
}

func (a NotificationEntry) Type() string {
	return NotificationType
}
