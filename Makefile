IGNITE_RUN = docker run --rm -ti --volume $(PWD):/apps ignitehq/cli:latest
UID := $(shell id --user)
GID := $(shell id --group)

.PHONY: build
build:
	go build -o build/sourcehubd cmd/sourcehubd/main.go

.PHONY: proto-ignite
proto-ignite:
	GOPRIVATE="github.com/sourcenetwork/*"
	$(IGNITE_RUN) generate proto-go

.PHONY: proto
proto:
	GOPRIVATE="github.com/sourcenetwork/*"
	docker image build --file proto/Dockerfile --tag sourcehub-proto-builder:latest proto/
	docker run --rm -it --workdir /app --user ${UID}:${GID} --volume $(PWD):/app sourcehub-proto-builder:latest buf generate --verbose
	# since gogoproto does not have a `module` argument as google's proto
	# the script has to do some cleaning up
	cp -r github.com/sourcenetwork/sourcehub/* .
	rm -rf github.com
	go mod tidy

.PHONY: test
test:
	go test ./...
