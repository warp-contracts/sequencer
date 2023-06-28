#!/bin/sh
# This script was used to initialize the chain for the first time and generate genesis.json
# Run in Linux environment
# set -ex

COIN="100000000000warp"
CHAIN_ID="sequencer-0"
TEMP_DIR="/tmp/init-chain"
rm -rf $TEMP_DIR
mkdir $TEMP_DIR
OUTFILE="$TEMP_DIR/out.txt"
touch $OUTFILE
REPO_DIR="$(dirname "$(realpath -- "$0"/..)")"
CONFIG_SUBDIR=".sequencer/config"
GENTX_SUBDIR="$CONFIG_SUBDIR/gentx"
SEQUENCER="$REPO_DIR/bin/sequencer"

log() {
    printf "$1\n"
}

genOut() {
    export HOME="$TEMP_DIR/all"
    MAIN_PASSWORD=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

    mkdir -p "$HOME/$GENTX_SUBDIR"

    $SEQUENCER init warp-sequencer --chain-id $CHAIN_ID  
    $SEQUENCER config keyring-backend file 

    # Modify genesis.json
    sed -i'' -e 's/"stake"/"warp"/g' $HOME/$CONFIG_SUBDIR/genesis.json
}

gen() {
    mkdir -p $1
    NAME=$1
    export HOME="$TEMP_DIR/$1"
    PASSWORD=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

    $SEQUENCER init warp-sequencer --chain-id $CHAIN_ID 
    $SEQUENCER config keyring-backend file 

    # Validator's account
    (
        echo $PASSWORD
        echo $PASSWORD
    ) | $SEQUENCER keys add $1 

    ADDRESS=$(echo $PASSWORD | $SEQUENCER keys show $1 -a --keyring-backend file)
    $SEQUENCER add-genesis-account $ADDRESS $COIN 

    # Register account as validator operator
    $(echo $PASSWORD | $SEQUENCER gentx $1 $COIN \
        --from $ADDRESS \
        --moniker $1 \
        --chain-id $CHAIN_ID \
        --keyring-backend file \
        --details "Warp Validator" \
        --website "warp.cc")  
        

    # Collect genesis tx to the output genesis.json
    cp $TEMP_DIR/$1/$GENTX_SUBDIR/* $TEMP_DIR/all/$GENTX_SUBDIR
    $SEQUENCER add-genesis-account $ADDRESS $COIN  --home $TEMP_DIR/all/.sequencer 

    log "Password: $1 : $PASSWORD"
}

copy() {
    cp $TEMP_DIR/all/$CONFIG_SUBDIR/*toml $REPO_DIR/network/local/$1/config
    cp $TEMP_DIR/all/$CONFIG_SUBDIR/*json $REPO_DIR/network/local/$1/config
    cp $TEMP_DIR/$1/$CONFIG_SUBDIR/node_key.json $REPO_DIR/network/local/$1/config
    cp $TEMP_DIR/$1/$CONFIG_SUBDIR/priv_validator_key.json $REPO_DIR/network/local/$1/config
    rm -f $REPO_DIR/network/local/$1/keyring-file/*
    cp $TEMP_DIR/$1/.sequencer/keyring-file/* $REPO_DIR/network/local/$1/keyring-file
}

run() {
    # Out directory is only used for generating genesis.json
    genOut

    # Initialize validators, each validator has its own directory
    # this directory will later be used to run the validator node
    # for NAME in "warp-pike" "warp-kirk" "warp-picard"; do
    for NAME in "sequencer-0" "sequencer-1" "sequencer-2"; do
        gen $NAME
    done

    export HOME="$TEMP_DIR/all"
    $SEQUENCER collect-gentxs 

    for NAME in "sequencer-0" "sequencer-1" "sequencer-2"; do
        copy $NAME
    done
}

run > $OUTFILE 2>&1
