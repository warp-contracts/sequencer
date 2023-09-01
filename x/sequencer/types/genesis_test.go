package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &GenesisState{

				LastArweaveBlock: &ArweaveBlockInfo{},
				NextArweaveBlockList: []NextArweaveBlock{
					{
						BlockInfo: &ArweaveBlockInfo{
							Height: 0,
						},
					},
					{
						BlockInfo: &ArweaveBlockInfo{
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
			genState: &GenesisState{
				NextArweaveBlockList: []NextArweaveBlock{
					{
						BlockInfo: &ArweaveBlockInfo{
							Height: 0,
						},
					},
					{
						BlockInfo: &ArweaveBlockInfo{
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
