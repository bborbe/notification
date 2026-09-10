# KV - Key-Value Store Abstraction Library

A Go library providing unified interfaces for key-value stores. This library abstracts different KV database implementations (BadgerDB, BoltDB, in-memory) behind common interfaces, enabling easy switching between backends without code changes.

## Features

- **Unified Interface**: Common API for different key-value stores
- **Generic Type Safety**: Full Go 1.18+ generics support for type-safe operations
- **Transaction Support**: Read and write transactions with proper error handling
- **Bucket Management**: Organized data storage with bucket abstraction
- **Metrics Integration**: Built-in Prometheus metrics support
- **Relation Store**: Bidirectional 1:N relationship management
- **Benchmarking**: Performance testing framework included

## Quick Start

### Installation

```bash
go get github.com/bborbe/kv
```

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "github.com/bborbe/kv"
    "github.com/bborbe/badgerkv" // or boltkv, memorykv
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {
    ctx := context.Background()
    
    // Create database instance (example with badgerkv)
    db, err := badgerkv.Open("/tmp/mydb")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
    // Create a type-safe store
    userStore := kv.NewStore[string, User](db, kv.BucketName("users"))
    
    // Add a user
    user := User{ID: "123", Name: "Alice", Email: "alice@example.com"}
    if err := userStore.Add(ctx, user.ID, user); err != nil {
        panic(err)
    }
    
    // Get a user
    retrievedUser, err := userStore.Get(ctx, "123")
    if err != nil {
        panic(err)
    }
    fmt.Printf("User: %+v\n", *retrievedUser)
    
    // Check if user exists
    exists, err := userStore.Exists(ctx, "123")
    if err != nil {
        panic(err)
    }
    fmt.Printf("User exists: %v\n", exists)
}
```

### Using Transactions

```go
// Manual transaction control
err := db.Update(ctx, func(ctx context.Context, tx kv.Tx) error {
    bucket, err := tx.CreateBucket(kv.BucketName("users"))
    if err != nil {
        return err
    }
    
    userData, _ := json.Marshal(user)
    return bucket.Put(ctx, []byte("123"), userData)
})
```

### Iterating Over Data

```go
// Stream all users
userCh := make(chan User, 10)
go func() {
    defer close(userCh)
    userStore.Stream(ctx, userCh)
}()

for user := range userCh {
    fmt.Printf("User: %+v\n", user)
}

// Or map over all users
userStore.Map(ctx, func(ctx context.Context, key string, user User) error {
    fmt.Printf("Key: %s, User: %+v\n", key, user)
    return nil
})
```

## Architecture

### Interface Hierarchy

The library uses a layered interface approach:

1. **DB Interface** - Database lifecycle management
   - `Update()` - Write transactions  
   - `View()` - Read-only transactions
   - `Sync()`, `Close()`, `Remove()` - Database management

2. **Transaction Interface** - Bucket operations within transactions
   - `Bucket()`, `CreateBucket()`, `DeleteBucket()` - Bucket management
   - `ListBucketNames()` - Bucket enumeration

3. **Bucket Interface** - Key-value operations
   - `Put()`, `Get()`, `Delete()` - Basic operations
   - `Iterator()`, `IteratorReverse()` - Iteration support

4. **Generic Store Layer** - Type-safe operations
   - Uses Go generics: `Store[KEY ~[]byte | ~string, OBJECT any]`
   - Automatic JSON marshaling/unmarshaling
   - Composed interfaces: `StoreAdder`, `StoreGetter`, `StoreRemover`, etc.

### Advanced Features

#### Metrics Wrapper
Monitor your database operations with Prometheus metrics:

```go
import "github.com/prometheus/client_golang/prometheus"

// Wrap your DB with metrics
metricsDB := kv.NewDBWithMetrics(db, prometheus.DefaultRegisterer, "myapp")
```

#### Relation Store
Manage bidirectional 1:N relationships:

```go
relationStore := kv.NewRelationStore[string, string](db, kv.BucketName("user_groups"))

// Add relationships
relationStore.Add(ctx, "user123", "group456")
relationStore.Add(ctx, "user123", "group789")

// Query relationships
groups, err := relationStore.GetByA(ctx, "user123") // Returns: ["group456", "group789"]
users, err := relationStore.GetByB(ctx, "group456") // Returns: ["user123"]
```

## Implementations

This library defines interfaces implemented by three concrete packages:

- **[badgerkv](https://github.com/bborbe/badgerkv)** - BadgerDB implementation (LSM-tree, high performance)
- **[boltkv](https://github.com/bborbe/boltkv)** - BoltDB implementation (B+ tree, ACID compliance)  
- **[memorykv](https://github.com/bborbe/memorykv)** - In-memory implementation (testing/development)

### Switching Implementations

```go
// BadgerDB for high-performance scenarios
db, err := badgerkv.Open("/path/to/badger")

// BoltDB for ACID compliance
db, err := boltkv.Open("/path/to/bolt.db")

// In-memory for testing
db := memorykv.New()
```

## Testing

The library includes comprehensive test suites that can be reused for implementations:

```go
import "github.com/bborbe/kv"

// Use the basic test suite for your implementation
var _ = Describe("MyKV", func() {
    kv.BasicTestSuite(func() kv.Provider {
        return mykvProvider{}
    })
})
```

## Development

### Prerequisites
- Go 1.18+ (for generics support)

### Commands

```bash
# Full pre-commit pipeline
make precommit

# Run tests
make test

# Format code
make format

# Generate mocks
make generate

# Check code quality
make check
```

## License

This project is licensed under the BSD-style license. See the [LICENSE](LICENSE) file for details.
