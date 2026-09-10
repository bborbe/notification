# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.


## v1.6.0

- feat: opt into `autoMerge.trivial` for mechanically-trivial update PRs

## v1.5.21

- chore: update Go to 1.27.0

## v1.5.20

- chore: Run gofmt last in `format` target so golines wrapping is normalized before the gofmt lint check
- chore: Bump golangci-lint to v2.13.1 and errcheck to v1.20.0 for Go 1.27 toolchain compatibility

## v1.5.19

- bump ginkgo to v2.32.1
- update go.sum

## v1.5.18

- chore(security): bump Go 1.26.5 -> 1.26.6 (stdlib GO-2026-5026 / GO-2026-5972 / GO-2026-6090)

- chore: opt into `goUpdate.autoUpdate` in `.maintainer.yaml` so github-update-go-watcher may file Go-version update tasks for this repo

- chore(security): bump `golang.org/x/mod` v0.37.0 -> v0.40.0 (GO-2026-6179 / GO-2026-6180, CVE-2026-56864 / CVE-2026-56865)

## v1.5.17

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.5.16

- Bump Go toolchain to 1.26.5

## v1.5.15

- bump ginkgo/gomega test deps
- bump golang.org/x/mod, x/net, x/sync, x/sys, x/text, x/tools

## v1.5.14

- bump go 1.26.3 → 1.26.4
- bump ginkgo v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0
- bump golang.org/x/net, x/sys, x/text for vuln fixes
- drop errcheck + gosec; improve vulncheck output formatting
- add .maintainer.yaml for autoRelease/autoApprove

## v1.5.13

- bump Go version to 1.26.3

## v1.5.12

- Remove replace directives for charmbracelet/x/cellbuf, denis-tingaikin/go-header, diskfs/go-diskfs, nunnatsa/ginkgolinter

## v1.5.11

- Bump ginkgo/v2 to v2.28.3 and gomega to v1.40.0
- Update golang.org/x/{mod,net,sys,text,tools} dependencies
- Remove tools.go dev tool imports file
- Clean up go.mod by removing dev tool dependencies

## v1.5.10

- bump go 1.26.1 → 1.26.2
- upgrade golangci-lint v2.11.3 → v2.11.4
- upgrade osv-scanner v2.3.4 → v2.3.5
- upgrade counterfeiter v6.12.1 → v6.12.2, go-modtool v0.6.0 → v0.7.1
- add OSV vulnerability ignores for bbolt and aws-sdk-go-v2

## v1.5.9

- Update indirect dependencies (docker, containerd, prometheus, opentelemetry)
- Add replace directives for charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Bump moby/buildkit v0.23.2 → v0.29.0, docker/docker v28.3.3 → v28.5.2
- Update go-git, klauspost/compress, prometheus stack

## v1.5.8

- fix incompatible charmbracelet/x/cellbuf dependency version

## v1.5.7

- chore: verify project health — all tests pass, linting clean, precommit succeeds

## v1.5.6

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- setup dark-factory config

## v1.5.5

- upgrade golangci-lint from v1 to v2
- update vulnerable deps (go-sdk, grpc)

## v1.5.4

- go mod update

## v1.5.3

- Update Go version from 1.25.7 to 1.26.0
- Update google/osv-scanner from v2.3.2 to v2.3.3
- Update securego/gosec from v2.22.11 to v2.23.0
- Update various indirect dependencies including anthropic-sdk-go, openai-go, and golang.org/x/* packages

## v1.5.2

- Update Go from 1.25.5 to 1.25.6
- Update ginkgo/v2 from 2.27.5 to 2.28.1
- Update gomega from 1.39.0 to 1.39.1
- Update osv-scanner/v2 and related security dependencies

## v1.5.1

- Update Go to 1.25.5
- Update dependencies
## v1.5.0

**Breaking Changes:**
- Change `HasData` interface from `map[string]string` to `map[string]any`
- Change `AddDataToError` parameter type from `map[string]string` to `map[string]any`
- Change `DataFromError` return type from `map[string]string` to `map[string]any`
- Change `AddToContext` value parameter from `string` to `any`
- Change `DataFromContext` return type from `map[string]string` to `map[string]any`

**Features:**
- Support rich error details including arrays, numbers, booleans, and nested objects
- Enable JSON error handlers to return structured data instead of comma-separated strings

## v1.4.0

- update go and deps

## v1.3.1

- Add GitHub Actions workflows for CI, code review, and Claude integration
- Add comprehensive test files for error handling patterns
- Add golangci-lint configuration
- Update dependencies
- Remove vendor directory
- Improve Makefile

## v1.3.0

- refactor
- errors.As unwrap error if not matching
- go mod update

## v1.2.0

-add errors is

## v1.1.1

-add errors is

## v1.1.0

- add errors wrap

## v1.0.0

- Initial Version
