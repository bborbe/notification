# Changelog

All notable changes to this project will be documented in this file.

## v1.11.15

- chore: update github.com/bborbe/collection to v1.20.25, github.com/bborbe/errors to v1.6.0, github.com/bborbe/kv to v1.21.12, github.com/onsi/gomega to v1.43.0

## v1.11.14

- chore: update github.com/bborbe/collection to v1.20.24

## v1.11.13

- chore: update Go to 1.27.0 and github.com/bborbe/collection to v1.20.23, github.com/bborbe/errors to v1.5.20, github.com/bborbe/kv to v1.21.11

## v1.11.12

- chore: Make `format` run golines before `gofmt -w` and bump golangci-lint to v2.13.1 + errcheck to v1.20.0 for Go 1.27 tooling compatibility

## v1.11.11

- chore: update dependencies (bborbe/collection v1.20.21, bborbe/errors v1.5.18, bborbe/kv v1.21.10, bborbe/run v1.9.35)

## v1.11.10

- update Go to 1.26.6 and update dependencies, fixing GO-2026-5972, GO-2026-6090, GO-2026-6179, GO-2026-6180, CVE-2026-56864, CVE-2026-56865

## v1.11.9
- fix: add `ctx.Done()` guards to two unbounded I/O loops. `tx.DeleteBucket` iterates and deletes **every key in a bucket** one at a time, and `statsImpl` walks every bucket calling `tx.Bucket(ctx, name)` per entry. Both had `ctx` in scope and neither could be interrupted.

## v1.11.8

- chore: add the missing root `LICENSE` (BSD-3-Clause, matching the rest of the fleet). The repo is public and had no license file.

## v1.11.7

- chore: update Go to 1.26.5 and update dependencies, fixing GO-2026-5841

## v1.11.6

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.11.5

- Bump github.com/dgraph-io/badger/v4 to v4.9.4
- Bump github.com/bborbe/kv to v1.21.5
- Bump github.com/bborbe/errors to v1.5.16
- Bump github.com/bborbe/collection to v1.20.17
- Bump go toolchain to 1.26.5

## v1.11.4

- Bump github.com/bborbe/collection to v1.20.16

## v1.11.3

- Bump bborbe/collection to v1.20.15
- Bump bborbe/errors to v1.5.15
- Bump bborbe/kv to v1.21.4
- Bump bborbe/run and getsentry/sentry-go

## v1.11.2

- bump go to 1.26.4
- bump bborbe/kv to v1.21.2
- bump badger/v4 to v4.9.2
- bump all remaining dependencies
- add exclude directive for cloud.google.com/go v0.26.0

## v1.11.1

- bump go 1.26.3 → 1.26.4
- bump golang.org/x/net v0.53.0 → v0.55.0, x/sys v0.43.0 → v0.45.0 (vuln fixes)
- bump bborbe/collection v1.20.12 → v1.20.13, bborbe/run v1.9.24 → v1.9.27
- drop standalone errcheck/gosec targets; move config into golangci.yml
- add cloud.google.com/go v0.26.0 exclude

## v1.11.0

- **BREAKING**: `Stats(ctx)` now returns `*libkv.Stats` instead of `libkv.Stats` (matches bborbe/kv v1.21.0 interface)
- Fast `Stats(ctx)` now lists bucket NAMES only (no per-prefix key scan)
- Implement `StatsDetailed(ctx) (*libkv.Stats, error)` — adds per-bucket `KeyCount` via `kv.Count`; O(total keys), do not poll hot
- Bump bborbe/kv v1.20.0 → v1.21.1

## v1.10.0

- implement `Stats(ctx) (Stats, error)` to satisfy bborbe/kv v1.20.0 `DB` interface; uses `db.Size()` for total (LSM + value-log) and `ListBucketNames` + `kv.Count` per bucket — O(n), do not poll hot
- bump bborbe/kv v1.19.7 → v1.20.0
- bump ginkgo v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0

## v1.9.13

- bump go 1.26.2 → 1.26.3
- bump bborbe/collection, errors, kv, run
- bump getsentry/sentry-go v0.46.2
- bump opentelemetry otel v1.43.0

## v1.9.12

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 452 to 50 lines.

## v1.9.11

- bump bborbe/collection, errors, kv dependencies
- bump go 1.26.1 → 1.26.2
- bump indirect deps (otel, moby/buildkit, docker/cli, etc.)
- add vuln ignores for bbolt/aws-sdk CVEs
- update vulncheck to filter known false positives

## v1.9.10

- Update go-git/go-git to v5.17.1 (fix security vulnerabilities)

## v1.9.9

- bump bborbe/errors, bborbe/kv, bborbe/run dependencies
- bump golangci-lint v2.11.4, osv-scanner v2.3.5
- bump docker, containerd, moby toolchain deps
- add runtime-spec replace directive for opencontainers/runtime-spec

## v1.9.8

- Update bborbe/collection to v1.20.7, bborbe/errors to v1.5.7, bborbe/kv to v1.19.2
- Update shoenig/go-modtool to v0.6.0
- Update bbolt to v1.4.3, go-yaml/v3 to v3.0.4
- Remove replace/exclude directives from go.mod

## v1.9.7

- chore: enable golangci-lint in Makefile check target and update .golangci.yml to standard config with nestif, errname, unparam, bodyclose, forcetypeassert, asasalint, prealloc linters
- refactor: extract runTx helper in badgerdb to eliminate dupl violation between Update and View
- fix: simplify bool comparisons and use bytes.Equal in badgerkv_tx.go to resolve staticcheck violations

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.9.6

- standardize Makefile: multiline trivy format

## v1.9.5

- chore: fix Go module cache corruption to restore passing precommit

## v1.9.4

- go mod update

## v1.9.3

- Update Go to 1.26.0

## v1.9.2

- Update Go to 1.25.7
- Update github.com/dgraph-io/badger/v4 to v4.9.1
- Update github.com/onsi/ginkgo/v2 to v2.28.1 and gomega to v1.39.1
- Update github.com/google/osv-scanner/v2 to v2.3.2
- Update numerous indirect dependencies

## v1.9.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.9.0

- update go and deps

## v1.8.4

- add golangci-lint configuration
- enhance CI with Trivy security scanning
- update Makefile with additional security tools (osv-scanner, gosec, trivy)
- update Go version to 1.25.2
- improve code formatting for long function signatures
- go mod update

## v1.8.3

- improve README with usage example and installation instructions
- go mod update

## v1.8.2

- add github workflow
- go mod update

## v1.8.1

- add tests
- go mod update

## v1.8.0

- OpenDB and OpenMemory return badgerkv.DB
- remove vendor
- go mod update

## v1.7.3

- go mod update

## v1.7.2

- go mod update

## v1.7.1

- fix ListBucketNames

## v1.7.0

- implement ListBucketNames
- go mod update

## v1.6.0

- add remove db files
- go mod update

## v1.5.2

- go mod update

## v1.5.1

- go mod update

## v1.5.0

- cache buckets per tx
- go mod update

## v1.4.2

- fix bucket name problem
- go mod update

## v1.4.1

- add interface to access bolt db, tx, bucket if needed

## v1.4.0

- prevent transaction open second transaction

## v1.3.0

- fulfill bucket testsuite

## v1.2.0

- use new testsuite

## v1.1.1

- update libkv

## v1.1.0

- Add context to update and view

## v1.0.0

- Initial Version
