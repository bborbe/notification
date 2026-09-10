// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	stderrors "errors"
)

// ErrResultChannelClosed is returned when a Kubernetes watch result channel is closed.
// This typically occurs when the watch connection times out or is terminated by the server.
// It is a normal occurrence in long-running watches and should trigger a reconnection.
var ErrResultChannelClosed = stderrors.New("watch result channel closed")

// ErrUnknownEventType is returned when a Kubernetes watch receives an event type
// that is not recognized or handled by the watcher implementation.
var ErrUnknownEventType = stderrors.New("unknown event type")
