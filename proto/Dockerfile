FROM bufbuild/buf:1.1.0 as BUILDER

FROM golang:1.19-alpine

ENV GOLANG_PROTOBUF_VERSION=1.28.0
ENV COSMOS_GOGO_VERSION=1.4.10
ENV GRPC_GATEWAY_VERSION=1.16.0

RUN apk add --no-cache git make

RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION} &&\
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v${GRPC_GATEWAY_VERSION} &&\
    go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@v${COSMOS_GOGO_VERSION}

COPY --from=BUILDER /usr/local/bin /usr/local/bin
