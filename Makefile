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
	docker run --rm -it --workdir /app -v $(PWD):/app sourcehub-proto-builder:latest buf generate --verbose
	cp -r github.com/sourcenetwork/sourcehub/* .
	sudo rm -rf github.com
	gofmt -w -r '"github.com/gogo/protobuf/grpc" -> "github.com/cosmos/gogoproto/grpc"' .
	go mod tidy
