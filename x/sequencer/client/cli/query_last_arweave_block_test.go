package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/warp-contracts/sequencer/testutil/network"
	"github.com/warp-contracts/sequencer/testutil/nullify"
	"github.com/warp-contracts/sequencer/x/sequencer/client/cli"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func networkWithLastArweaveBlockObjects(t *testing.T) (*network.Network, types.ArweaveBlockInfo) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	lastArweaveBlock := &types.ArweaveBlockInfo{}
	nullify.Fill(&lastArweaveBlock)
	state.LastArweaveBlock = lastArweaveBlock
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.LastArweaveBlock
}

func TestShowLastArweaveBlock(t *testing.T) {
	net, obj := networkWithLastArweaveBlockObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.ArweaveBlockInfo
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowLastArweaveBlock(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetLastArweaveBlockResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.LastArweaveBlock)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.LastArweaveBlock),
				)
			}
		})
	}
}
