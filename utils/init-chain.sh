#!/bin/sh
# This script was used to initialize the chain for the first time and generate genesis.json
# set -ex

COIN="100000000000warp"
CHAIN_ID="sequencer-0"
OUTFILE="./out.txt"

touch $OUTFILE
log() {
    printf "${RED}$1${NC}\n"
}

genOut() {
    export HOME="./out"
    MAIN_PASSWORD=$(cat /dev/random | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

    mkdir -p $HOME/.sequencer/config/gentx

    ./sequencer init warp-sequencer --chain-id $CHAIN_ID  
    ./sequencer config keyring-backend file 

    # Modify genesis.json
    sed -i 's/"stake"/"warp"/g' $HOME/.sequencer/config/genesis.json
}

gen() {
    mkdir -p $1
    NAME=$1
    export HOME="./$1"
    PASSWORD=$(cat /dev/random | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

    ./sequencer init warp-sequencer --chain-id $CHAIN_ID 
    ./sequencer config keyring-backend file 

    # Validator's account
    (
        echo $PASSWORD
        echo $PASSWORD
    ) | ./sequencer keys add $1 
    # echo $PASSWORD | ./sequencer keys show $1

    # Set up the genesis account
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

}

run() {
    # Out directory is only used for generating genesis.json
    genOut

    # Initialize validators, each validator has its own directory
    # this directory will later be used to run the validator node
    for NAME in "warp-pike" "warp-kirk" "warp-picard"; do
        gen $NAME
    done

    export HOME="./out"
    ./sequencer collect-gentxs 
}

run > $OUTFILE 2>&1
