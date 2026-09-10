# Changelog

All notable changes to this project will be documented in this file.

## v1.20.25

- chore: update github.com/bborbe/errors to v1.5.21

## v1.20.24

- chore: update Go to 1.27.0 and github.com/bborbe/errors to v1.5.20, github.com/bborbe/run to v1.9.37, github.com/onsi/ginkgo/v2 to v2.32.1

## v1.20.23

- chore: Bump golangci-lint to v2.13.1 and errcheck to v1.20.0 for Go 1.27 toolchain compatibility
- chore: Run `gofmt -w` last in the `format` target so golines wrapping is normalized before the gofmt lint check

## v1.20.22

- chore: update dependencies (bborbe/errors v1.5.18, bborbe/run v1.9.35)

## v1.20.21

- update Go to 1.26.6 and update dependencies (GO-2026-5972, GO-2026-6090)

## v1.20.20

- update Go to 1.26.5 and update dependencies

## v1.20.19

- fix: Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.20.18

- Bump github.com/bborbe/errors to v1.5.16
- Bump Go toolchain to 1.26.5

## v1.20.17

- Update github.com/bborbe/run to v1.9.30

## v1.20.16

- Bump github.com/bborbe/run to v1.9.29
- Bump github.com/getsentry/sentry-go to v0.47.0

## v1.20.15

- Update bborbe/errors dependency to v1.5.15

## v1.20.14

