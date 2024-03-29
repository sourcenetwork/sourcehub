package acp

import (
	"math/rand"

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
	_ = acpsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePolicy = "op_weight_msg_create_policy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePolicy int = 100

	opWeightMsgSetRelationship = "op_weight_msg_set_relationship"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetRelationship int = 100

	opWeightMsgDeleteRelationship = "op_weight_msg_delete_relationship"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRelationship int = 100

	opWeightMsgRegisterObject = "op_weight_msg_register_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterObject int = 100

	opWeightMsgUnregisterObject = "op_weight_msg_unregister_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterObject int = 100

	opWeightMsgCheckAccess = "op_weight_msg_check_access"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCheckAccess int = 100

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
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePolicy int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePolicy, &weightMsgCreatePolicy, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePolicy = defaultWeightMsgCreatePolicy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePolicy,
		acpsimulation.SimulateMsgCreatePolicy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetRelationship int
	simState.AppParams.GetOrGenerate(opWeightMsgSetRelationship, &weightMsgSetRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgSetRelationship = defaultWeightMsgSetRelationship
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetRelationship,
		acpsimulation.SimulateMsgSetRelationship(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRelationship int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteRelationship, &weightMsgDeleteRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRelationship = defaultWeightMsgDeleteRelationship
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRelationship,
		acpsimulation.SimulateMsgDeleteRelationship(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterObject int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterObject, &weightMsgRegisterObject, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterObject = defaultWeightMsgRegisterObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterObject,
		acpsimulation.SimulateMsgRegisterObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnregisterObject int
	simState.AppParams.GetOrGenerate(opWeightMsgUnregisterObject, &weightMsgUnregisterObject, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterObject = defaultWeightMsgUnregisterObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnregisterObject,
		acpsimulation.SimulateMsgUnregisterObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCheckAccess int
	simState.AppParams.GetOrGenerate(opWeightMsgCheckAccess, &weightMsgCheckAccess, nil,
		func(_ *rand.Rand) {
			weightMsgCheckAccess = defaultWeightMsgCheckAccess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCheckAccess,
		acpsimulation.SimulateMsgCheckAccess(am.accountKeeper, am.bankKeeper, am.keeper),
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
			opWeightMsgSetRelationship,
			defaultWeightMsgSetRelationship,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgSetRelationship(am.accountKeeper, am.bankKeeper, am.keeper)
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
		simulation.NewWeightedProposalMsg(
			opWeightMsgCheckAccess,
			defaultWeightMsgCheckAccess,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				acpsimulation.SimulateMsgCheckAccess(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
