package addressbook

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "addressBook/api/addressbook/addressbook"
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
					RpcMethod:      "ShowContact",
					Use:            "show-contact [id]",
					Short:          "Query show-contact",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListContacts",
					Use:            "list-contacts",
					Short:          "Query list-contacts",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "CreateContact",
					Use:            "create-contact [first-name] [last-name] [phone] [email] [address]",
					Short:          "Send a create-contact tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "firstName"}, {ProtoField: "lastName"}, {ProtoField: "phone"}, {ProtoField: "email"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateContact",
					Use:            "update-contact [first-name] [last-name] [phone] [email] [address] [id]",
					Short:          "Send a update-contact tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "firstName"}, {ProtoField: "lastName"}, {ProtoField: "phone"}, {ProtoField: "email"}, {ProtoField: "address"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteContact",
					Use:            "delete-contact [id]",
					Short:          "Send a delete-contact tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
