# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.14.17

- chore: update github.com/bborbe/collection to v1.20.25, github.com/bborbe/errors to v1.6.0, github.com/bborbe/time to v1.27.11, github.com/bborbe/validation to v1.4.23, github.com/onsi/gomega to v1.43.0, k8s.io/api to v0.37.0, k8s.io/apiextensions-apiserver to v0.37.0, k8s.io/apimachinery to v0.37.0, k8s.io/client-go to v0.37.0

## v1.14.16

- chore: update Go to 1.27.0 and github.com/bborbe/collection to v1.20.23, github.com/bborbe/errors to v1.5.20, github.com/bborbe/time to v1.27.10, github.com/bborbe/validation to v1.4.22, github.com/onsi/ginkgo/v2 to v2.32.1, k8s.io/api to v0.36.4, k8s.io/apiextensions-apiserver to v0.36.4, k8s.io/apimachinery to v0.36.4, k8s.io/client-go to v0.36.4

## v1.14.15

- chore: update github.com/bborbe/errors to v1.5.21

## v1.14.14

- chore: update github.com/bborbe/collection to v1.20.24

## v1.14.13

- chore: update Go to 1.27.0 and github.com/bborbe/collection to v1.20.23, github.com/bborbe/errors to v1.5.20, github.com/bborbe/time to v1.27.10, github.com/bborbe/validation to v1.4.22, github.com/onsi/ginkgo/v2 to v2.32.1, k8s.io/api to v0.36.4, k8s.io/apiextensions-apiserver to v0.36.4, k8s.io/apimachinery to v0.36.4, k8s.io/client-go to v0.36.4

## v1.14.12

- fix: StatefulSet deployer update path merges only mutable spec fields into the live object and preserves immutable fields (selector, serviceName, volumeClaimTemplates) — the previous full-spec Update was rejected by the API server with "updates to statefulset spec for fields other than ... are forbidden"

## v1.14.11

- chore: Fix make precommit on Go 1.27 — run gofmt last in the format target, bump golangci-lint to v2.13.1 and errcheck to v1.20.0
- fix: Check context cancellation in eventHandlerAlert add/delete so errors surface instead of always returning nil

## v1.14.10

- update Go to 1.26.6 and update dependencies (GO-2026-5026, GO-2026-5972, GO-2026-6090, GO-2026-6218)

## v1.14.9

- update Go to 1.26.5 and update dependencies

## v1.14.8

- fix: make the in-cluster CreateConfig spec environment-aware so the suite passes both inside and outside a Kubernetes pod

## v1.14.7

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.14.6

- Bump go toolchain to 1.26.5
- Update bborbe/collection, errors, time, math, parse, run deps

## v1.14.5

- Bump github.com/bborbe/collection to v1.20.16
- Bump github.com/bborbe/time to v1.27.4
- Bump github.com/bborbe/validation to v1.4.16
- Bump indirect deps (math, parse, run, sentry-go)

## v1.14.4

- Bump bborbe/collection, errors, time, validation deps
- Bump k8s.io libs to v0.36.2
- Bump ginkgo/gomega and golang.org/x deps

## v1.14.3

- bump k8s.io/{api,apimachinery,client-go,apiextensions-apiserver} to v0.36.1
- bump golang.org/x/{net,sys,term,text} for vulnerability fixes
- bump bborbe/time to v1.27.0, onsi/ginkgo v2.29.0, onsi/gomega v1.41.0
- drop standalone errcheck/gosec targets; move rules into golangci config
- add .maintainer.yaml for autoRelease/autoApprove

## v1.14.2

- bump bborbe/collection v1.20.11→v1.20.12
- bump bborbe/errors v1.5.11→v1.5.13
- bump bborbe/validation v1.4.12→v1.4.13
- bump go 1.26.2→1.26.3

## v1.14.1

- chore: Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 491 to 76 lines.

## v1.14.0

