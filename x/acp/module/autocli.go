package acp

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/sourcenetwork/sourcehub/api/sourcehub/acp"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Query_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Policy",
					Use:            "policy [id]",
					Short:          "Queries for a Policy by its ID",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "PolicyIds",
					Use:            "policy-ids",
					Short:          "Lists Registered Policies IDs",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod: "ValidatePolicy",
					Use:       "validate-policy [policy] [marshal_type]",
					Short:     "Validates the Payload of a Policy",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "policy"},
						{ProtoField: "marshal_type"},
					},
				},

				{
					RpcMethod:      "AccessDecision",
					Use:            "access-decision [id]",
					Short:          "Returns an AccessDecision by its id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
				{
			RpcMethod: "PolicyCmd",
			Use: "policy-cmd",
			Short: "Send a policy-cmd tx",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
		},
		// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
