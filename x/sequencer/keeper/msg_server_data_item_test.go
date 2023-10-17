package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
)

func TestDataItemMsgServer(t *testing.T) {
	msg := test.ArweaveL2Interaction(t)
	msg.SortKey = "1,2,3"
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.DataItem(wctx, &msg)
	require.NoError(t, err)

	result, found := k.GetPrevSortKey(ctx, "abc")
	require.True(t, found)
	require.Equal(t, "abc", result.Contract)
	require.Equal(t, "1,2,3", result.SortKey)
}