- Update github.com/bborbe/run from v1.9.27 to v1.9.28
- Update ginkgo from v2.29.0 to v2.32.0 and gomega from v1.41.0 to v1.42.1
- Bump golang.org/x/* indirect dependencies

## v1.20.13

- bump go 1.26.3 → 1.26.4
- bump bborbe/run v1.9.23 → v1.9.27
- bump ginkgo/v2 v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0
- bump golang.org/x/net, x/sys, x/text, getsentry/sentry-go
- add cloud.google.com/go v0.26.0 exclude

## v1.20.12

- bump go 1.26.2 → 1.26.3
- bump github.com/bborbe/errors v1.5.11 → v1.5.13

## v1.20.11

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 445 to ~38 lines

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.20.10

- Update bborbe/errors to v1.5.10 and bborbe/run to v1.9.21
- Update golang.org/x dependencies (crypto, mod, net, text, tools, vuln)
- Update indirect deps: getsentry, go-git, telemetry, term

## v1.20.9

- update go to 1.26.2
- update bborbe/errors to v1.5.9, bborbe/run to v1.9.17
- update golangci-lint to v2.11.4, osv-scanner to v2.3.5
- update shoenig/go-modtool to v0.7.1, counterfeiter to v6.12.2
- update various indirect dependencies (golang.org/x/sys, fatih/color, sentry-go, etc.)

## v1.20.8

- Update multiple indirect dependencies (docker, containerd, opentelemetry, moby)
- Replace k8s.io/kube-openapi replace directive with charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Remove large exclude block for k8s and sigs.k8s.io versions
- Add --allow-parallel-runners flag to golangci-lint invocation

## v1.20.7

- standardize Makefile: add mocks mkdir, reorder lint target, multiline trivy format

## v1.20.6

- chore: fix corrupted go module cache entries (osv-scalibr, containerd) to restore precommit health

## v1.20.5

- upgrade golangci-lint from v1 to v2
- update bborbe/errors to v1.5.5
- update bborbe/run to v1.9.8

## v1.20.4

- go mod update

## v1.20.3

- Update Go to 1.26.0

## v1.20.2

- Update Go from 1.25.5 to 1.25.7
- Update github.com/bborbe/errors to v1.5.2
- Update github.com/onsi/ginkgo/v2 to v2.28.1
- Update github.com/onsi/gomega to v1.39.1
- Update various transitive dependencies

## v1.20.1

- Update Go to 1.25.5
- Update dependencies

## v1.20.0

- update go and deps

## v1.19.0
- **BREAKING**: Rename Map to Each for slice operations (side effects, no transformation)
- **BREAKING**: Add new Map function with type transformation signature `Map[A, B any](ctx, []A, fn) ([]B, error)`
- Add context.Context support to Each function for cancellation
- Add context.Context support to Map function for cancellation
- Add Each() method to Set interface with context support
- Add Each() method to SetEqual interface with context support
- Add Each() method to SetHashCode interface with context support
- Add golangci-lint exclusion for interface duplication between SetEqual and SetHashCode
- Update copyright year to 2025 for new and modified files
- Add 14 new test cases for Set Each methods
- Maintain 97.5% test coverage

## v1.18.0
- Add Clone() method to Set, SetEqual, and SetHashCode for creating independent copies
- Add Without() method to Set, SetEqual, and SetHashCode for immutable element removal
- Add ErrNotFound as standard error name (NotFoundError kept as deprecated alias)
- Improve error wrapping consistency (use errors.Wrap instead of errors.Wrapf without format params)
- Optimize Copy() and Join() functions using direct append operations
- Fix unchecked type assertions in test files (13 occurrences)
- Fix test description typos (lenght → length)
- Enable lint target in Makefile check (was previously disabled)
- Enhance golangci-lint configuration with additional linters (funlen, gocognit, nestif, maintidx, errname, forcetypeassert, bodyclose, asasalint, prealloc)
- Update Go version from 1.25.3 to 1.25.4 in CI workflow
- Update dependencies (github.com/bborbe/run 1.8.1 → 1.8.2, and others)
- Remove deprecated golang.org/x/lint from tools.go
- Add 102 new test cases for Clone() and Without() methods
- Maintain 98.6% test coverage

## v1.17.0
- Add Intersect function to find common elements between two slices
- Intersect preserves order from first slice and handles duplicates automatically
- Add comprehensive test suite with 21 tests covering strings, integers, and custom types
- Update Go version from 1.25.2 to 1.25.3
- Update dependencies (github.com/bborbe/run, github.com/onsi/ginkgo/v2, and others)

## v1.16.0

- Add MarshalJSON and UnmarshalJSON methods to Set interface for JSON serialization
- Add MarshalJSON and UnmarshalJSON methods to SetEqual interface for JSON serialization
- Add MarshalJSON and UnmarshalJSON methods to SetHashCode interface for JSON serialization
- Support JSON arrays for all Set types (primitives, structs, maps, nested objects)
- Add comprehensive test coverage for JSON marshaling (39 new tests)
- Add tests for Sets embedded in structs with single and multiple Set fields
- Add tests for nested structs containing Set fields
- Add tests for round-trip JSON marshal/unmarshal operations

## v1.15.0

- Add UnmarshalText and MarshalText methods to Set interface for text encoding/decoding support
- Add comprehensive GoDoc comments for UnmarshalText and MarshalText methods

## v1.14.2

- Fix Set.UnmarshalText to properly support custom string-based types using unsafe.Pointer conversion
- Add comprehensive tests for Set with custom string types and type aliases

## v1.14.1

- Add MarshalText implementation for Set to enable automatic parsing with github.com/bborbe/argument

## v1.14.0

- Add ParseSetFromStrings function for converting string slices to Set with ~string constraint
- Add ParseSetFromString function for parsing comma-separated strings to Set with ~string constraint
- Add UnmarshalText implementation for Set to enable automatic parsing with github.com/bborbe/argument
- Add comprehensive test coverage for UnmarshalText and parsing functions

## v1.13.1

- refactor Strings() method to eliminate code duplication using elementToString helper
- refactor String() method to use formatSetString helper for consistent output
- optimize memory usage: change map[T]bool to map[T]struct{} in ContainsAny and ContainsAll
- standardize map lookup pattern to idiomatic two-value form
- add package-level documentation (doc.go) describing library features and architecture
- enhance HasHashCode documentation with security guidance for hash collision prevention
- add performance notes to ContainsAny and ContainsAll functions
- update copyright headers to 2025
- improve test coverage to 99.3%

## v1.13.0

- add ContainsAny function for checking if any element from one slice exists in another
- add ContainsAll and ContainsAny methods to Set, SetEqual, and SetHashCode interfaces
- add Strings() method to all Set implementations for sorted string slice output
- enhance Remove methods to accept variadic parameters for batch removal with single mutex lock
- improve String() methods to use sorted output for deterministic debugging
- optimize Strings() implementation with type switch for better performance (Stringer, string, default)
- add comprehensive test coverage for new ContainsAll and ContainsAny methods

## v1.12.0

- add String() method to Set, SetEqual, and SetHashCode interfaces for human-readable output
- implement efficient string representation using strings.Builder
- add comprehensive test coverage for String() methods with 100% code coverage
- document non-deterministic ordering behavior for map-based sets
- document insertion order preservation for SetEqual string representation

## v1.11.1

- update dependencies to resolve compatibility issues
- add exclude directives for cloud.google.com/go v0.26.0 and golang.org/x/tools v0.38.0

## v1.11.0

- add variadic constructors to Set, SetEqual, and SetHashCode for initialization with elements
- enhance Add methods to accept variadic parameters for batch operations with single mutex lock
- add comprehensive test coverage with 25 new test cases for variadic functionality
- improve GoDoc comments with performance characteristics and complexity analysis
- add performance comparison documentation between Set implementations
- add usage examples to all constructor documentation
- add go-modtool to Makefile for go.mod formatting
- update Go version to 1.25.2
- update dependencies (bborbe/errors, bborbe/run, and security tools)

## v1.10.2

- add GitHub Actions CI workflow for automated testing
- add GitHub Actions workflows for Claude Code integration
- add golangci-lint configuration
- enhance Makefile with additional quality checks (osv-scanner, gosec, trivy)
- improve test code formatting with golines
- update Go version to 1.24.6
- add comprehensive linting and security scanning tools

## v1.10.1

- add comprehensive documentation comments to all public functions, types, and interfaces
- improve documentation following Go doc best practices guidelines

## v1.10.0

**BREAKING CHANGES:**
- ContainsAll behavior changed: now only checks if all elements of second argument are present in first argument (superset check), instead of bidirectional equality check

## v1.9.1

- add tests
- go mod update

## v1.9.0

- remove vendor
- go mod update

## v1.8.0

- add StreamList
- go mod update

## v1.7.0

- add join
- go mod update

## v1.6.0

- add compare

## v1.5.0

- add Map helper
- go mod update

## v1.4.1

- go mod update

## v1.4.0

- add length method to sets
- go mod update

## v1.3.2

- go mod update

## v1.3.1

- add SetHashCode
- add SetEqual

## v1.3.0

- add Set

## v1.2.0

- add Equal
- add ContainsAll

## v1.1.0

- add ChannelFnCount

## v1.0.0

- Initial Version
