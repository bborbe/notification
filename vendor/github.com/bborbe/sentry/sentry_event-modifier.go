// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentry

import "github.com/getsentry/sentry-go"

// EventModifier provides an interface for modifying Sentry events before they are sent.
// It wraps the sentry.EventModifier interface to allow custom event processing.
type EventModifier interface {
	sentry.EventModifier
}

var _ EventModifier = EventModifierFunc(
	func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
		return event
	},
)

// EventModifierFunc is a function type that implements the EventModifier interface.
// It allows using functions as event modifiers without creating a separate struct.
type EventModifierFunc func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event

// ApplyToEvent implements the EventModifier interface by calling the function.
func (e EventModifierFunc) ApplyToEvent(
	event *sentry.Event,
	hint *sentry.EventHint,
	client *sentry.Client,
) *sentry.Event {
	return e(event, hint, client)
}

var _ EventModifier = EventModifierList{}

// EventModifierList is a slice of EventModifiers that applies all modifiers in sequence.
// Each modifier receives the event modified by the previous modifier in the list.
type EventModifierList []EventModifier

// ApplyToEvent implements the EventModifier interface by applying all modifiers in the list sequentially.
// The event is passed through each modifier in order, with each modifier receiving the result of the previous one.
func (e EventModifierList) ApplyToEvent(
	event *sentry.Event,
	hint *sentry.EventHint,
	client *sentry.Client,
) *sentry.Event {
	for _, modifier := range e {
		event = modifier.ApplyToEvent(event, hint, client)
	}
	return event
}
