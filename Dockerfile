FROM golang:1.21-bookworm as builder

ARG IGNITE_VERSION="28.1.0"

# Ignite CLI prompts users for tracking info which hangs the tty
# See https://github.com/ignite/cli/blob/main/ignite/internal/analytics/analytics.go#L71
ENV DO_NOT_TRACK=1

# Setup env and dev tools
RUN apt update &&\
    apt-get install --yes git make curl jq &&\
    mkdir /env &&\
    cd /env &&\
    curl -L https://github.com/ignite/cli/releases/download/v${IGNITE_VERSION}/ignite_${IGNITE_VERSION}_linux_amd64.tar.gz > ignite.tar.gz &&\
    tar xvf ignite.tar.gz &&\
    mv ignite /usr/local/bin/ignite &&\
    cd / && rm -rf /env
WORKDIR /app

# Cache deps
COPY go.* /app
COPY ./submodules /app/submodules
RUN go mod download

# Build
COPY . /app
ENV GOFLAGS='-buildvcs=false'
RUN --mount=type=cache,target=/root/.cache ignite chain build --skip-proto && ignite chain init --skip-proto

# Dev image entrypoint
ENTRYPOINT ["scripts/dev-entrypoint.sh"]
CMD ["start"]


# Deployment entrypoint
FROM debian:bookworm-slim

RUN useradd -ms /bin/bash tendermint
USER operator

COPY --from=builder /root/go/bin/sourcehubd /usr/local/bin/sourcehubd

EXPOSE 8080

ENTRYPOINT ["sourcehubd"]
