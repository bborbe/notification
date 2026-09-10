// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "github.com/bborbe/cqrs/cdb"

var DiscordV1SchemaID = cdb.SchemaID{
	Group:   "core",
	Kind:    "discord",
	Version: "v1",
}

var NotificationV1SchemaID = cdb.SchemaID{
	Group:   "core",
	Kind:    "notification",
	Version: "v1",
}
