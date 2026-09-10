# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.4.23

- chore: update github.com/bborbe/errors to v1.5.21

## v1.4.22

- chore: update Go to 1.27.0 and update dependencies

## v1.4.21

- chore: Run gofmt last in the `format` target so golines' wrapping is normalized before the gofmt lint check
- chore: Bump golangci-lint to v2.13.1 (fixes staticcheck `buildir` panic on Go 1.27 AST)
- chore: Bump errcheck to v1.20.0 (fixes missing `context` package types on Go 1.27)
- fix: Use `reflect.Pointer` instead of deprecated `reflect.Ptr` in Nil validation (new govet inline check under golangci-lint v2.13.1)

## v1.4.20

- Update github.com/bborbe/errors to v1.5.18
- Update Ginkgo to v2.32.1
- Enable automatic Go dependency updates

## v1.4.19

- update Go to 1.26.6 and update dependencies

## v1.4.18

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.4.17

- bump Go toolchain version to 1.26.5

## v1.4.16

- Bump github.com/bborbe/errors to v1.5.15

## v1.4.15

- Bump github.com/onsi/ginkgo/v2 from v2.29.0 to v2.32.0
- Bump github.com/onsi/gomega from v1.41.0 to v1.42.1
- Bump golang.org/x/* indirect dependencies

## v1.4.14

- bump go 1.26.3 → 1.26.4
- bump bborbe/errors v1.5.12 → v1.5.13
- bump ginkgo v2.28.3 → v2.29.0, gomega v1.40.0 → v1.41.0
- bump golang.org/x/net, sys, text to latest
- drop standalone errcheck/gosec; move config into golangci.yml

## v1.4.13

- bump go 1.26.2 -> 1.26.3
- bump bborbe/errors v1.5.11 -> v1.5.12

## v1.4.12

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 441 to 24 lines.

## v1.4.11

- update dependencies (errors, golangci-lint, osv-scanner, counterfeiter)
- bump Go toolchain to 1.26.2
- add vuln ignores for bbolt and aws-sdk-go-v2
- filter known CVEs in trivyignore
- improve vulncheck to filter known findings via jq

## v1.4.10

- bump dependency versions (docker, containerd, prometheus, otel, moby)
- add replace directives for charmbracelet/x/cellbuf, denis-tingaikin/go-header, nunnatsa/ginkgolinter, opencontainers/runtime-spec
- update golang.org/x/exp, go-git, klauspost/compress

## v1.4.9

- fix charmbracelet/x/cellbuf dependency incompatibility with ansi package

## v1.4.8

- chore: verify project health, all precommit checks pass with no issues

## v1.4.7

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- setup dark-factory config

## v1.4.6

- upgrade golangci-lint from v1 to v2
- update bborbe/errors to v1.5.5

## v1.4.5

- go mod update

## v1.4.4

- Update Go to 1.26.0

## v1.4.3

- Update Go from 1.25.5 to 1.25.7
- Update github.com/bborbe/errors from v1.5.1 to v1.5.2
- Update github.com/onsi/ginkgo/v2 from v2.27.5 to v2.28.1
- Update github.com/onsi/gomega from v1.39.0 to v1.39.1
- Update numerous indirect dependencies

## v1.4.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.4.0

- update go and deps

## v1.3.3

- Update Go version to 1.25.3

## v1.3.2

- Update Go version to 1.25.2
- Add golangci-lint configuration file (.golangci.yml)
- Enhance Makefile with additional quality checks (lint, gosec, trivy, osv-scanner)
- Add Trivy security scanner to GitHub Actions CI workflow
- Update development dependencies (golangci-lint, gosec, osv-scanner)
- Improve format target to handle vendor directories correctly
- Update copyright year to 2025

## v1.3.1

- Add comprehensive Go doc comments to all exported functions and types
- Create package documentation in doc.go with usage examples
- Enhance README.md with detailed API documentation and usage examples
- Improve .gitignore with comprehensive Go-specific entries
- Apply Go doc best practices throughout the codebase

## v1.3.0

- remove vendor
- go mod update

## v1.2.0

- add NotEmptySlice validation
- add NotEmptyString validation 

## v1.1.0

- add slice length validations
- go mod update

## v1.0.0

- Initial Version
