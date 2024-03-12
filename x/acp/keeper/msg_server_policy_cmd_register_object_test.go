package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/policy_cmd"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setupTestPolicyCmdRegisterObject(t *testing.T) (sdk.Context, types.MsgServer, *types.Policy, *testutil.AccountKeeperStub) {
	policy := `
    name: policy
    resources:
      resource:
        relations:
          owner:
            types:
              - actor
    `
	ctx, srv, accKeep := setupMsgServer(t)
	acc := accKeep.FirstAcc()

	resp, err := srv.CreatePolicy(ctx, &types.MsgCreatePolicy{
		Creator:      acc.GetAddress().String(),
		CreationTime: timestamp,
		Policy:       policy,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
	})
	require.Nil(t, err)

	return ctx, srv, resp.Policy, accKeep
}

func TestPolicyCmd_RegisterObject_RegisteringNewObjectIsSucessful(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	creator := accKeep.FirstAcc().GetAddress().String()

	bob, signer, err := did.ProduceDID()
	require.Nil(t, err)
	builder := policy_cmd.CmdBuilder{}
	builder.Actor(bob)
	builder.PolicyID(pol.Id)
	builder.CreationTimestamp(timestamp)
	builder.RegisterObject(types.NewObject("resource", "foo"))

	cmd, err := builder.Build()
	require.Nil(t, err)
	jws, err := policy_cmd.SignPayload(cmd, signer)
	require.Nil(t, err)

	msg := types.NewMsgPolicyCmdFromJWS(creator, jws)
	got, err := srv.PolicyCmd(ctx, msg)

	require.Nil(t, err)
	want := &types.MsgPolicyCmdResponse{
		Result: &types.PolicyCmdResult{
			Result: &types.PolicyCmdResult_RegisterObjectResult{
				RegisterObjectResult: &types.RegisterObjectCmdResult{
					Result: types.RegistrationResult_Registered,
					Record: &types.RelationshipRecord{
						CreationTime: timestamp,
						Creator:      creator,
						PolicyId:     pol.Id,
						Relationship: types.NewActorRelationship("resource", "foo", "owner", bob),
						Archived:     false,
						Actor:        bob,
					},
				},
			},
		},
	}
	require.Equal(t, want, got)

	event := &types.EventObjectRegistered{
		Actor:          bob,
		PolicyId:       pol.Id,
		ObjectId:       "foo",
		ObjectResource: "resource",
	}
	testutil.AssertEventEmmited(t, ctx, event)
}
