#!/bin/bash
# Run this script to quickly install, setup, and run the current version of the network.
#
# Examples:
# CHAIN_ID="localchain-1" HOME_DIR="~/.dungeonchain" BLOCK_TIME="1000ms" CLEAN=true sh scripts/run-node-test.sh
# CHAIN_ID="localchain-2" HOME_DIR="~/.dungeonchain" CLEAN=true RPC=36657 REST=2317 PROFF=6061 P2P=36656 GRPC=8090 GRPC_WEB=8091 ROSETTA=8081 BLOCK_TIME="500ms" sh scripts/test_ics_node.sh

set -eu

export KEY="acc0"
export KEY2="acc1"

export CHAIN_ID=${CHAIN_ID:-"localdungeon"}
export MONIKER="localvalidator"
export KEYALGO="secp256k1"
export KEYRING=${KEYRING:-"test"}
export HOME_DIR="mytestnet"
BINARY=$1
export DENOM=${DENOM:-udgn}

echo $BINARY

export CLEAN=${CLEAN:-"false"}
export RPC=${RPC:-"26657"}
export REST=${REST:-"1317"}
export PROFF=${PROFF:-"6060"}
export P2P=${P2P:-"26656"}
export GRPC=${GRPC:-"9090"}
export GRPC_WEB=${GRPC_WEB:-"9091"}
export ROSETTA=${ROSETTA:-"8080"}
export BLOCK_TIME=${BLOCK_TIME:-"5s"}

config_toml="${HOME_DIR}/config/config.toml"
client_toml="${HOME_DIR}/config/client.toml"
app_toml="${HOME_DIR}/config/app.toml"
genesis_json="${HOME_DIR}/config/genesis.json"

# if which binary does not exist, install it
if [ -z `which $BINARY` ]; then
  make install

  if [ -z `which $BINARY` ]; then
    echo "Ensure $BINARY is installed and in your PATH"
    exit 1
  fi
fi

command -v $BINARY > /dev/null 2>&1 || { echo >&2 "$BINARY command not found. Ensure this is setup / properly installed in your GOPATH (make install)."; exit 1; }
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }


add_key() {
  key=$1
  mnemonic=$2
  echo $mnemonic | $BINARY keys add $key --home $HOME_DIR --keyring-backend $KEYRING --algo $KEYALGO --recover
}

echo "adding keys"



## submit proposal to update params
$BINARY tx gov submit-proposal ./draft_proposal.json --from $KEY --chain-id $CHAIN_ID --home $HOME_DIR -y

## vote on proposal
sleep 1
$BINARY tx gov vote 1 yes --from $KEY --chain-id $CHAIN_ID --home $HOME_DIR -y

## validator 2 vote yes
sleep 1
$BINARY tx gov vote 1 yes --from $KEY2 --chain-id $CHAIN_ID --home $HOME_DIR -y


