# Go Log Utilities

[![CI](https://github.com/bborbe/log/actions/workflows/ci.yml/badge.svg)](https://github.com/bborbe/log/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/log)](https://goreportcard.com/report/github.com/bborbe/log)
[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/log.svg)](https://pkg.go.dev/github.com/bborbe/log)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

A Go library providing advanced logging utilities focused on log sampling and dynamic log level management, designed to integrate seamlessly with Google's `glog` library.

## Features

- **Log Sampling**: Reduce log volume with intelligent sampling mechanisms
- **Dynamic Log Level Management**: Change log levels at runtime via HTTP endpoints
- **Multiple Sampler Types**: Counter-based, time-based, glog-level-based, and custom samplers
- **Thread-Safe**: All components are designed for concurrent use
- **Extensible**: Factory pattern and interface-based design for easy customization

---

## Requirements

- Go 1.25 or later

## Installation

```bash
go get github.com/bborbe/log
```

## Quick Start

### Basic Log Sampling

```go
package main

import (
    "time"
    "github.com/bborbe/log"
    "github.com/golang/glog"
)

func main() {
    // Sample every 10th log entry
    modSampler := log.NewSampleMod(10)
    
    // Sample once every 10 seconds
    timeSampler := log.NewSampleTime(10 * time.Second)
    
    // Use in your logging code
    if modSampler.IsSample() {
        glog.V(2).Infof("This will be logged every 10th time")
    }
    
    if timeSampler.IsSample() {
        glog.V(2).Infof("This will be logged at most once every 10 seconds")
    }
}
```

### Dynamic Log Level Management

```go
package main

import (
    "context"
    "net/http"
    "time"
    
    "github.com/bborbe/log"
    "github.com/golang/glog"
    "github.com/gorilla/mux"
)

func main() {
    ctx := context.Background()
    
    // Create log level setter that auto-resets after 5 minutes
    logLevelSetter := log.NewLogLevelSetter(
        glog.Level(1), // default level
        5*time.Minute, // auto-reset duration
    )
    
    // Set up HTTP handler for dynamic log level changes
    router := mux.NewRouter()
    router.Handle("/debug/loglevel/{level}", 
        log.NewSetLoglevelHandler(ctx, logLevelSetter))
    
    http.ListenAndServe(":8080", router)
}
```

Now you can change log levels at runtime:
```bash
curl http://localhost:8080/debug/loglevel/4
```

## Sampler Types

### ModSampler
Samples every Nth log entry based on a counter:
```go
sampler := log.NewSampleMod(100) // Sample every 100th log
```

### TimeSampler
Samples based on time intervals:
```go
sampler := log.NewSampleTime(30 * time.Second) // Sample at most once per 30 seconds
```

### GlogLevelSampler
Samples based on glog verbosity levels:
```go
sampler := log.NewSamplerGlogLevel(3) // Sample when glog level >= 3
```

### ListSampler
Combines multiple samplers with OR logic:
```go
sampler := log.SamplerList{
    log.NewSampleTime(10 * time.Second),
    log.NewSamplerGlogLevel(4),
}
```

### FuncSampler
Create custom sampling logic:
```go
sampler := log.SamplerFunc(func() bool {
    // Your custom sampling logic
    return shouldSample()
})
```

### TrueSampler
Always samples (useful for testing or special cases):
```go
sampler := log.SamplerTrue{}
```

## Factory Pattern

Use the factory pattern for dependency injection:

```go
// Use the default factory
factory := log.DefaultSamplerFactory
sampler := factory.Sampler()

// Or create a custom factory
customFactory := log.SamplerFactoryFunc(func() log.Sampler {
    return log.NewSampleMod(50)
})
```

---

## Full Example

Here's a complete, runnable example demonstrating multiple features working together in a production-like scenario:

```go
package main

import (
    "context"
    "net/http"
    "time"

    "github.com/bborbe/log"
    "github.com/golang/glog"
    "github.com/gorilla/mux"
)

func main() {
    ctx := context.Background()

    // Combine time-based and level-based sampling
    // This will sample if EITHER condition is met (OR logic):
    // - At most once every 10 seconds, OR
    // - When glog verbosity is >= 4
    sampler := log.SamplerList{
        log.NewSampleTime(10 * time.Second),
        log.NewSamplerGlogLevel(4),
    }

    // Set up dynamic log level management
    // Default level: 1, auto-resets after 5 minutes
    logLevelSetter := log.NewLogLevelSetter(glog.Level(1), 5*time.Minute)

    // Create HTTP server with debug endpoint
    router := mux.NewRouter()
    router.Handle("/debug/loglevel/{level}",
        log.NewSetLoglevelHandler(ctx, logLevelSetter))

    // Start HTTP server in background
    go func() {
        if err := http.ListenAndServe(":8080", router); err != nil {
            glog.Fatalf("HTTP server failed: %v", err)
        }
    }()

    // Example application loop with sampled logging
    for i := 0; i < 1000; i++ {
        // High-frequency operation
        processItem(i)

        // Sampled logging to avoid log spam
        if sampler.IsSample() {
            glog.V(2).Infof("Processed item %d", i)
        }

        time.Sleep(100 * time.Millisecond)
    }
}

func processItem(i int) {
    // Your application logic here
    _ = i
}
```

You can change log levels at runtime:
```bash
# Increase verbosity to see more logs
curl http://localhost:8080/debug/loglevel/4

# Response: set loglevel to 4 completed
# Log level will auto-reset to 1 after 5 minutes
```

---

## HTTP Log Level Management

The library provides built-in HTTP handlers for runtime log level changes:

- **Endpoint**: `GET/POST /debug/loglevel/{level}`
- **Auto-reset**: Automatically reverts to default level after specified duration
- **Thread-safe**: Safe for concurrent access

Example integration with gorilla/mux:
```go
router := mux.NewRouter()
logLevelSetter := log.NewLogLevelSetter(glog.Level(1), 5*time.Minute)
router.Handle("/debug/loglevel/{level}",
    log.NewSetLoglevelHandler(context.Background(), logLevelSetter))
```

---

## Development

### Running Tests
```bash
make test
```

### Code Generation (Mocks)
```bash
make generate
```

### Full Development Workflow
```bash
make precommit  # Format, test, lint, and check
```

## Testing Framework

The library uses:
- **Ginkgo v2** for BDD-style testing
- **Gomega** for assertions
- **Counterfeiter** for mock generation

---

## Testing

### Testing Code That Uses Log Samplers

When testing code that uses this library, you have several options for controlling sampling behavior:

#### Option 1: Use TrueSampler for Always Sampling

```go
import (
    "testing"
    "github.com/bborbe/log"
    "github.com/golang/glog"
)

func TestYourCode(t *testing.T) {
    // Use TrueSampler to ensure logs always sample during tests
    sampler := log.NewSamplerTrue()

    // Your test code here
    if sampler.IsSample() {
        glog.V(2).Infof("This will always log in tests")
    }
}
```

#### Option 2: Use SamplerFunc for Controlled Testing

```go
func TestWithControlledSampling(t *testing.T) {
    shouldSample := true
    sampler := log.SamplerFunc(func() bool {
        return shouldSample
    })

    // Test when sampling is enabled
    if sampler.IsSample() {
        glog.V(2).Infof("Sampled log")
    }

    // Test when sampling is disabled
    shouldSample = false
    if sampler.IsSample() {
        t.Error("Should not sample")
    }
}
```

#### Option 3: Use Mock Samplers (Counterfeiter)

```go
import (
    "testing"
    "github.com/bborbe/log/mocks"
)

func TestWithMockSampler(t *testing.T) {
    mockSampler := &mocks.LogSampler{}

    // Configure mock behavior
    mockSampler.IsSampleReturns(true)

    // Your test code using the mock
    result := mockSampler.IsSample()
    if !result {
        t.Error("Expected sampling to be enabled")
    }

    // Verify mock was called
    if mockSampler.IsSampleCallCount() != 1 {
        t.Error("Expected IsSample to be called once")
    }
}
```

#### Testing with Ginkgo/Gomega

```go
import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/bborbe/log"
)

var _ = Describe("YourComponent", func() {
    var sampler log.Sampler

    BeforeEach(func() {
        sampler = log.NewSamplerTrue()
    })

    It("should sample logs", func() {
        Expect(sampler.IsSample()).To(BeTrue())
    })
})
```

---

## License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Add tests for your changes
4. Run `make precommit` to ensure code quality
5. Submit a pull request

## Dependencies

- [glog](https://github.com/golang/glog) - Core logging functionality
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP routing for log level endpoints
- [github.com/bborbe/time](https://github.com/bborbe/time) - Time utilities
