# memorykv

In-memory key-value store implementing `github.com/bborbe/kv` interface.

## Installation

```bash
go get github.com/bborbe/memorykv
```

## Usage

```go
package main

import (
    "context"
    "github.com/bborbe/kv/libkv"
    "github.com/bborbe/memorykv"
)

func main() {
    ctx := context.Background()
    
    // Open in-memory database
    db, err := memorykv.OpenMemory(ctx)
    if err != nil {
        panic(err)
    }
    defer db.Close(ctx)
    
    // Use within transactions
    err = db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(ctx, []byte("my-bucket"))
        if err != nil {
            return err
        }
        return bucket.Put(ctx, []byte("key"), []byte("value"))
    })
}
```

Perfect for testing and development where you need a lightweight, ephemeral key-value store that's compatible with other `github.com/bborbe/kv` implementations.

## License

This project is licensed under the BSD-style license. See the [LICENSE](LICENSE) file for details.
