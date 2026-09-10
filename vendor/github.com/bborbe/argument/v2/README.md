# Argument

[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/argument/v2.svg)](https://pkg.go.dev/github.com/bborbe/argument/v2)
[![CI](https://github.com/bborbe/argument/actions/workflows/ci.yml/badge.svg)](https://github.com/bborbe/argument/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/argument/v2)](https://goreportcard.com/report/github.com/bborbe/argument/v2)
[![License](https://img.shields.io/github/license/bborbe/argument)](https://github.com/bborbe/argument/blob/master/LICENSE)

A declarative Go library for parsing command-line arguments and environment variables into structs using struct tags. Perfect for building CLI applications with clean configuration management.

## Features

- 🏷️ **Declarative**: Use struct tags to define argument names, environment variables, and defaults
- 🔄 **Multiple Sources**: Supports command-line arguments, environment variables, and default values
- 📋 **Slice Support**: Parse comma-separated values into slices with configurable separators
- 🔧 **Custom Parsing**: Implement `encoding.TextUnmarshaler` for complex parsing logic
- ⚡ **Zero Dependencies**: Minimal external dependencies for core functionality
- ✅ **Type Safe**: Supports all common Go types including pointers for optional values
- 🕐 **Extended Duration**: Enhanced `time.Duration` parsing with support for days and weeks
- 🧪 **Well Tested**: Comprehensive test suite with BDD-style tests

## Installation

```bash
go get github.com/bborbe/argument/v2
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/bborbe/argument/v2"
)

func main() {
    var config struct {
        Username string `arg:"username" env:"USERNAME" default:"guest"`
        Password string `arg:"password" env:"PASSWORD"`
        Port     int    `arg:"port" env:"PORT" default:"8080"`
        Debug    bool   `arg:"debug" env:"DEBUG"`
    }
    
    ctx := context.Background()
    if err := argument.Parse(ctx, &config); err != nil {
        log.Fatalf("Failed to parse arguments: %v", err)
    }
    
    fmt.Printf("Starting server on port %d for user %s\n", config.Port, config.Username)
}
```

## Usage Examples

### Basic Configuration

```go
type Config struct {
    Host     string `arg:"host" env:"HOST" default:"localhost"`
    Port     int    `arg:"port" env:"PORT" default:"8080"`
    LogLevel string `arg:"log-level" env:"LOG_LEVEL" default:"info"`
}

var config Config
err := argument.Parse(context.Background(), &config)
```

Run with: `./app -host=0.0.0.0 -port=9090 -log-level=debug`

### Optional Values with Pointers

```go
type DatabaseConfig struct {
    Host     string  `arg:"db-host" env:"DB_HOST" default:"localhost"`
    Port     int     `arg:"db-port" env:"DB_PORT" default:"5432"`
    Timeout  *int    `arg:"timeout" env:"DB_TIMEOUT"`        // Optional
    MaxConns *int    `arg:"max-conns" env:"DB_MAX_CONNS"`    // Optional
}
```

### Duration with Extended Parsing

```go
type ServerConfig struct {
    ReadTimeout  time.Duration `arg:"read-timeout" env:"READ_TIMEOUT" default:"30s"`
    WriteTimeout time.Duration `arg:"write-timeout" env:"WRITE_TIMEOUT" default:"1m"`
    IdleTimeout  time.Duration `arg:"idle-timeout" env:"IDLE_TIMEOUT" default:"2h"`
    Retention    time.Duration `arg:"retention" env:"RETENTION" default:"30d"`  // 30 days
    BackupFreq   time.Duration `arg:"backup-freq" env:"BACKUP_FREQ" default:"1w"` // 1 week
}
```

Supported duration units: `ns`, `us`, `ms`, `s`, `m`, `h`, `d` (days), `w` (weeks)

### Required Fields

```go
type APIConfig struct {
    APIKey   string `arg:"api-key" env:"API_KEY"`           // Required (no default)
    Endpoint string `arg:"endpoint" env:"ENDPOINT"`         // Required (no default)
    Region   string `arg:"region" env:"REGION" default:"us-east-1"`
}

// This will return an error if APIKey or Endpoint are not provided
err := argument.Parse(context.Background(), &config)
```

### Custom Types

You can use custom types (named types with underlying primitive types) for better type safety:

```go
type Username string
type Port int
type IsEnabled bool
type Rate float64

type AppConfig struct {
    Username Username  `arg:"user" env:"USERNAME" default:"guest"`
    Port     Port      `arg:"port" env:"PORT" default:"8080"`
    Debug    IsEnabled `arg:"debug" env:"DEBUG" default:"false"`
    Rate     Rate      `arg:"rate" env:"RATE" default:"1.5"`
}

var config AppConfig
err := argument.Parse(context.Background(), &config)

// Access values with type safety
fmt.Printf("Username: %s\n", string(config.Username))
fmt.Printf("Port: %d\n", int(config.Port)) 
fmt.Printf("Debug: %t\n", bool(config.Debug))
fmt.Printf("Rate: %f\n", float64(config.Rate))
```

Custom types work with all supported underlying types:
- `string` → `type Username string`
- `int`, `int32`, `int64`, `uint`, `uint64` → `type Port int`
- `bool` → `type IsEnabled bool` 
- `float64` → `type Rate float64`

### Slice Types

Parse comma-separated values into slices automatically:

```go
type Config struct {
    Hosts    []string  `arg:"hosts" env:"HOSTS" default:"localhost,127.0.0.1"`
    Ports    []int     `arg:"ports" env:"PORTS" separator:":" default:"8080:8081:8082"`
    Prices   []float64 `arg:"prices" default:"9.99,19.99,29.99"`
    Flags    []bool    `arg:"flags" default:"true,false,true"`

    // Custom separator for specific use cases
    Tags     []string  `arg:"tags" separator:"|" default:"prod|api|web"`
}

var config Config
err := argument.Parse(context.Background(), &config)
```

Run with: `./app -hosts=server1,server2,server3 -ports=8080:9000:9001`

Features:
- Default comma separator (`,`) can be customized with `separator:` tag
- Automatic whitespace trimming around elements
- Empty string creates empty slice
- Works with custom types: `[]Username`, `[]Environment`, etc.

### Custom Parsing with TextUnmarshaler

Implement `encoding.TextUnmarshaler` for complex parsing logic:

```go
import "encoding"

type Broker string

func (b *Broker) UnmarshalText(text []byte) error {
    value := string(text)
    if !strings.Contains(value, "://") {
        value = "plain://" + value  // Add default schema
    }
    *b = Broker(value)
    return nil
}

type KafkaConfig struct {
    Broker  Broker   `arg:"broker" default:"localhost:9092"`
    Brokers []Broker `arg:"brokers" env:"KAFKA_BROKERS"`  // Works in slices too!
}

var config KafkaConfig
err := argument.Parse(context.Background(), &config)
// broker "localhost:9092" becomes "plain://localhost:9092"
// brokers "kafka1:9092,ssl://kafka2:9093" becomes ["plain://kafka1:9092", "ssl://kafka2:9093"]
```

Use cases:
- URL validation and normalization
- Default schema/prefix handling
- Complex validation logic
- Format conversion

## Supported Types

- **Strings**: `string`
- **Integers**: `int`, `int32`, `int64`, `uint`, `uint64`
- **Floats**: `float64`
- **Booleans**: `bool`
- **Durations**: `time.Duration` (with extended parsing)
- **Pointers**: `*string`, `*int`, `*float64`, etc. (for optional values)
- **Slices**: `[]string`, `[]int`, `[]int64`, `[]uint`, `[]uint64`, `[]float64`, `[]bool`
- **Custom Types**: Named types with underlying primitive types
- **Custom Parsing**: Any type implementing `encoding.TextUnmarshaler`

## Priority Order

Values are applied with the following precedence (highest priority first):

1. **Command-line arguments** (from `arg:` tag) - **Highest priority**
2. **Environment variables** (from `env:` tag)
3. **Default values** (from `default:` tag) - Lowest priority

Command-line arguments override environment variables, which override default values.

### Priority Example

Here's a complete example demonstrating how the priority works:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/bborbe/argument/v2"
)

func main() {
    var config struct {
        Port int `arg:"port" env:"PORT" default:"8080"`
    }

    ctx := context.Background()
    if err := argument.Parse(ctx, &config); err != nil {
        log.Fatalf("Failed to parse: %v", err)
    }

    fmt.Printf("Port: %d\n", config.Port)
}
```

Running this program with different inputs shows the priority order:

```bash
# 1. Only default value (no env, no arg)
$ go run main.go
Port: 8080

# 2. Environment variable overrides default
$ PORT=9000 go run main.go
Port: 9000

# 3. Command-line argument overrides environment variable
$ PORT=9000 go run main.go -port=7000
Port: 7000

# 4. Command-line argument overrides default (no env set)
$ go run main.go -port=7000
Port: 7000
```

**Summary**: The final value is **7000** (from `-port=7000`) when both env and arg are provided, demonstrating that command-line arguments have the highest priority.

**Note**: If no `default` tag is specified and neither environment variable nor argument is provided, the field will contain its Go zero value:
- Numbers: `0` (int, float64, etc.)
- Strings: `""` (empty string)
- Booleans: `false`
- Pointers: `nil`
- Slices: `nil`

## API Documentation

For complete API documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/bborbe/argument/v2).

### Main Functions

- `Parse(ctx context.Context, data interface{}) error` - Parse arguments and environment variables (quiet mode)
- `ParseAndPrint(ctx context.Context, data interface{}) error` - Parse and print the final configuration values
- `ValidateRequired(ctx context.Context, data interface{}) error` - Check that all required fields are set

## Command-Line Usage

Your application will automatically support standard Go flag syntax:

```bash
# Long form with equals
./app -host=localhost -port=8080

# Long form with space  
./app -host localhost -port 8080

# Boolean flags
./app -debug        # sets debug=true
./app -debug=false  # sets debug=false
```

## Environment Variables

Set environment variables to configure your application:

```bash
export HOST=0.0.0.0
export PORT=9090
export DEBUG=true
./app
```

## Error Handling

The library provides detailed error messages for common issues:

```go
err := argument.Parse(ctx, &config)
if err != nil {
    // Errors include context about what failed:
    // - Missing required fields
    // - Type conversion errors  
    // - Invalid duration formats
    log.Fatal(err)
}
```

## License

This project is licensed under the BSD-style license. See the LICENSE file for details.
