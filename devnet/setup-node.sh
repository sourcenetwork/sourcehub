#!/usr/bin/sh
# setup-node.sh initializes a sourcehub node,
# creates a validator account on the test keyring
# and self delegates some amount of sourcebucks.
#
# The signed tx is copied by make-update-gentx.

set -e


PATH="/app/build/:$PATH"

CHAIN_ID="sourcehub-localnet"
NODE_NAME="$(cat /etc/hostname)"
CHAIN_STATE="$HOME/sourcehub"
VALIDATOR_NAME="validator"
AMOUNT="100000000000sourcebucks"
CONFIGS_DIR="/app/devnet/configs"

if [ -d $CHAIN_STATE/config ];
then
    echo "Chain has been previously initialize; exiting"
    return 0
fi


sourcehubd init $NODE_NAME --chain-id $CHAIN_ID --home $CHAIN_STATE

cp $CONFIGS_DIR/*.toml $CHAIN_STATE/config/

sourcehubd keys add $VALIDATOR_NAME --home $CHAIN_STATE --keyring-backend test

VALIDATOR_ADDR=$(sourcehubd keys show $VALIDATOR_NAME --address --home $CHAIN_STATE --keyring-backend test) 
echo "Validator: $VALIDATOR_ADDR"
echo -n $VALIDATOR_ADDR > $CHAIN_STATE/validator-addr

sourcehubd genesis add-genesis-account $VALIDATOR_ADDR $AMOUNT --home $CHAIN_STATE

sourcehubd genesis gentx $VALIDATOR_NAME $AMOUNT --chain-id $CHAIN_ID --home $CHAIN_STATE --keyring-backend test
