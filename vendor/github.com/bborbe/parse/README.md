# Parse

[![CI](https://github.com/bborbe/parse/workflows/CI/badge.svg)](https://github.com/bborbe/parse/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/parse)](https://goreportcard.com/report/github.com/bborbe/parse)
[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/parse.svg)](https://pkg.go.dev/github.com/bborbe/parse)
[![Coverage Status](https://coveralls.io/repos/github/bborbe/parse/badge.svg?branch=master)](https://coveralls.io/github/bborbe/parse?branch=master)

A robust Go utility library for parsing and converting various data types from `interface{}` values with comprehensive fallback mechanisms and error handling.

## Features

- **Type-safe parsing** with context support
- **Fallback values** via `ParseXDefault` functions
- **Subtype handling** for custom types derived from basic types
- **Interface support** for `fmt.Stringer` implementations
- **Comprehensive error handling** with structured errors
- **Zero dependencies** for runtime (only `github.com/bborbe/errors`)

## Installation

```bash
go get github.com/bborbe/parse
```

## Usage

### Basic String Parsing

```go
import (
    "context"
    "github.com/bborbe/parse"
)

// Parse with error handling
value, err := parse.ParseString(context.Background(), 42)
if err != nil {
    // Handle error
}
fmt.Println(value) // "42"

// Parse with default fallback
value := parse.ParseStringDefault(context.Background(), invalidValue, "default")
fmt.Println(value) // "default"
```

### Integer Parsing

```go
// Parse int
num, err := parse.ParseInt(context.Background(), "123")
if err != nil {
    // Handle error
}
fmt.Println(num) // 123

// Parse int64 with default
num64 := parse.ParseInt64Default(context.Background(), "invalid", 0)
fmt.Println(num64) // 0
```

### Boolean Parsing

```go
// Parse bool
flag, err := parse.ParseBool(context.Background(), "true")
if err != nil {
    // Handle error
}
fmt.Println(flag) // true

// Parse with default
flag := parse.ParseBoolDefault(context.Background(), "invalid", false)
fmt.Println(flag) // false
```

### Float Parsing

```go
// Parse float64
f, err := parse.ParseFloat64(context.Background(), "3.14")
if err != nil {
    // Handle error
}
fmt.Println(f) // 3.14
```

### Array Parsing

```go
// Parse string array
strs, err := parse.ParseStrings(context.Background(), []interface{}{"a", "b", "c"})
if err != nil {
    // Handle error
}
fmt.Println(strs) // ["a", "b", "c"]

// Parse int array with default
ints := parse.ParseIntArrayDefault(context.Background(), "invalid", []int{1, 2, 3})
fmt.Println(ints) // [1, 2, 3]
```

### Time Parsing

```go
// Parse time with custom format
t, err := parse.ParseTime(context.Background(), "2023-12-25", "2006-01-02")
if err != nil {
    // Handle error
}
fmt.Println(t) // 2023-12-25 00:00:00 +0000 UTC
```

### ASCII Conversion

```go
// Convert to ASCII (removes diacritics)
ascii, err := parse.ParseASCII(context.Background(), "café")
if err != nil {
    // Handle error
}
fmt.Println(ascii) // "cafe"
```

### Custom Types

The library supports custom types derived from basic types:

```go
type MyString string

value, err := parse.ParseString(context.Background(), MyString("hello"))
// Works seamlessly with custom types
```

## API Reference

### Core Functions

- `ParseString(ctx, value) (string, error)` - Parse to string
- `ParseInt(ctx, value) (int, error)` - Parse to int
- `ParseInt64(ctx, value) (int64, error)` - Parse to int64
- `ParseBool(ctx, value) (bool, error)` - Parse to bool
- `ParseFloat64(ctx, value) (float64, error)` - Parse to float64
- `ParseTime(ctx, value, format) (time.Time, error)` - Parse to time.Time
- `ParseASCII(ctx, value) (string, error)` - Convert to ASCII

### Array Functions

- `ParseStrings(ctx, value) ([]string, error)` - Parse to string array
- `ParseIntArray(ctx, value) ([]int, error)` - Parse to int array
- `ParseInt64Array(ctx, value) ([]int64, error)` - Parse to int64 array

### Default Functions

All parse functions have corresponding `ParseXDefault` variants that return a fallback value on error:

- `ParseStringDefault(ctx, value, defaultValue) string`
- `ParseIntDefault(ctx, value, defaultValue) int`
- `ParseBoolDefault(ctx, value, defaultValue) bool`
- etc.

Full API documentation: [pkg.go.dev/github.com/bborbe/parse](https://pkg.go.dev/github.com/bborbe/parse)

## License

This project is licensed under the BSD-style license. See the [LICENSE](LICENSE) file for details.
