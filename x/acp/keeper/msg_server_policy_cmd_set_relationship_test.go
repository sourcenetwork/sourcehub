package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/policy_cmd"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setupTestPolicyCmdSetRelationship(t *testing.T) (ctx sdk.Context, srv types.MsgServer, pol *types.Policy, creator string, alice string, builder policy_cmd.CmdBuilder) {
	policy := `
    name: policy
    resources:
      file:
        relations:
          owner:
            types:
              - actor
          admin:
            manages:
              - reader
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

	builder = policy_cmd.CmdBuilder{}
	builder.Actor(alice)
	builder.PolicyID(pol.Id)
	builder.CreationTimestamp(timestamp)
	builder.SetSigner(signer)
	return
}

func TestPolicyCmd_SetRelationship_OwnerCanSetRelationshipForObjectTheyOwn(t *testing.T) {
	ctx, srv, pol, creator, alice, builder := setupTestPolicyCmdSetRelationship(t)

	builder.SetRelationship(types.NewActorRelationship("file", "foo", "reader", alice))
	jws, _ := builder.BuildJWS()
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.MsgPolicyCmdResponse_SetRelationshipResult{
			SetRelationshipResult: &types.SetRelationshipCmdResult{
				RecordExisted: false,
				Record: &types.RelationshipRecord{
					Actor:        alice,
					CreationTime: timestamp,
					Creator:      creator,
					PolicyId:     pol.Id,
					Relationship: types.NewActorRelationship("file", "foo", "reader", alice),
					Archived:     false,
				},
			},
		},
	}
	require.Nil(t, err)
	require.Equal(t, want, got)
}

func TestPolicyCmd_SetRelationship_ActorCannotSetRelationshipForUnregisteredObject(t *testing.T) {
	ctx, srv, _, creator, alice, builder := setupTestPolicyCmdSetRelationship(t)

	builder.SetRelationship(types.NewActorRelationship("file", "does-not-exist", "reader", alice))
	jws, _ := builder.BuildJWS()
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrObjectNotFound)
}

func TestPolicyCmd_SetRelationship_ActorCannotSetRelationshipForObjectTheyDoNotOwn(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdSetRelationship(t)

	bob, bobSigner := mustGenerateActor()
	builder.Actor(bob)
	builder.SetSigner(bobSigner)
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "reader", bob))
	jws, _ := builder.BuildJWS()
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPolicyCmd_SetRelationship_ManagerActorCanSetRelationshipForARelationTheyManage(t *testing.T) {
	ctx, srv, pol, creator, _, builder := setupTestPolicyCmdSetRelationship(t)

	// Given object foo and Bob as a manager
	bob, bobSigner := mustGenerateActor()
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "admin", bob))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	// When Bob tries to create a relationship to rel they manage
	builder.SetSigner(bobSigner)
	builder.Actor(bob)
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "reader", bob))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	want := &types.MsgPolicyCmdResponse{
		Result: &types.MsgPolicyCmdResponse_SetRelationshipResult{
			SetRelationshipResult: &types.SetRelationshipCmdResult{
				RecordExisted: false,
				Record: &types.RelationshipRecord{
					Actor:        bob,
					CreationTime: timestamp,
					Creator:      creator,
					PolicyId:     pol.Id,
					Relationship: types.NewActorRelationship("file", "foo", "reader", bob),
					Archived:     false,
				},
			},
		},
	}
	require.Equal(t, want, got)
	require.Nil(t, err)
}

func TestPolicyCmd_SetRelationship_ManagerActorCannotSetRelationshipToRelationTheyDoNotManage(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdSetRelationship(t)

	// Given object foo and Bob as a manager
	bob, bobSigner := mustGenerateActor()
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "admin", bob))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder.Actor(bob)
	builder.SetSigner(bobSigner)
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "admin", bob))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPolicyCmd_SetRelationship_ActorIsNotAllowedToSetAnOwnerRelationshipDirectly(t *testing.T) {
	ctx, srv, _, creator, _, builder := setupTestPolicyCmdSetRelationship(t)

	// Given object foo anp Bob as a manager
	bob, bobSigner := mustGenerateActor()
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "admin", bob))
	jws, err := builder.BuildJWS()
	require.Nil(t, err)
	_, err = srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))
	require.Nil(t, err)

	builder.Actor(bob)
	builder.SetSigner(bobSigner)
	builder.SetRelationship(types.NewActorRelationship("file", "foo", "owner", bob))
	jws, err = builder.BuildJWS()
	require.Nil(t, err)
	got, err := srv.PolicyCmd(ctx, types.NewMsgPolicyCmdFromJWS(creator, jws))

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrAcpProtocolViolation)
}
