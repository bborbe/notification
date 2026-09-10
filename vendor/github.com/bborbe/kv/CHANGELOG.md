# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.21.13

- chore: update github.com/bborbe/errors to v1.6.0, github.com/bborbe/http to v1.26.25, github.com/bborbe/run to v1.10.2, github.com/onsi/gomega to v1.43.0

## v1.21.12

- chore: update Go to 1.27.0 and github.com/bborbe/errors to v1.5.21, github.com/bborbe/http to v1.26.24, github.com/bborbe/log to v1.6.25, github.com/bborbe/run to v1.9.37

## v1.21.11

- chore: Run gofmt last in the `format` target and bump golangci-lint to v2.13.1 + errcheck to v1.20.0 for Go 1.27 compatibility

## v1.21.10

- update Go to 1.26.6 and update dependencies (fixes GO-2026-6179, GO-2026-6180)

## v1.21.9

- fix: `relationStoreTx.Add` and `.Remove` now check `ctx.Done()` before each iteration. Both loop over the caller-supplied `relatedIDs` doing two bucket operations per element (an `IDs()` read and a `relationIDBucket.Add()` write), so a long slice could not be interrupted if the underlying store does not honour the context. Returning early leaves the caller's `Tx` to roll back, which is the correct behaviour.

## v1.21.8

- update Go to 1.26.5 and update dependencies

## v1.21.7

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.21.6

- Bump github.com/bborbe/errors to v1.5.16
- Bump github.com/bborbe/http to v1.26.16
- Bump github.com/bborbe/log to v1.6.17
- Bump go toolchain to 1.26.5
- Update transitive bborbe dependencies

## v1.21.5

- Bump bborbe/log to v1.6.15
- Bump bborbe/run to v1.9.30

## v1.21.4

- Bump github.com/bborbe/errors to v1.5.15
- Bump github.com/bborbe/run to v1.9.29
- Bump github.com/getsentry/sentry-go to v0.47.0

## v1.21.3