- Update k8s.io/api, apimachinery, client-go v0.35.3 → v0.36.0
- Remove deprecated AutoscalingV2beta1/V2beta2 from K8s interface mock
- Replace SchedulingV1alpha1 with SchedulingV1alpha2 in mock
- Update bborbe/* and golang.org/x/* dependencies
- Remove golang.org/x/lint/golint from tools

## v1.13.6

- update k8s.io/api, apimachinery, client-go to v0.35.3
- update bborbe/* dependencies (collection, errors, time, validation, parse)
- update golangci-lint to v2.11.4, counterfeiter to v6.12.2
- update go-modtool to v0.7.1, osv-scanner to v2.3.5
- bump Go toolchain to 1.26.2

## v1.13.5

- Update indirect dependencies (docker, containerd, opentelemetry, go-openapi, moby/buildkit)
- Add replace directives for charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Bump golang.org/x/{exp,time} and go-git/go-git

## v1.13.4

- chore: enable golangci-lint in check target and fix all lint violations (depguard, dupl, prealloc, revive, staticcheck)

## v1.13.3

- chore: add missing license header to mocks/mocks.go

## v1.13.2

- upgrade golangci-lint from v1 to v2
- standardize Makefile: add .PHONY declarations, multiline trivy, mocks mkdir
- update .golangci.yml to v2 format
- setup dark-factory config

## v1.13.1

- update dependencies (collection, errors, time, run)
- update docker/cli, getsentry/sentry-go, go.yaml.in/yaml/v3
- remove outdated exclude directives from go.mod

## v1.13.0

- upgrade k8s.io deps to v0.35.2 and structured-merge-diff to v6
- remove kube-openapi replace directive and k8s version excludes
- fix CronJob builder test for k8s v0.35 serialization change

## v1.12.4

- go mod update

## v1.12.3

- go mod update

## v1.12.2

- Update Go to 1.26.0

## v1.12.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.12.0

- update go and deps

## v1.11.0

- Add PodWatcher interface and implementation for watching Kubernetes pod changes
- Add PodEventProcessor interface for handling pod update and delete events
- Add PodEventProcessorSkipError wrapper for gracefully skipping errors in pod event processing
- Add PodInterface mock for testing
- Add NewPodWatcherRetry wrapper for automatic retry with configurable wait duration
- Add ErrResultChannelClosed error constant for watch channel closure detection
- Add ErrUnknownEventType error constant for unknown K8s event types
- **Improve error semantics**: Watchers now return ErrResultChannelClosed instead of nil when watch connection closes
- Update ServiceWatcher and SecretWatcher to return ErrResultChannelClosed on channel closure
- Add comprehensive test coverage for PodWatcher with all event types and error scenarios

## v1.10.0

- Add SecretWatcher interface and implementation for watching Kubernetes secret changes
- Add SecretEventProcessor interface for handling secret update and delete events
- Add SecretEventProcessorSkipError wrapper for gracefully skipping errors in secret event processing
- Add SecretInterface mock for testing
- **Refactor watchers with decorator pattern**: Separate retry logic from base watchers
- Add NewServiceWatcherRetry wrapper for automatic retry with configurable wait duration
- Add NewSecretWatcherRetry wrapper for automatic retry with configurable wait duration
- Add godoc comments to exported constructors (NewServiceWatcher, NewSecretWatcher)
- **BREAKING CHANGE**: NewServiceWatcher and NewSecretWatcher constructors no longer accept waiterDuration parameter - use retry wrappers instead
- Update dependencies (go.mod and go.sum)

## v1.9.3

- Update Go version to 1.25.2 in CI workflow
- Upgrade osv-scanner to v2 with improved config file support
- Update dependencies (go.mod and go.sum)

## v1.9.2

- rename NewServiceManagerSkipError -> NewServiceEventProcessorSkipError

## v1.9.1

- Add ServiceWatcher interface and implementation for watching Kubernetes service changes
- Add ServiceEventProcessor interface for handling service update and delete events
- Add ServiceEventProcessorSkipError wrapper for gracefully skipping errors in service event processing
- Add comprehensive test coverage for ServiceWatcher with all event types and error scenarios

## v1.9.0

- Add comprehensive package-level documentation (doc.go)
- Add godoc comments to all exported types and interfaces
- Improve README.md with installation guide, features, quick start examples, and API documentation
- Clean up CLAUDE.md to reference coding guidelines instead of duplicating content
- Update mocks and regenerate after adding documentation
- **WARNING**: Breaking changes in CronJobBuilder - SetParallelism, SetBackoffLimit, and SetCompletions now accept int32 instead of int

## v1.8.9

- go mod update

## v1.8.8

- add ApiextensionsClientset 

## v1.8.7

- add github workflow

## v1.8.6

- go mod update

## v1.8.5

- add tests
- go mod update
  `
## v1.8.4

- add JobAlreadyExistsError in JobDeployer  

## v1.8.3

- add MarshalAsYaml for testing

## v1.8.2

- remove undeploy from job deployer

## v1.8.1

- add CronScheduleExpression type with validation

## v1.8.0

- add CronJobBuilder and CronJobDeployer 

## v1.7.4

- add JobBuilder setters for TTLSecondsAfterFinished, CompletionMode and PodReplacementPolicy

## v1.7.3

- increase test coverage for all deployer components and missing builder components

## v1.7.2

- add Job restart policy validation and tests

## v1.7.1

- add ObjectMetaBuilder.SetLabels and ObjectMetaBuilder.SetAnnotations

## v1.7.0

- add HasBuild interfaces and HasBuildFunc funcs

## v1.6.3

- fix job builder app label

## v1.6.2

- fix job api version

## v1.6.1
 
- allow set PriorityClassName in PodSpec
- allow set parallelism, completions and backoffLimit in JobBuilder
- go mod update

## v1.6.0

- remove vendor
- go mod update

## v1.5.2

- add SetImagePullSecrets to PodSpecBuilder 
- go mod update

## v1.5.1

- fix DeploymentBuilder and StatefulSetBuilder

## v1.5.0

- allow define imagePullSecrets
- go mod update

## v1.4.0

- add ResourceEventHandler
- add EventHandler

## v1.3.8

- improve error message
- go mod update

## v1.3.7

- add k8s_ prefix to all go files

## v1.3.6

- allow BuildName with number at the end
- go mod update
- remove deprecated golint

## v1.3.5

- add name from pod function
- go mod update

## v1.3.4

- go mod update

## v1.3.3

- set some defaults for jobs

## v1.3.2

- allow set affinity in podSpec

## v1.3.1

- allow set restartPolicy on containerBuilder
- go mod update

## v1.3.0

- update to k8s v0.31.0
- go mod update

## v1.2.0

- add Name type
- use Name in services
- move validation

## v1.1.0

- skip ingress update if equal
- go mod update

## v1.0.0

- Initial Version
