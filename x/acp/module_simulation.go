package acp

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sourcenetwork/sourcehub/testutil/sample"
	acpsimulation "github.com/sourcenetwork/sourcehub/x/acp/simulation"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = acpsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreatePolicy = "op_weight_msg_create_policy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePolicy int = 100

	opWeightMsgCreateRelationship = "op_weight_msg_create_relationship"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRelationship int = 100

	opWeightMsgDeleteRelationship = "op_weight_msg_delete_relationship"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRelationship int = 100

	opWeightMsgRegisterObject = "op_weight_msg_register_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterObject int = 100

	opWeightMsgUnregisterObject = "op_weight_msg_unregister_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterObject int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	acpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&acpGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePolicy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePolicy, &weightMsgCreatePolicy, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePolicy = defaultWeightMsgCreatePolicy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePolicy,
		acpsimulation.SimulateMsgCreatePolicy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRelationship int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRelationship, &weightMsgCreateRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRelationship = defaultWeightMsgCreateRelationship
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRelationship,
		acpsimulation.SimulateMsgCreateRelationship(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRelationship int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRelationship, &weightMsgDeleteRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRelationship = defaultWeightMsgDeleteRelationship
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRelationship,
		acpsimulation.SimulateMsgDeleteRelationship(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterObject, &weightMsgRegisterObject, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterObject = defaultWeightMsgRegisterObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterObject,
		acpsimulation.SimulateMsgRegisterObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnregisterObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnregisterObject, &weightMsgUnregisterObject, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterObject = defaultWeightMsgUnregisterObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnregisterObject,
		acpsimulation.SimulateMsgUnregisterObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePolicy,
			defaultWeightMsgCreatePolicy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgCreatePolicy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRelationship,
			defaultWeightMsgCreateRelationship,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgCreateRelationship(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRelationship,
			defaultWeightMsgDeleteRelationship,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgDeleteRelationship(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterObject,
			defaultWeightMsgRegisterObject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgRegisterObject(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnregisterObject,
			defaultWeightMsgUnregisterObject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgUnregisterObject(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
