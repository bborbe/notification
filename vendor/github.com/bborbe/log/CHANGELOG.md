# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.6.25

- chore: update go module dependencies

## v1.6.24

- chore: update Go to 1.26.6 and update dependencies

## v1.6.23

- chore: Reorder format target so gofmt -w runs last (after golines) so its wrapping is normalized before the gofmt lint check passes
- chore: Bump golangci-lint to v2.13.1 (fixes staticcheck `buildir` panic on Go 1.27 AST)
- chore: Bump errcheck to v1.20.0 (fixes `package "context" without types` on Go 1.27)

## v1.6.22

- update Go to 1.26.6 and update dependencies (GO-2026-5972, GO-2026-6090, CVE-2026-56864, CVE-2026-56865)

## v1.6.21

- fix: `memoryMonitor.LogMemoryUsage`, `logLevelSetter.Set` and `logLevelSetter.resetLogLevel` now use `libtime` instead of stdlib `time`, matching the convention `log_sampler-time.go` already documents ("uses github.com/bborbe/time for consistent time handling across the library"). All three drive time-threshold logic — log-if-interval-elapsed and auto-reset-after-duration — which could not be exercised deterministically against a frozen clock. The reset check uses `libtime.Now().Sub(lastSetTime)` rather than `time.Since(lastSetTime)`, so both the write and the read of `lastSetTime` come from the same clock.

## v1.6.20

- update Go to 1.26.5 and update dependencies

## v1.6.19

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.6.18

- Bump go toolchain to 1.26.5
- Update bborbe/time, collection, math, parse, run, validation deps
- Update getsentry/sentry-go to v0.47.0

## v1.6.17

- Bump github.com/bborbe/time to v1.27.4
- Bump github.com/bborbe/math to v1.3.11
- Bump github.com/bborbe/parse to v1.10.15
- Bump github.com/bborbe/validation to v1.4.15

## v1.6.16

- Bump github.com/bborbe/time to v1.27.3
- Bump github.com/bborbe/collection, errors, math, parse, run, validation deps
- Bump github.com/getsentry/sentry-go to v0.46.2

## v1.6.15

- bump github.com/bborbe/time to v1.27.1
- bump github.com/onsi/ginkgo/v2 to v2.32.0, github.com/onsi/gomega to v1.42.1
- bump golang.org/x/* indirect dependencies

## v1.6.14

- bump go toolchain to 1.26.4
- update bborbe/* deps (time, collection, errors, parse, validation)
- update ginkgo/gomega and golang.org/x/* for security fixes
- drop standalone errcheck/gosec; move checks into golangci-yml
- add .maintainer.yaml; set autoRelease=false in dark-factory

## v1.6.13

- bump go toolchain to 1.26.3

## v1.6.12

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 453 to ~43 lines.

## v1.6.11

- bump Go toolchain to 1.26.2
- update bborbe/* dependencies (time, collection, errors, parse, validation)
- update opentelemetry to v1.40.0
- update moby/buildkit, docker/cli, containerd deps
- add vuln/osv/trivy ignore entries for known indirect CVEs

## v1.6.10

- pin charmbracelet/x/cellbuf to v0.0.15 in go.mod

## v1.6.9

- Update dependencies: golangci-lint, osv-scanner, bborbe/* libs
- Update docker/moby deps: docker v28.5.2, buildkit v0.28.1
- Add --allow-parallel-runners flag to golangci-lint Makefile target
- Enable autoRelease in dark-factory config
- Clean up go.mod: remove exclude blocks, update replace directives

## v1.6.8

- chore: confirm project health — all tests pass, linting clean, no vulnerabilities found

## v1.6.7

- chore: verify project health — all tests pass, linting clean, no vulnerabilities found

## v1.6.6

- standardize Makefile: add mocks mkdir, reorder lint, multiline trivy, add .PHONY declarations

## v1.6.5

- upgrade golangci-lint from v1 to v2
- update bborbe/time to v1.25.1

## v1.6.4

- go mod update

## v1.6.3

- Update Go version to 1.26.0
- Fix race condition in log level reset mechanism
- Update dependencies (osv-scanner, gosec, bborbe packages)
- Add concurrent safety test for log level setter
- Enable linting in CI checks

## v1.6.2

- Update Go to 1.25.6
- Update dependencies (ginkgo v2.28.1, gomega v1.39.1, osv-scanner v2.3.2)
- Update indirect dependencies (BurntSushi/toml v1.6.0, grpc v1.78.0, protobuf v1.36.11)
- Add .update-logs/ and .mcp-* to .gitignore

## v1.6.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.6.0

- update go and deps

## v1.5.0
- Add comprehensive package-level documentation (doc.go) with examples and usage guidance
- Add Ginkgo v2 CLI to development tools for better test execution
- Enhance README with comprehensive Testing section showing multiple testing approaches
- Add Full Example section to README demonstrating production-like usage
- Update CI badge to new GitHub Actions syntax
- Remove deprecated golint tool (replaced by golangci-lint)
- Add horizontal rules to README for better visual section separation

## v1.4.4
- Update dependencies (github.com/bborbe/time v1.20.0, github.com/securego/gosec/v2 v2.22.10, and 19 indirect dependencies)
- Add exclusion for golang.org/x/tools v0.38.0 due to counterfeiter compatibility

## v1.4.3
- Update Go version from 1.25.2 to 1.25.3

## v1.4.2
- Fix integer overflow vulnerability in log level handler (G115)
- Add golangci-lint configuration and security scanning tools
- Update Go version from 1.24.5 to 1.25.2
- Add osv-scanner, gosec, and trivy security checks to Makefile
- Update CI workflow with Trivy installation
- Update dependencies and tooling

## v1.4.1

- Add LogMemoryUsagef method for formatted memory logging
- Refactor memory stats logging to reduce code duplication
- Extract MemoryStats utility function to separate file
- Remove trading-specific comments from memory monitor

## v1.4.0

- Add MemoryMonitor interface for runtime memory usage monitoring
- Implement memory monitoring with configurable logging intervals
- Add memory usage logging at start/end with garbage collection metrics
- Generate MemoryMonitor mock for testing

## v1.3.0

- Add comprehensive GoDoc documentation to all exported functions and types
- Enhance README.md with detailed usage examples, API documentation, and status badges
- Add GitHub Actions CI/CD workflow with automated testing and coverage reporting
- Improve .gitignore with comprehensive Go development patterns
- Add status badges for CI, Go Report Card, pkg.go.dev reference, and code coverage
- Enhance project structure following Go library best practices

## v1.2.1

- go mod update
- add tests

## v1.2.0

- generate mocks
- go mod update
- add tests

## v1.1.0

- remove vendor
- go mod update

## v1.0.1

- add license
- go mod update

## v1.0.0

- Initial Version
