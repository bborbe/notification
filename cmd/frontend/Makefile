include ../Makefile.common

SERVICE = notification-frontend

run:
	@go run -mod=mod ./cmd/frontend -listen="localhost:" -datadir="$(DATADIR)" -kafka-brokers="$(KAFKA_BROKERS)" -branch="$(BRANCH)" -v=2
