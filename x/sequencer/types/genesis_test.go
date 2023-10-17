package types

import (
	"github.com/stretchr/testify/require"
	"testing"
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

				LastArweaveBlock: &LastArweaveBlock{},
				PrevSortKeyList: []PrevSortKey{
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
			genState: &GenesisState{
				PrevSortKeyList: []PrevSortKey{
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
