# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.10.0

- feat: opt into `autoMerge.trivial` for mechanically-trivial update PRs

## v1.9.27

- chore: update Go to 1.27.0 and github.com/bborbe/errors to v1.5.20, github.com/bborbe/run to v1.9.37, github.com/getsentry/sentry-go to v0.48.0, github.com/onsi/ginkgo/v2 to v2.32.1

## v1.9.26

- chore: Bump golangci-lint to v2.13.1 and errcheck to v1.20.0; run gofmt last in format target for Go 1.27 toolchain compatibility

## v1.9.25

- chore(security): bump Go 1.26.5 -> 1.26.6 (stdlib GO-2026-5026 / GO-2026-5972 / GO-2026-6090)

- chore(security): bump `golang.org/x/mod` v0.37.0 -> v0.40.0 (GO-2026-6179 / GO-2026-6180, CVE-2026-56864 / CVE-2026-56865)

## v1.9.24

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.9.23

- Bump github.com/bborbe/errors to v1.5.16
- Bump Go toolchain to 1.26.5

## v1.9.22

- Bump github.com/bborbe/run to v1.9.30

## v1.9.21

- Bump github.com/bborbe/run to v1.9.29

## v1.9.20

- Bump github.com/bborbe/errors to v1.5.15

## v1.9.19

- Bump sentry-go to v0.47.0
- Bump ginkgo/v2 to v2.32.0 and gomega to v1.42.1
- Bump bborbe/run to v1.9.28 and indirect dependencies

## v1.9.18

- bump go 1.26.3 → 1.26.4
- bump bborbe/run v1.9.23 → v1.9.27
- bump ginkgo v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0
- bump golang.org/x/net, sys, text to latest
- add cloud.google.com/go v0.26.0 exclude

## v1.9.17

- bump go 1.26.2 → 1.26.3
- bump bborbe/errors v1.5.11 → v1.5.13
- bump sentry-go v0.46.1 → v0.46.2

## v1.9.16

- chore: update github.com/bborbe/errors to v1.5.10 and indirect dependencies
- chore: migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block; go.mod reduced from 443 to 38 lines

## v1.9.15

- Update bborbe/errors to v1.5.9, bborbe/run to v1.9.17
- Update getsentry/sentry-go to v0.45.0
- Update golangci-lint to v2.11.4, counterfeiter to v6.12.2
- Update golang.org/x/sys and other indirect deps
- Add new CVE ignores to osv-scanner and trivyignore

## v1.9.14

- Update multiple indirect dependencies (docker, containerd, moby, opentelemetry)
- Bump go-git, prometheus, klauspost/compress, and other transitive deps
- Remove stale exclude directives for k8s and other packages from go.mod
- Update replace directives for charmbracelet, denis-tingaikin, and opencontainers

## v1.9.13

- chore: verify project health — all tests pass, linting clean, precommit exits 0

## v1.9.12

- chore: verify project health — all tests pass, linting clean, precommit exits 0

## v1.9.11

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- setup dark-factory config

## v1.9.10

- upgrade golangci-lint from v1 to v2
- update bborbe/errors to v1.5.5
- update bborbe/run to v1.9.8

## v1.9.9

- go mod update

## v1.9.8

- Update Go to 1.26.0

## v1.9.7

- Update Go toolchain from 1.25.6 to 1.25.7
- Update github.com/bborbe/errors from v1.5.1 to v1.5.2

## v1.9.6

- reduce log verbosity for exception capture (V(2) → V(3))

## v1.9.5

- Update ginkgo/gomega test dependencies
- Update golang.org/x toolchain dependencies
- Update google/pprof profiling library

## v1.9.4

- refactor NewClient to reduce cognitive complexity by extracting helper functions
- replace interface{} with any
- use maps.Copy for map operations

## v1.9.3

- update Go to 1.25.6
- update github.com/getsentry/sentry-go to v0.42.0
- update github.com/google/osv-scanner/v2 to v2.3.2
- update google.golang.org/grpc to v1.78.0
- add .gitignore entries

## v1.9.2

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.9.1

- fix error tag values by converting all types to strings with fmt.Sprintf
- update dependencies (errors v1.4.0 → v1.5.0, run v1.8.3 → v1.9.0, ginkgo v2.27.2 → v2.27.3, gomega v1.38.2 → v1.38.3)
- update example to use map[string]any for error data

## v1.9.0

- update go and deps

## v1.8.5

- update Go version from 1.25.2 to 1.25.3
- update dependencies (Sentry SDK v0.35.3 → v0.36.0, bborbe/run v1.7.8 → v1.8.1, gosec v2.22.9 → v2.22.10)
- remove deprecated golang.org/x/lint dependency
- add comprehensive GoDoc documentation for Client interface and NewClient function
- add sensitive data warning to NewClient documentation
- enhance README with Development, Dependencies, and License sections
- improve error handling with nil checks for eventID in logging
- fix typos in logging messages (orginal → original, execption → exception)
- add defensive nil check for event.Tags initialization
- update copyright headers to include 2025
- update various indirect dependencies (crypto, net, sys, text, etc.)

## v1.8.4

- add golangci-lint configuration file
- enhance Makefile with additional linting and security scanning targets (lint, gosec, trivy, osv-scanner)
- update Go version to 1.25.2
- update dependencies (errors, run, ginkgo)
- add new development tools (golangci-lint, gosec, osv-scanner)
- improve GitHub Actions CI workflow with updated Go version
- streamline README by removing development sections

## v1.8.3

- add EventModifier interface for custom event processing
- add EventModifierFunc type for function-based event modifiers
- add EventModifierList type for sequential modifier application
- add comprehensive test coverage for event modifier functionality
- add detailed godoc documentation for all event modifier types

## v1.8.2

- enhance README.md with comprehensive documentation including status badges, features overview, installation instructions, and usage examples
- add detailed package-level GoDoc comment with usage examples and feature descriptions

## v1.8.1

- add LICENSE file  
- add GitHub workflows for CI and Claude Code integration
- update Go version from 1.24.1 to 1.24.5
- update dependencies (sentry-go, glog, counterfeiter, ginkgo, gomega)
- add golines tool for code formatting
- improve .gitignore patterns

## v1.8.0

- remove vendor
- go mod update

## v1.7.1

- add license
- go mod update 

## v1.7.0

- add data to tags

## v1.6.0

- allow exclude errors

## v1.5.0

- add example
- add tags from context and error

## v1.4.0

- add flush method

## v1.3.0

- simplify newClient

## v1.2.1

- rename proxy roundtripper

## v1.2.0

- add NewClientWithOptions
- add proxy roundTripper

## v1.1.0

- add skip error and report
- go mod update

## v1.0.0

- Initial Version
