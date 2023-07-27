FROM bufbuild/buf:1.1.0 as BUILDER

FROM golang:1.19-alpine

ARG COSMOS_GOGO_VERSION=1.4.10
ARG GRPC_GATEWAY_VERSION=1.16.0
ARG VALIDATE_VERSION=1.0.2

RUN apk add --no-cache git make && adduser --disabled-password builder &&\
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION} &&\
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v${GRPC_GATEWAY_VERSION} &&\
    go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@v${COSMOS_GOGO_VERSION} &&\
    go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@v${VALIDATE_VERSION}

COPY --from=BUILDER /usr/local/bin /usr/local/bin

USER builder
