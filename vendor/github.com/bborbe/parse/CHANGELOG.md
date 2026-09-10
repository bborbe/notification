# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.11.0

- feat: opt into `autoMerge.trivial` for mechanically-trivial update PRs

## v1.10.21

- chore: Run `gofmt -w` last in the `format` target so golines' wrapping is normalized before the gofmt lint check
- chore: Bump golangci-lint to v2.13.1 (fixes staticcheck `buildir` panic on Go 1.27 AST)
- chore: Bump errcheck to v1.20.0 (fixes `package "context" without types` on Go 1.27)
- fix: Replace deprecated `reflect.Ptr` with `reflect.Pointer` in parse-strings.go to satisfy the govet `inline` check on Go 1.27

## v1.10.20

- Bump Go toolchain to 1.26.6 and update dependencies
- Fixed vulnerabilities: GO-2026-5972, GO-2026-6090, CVE-2026-56864, CVE-2026-56865

## v1.10.19

- docs: add a License section to the README

## v1.10.18

- Bump bborbe/errors to v1.5.16
- Bump bborbe/math to v1.3.15
- Bump golang.org/x/text to v0.40.0
- Bump other transitive dependencies
- Bump Go toolchain to 1.26.5

## v1.10.17

- Bump github.com/bborbe/math to v1.3.13
- Bump github.com/bborbe/collection to v1.20.15
- Bump github.com/bborbe/run to v1.9.28

## v1.10.16

- Bump bborbe/math to v1.3.12
- Bump bborbe/collection to v1.20.13
- Bump bborbe/run to v1.9.27
- Bump getsentry/sentry-go to v0.46.2

## v1.10.15

- Bump github.com/bborbe/errors from v1.5.13 to v1.5.15

## v1.10.14

- Bump github.com/bborbe/math from v1.3.10 to v1.3.11
- Bump github.com/onsi/ginkgo/v2 from v2.29.0 to v2.32.0
- Bump github.com/onsi/gomega from v1.41.0 to v1.42.1
- Bump golang.org/x/text from v0.37.0 to v0.38.0

## v1.10.13

- bump go 1.26.3 → 1.26.4
- bump bborbe/math v1.3.9 → v1.3.10, ginkgo v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0
- bump golang.org/x/text v0.36.0 → v0.37.0, x/net v0.53.0 → v0.55.0, x/sys v0.43.0 → v0.45.0
- drop standalone errcheck/gosec targets; move config into .golangci.yml
- add .maintainer.yaml; set autoRelease=false in .dark-factory.yaml

## v1.10.12

- bump go 1.26.2 → 1.26.3
- bump github.com/bborbe/errors v1.5.11 → v1.5.13

## v1.10.11

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 447 lines to ~39 lines.

## v1.10.10

- Update go runtime to 1.26.2
- Bump bborbe/errors to v1.5.9, bborbe/math to v1.3.7
- Update golangci-lint to v2.11.4, osv-scanner to v2.3.5
- Update golang.org/x/text to v0.36.0
- Add new vuln ignores to osv-scanner and trivyignore

## v1.10.9

- Update multiple indirect dependencies (docker, containerd, moby, otel)
- Replace k8s.io/kube-openapi replace directive with charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Remove large exclude block for k8s and other packages
- Bump golang.org/x/exp, go-git, klauspost/compress, and yaml libraries

## v1.10.8

- chore: verify project health — all tests pass, linting clean, precommit exits 0

## v1.10.7

- chore: verify project health — all tests pass, linting clean, precommit exits 0

## v1.10.6

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- setup dark-factory config

## v1.10.5

- upgrade golangci-lint from v1 to v2
- update bborbe/errors to v1.5.5
- update bborbe/math to v1.3.5

## v1.10.4

- go mod update

## v1.10.3

- Update Go to 1.26.0

## v1.10.2

- Update Go to 1.25.7
- Update bborbe/errors to v1.5.2
- Update bborbe/math to v1.3.1
- Update testing frameworks (ginkgo v2.28.1, gomega v1.39.1)
- Update numerous indirect dependencies

## v1.10.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.10.0

- update go and deps

## v1.9.1

- Rename InvalidTypeError to ErrInvalidType to follow Go error naming conventions

## v1.9.0

- Add HasStrings and HasString interfaces for custom string conversion
- Add support for parsing slices of string subtypes via reflection
- Add support for parsing slices of types implementing String() method
- Improve Makefile with .PHONY declarations for all targets
- Enable race detection in test target
- Add go-modtool for go.mod formatting
- Remove deprecated golang.org/x/lint/golint dependency
- Enhance .golangci.yml with 8 additional linters (funlen, gocognit, nestif, maintidx, errname, unparam, bodyclose, forcetypeassert)
- Fix typo in .golangci.yml depguard rule
- Add standard ignore patterns to .gitignore

## v1.8.3

- Upgrade osv-scanner to v2.2.3 with configuration file support
- Enhance Makefile osv-scanner target to support optional .osv-scanner.toml config
- Update development tool dependencies

## v1.8.2

- Update Go version to 1.25.2
- Update CI workflow to use Go 1.25.2
- Update dependencies (golang.org/x/text, golang.org/x/crypto, golang.org/x/net, golang.org/x/sys, golang.org/x/term, golang.org/x/exp)
- Update prometheus client libraries
- Update internal dependencies (bborbe/collection, bborbe/run)
- Add exclude directive for incompatible golang.org/x/tools v0.38.0

## v1.8.1

- Enhance README.md with comprehensive documentation
- Add installation instructions and usage examples
- Add status badges for CI, coverage, and Go Report Card
- Add GoDoc comments to all exported functions
- Improve API documentation for pkg.go.dev

## v1.8.0

- Add ParseASCII function with correct capitalization
- Add golangci-lint configuration with revive linter
- Add GitHub Actions workflows for CI/CD and code review
- Enhance Makefile with lint target
- Update dependencies and test coverage
- Deprecate ParseAscii in favor of ParseASCII

## v1.7.1

- add tests
- go mod update

## v1.7.0

- remove vendor
- go mod update

## v1.6.1

- add license
- go mod update

## v1.6.0

- allow parse subtypes 

## v1.5.0

- add parse int64 array
- go mod update

## v1.4.2

- add InvalidTypeError to ParseString

## v1.4.1

- go mod update
- improve test

## v1.4.0

- add ParseAscii

## v1.3.2

- improve ParseStrings
- go mod update

## v1.3.1

- allow more types for ParseString
- go mod update

## v1.3.0

- add parse with default
- go mod update

## v1.2.0

- allow parse string to bool

## v1.1.0

- allow stringer in parseInt

## v1.0.1

- go mod update

## v1.0.0

- Initial Version
