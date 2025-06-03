BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"
CONFIG := "./configs/calendar_config.yaml"
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

migrate:
	go run ./cmd/calendar migrate

test:
	go test -race ./internal/...

run: build
	$(BIN) -config $(CONFIG)

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.57.2

lint: install-lint-deps

.PHONY: build run build-img run-img version test lint
