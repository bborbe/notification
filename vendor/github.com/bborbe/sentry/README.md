# Sentry

[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/sentry.svg)](https://pkg.go.dev/github.com/bborbe/sentry)
[![CI](https://github.com/bborbe/sentry/actions/workflows/ci.yml/badge.svg)](https://github.com/bborbe/sentry/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/sentry)](https://goreportcard.com/report/github.com/bborbe/sentry)

A Go library that provides an enhanced wrapper around the [Sentry Go SDK](https://github.com/getsentry/sentry-go) with additional functionality for error exclusion, automatic tag enrichment, and context data extraction.

## Features

- **Automatic Tag Enrichment**: Extracts data from context and errors to add as Sentry tags
- **Error Filtering**: Configurable error exclusion to prevent noise
- **Context Integration**: Uses `github.com/bborbe/errors` for context data extraction
- **Proxy Support**: HTTP transport wrapper for proxy configurations
- **Type-Safe Interface**: Clean interface abstraction with mock generation support

## Installation

```bash
go get github.com/bborbe/sentry
```

## Quick Start

```go
package main

import (
    "context"
    "time"

    "github.com/bborbe/errors"
    "github.com/bborbe/sentry"
    "github.com/getsentry/sentry-go"
)

func main() {
    ctx := context.Background()

    // Create client with options
    client, err := sentry.NewClient(ctx, sentry.ClientOptions{
        Dsn: "your-sentry-dsn-here",
        Tags: map[string]string{
            "service": "my-app",
            "version": "1.0.0",
        },
    })
    if err != nil {
        panic(err)
    }
    defer func() {
        client.Flush(2 * time.Second)
        client.Close()
    }()

    // Add context data
    ctx = errors.AddToContext(ctx, "user_id", "12345")
    
    // Create error with data
    err = errors.AddDataToError(
        errors.New("something went wrong"),
        map[string]string{"operation": "user_login"},
    )

    // Capture exception with automatic tag enrichment
    client.CaptureException(err, &sentry.EventHint{
        Context: ctx,
        Data: map[string]interface{}{
            "request_id": "req-abc123",
            "retries": 3,
        },
    }, sentry.NewScope())
}
```

## Core Components

### Client Interface

The main interface provides these methods:

```go
type Client interface {
    CaptureMessage(message string, hint *sentry.EventHint, scope sentry.EventModifier) *sentry.EventID
    CaptureException(exception error, hint *sentry.EventHint, scope sentry.EventModifier) *sentry.EventID
    Flush(timeout time.Duration) bool
    io.Closer
}
```

### Error Exclusion

Filter out specific errors to reduce noise:

```go
excludeFunc := func(err error) bool {
    return errors.Is(err, context.Canceled)
}

client, err := sentry.NewClient(ctx, clientOptions, excludeFunc)
```

### Automatic Tag Enrichment

The client automatically extracts and adds tags from:
- Context data (using `github.com/bborbe/errors`)
- Error data (attached to errors)
- Hint data (passed in EventHint)

## API Documentation

For detailed API documentation, visit [pkg.go.dev/github.com/bborbe/sentry](https://pkg.go.dev/github.com/bborbe/sentry).

## Dependencies

This library integrates closely with:
- **github.com/bborbe/errors**: Enhanced context and error data extraction
- **github.com/getsentry/sentry-go**: Official Sentry Go SDK (v0.36.0+)

The integration with `bborbe/errors` enables automatic extraction of context data and error metadata as Sentry tags.

## Development

### Running Tests
```bash
make test
```

### Code Generation (Mocks)
```bash
make generate
```

### Full Pre-commit Workflow
```bash
make precommit  # Format, generate, test, and lint
```

## License

This project is licensed under the BSD-style license. See the LICENSE file for details.
