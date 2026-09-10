# Collection

A comprehensive Go library providing generic, type-safe collection utilities and data structures. Leverages Go generics for efficient slice operations, set management, channel processing, and pointer utilities.

## Features

- **🔧 Slice Operations**: Filter, Map, Find, Contains, Unique, Reverse, Copy, Exclude, Join
- **📦 Set Data Structure**: Thread-safe generic set with mutex protection
- **⚡ Channel Processing**: Concurrent channel operations with context support
- **🎯 Pointer Utilities**: Generic Ptr/Unptr functions for easy pointer creation
- **⚖️ Comparison Utilities**: Generic equality and string comparison helpers

## Installation

```bash
go get github.com/bborbe/collection
```

## Quick Start

```go
import "github.com/bborbe/collection"

// Filter slice elements
numbers := []int{1, 2, 3, 4, 5}
evens := collection.Filter(numbers, func(n int) bool { return n%2 == 0 })
// Result: [2, 4]

// Create and use a Set
set := collection.NewSet[string]()
set.Add("hello")
set.Add("world")
fmt.Println(set.Contains("hello")) // true

// Create pointers easily
name := "John"
namePtr := collection.Ptr(name) // *string pointing to "John"
```

## API Reference

### Slice Operations

#### Filter
```go
func Filter[T any](list []T, match func(value T) bool) []T
```
Returns a new slice containing only elements that match the predicate.

```go
users := []User{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 17}}
adults := collection.Filter(users, func(u User) bool { return u.Age >= 18 })
```

#### Find
```go
func Find[T any](list []T, match func(value T) bool) (*T, error)
```
Returns the first element matching the predicate, or `NotFoundError`.

```go
user, err := collection.Find(users, func(u User) bool { return u.Name == "Alice" })
if err != nil {
    // Handle not found
}
```

#### Map
```go
func Map[T any](list []T, fn func(value T) error) error
```
Applies a function to each element in the slice, returning the first error encountered.

```go
err := collection.Map(users, func(u User) error {
    return validateUser(u)
})
```

#### Unique
```go
func Unique[T comparable](list []T) []T
```
Returns a new slice with duplicate elements removed, preserving order.

```go
numbers := []int{1, 2, 2, 3, 3, 3}
unique := collection.Unique(numbers) // [1, 2, 3]
```

#### Contains
```go
func Contains[T comparable](list []T, search T) bool
func ContainsAll[T comparable](list []T, search []T) bool
```
Check if slice contains specific element(s).

```go
fruits := []string{"apple", "banana", "cherry"}
hasApple := collection.Contains(fruits, "apple") // true
hasAll := collection.ContainsAll(fruits, []string{"apple", "banana"}) // true
```

#### Reverse
```go
func Reverse[T any](list []T) []T
```
Returns a new slice with elements in reverse order.

```go
numbers := []int{1, 2, 3}
reversed := collection.Reverse(numbers) // [3, 2, 1]
```

#### Copy
```go
func Copy[T any](list []T) []T
```
Creates a shallow copy of the slice.

#### Exclude
```go
func Exclude[T comparable](list []T, exclude []T) []T
```
Returns a new slice excluding specified elements.

```go
numbers := []int{1, 2, 3, 4, 5}
filtered := collection.Exclude(numbers, []int{2, 4}) // [1, 3, 5]
```

#### Join
```go
func Join[T ~string](list []T, separator string) string
```
Joins string-like elements with a separator.

```go
words := []string{"hello", "world", "!"}
sentence := collection.Join(words, " ") // "hello world !"
```

### Set Data Structure

#### Creating Sets
```go
func NewSet[T comparable]() Set[T]
```
Creates a new thread-safe generic set.

```go
stringSet := collection.NewSet[string]()
intSet := collection.NewSet[int]()
```

#### Set Operations
```go
type Set[T comparable] interface {
    Add(element T)
    Remove(element T) 
    Contains(element T) bool
    Slice() []T
    Length() int
}
```

Example usage:
```go
set := collection.NewSet[string]()
set.Add("apple")
set.Add("banana")
set.Add("apple") // Duplicate, ignored

fmt.Println(set.Length()) // 2
fmt.Println(set.Contains("apple")) // true
fmt.Println(set.Slice()) // ["apple", "banana"] (order not guaranteed)

set.Remove("apple")
fmt.Println(set.Length()) // 1
```

### Pointer Utilities

#### Ptr
```go
func Ptr[T any](value T) *T
```
Creates a pointer to the given value.

```go
name := "John"
namePtr := collection.Ptr(name) // *string

age := 25
agePtr := collection.Ptr(age) // *int
```

#### Unptr
```go
func Unptr[T any](ptr *T) T
func UnptrWithDefault[T any](ptr *T, defaultValue T) T
```
Dereferences pointers safely.

```go
var namePtr *string = collection.Ptr("John")
name := collection.Unptr(namePtr) // "John"

var nilPtr *string
name := collection.UnptrWithDefault(nilPtr, "Anonymous") // "Anonymous"
```

### Channel Processing

#### ChannelFnMap
```go
func ChannelFnMap[T interface{}](
    ctx context.Context,
    getFn func(ctx context.Context, ch chan<- T) error,
    mapFn func(ctx context.Context, t T) error,
) error
```
Processes data from a producer function through a mapper function concurrently.

```go
err := collection.ChannelFnMap(
    ctx,
    func(ctx context.Context, ch chan<- int) error {
        for i := 0; i < 100; i++ {
            select {
            case ch <- i:
            case <-ctx.Done():
                return ctx.Err()
            }
        }
        return nil
    },
    func(ctx context.Context, num int) error {
        fmt.Printf("Processing: %d\n", num)
        return nil
    },
)
```

#### ChannelFnList & ChannelFnCount
```go
func ChannelFnList[T any](ctx context.Context, getFn func(ctx context.Context, ch chan<- T) error) ([]T, error)
func ChannelFnCount[T any](ctx context.Context, getFn func(ctx context.Context, ch chan<- T) error) (int, error)
```

Collect channel data into a slice or count items.

### Comparison Utilities

#### Equal & Compare
```go
func Equal[T comparable](a, b T) bool
func Compare[T ~string](a, b T) int
```

Generic comparison functions for any comparable types.

## Advanced Examples

### Processing Large Datasets Concurrently
```go
ctx := context.Background()

// Process items concurrently with automatic batching
err := collection.ChannelFnMap(
    ctx,
    func(ctx context.Context, ch chan<- WorkItem) error {
        return loadWorkItems(ctx, ch) // Your data loading logic
    },
    func(ctx context.Context, item WorkItem) error {
        return processWorkItem(ctx, item) // Your processing logic
    },
)
```

### Building Complex Data Pipelines
```go
// Load raw data
rawData := loadRawData()

// Filter valid entries
validData := collection.Filter(rawData, func(item DataItem) bool {
    return item.IsValid()
})

// Remove duplicates
uniqueData := collection.Unique(validData)

// Find specific items
importantItem, err := collection.Find(uniqueData, func(item DataItem) bool {
    return item.Priority == "high"
})

// Create a set for fast lookups
processedIDs := collection.NewSet[string]()
for _, item := range uniqueData {
    processedIDs.Add(item.ID)
}
```

## Requirements

- Go 1.24.5 or later
- Support for Go generics

## Dependencies

- `github.com/bborbe/errors` - Enhanced error handling
- `github.com/bborbe/run` - Concurrent execution utilities

## Testing

```bash
go test ./...
```

## License

BSD-style license. See LICENSE file for details.
