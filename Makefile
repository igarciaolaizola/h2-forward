SHELL       =  /bin/bash
REPO        ?= docker.io/igarciaolaizola/h2-forward
DOCKER      = $(shell { command -v img || command -v docker; } 2>/dev/null)
VERSION     ?= $(shell git rev-parse --verify --short HEAD)

.PHONY: clean
clean:
	rm -f ./bin/*
	rm -f cmd/h2-forward/*-linux-amd64;
	rm -f build/container/*-linux-amd64;

.PHONY: build
build: app-build docker-build

.PHONY: app-build
app-build:
	@echo "Building app..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build \
		-a -x -tags netgo -installsuffix cgo -installsuffix netgo \
		-o ./bin/h2-forward-$(VERSION)-linux-amd64 \
		./cmd/h2-forward; \

.PHONY: docker-build
docker-build:
	@echo "Building docker image..."
	cp bin/h2-forward-$(VERSION)-linux-amd64 build/container/h2-forward-linux-amd64; \
	chmod 0755 build/container/h2-forward-linux-amd64; \
	"$(DOCKER)" build \
		-f build/container/Dockerfile \
		-t $(REPO):$(VERSION) \
	build/container/

.PHONY: docker-push
docker-push:
	@echo "Pushing docker image..."
	@DOCKER_CONFIG=$$(pwd)/.docker "$(DOCKER)" push $(REPO):$(VERSION); \

#
# Debug any makefile variable
# Usage: print-<VAR>
#
print-%  : ; @echo $* = $($*)