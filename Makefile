IGNITE_RUN = docker run --rm -ti --volume $(PWD):/apps ignitehq/cli:latest

.PHONY: build
build:
	go build -o build/sourcehubd cmd/sourcehubd/main.go

.PHONY: proto
proto:
	GOPRIVATE="github.com/sourcenetwork/*"
	$(IGNITE_RUN) generate proto-go
