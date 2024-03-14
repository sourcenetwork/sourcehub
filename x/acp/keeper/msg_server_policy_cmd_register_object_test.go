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

func setupTestPolicyCmdRegisterObject(t *testing.T) (ctx sdk.Context, srv types.MsgServer, pol *types.Policy, creator string, alice string, builder policy_cmd.CmdBuilder) {
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
	creator = accKeep.FirstAcc().GetAddress().String()

	resp, err := srv.CreatePolicy(ctx, &types.MsgCreatePolicy{
		Creator:      creator,
		CreationTime: timestamp,
		Policy:       policy,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
	})
	require.Nil(t, err)
	pol = resp.Policy

	alice, signer, err := did.ProduceDID()
	require.Nil(t, err)

	builder = policy_cmd.CmdBuilder{}
	builder.Actor(alice)
	builder.PolicyID(pol.Id)
	builder.CreationTimestamp(timestamp)
	builder.SetSigner(signer)

	return
}

func TestPolicyCmd_RegisterObject_RegisteringNewObjectIsSucessful(t *testing.T) {
	ctx, srv, pol, creator, alice, builder := setupTestPolicyCmdRegisterObject(t)

	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, err := builder.BuildJWS()
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
						Relationship: types.NewActorRelationship("resource", "foo", "owner", alice),
						Archived:     false,
						Actor:        alice,
					},
				},
			},
		},
	}
	require.Equal(t, want, got)

	event := &types.EventObjectRegistered{
		Actor:          alice,
		PolicyId:       pol.Id,
		ObjectId:       "foo",
		ObjectResource: "resource",
	}
	testutil.AssertEventEmmited(t, ctx, event)
}

func TestPolicyCmd_RegisterObject_RegisteringObjectRegisteredToAnotherUserErrors(t *testing.T) {
	ctx, srv, pol, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	// Given creator as Owner of foo
	msgReg := types.NewMsgRegisterObject(creator, pol.Id, types.NewObject("resource", "foo"), timestamp)
	_, err := srv.RegisterObject(ctx, msgReg)
	require.Nil(t, err)

	// When alice tries to Register Foo
	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	msg := types.NewMsgPolicyCmdFromJWS(creator, jws)
	result, err := srv.PolicyCmd(ctx, msg)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPolicyCmd_RegisterObject_ReregisteringObjectOwnedByUserIsNoop(t *testing.T) {
	ctx, srv, pol, creator, alice, builder := setupTestPolicyCmdRegisterObject(t)

	// Given alice as Owner of foo
	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, _ := builder.BuildJWS()
	_, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	// When alice tries to reregister Foo
	jws, _ = builder.BuildJWS()
	result, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.PolicyCmdResult{
			Result: &types.PolicyCmdResult_RegisterObjectResult{
				RegisterObjectResult: &types.RegisterObjectCmdResult{
					Result: types.RegistrationResult_NoOp,
					Record: &types.RelationshipRecord{
						CreationTime: timestamp,
						Creator:      creator,
						PolicyId:     pol.Id,
						Relationship: types.NewActorRelationship("resource", "foo", "owner", alice),
						Archived:     false,
						Actor:        alice,
					},
				},
			},
		},
	}
	require.Equal(t, want, result)
	require.Nil(t, err)
}

func TestPolicyCmd_RegisterObject_RegisteringAnotherUsersArchivedObjectErrors(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	// Given Alice as owner of archived obj foo
	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)
	builder.UnregisterObject(types.NewObject("resource", "foo"))
	jws, _ = builder.BuildJWS()
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	// When Bob attempt to register Foo
	bob, bobSigner, err := did.ProduceDID()
	require.Nil(t, err)
	builder.SetSigner(bobSigner)
	builder.Actor(bob)
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPolicyCmd_RegisterObject_RegisteringArchivedUserObjectUnarchivesObject(t *testing.T) {
	ctx, srv, pol, creator, alice, builder := setupTestPolicyCmdRegisterObject(t)

	// Given Alice as owner of archived obj foo
	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)
	builder.UnregisterObject(types.NewObject("resource", "foo"))
	jws, _ = builder.BuildJWS()
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	// When Alice attempt to reregister Foo
	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	resp, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.PolicyCmdResult{
			Result: &types.PolicyCmdResult_RegisterObjectResult{
				RegisterObjectResult: &types.RegisterObjectCmdResult{
					Result: types.RegistrationResult_Unarchived,
					Record: &types.RelationshipRecord{
						CreationTime: timestamp,
						Creator:      creator,
						PolicyId:     pol.Id,
						Relationship: types.NewActorRelationship("resource", "foo", "owner", alice),
						Archived:     false,
						Actor:        alice,
					},
				},
			},
		},
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)

	event := &types.EventObjectRegistered{
		Actor:          alice,
		PolicyId:       pol.Id,
		ObjectId:       "foo",
		ObjectResource: "resource",
	}
	testutil.AssertEventEmmited(t, ctx, event)
}

func TestPolicyCmd_RegisterObject_RegisteringObjectInAnUndefinedResourceErrors(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	builder.RegisterObject(types.NewObject("some-undefined-resource", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.NotNil(t, err) // Error should be issue by Zanzi, the internal error codes aren't stable yet
}

func TestPolicyCmd_RegisterObject_CreatorWithInvalidDID(t *testing.T) {
	ctx, srv, pol, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	payload := types.PolicyCmdPayload{
		Actor:           "some invalid did",
		CreationTime:    timestamp,
		IssuedHeight:    100,
		ExpirationDelta: 10,
		PolicyId:        pol.Id,
		Cmd: &types.PolicyCmdPayload_RegisterObjectCmd{
			RegisterObjectCmd: &types.RegisterObjectCmd{
				Object: types.NewObject("resource", "foo"),
			},
		},
	}
	jws, err := policy_cmd.SignPayload(payload, builder.GetSigner())
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.NotNil(t, err)
}

func TestPolicyCmd_RegisterObject_RegisteringToUnknownPolicyReturnsError(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	builder.PolicyID("bae123456")
	builder.RegisterObject(types.NewObject("resource", "foo"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrPolicyNotFound)
}

func TestPolicyCmd_RegisterObject_BlankResourceErrors(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	builder.RegisterObject(types.NewObject("", "obj"))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.NotNil(t, err)
}

func TestPolicyCmd_RegisterObject_BlankObjectIdErrors(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdRegisterObject(t)

	builder.RegisterObject(types.NewObject("resource", ""))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.NotNil(t, err)
}
