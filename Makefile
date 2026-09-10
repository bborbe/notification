include cmd/Makefile.common

IMAGES := notification-controller notification-discord notification-frontend
VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo v0.1.0)

.PHONY: precommit
precommit:
	go vet ./...
	go test ./...

.PHONY: buca
buca: ## build + push all service images from git tag (publish-only; deploy lives in quant)
	@for img in $(IMAGES); do \
		echo "==> docker.io/bborbe/$$img:$(VERSION)"; \
		docker build -t docker.io/bborbe/$$img:$(VERSION) -f cmd/$$img/Dockerfile . && \
		docker push docker.io/bborbe/$$img:$(VERSION) || exit 1; \
	done
