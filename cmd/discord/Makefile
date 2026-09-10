include ../Makefile.common

SERVICE = notification-discord

run:
	@go run -mod=mod ./cmd/discord -listen="localhost:" -datadir="$(DATADIR)" -kafka-brokers="$(KAFKA_BROKERS)" -branch="$(BRANCH)" -v=2
