package types

import (
	"fmt"
// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LastArweaveBlock: nil,
		PrevSortKeyList:  []PrevSortKey{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in prevSortKey
	prevSortKeyIndexMap := make(map[string]struct{})

	for _, elem := range gs.PrevSortKeyList {
		index := string(PrevSortKeyKey(elem.Contract))
		if _, ok := prevSortKeyIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for prevSortKey")
		}
		prevSortKeyIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
