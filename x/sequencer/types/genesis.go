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
		LastArweaveBlock:     nil,
		NextArweaveBlockList: []NextArweaveBlock{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in nextArweaveBlock
	nextArweaveBlockIndexMap := make(map[string]struct{})

	for _, elem := range gs.NextArweaveBlockList {
		index := string(NextArweaveBlockKey(elem.GetHeightString()))
		if _, ok := nextArweaveBlockIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nextArweaveBlock")
		}
		nextArweaveBlockIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
