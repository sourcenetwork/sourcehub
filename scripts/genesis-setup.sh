#!/usr/bin/sh
set -e

CHAIN_ID="sourcehub-dev"
VALIDATOR="validator"
NODE_NAME="node"
BIN="build/sourcehubd"

$BIN init $NODE_NAME --chain-id $CHAIN_ID

$BIN keys add $VALIDATOR --keyring-backend test
VALIDATOR_ADDR=$(build/sourcehubd keys show $VALIDATOR -a --keyring-backend test)
$BIN genesis add-genesis-account $VALIDATOR_ADDR 100000000000stake
$BIN genesis gentx $VALIDATOR 100000000stake --chain-id $CHAIN_ID --keyring-backend test

$BIN genesis collect-gentxs
