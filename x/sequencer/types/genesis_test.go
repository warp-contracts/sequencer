package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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
				LastArweaveBlock: &types.LastArweaveBlock{},
				PrevSortKeyList: []types.PrevSortKey{
					{
						Contract: "0",
					},
					{
						Contract: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated prevSortKey",
			genState: &types.GenesisState{
				PrevSortKeyList: []types.PrevSortKey{
					{
						Contract: "0",
					},
					{
						Contract: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
