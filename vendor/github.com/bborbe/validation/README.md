# Validation

[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/validation.svg)](https://pkg.go.dev/github.com/bborbe/validation)
[![CI](https://github.com/bborbe/validation/actions/workflows/ci.yml/badge.svg)](https://github.com/bborbe/validation/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/validation)](https://goreportcard.com/report/github.com/bborbe/validation)

A Go validation library that provides functional composition for validation rules using the `HasValidation` interface pattern.

## Features

- **Functional Composition**: Combine validators using `All` (AND logic) and `Any` (OR logic)
- **Interface-Based**: All validators implement `HasValidation` with `Validate(ctx context.Context) error`
- **Context-Aware**: All validation functions accept context for cancellation and timeout support
- **Error Wrapping**: Consistent error handling using `github.com/bborbe/errors`
- **Comprehensive Validators**: Basic, string, slice, length, and conditional validators

## Installation

```bash
go get github.com/bborbe/validation
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/bborbe/validation"
)

func main() {
    ctx := context.Background()
    
    // Basic validation
    validator := validation.All{
        validation.NotNil(someValue),
        validation.NotEmptyString(someString),
        validation.LengthGe(someSlice, 1),
    }
    
    if err := validator.Validate(ctx); err != nil {
        fmt.Printf("Validation failed: %v\n", err)
        return
    }
    
    fmt.Println("Validation passed!")
}
```

## Available Validators

### Basic Validators
- `NotNil(value)` - Validates that a value is not nil
- `Nil(value)` - Validates that a value is nil
- `True(value)` - Validates that a boolean is true
- `False(value)` - Validates that a boolean is false
- `Equal(actual, expected)` - Validates equality

### String/Slice Validators
- `NotEmptyString(value)` - Validates that a string is not empty
- `NotEmptySlice(value)` - Validates that a slice is not empty

### Length Validators
- `LengthEqual(value, length)` - Validates exact length
- `LengthGt(value, length)` - Validates length greater than
- `LengthGe(value, length)` - Validates length greater than or equal
- `LengthLt(value, length)` - Validates length less than
- `LengthLe(value, length)` - Validates length less than or equal

### Logical Operators
- `All{validators...}` - All validators must pass (AND logic)
- `Any{validators...}` - At least one validator must pass (OR logic)
- `Either{validator1, validator2}` - Exactly one validator must pass
- `Not{validator}` - Inverts validation result

### Conditional Validators
- `NilOrValid(value, validator)` - Pass if nil or if validator passes
- `NotNilAndValid(value, validator)` - Pass if not nil and validator passes

## Usage Examples

### Combining Validators

```go
validator := validation.All{
    validation.NotNil(user),
    validation.NotEmptyString(user.Name),
    validation.Any{
        validation.NotEmptyString(user.Email),
        validation.NotEmptyString(user.Phone),
    },
}
```

### Custom Validators

```go
customValidator := validation.HasValidationFunc(func(ctx context.Context) error {
    if someCondition {
        return validation.NewError("custom validation failed")
    }
    return nil
})
```

## API Documentation

For complete API documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/bborbe/validation).

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Run tests (`make test`)
4. Run checks (`make check`)
5. Commit your changes (`git commit -am 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## Development

```bash
# Install dependencies
make ensure

# Run tests
make test

# Run all checks (format, generate, test, check, addlicense)
make precommit
```

## License

This project is licensed under the BSD 2-Clause License - see the [LICENSE](LICENSE) file for details.
