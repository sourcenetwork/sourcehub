IGNITE_RUN = docker run --rm -ti --volume $(PWD):/apps ignitehq/cli:latest
UID := $(shell id --user)
GID := $(shell id --group)
BIN = build/sourcehubd
DEMO_SRC = cmd/token-protocol-demo/main.go
DEMO_BIN = build/token-protocol-demo

.PHONY: build
build:
	go build -o ${BIN} cmd/sourcehubd/main.go

.PHONY: proto
proto:
	ignite generate proto-go

.PHONY: test
test:
	go test ./...

.PHONY: simulate
simulate:
	ignite chain simulate
	

.PHONY: fmt
fmt:
	gofmt -w .
	buf format --write

.PHONY: run
run:
	${BIN} start

.PHONY: docs
docs:
	pkgsite -http 0.0.0.0:8080
