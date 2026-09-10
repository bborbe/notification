# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## Unreleased

### Added

- feat: extract generic notification core from trading monorepo (model, CDB publish command + sender, controller, NotificationHandlerTx, frontend store/API) into this repo
- feat: opaque metadata on the notification model (source-specific fields pass through as key/value pairs)
