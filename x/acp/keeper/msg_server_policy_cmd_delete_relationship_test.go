package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/policy_cmd"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setupTestPolicyCmdDeleteRelationship(t *testing.T) (ctx sdk.Context, srv types.MsgServer, pol *types.Policy, creator string, alice string, builder policy_cmd.CmdBuilder) {
	policy := `
    name: policy
    resources:
      file:
        relations:
          owner:
            types:
              - actor
          reader:
            types:
              - actor
          writer:
            types:
              - actor
          admin:
            types:
              - actor
            manages:
              - reader
    `
	ctx, srv, accKeep := setupMsgServer(t)
	creator = accKeep.FirstAcc().GetAddress().String()

	resp, err := srv.CreatePolicy(ctx, &types.MsgCreatePolicy{
		Creator:      creator,
		CreationTime: timestamp,
		Policy:       policy,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
	})
	require.Nil(t, err)
	pol = resp.Policy

	alice, signer := mustGenerateActor()

	builder = policy_cmd.CmdBuilder{}
	builder.Actor(alice)
	builder.PolicyID(pol.Id)
	builder.CreationTimestamp(timestamp)
	builder.SetSigner(signer)
	builder.RegisterObject(types.NewObject("file", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder.SetRelationship(types.NewActorRelationship("file", "foo", "reader", alice))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder.SetRelationship(types.NewActorRelationship("file", "foo", "writer", alice))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder = policy_cmd.CmdBuilder{}
	builder.Actor(alice)
	builder.PolicyID(pol.Id)
	builder.CreationTimestamp(timestamp)
	builder.SetSigner(signer)
	return
}

func TestPolicyCmd_DeleteRelationship_ObjectOwnerCanRemoveRelationship(t *testing.T) {
	ctx, srv, _, creator, alice, builder := setupTestPolicyCmdDeleteRelationship(t)

	builder.DeleteRelationship(types.NewActorRelationship("file", "foo", "reader", alice))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.PolicyCmdResult{
			Result: &types.PolicyCmdResult_DeleteRelationshipResult{
				DeleteRelationshipResult: &types.DeleteRelationshipCmdResult{
					RecordFound: true,
				},
			},
		},
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}

func TestPolicyCmd_DeleteRelationship_ObjectManagerCanRemoveRelationshipsForRelationTheyManage(t *testing.T) {
	ctx, srv, _, creator, alice, builder := setupTestPolicyCmdDeleteRelationship(t)

	// Given Bob as Manager
	bob, bobSigner := mustGenerateActor()
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "admin", bob))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder.SetSigner(bobSigner)
	builder.Actor(bob)
	builder.DeleteRelationship(types.NewActorRelationship("file", "foo", "reader", alice))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.PolicyCmdResult{
			Result: &types.PolicyCmdResult_DeleteRelationshipResult{
				DeleteRelationshipResult: &types.DeleteRelationshipCmdResult{
					RecordFound: true,
				},
			},
		},
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}

func TestPolicyCmd_DeleteRelationship_ObjectManagerCannotRemoveRelationshipForRelationTheyDontManage(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdDeleteRelationship(t)

	// Given Bob as Manager
	bob, bobSigner := mustGenerateActor()
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "admin", bob))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder.SetSigner(bobSigner)
	builder.Actor(bob)
	builder.DeleteRelationship(types.NewActorRelationship("file", "foo", "writer", bob))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}
