# Warp's sequencer
**sequencer** is a blockchain used for sequencing Warp interactions. 

## Local development
For local development you can use the usual `ignite chain serve` or a preconfigured local network of 3 nodes started with:
```
# Start the network
make docker-run

# Verify 3 validators are running
curl http://localhost:1317/cosmos/staking/v1beta1/validators 
```
