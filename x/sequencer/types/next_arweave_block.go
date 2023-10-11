package types

import (
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func IsArweaveBlockOldEnough(sequencerBlockHeader tmproto.Header, newBlockInfo *ArweaveBlockInfo) bool {
	arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
	sequencerBlockTimestamp := sequencerBlockHeader.Time

	return sequencerBlockTimestamp.After(arweaveBlockTimestamp.Add(time.Hour))
}
