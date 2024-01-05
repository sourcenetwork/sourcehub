#!/usr/bin/env sh
#
if [ -z $VALIDATOR_ADDRESS ];
then 
    echo 'Must set $VALIDATOR_ADDRESS'
    exit 1
fi

if [ -z $1 ];
then
    echo 'faucet.sh target-account'
    exit 1
fi

build/sourcehubd tx bank  send "$VALIDATOR_ADDRESS" "$1" 10000stake --from $VALIDATOR_ADDRESS --chain-id sourcehub-dev --keyring-backend test --yes
