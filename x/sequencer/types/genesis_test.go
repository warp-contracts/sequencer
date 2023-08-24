package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				LastArweaveBlock: &types.ArweaveBlockInfo{},
				NextArweaveBlockList: []types.NextArweaveBlock{
					{
						BlockInfo: &types.ArweaveBlockInfo{
							Height: 0,
						},
					},
					{
						BlockInfo: &types.ArweaveBlockInfo{
							Height: 1,
						},
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated nextArweaveBlock",
			genState: &types.GenesisState{
				NextArweaveBlockList: []types.NextArweaveBlock{
					{
						BlockInfo: &types.ArweaveBlockInfo{
							Height: 0,
						},
					},
					{
						BlockInfo: &types.ArweaveBlockInfo{
							Height: 0,
						},
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
