// Copyright (c) 2023-2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sentry provides an enhanced wrapper around the Sentry Go SDK with additional
// functionality for error exclusion, automatic tag enrichment, and context data extraction.
//
// The main Client interface wraps the official Sentry Go SDK and adds:
//   - Automatic tag extraction from context and errors
//   - Configurable error filtering to reduce noise
//   - Enhanced integration with github.com/bborbe/errors for context data
//   - Proxy support for HTTP transport
//
// Example usage:
//
//	client, err := sentry.NewClient(ctx, sentry.ClientOptions{
//	    Dsn: "your-dsn-here",
//	    Tags: map[string]string{"service": "my-app"},
//	})
//	if err != nil {
//	    return err
//	}
//	defer client.Close()
//
//	// Capture exception with automatic tag enrichment
//	client.CaptureException(err, &sentry.EventHint{Context: ctx}, nil)
package sentry

import (
	"context"
	"fmt"
	"io"
	"maps"
	stdtime "time"

	"github.com/bborbe/errors"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
)

//counterfeiter:generate -o mocks/sentry-client.go --fake-name SentryClient . Client

// Client provides an enhanced interface for interacting with Sentry error tracking.
// It wraps the official Sentry Go SDK and adds automatic tag enrichment from context
// and errors, configurable error filtering, and enhanced integration with github.com/bborbe/errors.
type Client interface {
	CaptureMessage(
		message string,
		hint *sentry.EventHint,
		scope sentry.EventModifier,
	) *sentry.EventID
	CaptureException(
		exception error,
		hint *sentry.EventHint,
		scope sentry.EventModifier,
	) *sentry.EventID
	Flush(timeout stdtime.Duration) bool
	io.Closer
}

// NewClient creates a new Sentry client with enhanced functionality including automatic
// tag enrichment and error filtering. It accepts standard Sentry ClientOptions and optional
// ExcludeError functions to filter out specific errors from being sent to Sentry.
//
// WARNING: Do not pass sensitive information (passwords, API keys, PII, tokens) in hint.Data,
// context data, or error data as these will be sent to Sentry as tags and may be stored or
// logged externally.
func NewClient(
	ctx context.Context,
	clientOptions sentry.ClientOptions,
	excludeErrors ...ExcludeError,
) (Client, error) {
	newClient, err := sentry.NewClient(clientOptions)
	if err != nil {
		return nil, errors.Wrap(ctx, err, "create sentry client failed")
	}
	newClient.AddEventProcessor(enrichEventTags)
	return &client{
		client:        newClient,
		excludeErrors: excludeErrors,
	}, nil
}

func enrichEventTags(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	if event.Tags == nil {
		event.Tags = make(map[string]string)
	}
	addContextTags(event, hint)
	addErrorTags(event, hint)
	addHintDataTags(event, hint)
	return event
}

func addContextTags(event *sentry.Event, hint *sentry.EventHint) {
	if hint.Context == nil {
		return
	}
	for k, v := range errors.DataFromContext(hint.Context) {
		event.Tags[k] = fmt.Sprintf("%v", v)
	}
}

func addErrorTags(event *sentry.Event, hint *sentry.EventHint) {
	if hint.OriginalException == nil {
		return
	}
	for k, v := range errors.DataFromError(hint.OriginalException) {
		event.Tags[k] = fmt.Sprintf("%v", v)
	}
}

func addHintDataTags(event *sentry.Event, hint *sentry.EventHint) {
	switch data := hint.Data.(type) {
	case map[string]any:
		for k, v := range data {
			if v == nil {
				continue
			}
			event.Tags[k] = fmt.Sprintf("%v", v)
		}
	case map[string]string:
		maps.Copy(event.Tags, data)
	}
}

type client struct {
	client        *sentry.Client
	excludeErrors ExcludeErrors
}

func (c *client) Flush(timeout stdtime.Duration) bool {
	return c.client.Flush(timeout)
}

func (c *client) CaptureMessage(
	message string,
	hint *sentry.EventHint,
	scope sentry.EventModifier,
) *sentry.EventID {
	eventID := c.client.CaptureMessage(message, hint, scope)
	if eventID != nil {
		glog.V(2).Infof("capture sentry message with id %s", *eventID)
	} else {
		glog.V(2).Infof("capture sentry message failed: eventID is nil")
	}
	return eventID
}

func (c *client) CaptureException(
	err error,
	hint *sentry.EventHint,
	scope sentry.EventModifier,
) *sentry.EventID {
	if c.excludeErrors.IsExcluded(err) {
		glog.V(4).Infof("capture error %v is excluded => skip", err)
		return nil
	}
	if scope == nil {
		scope = sentry.NewScope()
	}
	if hint == nil {
		hint = &sentry.EventHint{}
	}
	if hint.OriginalException == nil {
		hint.OriginalException = err
	}
	eventID := c.client.CaptureException(err, hint, scope)
	if eventID != nil {
		glog.V(3).Infof("capture sentry exception with id %s", *eventID)
	} else {
		glog.V(2).Infof("capture sentry exception failed: eventID is nil")
	}
	return eventID
}

func (c *client) Close() error {
	c.client.Flush(2 * stdtime.Second)
	return nil
}
