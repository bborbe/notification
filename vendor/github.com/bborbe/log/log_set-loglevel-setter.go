// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"context"
	"flag"
	"strconv"
	"sync"
	"time"

	libtime "github.com/bborbe/time"
	"github.com/golang/glog"
)

//counterfeiter:generate -o mocks/log-loglevel-setter.go --fake-name LogLevelSetter . LogLevelSetter

// LogLevelSetter provides an interface for dynamically changing log levels at runtime.
// This is particularly useful for debugging production systems without restarts.
//
//nolint:revive // LogLevelSetter is the canonical name; changing it would break the API
type LogLevelSetter interface {
	// Set changes the current log level to the specified value.
	// The implementation may automatically reset to a default level after a timeout.
	Set(ctx context.Context, logLevel glog.Level) error
}

// LogLevelSetterFunc is a function type that implements the LogLevelSetter interface.
// It allows regular functions to be used as log level setters.
//
//nolint:revive // LogLevelSetterFunc matches LogLevelSetter; changing it would break the API
type LogLevelSetterFunc func(ctx context.Context, logLevel glog.Level) error

// Set implements the LogLevelSetter interface by calling the underlying function.
func (l LogLevelSetterFunc) Set(ctx context.Context, logLevel glog.Level) error {
	return l(ctx, logLevel)
}

// NewLogLevelSetter creates a new LogLevelSetter that automatically resets to the
// default log level after the specified duration.
//
// Parameters:
//   - defaultLoglevel: The log level to reset to after the auto-reset duration
//   - autoResetDuration: How long to wait before automatically resetting the log level
//
// The setter is thread-safe and can handle concurrent log level changes.
func NewLogLevelSetter(
	defaultLoglevel glog.Level,
	autoResetDuration time.Duration,
) LogLevelSetter {
	return &logLevelSetter{
		defaultLoglevel:   defaultLoglevel,
		autoResetDuration: autoResetDuration,
	}

}

type logLevelSetter struct {
	autoResetDuration time.Duration
	defaultLoglevel   glog.Level

	mux             sync.Mutex
	lastSetTime     time.Time
	currentLogLevel glog.Level
}

func (l *logLevelSetter) Set(ctx context.Context, logLevel glog.Level) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	l.lastSetTime = libtime.Now()
	l.currentLogLevel = logLevel

	_ = flag.Set("v", strconv.Itoa(int(logLevel)))

	glog.V(l.defaultLoglevel).
		Infof("set loglevel to %d and reset in %v back to %d", logLevel, l.autoResetDuration, l.defaultLoglevel)
	go func() { // #nosec G118 -- intentional: reset must outlive caller's context
		ctx, cancel := context.WithTimeout(context.Background(), l.autoResetDuration)
		defer cancel()

		<-ctx.Done()
		l.resetLogLevel()
	}()
	return nil
}

func (l *logLevelSetter) resetLogLevel() {
	l.mux.Lock()
	defer l.mux.Unlock()

	if libtime.Now().Sub(l.lastSetTime) <= l.autoResetDuration {
		glog.V(l.defaultLoglevel).Infof("time since lastSet is too short => skip reset loglevel")
		return
	}

	_ = flag.Set("v", strconv.Itoa(int(l.defaultLoglevel)))
	glog.V(l.defaultLoglevel).Infof("loglevel set back to %d", l.defaultLoglevel)
}
