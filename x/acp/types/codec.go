package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePolicy{}, "acp/CreatePolicy", nil)
	cdc.RegisterConcrete(&MsgDeletePolicy{}, "acp/DeletePolicy", nil)
	cdc.RegisterConcrete(&MsgCreateRelationship{}, "acp/CreateRelationship", nil)
	cdc.RegisterConcrete(&MsgDeleteRelationship{}, "acp/DeleteRelationship", nil)
	cdc.RegisterConcrete(&MsgRegisterObject{}, "acp/RegisterObject", nil)
	cdc.RegisterConcrete(&MsgUnregisterObject{}, "acp/UnregisterObject", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePolicy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeletePolicy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRelationship{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteRelationship{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterObject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterObject{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
