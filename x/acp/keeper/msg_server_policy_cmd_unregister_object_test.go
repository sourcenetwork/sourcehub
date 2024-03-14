package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/policy_cmd"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setupTestPolicyCmdUnregisterObject(t *testing.T) (ctx sdk.Context, srv types.MsgServer, pol *types.Policy, creator string, alice string, builder policy_cmd.CmdBuilder) {
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

	builder = policy_cmd.CmdBuilder{}
	builder.Actor(alice)
	builder.PolicyID(pol.Id)
	builder.CreationTimestamp(timestamp)
	builder.SetSigner(signer)
	return
}

func TestPolicyCmd_UnregisterObject_RegisteredObjectCanBeUnregisteredByAuthor(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdUnregisterObject(t)

	builder.UnregisterObject(types.NewObject("file", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.MsgPolicyCmdResponse_UnregisterObjectResult{
			UnregisterObjectResult: &types.UnregisterObjectCmdResult{
				Found:                true,
				RelationshipsRemoved: 2,
			},
		},
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}

func TestPolicyCmd_UnregisterObject_ActorCannotUnregisterObjectTheyDoNotOwn(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdUnregisterObject(t)

	bob, bobSigner := mustGenerateActor()
	builder.UnregisterObject(types.NewObject("file", "foo"))
	builder.Actor(bob)
	builder.SetSigner(bobSigner)
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPolicyCmd_UnregisterObject_UnregisteringAnObjectThatDoesNotExistReturnsUnauthorized(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdUnregisterObject(t)

	builder.UnregisterObject(types.NewObject("file", "file-that-isnt-registered"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPolicyCmd_UnregisterObject_UnregisteringAnAlreadyArchivedObjectIsANoop(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdUnregisterObject(t)

	// Given an Archived Object
	builder.UnregisterObject(types.NewObject("file", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	// When the Creator Unregisters it
	builder.UnregisterObject(types.NewObject("file", "foo"))
	jws, _ = builder.BuildJWS()
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.MsgPolicyCmdResponse_UnregisterObjectResult{
			UnregisterObjectResult: &types.UnregisterObjectCmdResult{
				Found: true,
			},
		},
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}

func TestPolicyCmd_UnregisterObject_SendingInvalidPolicyIdErrors(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdUnregisterObject(t)

	builder.UnregisterObject(types.NewObject("file", "foo"))
	builder.PolicyID("abc12345")
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrPolicyNotFound)
}
