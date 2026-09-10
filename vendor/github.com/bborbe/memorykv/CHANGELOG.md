# Changelog

All notable changes to this project will be documented in this file.

## v1.6.11

- chore: update go module dependencies

## v1.6.10

- chore: update github.com/bborbe/errors to v1.6.0, github.com/bborbe/kv to v1.21.12, github.com/onsi/gomega to v1.43.0, golang.org/x/exp to v0.0.0-20260824195058-e88cd73687aa

## v1.6.9

- chore: update Go to 1.27.0 and update dependencies
- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.6.8

- update Go to 1.26.6 and update dependencies (GO-2026-6179, GO-2026-6180, GO-2026-5972, GO-2026-6090, CVE-2026-56864, CVE-2026-56865)

## v1.6.7

- docs: add a License section to the README

## v1.6.6

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.6.5

- Bump go to 1.26.5
- Bump bborbe/errors to v1.5.16
- Bump bborbe/kv to v1.21.5

## v1.6.4

- Bump github.com/bborbe/kv from v1.21.2 to v1.21.4

## v1.6.3

- Bump github.com/bborbe/errors to v1.5.15

## v1.6.2

- bump github.com/bborbe/kv from v1.21.1 to v1.21.2
- bump github.com/onsi/ginkgo/v2 from v2.29.0 to v2.32.0
- bump github.com/onsi/gomega from v1.41.0 to v1.42.1

## v1.6.1

- bump go 1.26.3 → 1.26.4
- update golang.org/x/{exp,mod,net,sys,text,tools}
- drop kisielk/errcheck indirect dep
- exclude cloud.google.com/go v0.26.0

## v1.6.0

- **BREAKING**: `Stats(ctx)` now returns `*libkv.Stats` instead of `libkv.Stats` (matches bborbe/kv v1.21.0 interface)
- Implement `StatsDetailed(ctx) (*libkv.Stats, error)` — adds per-bucket `KeyCount`; `Stats` (fast) leaves it zero
- Bump bborbe/kv v1.20.0 → v1.21.1

## v1.5.0

- implement `Stats(ctx) (Stats, error)` to satisfy bborbe/kv v1.20.0 `DB` interface
- bump bborbe/kv v1.19.7 → v1.20.0
- bump ginkgo v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0

## v1.4.12

- bump bborbe/errors v1.5.11 → v1.5.13
- bump bborbe/kv v1.19.6 → v1.19.7
- bump golang.org/x/exp to 20260410
- bump go 1.26.2 → 1.26.3

## v1.4.11

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 450 to 38 lines.

## v1.4.10

- bump go 1.26.2
- update bborbe/errors v1.5.9, bborbe/kv v1.19.4
- update moby/buildkit v0.29.0, otel v1.40.0
- update docker/cli, go-git, klauspost/compress and other indirect deps
- add osv/trivy ignores for new CVEs

## v1.4.9

- Update go-git/go-git to v5.17.1 (fix security vulnerabilities)

## v1.4.8

- Update bborbe/errors to v1.5.8 and bborbe/kv to v1.19.3
- Update golangci-lint to v2.11.4 and osv-scanner to v2.3.5
- Update docker, moby/buildkit, containerd dependencies
- Remove large exclude block and update replace directive in go.mod
- Update multiple indirect dependencies

## v1.4.7

- chore: verify project health — all tests pass, linting clean, precommit succeeds at exit code 0

## v1.4.6

- chore: verify project health — all tests pass, linting clean, precommit succeeds

## v1.4.5

- upgrade golangci-lint from v1 to v2
- standardize Makefile: add .PHONY declarations, multiline trivy, mocks mkdir
- update .golangci.yml to v2 format
- setup dark-factory config

## v1.4.4

- go mod update

## v1.4.3

- Update Go to 1.26.0

## v1.4.2

- Update Go from 1.25.5 to 1.25.7
- Update github.com/bborbe/errors from v1.5.1 to v1.5.2
- Update github.com/bborbe/kv from v1.18.0 to v1.18.1
- Update testing dependencies (ginkgo v2.28.1, gomega v1.39.1)
- Update indirect dependencies and tooling

## v1.4.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.4.0

- update go and deps

## v1.3.6

- add golangci-lint configuration
- update CI workflow
- update Makefile with improved linting targets
- go mod update

## v1.3.5

- improve README with usage example and installation instructions
- go mod update

## v1.3.4

- add github workflow
- go mod update

## v1.3.1

- add tests
- go mod update

## v1.3.0

- remove vendor
- go mod update

## v1.2.3

- go mod update

## v1.2.2

- go mod update

## v1.2.1

- improve lock
- add license
- go mod update

## v1.2.0

- implement ListBucketNames
- go mod update

## v1.1.3

- go mod update

## v1.1.2

- go mod update

## v1.1.1

- go mod update

## v1.1.0

- cache buckets per tx
- go mod update

## v1.0.0

- Initial Version
