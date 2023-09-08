package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsArweaveBlockOldEnough(ctx sdk.Context, newBlockInfo *ArweaveBlockInfo) bool {
	arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
	cosmosBlockTimestamp := ctx.BlockHeader().Time

	return cosmosBlockTimestamp.After(arweaveBlockTimestamp.Add(time.Hour))
}
