# Changelog

All notable changes to this project will be documented in this file.

## v1.4.8

- chore: update github.com/bborbe/collection to v1.20.25

## v1.4.7

- chore: update go module dependencies

## v1.4.6

- chore: update github.com/bborbe/collection to v1.20.24

## v1.4.5

- chore: update go module dependencies

## v1.4.4

- chore: update Go to 1.27.0 and github.com/bborbe/collection to v1.20.23, github.com/onsi/ginkgo/v2 to v2.32.1

## v1.3.21

- chore: Reorder format target so gofmt -w runs last after golines (Go 1.27 tooling compatibility)
- chore: Bump golangci-lint to v2.13.1 (fixes staticcheck buildir panic on Go 1.27 AST)
- chore: Bump errcheck to v1.20.0 (fixes `package "context" without types` on Go 1.27)

## v1.3.20

- update dependencies: github.com/bborbe/collection to v1.20.21, github.com/bborbe/errors to v1.5.18, github.com/bborbe/run to v1.9.35

## v1.3.19

- update Go to 1.26.6 and dependencies (fixes GO-2026-5972, GO-2026-6090, CVE-2026-56864, CVE-2026-56865)

## v1.3.18

- docs: add a License section to the README

## v1.3.17

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.3.16

- bump Go toolchain to 1.26.5

## v1.3.15

- Update github.com/bborbe/collection to v1.20.17
- Update github.com/bborbe/run to v1.9.30

## v1.3.14

- Bump github.com/bborbe/collection to v1.20.16
- Bump github.com/bborbe/run to v1.9.29
- Bump github.com/getsentry/sentry-go to v0.47.0

## v1.3.13

- Bump github.com/bborbe/collection to v1.20.15
- Bump github.com/bborbe/errors to v1.5.15
- Bump github.com/bborbe/run to v1.9.28

## v1.3.12

- Bump github.com/bborbe/collection to v1.20.13
- Bump test dependencies: ginkgo v2.32.0, gomega v1.42.1
- Bump indirect dependencies: run, sentry-go, golang.org/x/*

## v1.3.11

- bump go toolchain to 1.26.4
- bump bborbe/collection, bborbe/errors, onsi/ginkgo, onsi/gomega
- bump golang.org/x/net, x/sys, x/text for security fixes
- drop standalone errcheck + gosec; move config into golangci.yml
- add cloud.google.com/go v0.26.0 exclude

## v1.3.10

- bump Go toolchain from 1.26.2 to 1.26.3

## v1.3.9

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 447 lines to 38 lines.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.3.8

- bump go 1.26.1 → 1.26.2
- bump bborbe/collection, golangci-lint, osv-scanner, counterfeiter, go-modtool
- bump fatih/color, sqlclosecheck, sonatard/noctx and other indirect deps
- add vuln ignore rules for GO-2026-4923, bbolt, aws-sdk-go-v2

## v1.3.7

- Update multiple indirect dependencies (docker, containerd, moby, opentelemetry)
- Upgrade go-git, klauspost/compress, opencontainers packages
- Replace k8s.io/kube-openapi replace directive with charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Remove large exclude block for k8s and other packages

## v1.3.6

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations
- setup dark-factory config

## v1.3.5

- upgrade golangci-lint from v1 to v2
- update bborbe/collection to v1.20.5

## v1.3.4

- go mod update

## v1.3.3

- Update Go to 1.26.0

## v1.3.2

- Update Go to 1.25.7
- Update github.com/google/osv-scanner/v2 to v2.3.2
- Update github.com/onsi/ginkgo/v2 to v2.28.1
- Update github.com/onsi/gomega to v1.39.1
- Update various indirect dependencies for security and stability

## v1.3.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.3.0

- update go and deps

## v1.2.1
- Add GitHub Actions workflows (CI, Claude Code Review, Claude)
- Add golangci-lint configuration
- Enhance Makefile with additional quality checks (lint, gosec, trivy, osv-scanner)
- Update dependencies and Go version to 1.25.2

## v1.2.0

- remove vendor
- go mod update
 
## v1.1.1

- add license
- go mod update

## v1.1.0

- add Sqrt
- go mod update

## v1.0.0

- Initial Version
