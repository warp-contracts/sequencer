package sequencer

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
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
	_ = sample.AccAddress
	_ = sequencersimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgDataItem          = "op_weight_msg_data_item"
	defaultWeightMsgDataItem int = 100

	opWeightMsgCreateLastArweaveBlock = "op_weight_msg_last_arweave_block"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLastArweaveBlock int = 100

	opWeightMsgUpdateLastArweaveBlock = "op_weight_msg_last_arweave_block"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLastArweaveBlock int = 100

	opWeightMsgDeleteLastArweaveBlock = "op_weight_msg_last_arweave_block"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLastArweaveBlock int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
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
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDataItem int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDataItem, &weightMsgDataItem, nil,
		func(_ *rand.Rand) {
			weightMsgDataItem = defaultWeightMsgDataItem
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDataItem,
		sequencersimulation.SimulateMsgDataItem(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateLastArweaveBlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLastArweaveBlock, &weightMsgCreateLastArweaveBlock, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLastArweaveBlock = defaultWeightMsgCreateLastArweaveBlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLastArweaveBlock,
		sequencersimulation.SimulateMsgCreateLastArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLastArweaveBlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateLastArweaveBlock, &weightMsgUpdateLastArweaveBlock, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLastArweaveBlock = defaultWeightMsgUpdateLastArweaveBlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLastArweaveBlock,
		sequencersimulation.SimulateMsgUpdateLastArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLastArweaveBlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteLastArweaveBlock, &weightMsgDeleteLastArweaveBlock, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLastArweaveBlock = defaultWeightMsgDeleteLastArweaveBlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLastArweaveBlock,
		sequencersimulation.SimulateMsgDeleteLastArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper),
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
			opWeightMsgCreateLastArweaveBlock,
			defaultWeightMsgCreateLastArweaveBlock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sequencersimulation.SimulateMsgCreateLastArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLastArweaveBlock,
			defaultWeightMsgUpdateLastArweaveBlock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sequencersimulation.SimulateMsgUpdateLastArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteLastArweaveBlock,
			defaultWeightMsgDeleteLastArweaveBlock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sequencersimulation.SimulateMsgDeleteLastArweaveBlock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
