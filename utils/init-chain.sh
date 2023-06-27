#!/bin/sh
# This script was used to initialize the chain for the first time and generate genesis.json
# Run in Linux environment
# set -ex

COIN="100000000000warp"
CHAIN_ID="sequencer-0"
OUTFILE="./out.txt"
touch $OUTFILE

log() {
    printf "$1\n"
}

genOut() {
    export HOME="./out"
    MAIN_PASSWORD=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

    mkdir -p $HOME/.sequencer/config/gentx

    ./sequencer init warp-sequencer --chain-id $CHAIN_ID  
    ./sequencer config keyring-backend file 

    # Modify genesis.json
    sed -i'' -e 's/"stake"/"warp"/g' $HOME/.sequencer/config/genesis.json
}

gen() {
    mkdir -p $1
    NAME=$1
    export HOME="./$1"
    PASSWORD=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

    ./sequencer init warp-sequencer --chain-id $CHAIN_ID 
    ./sequencer config keyring-backend file 

    # Validator's account
    (
        echo $PASSWORD
        echo $PASSWORD
    ) | ./sequencer keys add $1 

    ADDRESS=$(echo $PASSWORD | ./sequencer keys show $1 -a --keyring-backend file)
    ./sequencer add-genesis-account $ADDRESS $COIN 

    # Register account as validator operator
    $(echo $PASSWORD | ./sequencer gentx $1 $COIN \
        --from $ADDRESS \
        --moniker $1 \
        --chain-id $CHAIN_ID \
        --keyring-backend file \
        --details "Warp Validator" \
        --website "warp.cc")  
        

    # Collect genesis tx to the output genesis.json
    cp $1/.sequencer/config/gentx/* out/.sequencer/config/gentx

    ./sequencer add-genesis-account $ADDRESS $COIN  --home out/.sequencer 

    log "Password: $1 : $PASSWORD"
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

    export HOME="./out"
    ./sequencer collect-gentxs 
}

run > $OUTFILE 2>&1
