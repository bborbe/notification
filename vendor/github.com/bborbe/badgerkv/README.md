# BadgerKV

[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/badgerkv.svg)](https://pkg.go.dev/github.com/bborbe/badgerkv)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/badgerkv)](https://goreportcard.com/report/github.com/bborbe/badgerkv)

BadgerKV is a Go library that provides a standardized key-value store interface built on top of [BadgerDB](https://github.com/dgraph-io/badger). It implements the `github.com/bborbe/kv` interface, offering a clean and consistent API for database operations including transactions, buckets, and key-value operations.

## Features

- **Standardized Interface**: Implements the `github.com/bborbe/kv` interface for consistent database operations
- **Transaction Support**: Full ACID transaction support with automatic rollback on errors
- **Bucket-based Organization**: Organize data into logical buckets within transactions
- **Memory & Disk Storage**: Support for both file-based and in-memory databases
- **Iterator Support**: Forward and reverse iteration over bucket contents
- **Context-aware**: Full context support for cancellation and deadlines
- **Thread-safe**: Safe for concurrent use across multiple goroutines

## Installation

```bash
go get github.com/bborbe/badgerkv
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/bborbe/badgerkv"
)

func main() {
    ctx := context.Background()
    
    // Open a file-based database
    db, err := badgerkv.OpenPath(ctx, "/tmp/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Start a transaction
    err = db.Update(ctx, func(ctx context.Context, tx badgerkv.Tx) error {
        // Get or create a bucket
        bucket, err := tx.CreateBucketIfNotExists([]byte("users"))
        if err != nil {
            return err
        }
        
        // Store a key-value pair
        return bucket.Put([]byte("user:1"), []byte(`{"name": "John", "age": 30}`))
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // Read data
    err = db.View(ctx, func(ctx context.Context, tx badgerkv.Tx) error {
        bucket := tx.Bucket([]byte("users"))
        if bucket == nil {
            return nil // bucket doesn't exist
        }
        
        value, err := bucket.Get([]byte("user:1"))
        if err != nil {
            return err
        }
        
        log.Printf("User data: %s", value)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### In-Memory Database

```go
// Create an in-memory database (useful for testing)
db, err := badgerkv.OpenMemory(ctx)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

### Memory Optimization

```go
// Use minimal memory settings for resource-constrained environments
db, err := badgerkv.OpenPath(ctx, "/tmp/mydb", badgerkv.MinMemoryUsageOptions)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

### Iteration

```go
err = db.View(ctx, func(ctx context.Context, tx badgerkv.Tx) error {
    bucket := tx.Bucket([]byte("users"))
    if bucket == nil {
        return nil
    }
    
    // Forward iteration
    iter := bucket.Iterator()
    defer iter.Close()
    
    for iter.First(); iter.Valid(); iter.Next() {
        key := iter.Key().Data()
        value := iter.Item().Value()
        log.Printf("Key: %s, Value: %s", key, value)
    }
    
    return nil
})
```

### Custom BadgerDB Options

```go
import "github.com/dgraph-io/badger/v4"

// Define custom options
customOptions := func(opts *badger.Options) {
    opts.Logger = nil // Disable logging
    opts.SyncWrites = true // Force sync on writes
}

db, err := badgerkv.OpenPath(ctx, "/tmp/mydb", customOptions)
```

## API Overview

### Database Operations

- `OpenPath(ctx, path, ...options)` - Open file-based database
- `OpenMemory(ctx, ...options)` - Open in-memory database
- `DB.View(ctx, fn)` - Read-only transaction
- `DB.Update(ctx, fn)` - Read-write transaction
- `DB.Close()` - Close database

### Transaction Operations

- `Tx.Bucket(name)` - Get existing bucket
- `Tx.CreateBucket(name)` - Create new bucket (fails if exists)
- `Tx.CreateBucketIfNotExists(name)` - Get or create bucket
- `Tx.DeleteBucket(name)` - Delete bucket and all contents

### Bucket Operations

- `Bucket.Get(key)` - Retrieve value by key
- `Bucket.Put(key, value)` - Store key-value pair
- `Bucket.Delete(key)` - Delete key
- `Bucket.Iterator()` - Create iterator for bucket contents

### Iterator Operations

- `Iterator.First()` - Move to first item
- `Iterator.Last()` - Move to last item
- `Iterator.Next()` - Move to next item
- `Iterator.Prev()` - Move to previous item
- `Iterator.Valid()` - Check if iterator position is valid
- `Iterator.Key()` - Get current key
- `Iterator.Item()` - Get current item

## Transaction Context

BadgerKV prevents nested transactions by tracking transaction state in the context. Use `IsTransactionOpen(ctx)` to check if a transaction is already active.

```go
if badgerkv.IsTransactionOpen(ctx) {
    // Already in a transaction, cannot start nested transaction
    return errors.New("nested transactions not supported")
}
```

## Error Handling

BadgerKV uses the `github.com/bborbe/errors` package for enhanced error handling with context preservation:

```go
err = db.Update(ctx, func(ctx context.Context, tx badgerkv.Tx) error {
    bucket, err := tx.CreateBucket([]byte("test"))
    if err != nil {
        return errors.Wrapf(ctx, err, "create bucket failed")
    }
    
    return bucket.Put([]byte("key"), []byte("value"))
})

if err != nil {
    log.Printf("Transaction failed: %v", err)
}
```

## Testing

BadgerKV is thoroughly tested using Ginkgo v2 and Gomega:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestSpecificFunction ./...
```

## Dependencies

- **BadgerDB v4**: High-performance key-value database
- **github.com/bborbe/kv**: Common key-value interface
- **github.com/bborbe/errors**: Enhanced error handling
- **github.com/bborbe/collection**: Utility functions

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Run `make precommit` to ensure code quality
5. Submit a pull request

## License

This project is licensed under the BSD-style license. See the LICENSE file for details.
