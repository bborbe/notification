# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v0.7.10

- chore: update github.com/bborbe/badgerkv to v1.11.27, github.com/bborbe/kafka to v1.29.6, github.com/bborbe/run to v1.11.8

## v0.7.9

- chore: update github.com/bborbe/badgerkv to v1.11.26, github.com/bborbe/kafka to v1.29.5, github.com/bborbe/kv to v1.21.24, github.com/bborbe/memorykv to v1.6.24, github.com/bborbe/run to v1.11.7, github.com/bborbe/time to v1.27.21

## v0.7.8

- chore: update github.com/bborbe/badgerkv to v1.11.24, github.com/bborbe/kafka to v1.29.3, github.com/bborbe/kv to v1.21.21, github.com/bborbe/memorykv to v1.6.22, github.com/bborbe/run to v1.11.6, github.com/bborbe/time to v1.27.20

## v0.7.7

- chore: update github.com/bborbe/badgerkv to v1.11.23, github.com/bborbe/boltkv to v1.16.5, github.com/bborbe/cqrs to v0.6.15, github.com/bborbe/http to v1.26.32, github.com/bborbe/kafka to v1.29.2, github.com/bborbe/kv to v1.21.20, github.com/bborbe/memorykv to v1.6.21, github.com/bborbe/time to v1.27.19

## v0.7.6

- chore: update github.com/bborbe/boltkv to v1.16.4, github.com/bborbe/collection to v1.21.4, github.com/bborbe/cqrs to v0.6.14, github.com/bborbe/errors to v1.6.5, github.com/bborbe/kafka to v1.29.1, github.com/bborbe/run to v1.11.5

## v0.7.5

- chore: update Go to 1.27.3 and github.com/bborbe/badgerkv to v1.11.22, github.com/bborbe/collection to v1.21.3, github.com/bborbe/cqrs to v0.6.13, github.com/bborbe/errors to v1.6.4, github.com/bborbe/http to v1.26.31, github.com/bborbe/kafka to v1.28.6, github.com/bborbe/kv to v1.21.19, github.com/bborbe/log to v1.7.5, github.com/bborbe/memorykv to v1.6.20, github.com/bborbe/parse to v1.12.4, github.com/bborbe/time to v1.27.18, github.com/bborbe/validation to v1.5.6

## v0.7.4

- chore: update github.com/bborbe/kafka to v1.28.4

## v0.7.3

- chore: update github.com/bborbe/badgerkv to v1.11.21, github.com/bborbe/boltkv to v1.16.2, github.com/bborbe/collection to v1.21.2, github.com/bborbe/cqrs to v0.6.12, github.com/bborbe/http to v1.26.30, github.com/bborbe/kafka to v1.28.3, github.com/bborbe/kv to v1.21.18, github.com/bborbe/log to v1.7.4, github.com/bborbe/memorykv to v1.6.19, github.com/bborbe/parse to v1.12.3, github.com/bborbe/time to v1.27.17, github.com/bborbe/validation to v1.5.5

## v0.7.2

- chore: update Go to 1.27.2 and github.com/bborbe/badgerkv to v1.11.18, github.com/bborbe/boltkv to v1.16.1, github.com/bborbe/errors to v1.6.3, github.com/bborbe/http to v1.26.28, github.com/bborbe/kafka to v1.27.0, github.com/bborbe/kv to v1.21.17, github.com/bborbe/log to v1.7.3, github.com/bborbe/memorykv to v1.6.15, github.com/bborbe/parse to v1.12.2, github.com/bborbe/run to v1.11.2, github.com/bborbe/time to v1.27.16, github.com/onsi/ginkgo/v2 to v2.33.1, github.com/prometheus/client_golang to v1.25.0

## v0.7.1

- chore: update Go to 1.27.1 and github.com/IBM/sarama to v1.61.1, github.com/bborbe/badgerkv to v1.11.17, github.com/bborbe/boltkv to v1.15.5, github.com/bborbe/collection to v1.21.0, github.com/bborbe/cqrs to v0.6.11, github.com/bborbe/errors to v1.6.1, github.com/bborbe/http to v1.26.26, github.com/bborbe/kafka to v1.26.0, github.com/bborbe/kv to v1.21.16, github.com/bborbe/log to v1.7.1, github.com/bborbe/memorykv to v1.6.14, github.com/bborbe/parse to v1.12.0, github.com/bborbe/run to v1.11.0, github.com/bborbe/time to v1.27.14, github.com/bborbe/validation to v1.5.2, github.com/onsi/ginkgo/v2 to v2.33.0, github.com/onsi/gomega to v1.44.0

## v0.7.0

- feat: add `telegram.Bot` and a `Bot` field on `command/telegram.SendCommand`, naming which bot delivers a message. Optional and omitempty: an absent field unmarshals to the empty default bot, so commands produced before this change keep reaching the bot that has always handled them. A chat id cannot identify a bot — for a private chat it is the recipient's own user id — so consumers filter on `Bot` to let several bots serve one chat.

## v0.6.1

- fix: bind the notification controller's initiator to a telegram role. The telegram permissions landed in v0.6.0 but nothing granted them to anyone: `core-notification-controller` was bound only to `DiscordUser`, so the deployed telegram service rejected every send command with `permissions([discord.send]) does not contains any of permissions([telegram.send telegram.admin])` — the handler routed the notification correctly and the delivery never happened. `TelegramUser` (send only) mirrors `DiscordUser`, `TelegramAdmin` mirrors `DiscordAdmin`, and the controller's initiator joins `ApiInitiator` on the user role exactly as it does for discord.

## v0.6.0

- feat: register telegram send and admin permissions for the telegram notification handler

## v0.5.0

- feat: add telegram send command and sender for the telegram notification handler

## v0.4.0

- feat: register agent-escalation notification type for the agent escalation producer
- feat: register gchat-relevant notification type for the gchat-watcher producer

## v0.3.0

- feat: register mantra notification type for the mantra-watcher producer

## v0.2.0

- feat: register go-release notification type for the go-version-watcher producer

## v0.1.2

- refactor: scope repo to library only — services moved to dedicated repos (notification-controller, notification-discord), each with own go.mod and Dockerfile

## v0.1.1

- refactor: drop frontend read-model from delivery core (read-model stays in trading as the trading UI query service)

## v0.1.0

- feat: extract generic notification core from trading monorepo (model, CDB publish command + sender, controller, NotificationHandlerTx) into this repo
- feat: opaque metadata on the notification model (source-specific fields pass through as key/value pairs)
