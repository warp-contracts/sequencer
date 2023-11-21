package types

import (
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// how many Arweave blocks must pass before the Arweave block is downloaded by the sequencer nodes
const ARWEAVE_BLOCK_CONFIRMATIONS = 10

// how much older the Arweave block should be compared to the sequencer block to be included in the blockchain
const ARWEAVE_BLOCK_DELAY = 30 * time.Minute

func IsArweaveBlockOldEnough(sequencerBlockHeader tmproto.Header, newBlockInfo *ArweaveBlockInfo) bool {
	arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
	sequencerBlockTimestamp := sequencerBlockHeader.Time

	return sequencerBlockTimestamp.After(arweaveBlockTimestamp.Add(ARWEAVE_BLOCK_DELAY))
}
