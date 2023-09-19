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
		LastSortKeyList:  []LastSortKey{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in lastSortKey
	lastSortKeyIndexMap := make(map[string]struct{})

	for _, elem := range gs.LastSortKeyList {
		index := string(LastSortKeyKey(elem.Contract))
		if _, ok := lastSortKeyIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lastSortKey")
		}
		lastSortKeyIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
