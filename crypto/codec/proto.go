package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/warp-contracts/sequencer/crypto/keys/arweave"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	var pk *cryptotypes.PubKey
	registry.RegisterImplementations(pk, &arweave.PubKey{})

	var priv *cryptotypes.PrivKey
	registry.RegisterImplementations(priv, &arweave.PrivKey{})
}