- Update github.com/onsi/ginkgo/v2 to v2.32.0
- Update github.com/onsi/gomega to v1.42.1
- Bump indirect dependencies (bborbe libs, golang.org/x/*)

## v1.21.2

- bump go 1.26.4
- bump bborbe/log v1.6.14, bborbe/run v1.9.28, x/net v0.55.0, x/sys v0.45.0, x/text v0.37.0
- drop standalone errcheck/gosec tools; move config into golangci.yml
- add .maintainer.yaml (autoRelease + autoApprove)
- exclude cloud.google.com/go v0.26.0

## v1.21.1

- Add tests for `BucketName.MarshalJSON` / `UnmarshalJSON` covering plain encoding, empty, special chars, round-trip, struct fields, and non-string error case

## v1.21.0

- **BREAKING**: `DB.Stats(ctx)` now returns `*Stats` instead of `Stats` (nil on error)
- Add `DB.StatsDetailed(ctx) (*Stats, error)` — slow variant that includes per-bucket `KeyCount` and `SizeB`
- Add `Stats.Detailed bool` field indicating whether per-bucket counts are populated
- `BucketStats.KeyCount` and `BucketStats.SizeB` now use `omitempty` (skipped in fast Stats output)
- Add `BucketName.MarshalJSON` / `UnmarshalJSON` — emits the plain string name instead of base64

## v1.20.0

- Add `Stats(ctx) (Stats, error)` method to `DB` interface
- Add `Stats` and `BucketStats` types for database statistics
- Implement `Stats` in metrics wrapper (`dbWithMetrics`)
- Bump ginkgo v2.29.0, gomega v1.41.0, bborbe/time v1.27.0

## v1.19.8

- bump github.com/bborbe/log v1.6.12 → v1.6.13

## v1.19.7

- Bump Go version to 1.26.3
- Update bborbe/errors to v1.5.13, bborbe/http to v1.26.11, bborbe/run to v1.9.24
- Update bborbe/sentry to v1.9.16, getsentry/sentry-go to v0.46.2
- Clean up indirect dependencies in go.mod

## v1.19.6

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod direct deps reduced from 20 to 9.
- fix: Upgrade go-git/go-git to v5.18.0 (GHSA-3xc5-wrhm-f963)
- chore: Remove stale osv-scanner suppressions for bbolt and aws-sdk (no longer in dep graph); update docker suppression IDs to primary GO- identifiers

## v1.19.5

- Update Go to 1.26.2
- Update bborbe/* deps (errors, http, log, run, collection, parse, time, validation)
- Update third-party deps (moby/buildkit, otel, docker/cli, go-git, sentry, etc.)
- Re-enable lint in make check target
- Add new osv-scanner/trivy ignores for known indirect vulns

## v1.19.4

- Update go-git/go-git to v5.17.1 (fix security vulnerabilities)

## v1.19.3

- Update bborbe/* dependencies (errors, http, log, run, sentry, time, math)
- Update golangci-lint v2.11.4 and osv-scanner v2.3.5
- Update docker, containerd, moby, and opencontainers deps
- Add .osv-scanner.toml with vulnerability ignores for docker indirect deps
- Clean up go.mod: remove exclude blocks, update replace directive

## v1.19.2

- chore: verify project health — all tests pass, linting clean, precommit exits 0

## v1.19.1

- chore: verify project health — all tests pass, linting clean, no vulnerabilities

## v1.18.7

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- use go mod tidy -e for transitive dep compatibility

## v1.18.6

- upgrade golangci-lint from v1 to v2
- add trivy ghcr.io db-repository
- update bborbe/errors to v1.5.5
- update bborbe/run to v1.9.8
- update bborbe/log to v1.6.5

## v1.18.5

- go mod update

## v1.18.4

- go mod update

## v1.18.3

- Update Go to 1.26.0

## v1.18.2

- Update Go to 1.25.7
- Update github.com/bborbe dependencies
- Update testing dependencies (ginkgo, gomega)
- Update google/osv-scanner to v2.3.2
- Add .update-logs/ and .mcp-* to .gitignore

## v1.18.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.18.0

- update go and deps

## v1.17.0

- Refactor error variables to follow Go ErrFoo naming convention
- Add ErrTransactionAlreadyOpen, ErrBucketNotFound, ErrBucketAlreadyExists with backward-compatible deprecation
- Rename StoreStream/StoreList interfaces to StoreStreamer/StoreLister with backward-compatible type aliases
- Update Go version from 1.25.2 to 1.25.4
- Update dependencies (github.com/bborbe/http, github.com/bborbe/log, github.com/bborbe/run, github.com/onsi/ginkgo/v2)
- Enhance golangci-lint configuration with additional linters (errname, unparam, bodyclose, forcetypeassert, asasalint, prealloc, nestif)
- Add exclusion rules for deprecated error variables in linter configuration
- Fix deprecated Go stdlib usage (io/ioutil) in depguard rules

## v1.16.1

- Add comprehensive test suite for benchmark package (26 tests, 73% coverage)
- Add tests for RandString and ShuffleSlice utility functions
- Add tests for Benchmark core functionality with mock validation
- Add tests for HTTP handler with parameter parsing
- Update github.com/bborbe/errors from v1.3.0 to v1.3.1
- Update github.com/bborbe/run from v1.7.7 to v1.8.0 (adds FuncRunner interface and mock)
- Update .gitignore patterns for coverage output files

## v1.16.0

- Add golangci-lint configuration and integration
- Update Makefile with new lint target and improved formatting
- Integrate golines for automatic line length formatting (max 100 chars)
- Update goimports-reviser to v3
- Update github.com/bborbe/http from v1.14.2 to v1.15.2
- Update github.com/onsi/ginkgo/v2 from v2.25.3 to v2.26.0
- Apply code formatting improvements across codebase
- Improve error checking exclusions in Makefile

## v1.15.3

- Update Go version to 1.25.2
- Update CI workflow to use Go 1.25.2

## v1.15.2

- Upgrade osv-scanner from v1 to v2
- Add config file support for osv-scanner in Makefile
- go mod update

## v1.15.1

- go mod update

## v1.15.0

- Add StoreList interface with List method for retrieving all objects as slice
- Add StoreListTx interface for transaction-based list operations
- Implement List methods in Store and StoreTx concrete implementations
- Add comprehensive test coverage for new List functionality
- Fix security warning in benchmark random data generator

## v1.14.4

- Add comprehensive GoDoc documentation for all exported interfaces, types, and functions
- Improve API documentation coverage for better developer experience
- Document bucket operations, store interfaces, transaction handling, and metrics
- Add usage examples and parameter descriptions to public APIs

## v1.14.3

- improve README with usage example and installation instructions
- go mod update

## v1.14.2

- add github workflow
- go mod update

## v1.14.1

- add tests
- go mod update

## v1.14.0

- add RunnableTx and FuncTx

## v1.13.2

- add lock to reset handlers and improve logging
- go mod update

## v1.13.1

- add NewStoreFromTx
- go mod update

## v1.13.0

- add DBWithMetrics
- remove vendor files
- go mod update

## v1.12.2

- go mod update

## v1.12.1

- go mod update
- add test for relation store mocks

## v1.12.0

- add Invert for the RelationStore and RelationStoreTx

## v1.11.5

- add MapIDRelations and MapRelationIDs

## v1.11.4

- remove performance bug in relationStoreTx delete

## v1.11.3

- move JsonHandlerTx to github.com/bborbe/http
- go mod update

## v1.11.2

- add missing license file
- go mod update

## v1.11.1

- rename NewUpdateHandlerViewTx -> NewJsonHandlerUpdateTx

## v1.11.0

- add JsonHandlerTx
- go mod update

## v1.10.0

- add ListBucketNames
- go mod update

## v1.9.1

- ignore BucketNotFoundError on Map, Remove and Exists
- go mod update

## v1.9.0

- add remove to DB to delete the complete database
- add handler for reset bucket and complete database
- go mod update

## v1.8.2

- fix replace in relationStore
- go mod update

## v1.8.1

- add simple benchmark

## v1.8.0

- add relation store
- go mod update

## v1.7.0

- expect same tx returns same bucket
- go mod update

## v1.6.0

- add stream and exists to store

## v1.5.0

- add JSON store
- go mod update

## v1.4.2

- add KeyNotFoundError

## v1.4.1

- add mocks

## v1.4.0

- expect error if transaction open second transaction

## v1.3.1

- improve iterator testsuite

## v1.3.0

- add bucket testsuite

## v1.2.0

- add provider
- improve testsuite

## v1.1.1

- add test for iterator seek not found

## v1.1.0

- Add context to update and view

## v1.0.0

- Initial Version
