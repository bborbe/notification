# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v0.3.0

- feat: register mantra notification type for the mantra-watcher producer

## Unreleased

- feat: register gchat-relevant notification type for the gchat-watcher producer

## v0.2.0

- feat: register go-release notification type for the go-version-watcher producer

## v0.1.2

- refactor: scope repo to library only — services moved to dedicated repos (notification-controller, notification-discord), each with own go.mod and Dockerfile

## v0.1.1

- refactor: drop frontend read-model from delivery core (read-model stays in trading as the trading UI query service)

## v0.1.0

- feat: extract generic notification core from trading monorepo (model, CDB publish command + sender, controller, NotificationHandlerTx) into this repo
- feat: opaque metadata on the notification model (source-specific fields pass through as key/value pairs)
