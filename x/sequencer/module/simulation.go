package sequencer

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/warp-contracts/sequencer/testutil/sample"
	sequencersimulation "github.com/warp-contracts/sequencer/x/sequencer/simulation"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// avoid unused import issue
var (
	_ = sequencersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgDataItem = "op_weight_msg_data_item"
	defaultWeightMsgDataItem int = 100

	opWeightMsgArweaveBlock = "op_weight_msg_arweave_block"
	defaultWeightMsgArweaveBlock int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	sequencerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&sequencerGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDataItem int
	simState.AppParams.GetOrGenerate(opWeightMsgDataItem, &weightMsgDataItem, nil,
		func(_ *rand.Rand) {
			weightMsgDataItem = defaultWeightMsgDataItem
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDataItem,
		sequencersimulation.SimulateMsgDataItem(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgArweaveBlock int
	simState.AppParams.GetOrGenerate(opWeightMsgArweaveBlock, &weightMsgArweaveBlock, nil,
		func(_ *rand.Rand) {
			weightMsgArweaveBlock = defaultWeightMsgArweaveBlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgArweaveBlock,
		sequencersimulation.SimulateMsgArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgDataItem,
			defaultWeightMsgDataItem,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sequencersimulation.SimulateMsgDataItem(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgArweaveBlock,
			defaultWeightMsgArweaveBlock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sequencersimulation.SimulateMsgArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}