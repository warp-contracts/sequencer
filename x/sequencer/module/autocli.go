package sequencer

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/warp-contracts/sequencer/api/sequencer/sequencer"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "LastArweaveBlock",
					Use:       "show-last-arweave-block",
					Short:     "show last-arweave-block",
				},
				{
					RpcMethod: "PrevSortKeyAll",
					Use:       "list-prev-sort-key",
					Short:     "List all prev-sort-key",
				},
				{
					RpcMethod:      "PrevSortKey",
					Use:            "show-prev-sort-key [id]",
					Short:          "Shows a prev-sort-key",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
