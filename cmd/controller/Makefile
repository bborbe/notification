include ../Makefile.common

SERVICE = notification-controller

run:
	@go run -mod=mod ./cmd/controller -listen="localhost:" -datadir="$(DATADIR)" -kafka-brokers="$(KAFKA_BROKERS)" -branch="$(BRANCH)" -v=2
