# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## Unreleased

- chore: update Go to 1.27.1 and github.com/bborbe/badgerkv to v1.11.17, github.com/bborbe/boltkv to v1.15.5, github.com/bborbe/collection to v1.21.0, github.com/bborbe/cqrs to v0.6.11, github.com/bborbe/errors to v1.6.1, github.com/bborbe/http to v1.26.26, github.com/bborbe/kafka to v1.26.0, github.com/bborbe/kv to v1.21.16, github.com/bborbe/log to v1.7.1, github.com/bborbe/memorykv to v1.6.14, github.com/bborbe/parse to v1.12.0, github.com/bborbe/run to v1.11.0, github.com/bborbe/time to v1.27.14, github.com/bborbe/validation to v1.5.2, github.com/onsi/ginkgo/v2 to v2.32.2

## v0.1.2

- refactor: scope repo to library only — services moved to dedicated repos (notification-controller, notification-discord), each with own go.mod and Dockerfile

## v0.1.1

- refactor: drop frontend read-model from delivery core (read-model stays in trading as the trading UI query service)

## v0.1.0

- feat: extract generic notification core from trading monorepo (model, CDB publish command + sender, controller, NotificationHandlerTx) into this repo
- feat: opaque metadata on the notification model (source-specific fields pass through as key/value pairs)
