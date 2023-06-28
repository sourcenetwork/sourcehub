# About

Protobuff definitions for SourceHub.

# Setup

SourceHub uses the `buff` ecosystem to manage its protobuff types.

The public Tx and Query types used by SourceHub are defined within a single buff module.

Go code genration uses [Cosmo's gogoproto](https://github.com/cosmos/gogoproto) (fork of gogo protobuff) "protoc" plugin.

Note that [buf does not use protoc](https://buf.build/docs/reference/internal-compiler/).

# Generating

The protoc environment is built in a [Dockerfile](./Dockerfile).
run `make proto` in the root project directory to regenreate the protobuff files.
