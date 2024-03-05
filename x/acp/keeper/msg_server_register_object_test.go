package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func registerObjectTestSetup(t *testing.T) (sdk.Context, types.MsgServer, *types.Policy, *testutil.AccountKeeperStub) {
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

func TestRegisterObject_RegisteringNewObjectIsSucessful(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	creator := accKeep.FirstAcc()
	addr := creator.GetAddress().String()
	did, _ := did.IssueDID(creator)

	msg := types.NewMsgRegisterObject(addr, pol.Id, types.NewObject("resource", "foo"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, err)
	want := &types.MsgRegisterObjectResponse{
		Result: types.RegistrationResult_Registered,
		Record: &types.RelationshipRecord{
			CreationTime: timestamp,
			Creator:      addr,
			PolicyId:     pol.Id,
			Relationship: types.NewActorRelationship("resource", "foo", "owner", did),
			Archived:     false,
			Actor:        did,
		},
	}
	require.Equal(t, want, got)

	event := &types.EventObjectRegistered{
		Actor:          did,
		PolicyId:       pol.Id,
		ObjectId:       "foo",
		ObjectResource: "resource",
	}
	testutil.AssertEventEmmited(t, ctx, event)
}

func TestRegisterObject_RegisteringObjectRegisteredToAnotherUserErrors(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	// Given Bob as Owner of foo
	bob := accKeep.GenAccount().GetAddress().String()
	msg := types.NewMsgRegisterObject(bob, pol.Id, types.NewObject("resource", "foo"), timestamp)
	_, err := srv.RegisterObject(ctx, msg)
	require.Nil(t, err)

	alice := accKeep.GenAccount().GetAddress().String()
	msg = types.NewMsgRegisterObject(alice, pol.Id, types.NewObject("resource", "foo"), timestamp)
	result, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}
func TestRegisterObject_ReregisteringObjectOwnedByUserIsNoop(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	// Given Bob as Owner of foo
	bobAcc := accKeep.FirstAcc()
	bob := bobAcc.GetAddress().String()
	bobDID, _ := did.IssueDID(bobAcc)
	msg := types.NewMsgRegisterObject(bob, pol.Id, types.NewObject("resource", "foo"), timestamp)
	_, err := srv.RegisterObject(ctx, msg)
	require.Nil(t, err)

	// When Bob attempt to reregister Foo
	msg = types.NewMsgRegisterObject(bob, pol.Id, types.NewObject("resource", "foo"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	want := &types.MsgRegisterObjectResponse{
		Result: types.RegistrationResult_NoOp,
		Record: &types.RelationshipRecord{
			CreationTime: timestamp,
			Creator:      bob,
			PolicyId:     pol.Id,
			Relationship: types.NewActorRelationship("resource", "foo", "owner", bobDID),
			Archived:     false,
			Actor:        bobDID,
		},
	}
	t.Logf("err %v", err)
	require.Equal(t, want, got)
	require.Nil(t, err)
}

func TestRegisterObject_RegisteringAnotherUsersArchivedObjectErrors(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	// Given Bob as Owner of archived obj foo
	bobAcc := accKeep.FirstAcc()
	bob := bobAcc.GetAddress().String()
	msg := types.NewMsgRegisterObject(bob, pol.Id, types.NewObject("resource", "foo"), timestamp)
	_, err := srv.RegisterObject(ctx, msg)
	require.Nil(t, err)
	_, err = srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(bob, pol.Id, types.NewObject("resource", "foo")))
	require.Nil(t, err)

	// When Alice attempt to register Foo
	alice := accKeep.GenAccount().GetAddress().String()
	msg = types.NewMsgRegisterObject(alice, pol.Id, types.NewObject("resource", "foo"), timestamp)
	resp, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestRegisterObject_RegisteringArchivedUserObjectUnarchivesObject(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	// Given Bob as Owner of archived obj foo
	bobAcc := accKeep.FirstAcc()
	bob := bobAcc.GetAddress().String()
	bobDID, _ := did.IssueDID(bobAcc)
	msg := types.NewMsgRegisterObject(bob, pol.Id, types.NewObject("resource", "foo"), timestamp)
	_, err := srv.RegisterObject(ctx, msg)
	require.Nil(t, err)
	_, err = srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(bob, pol.Id, types.NewObject("resource", "foo")))
	require.Nil(t, err)

	// When Bob attempt to reregister Foo
	msg = types.NewMsgRegisterObject(bob, pol.Id, types.NewObject("resource", "foo"), timestamp)
	resp, err := srv.RegisterObject(ctx, msg)

	want := &types.MsgRegisterObjectResponse{
		Result: types.RegistrationResult_Unarchived,
		Record: &types.RelationshipRecord{
			CreationTime: timestamp,
			Creator:      bob,
			PolicyId:     pol.Id,
			Relationship: types.NewActorRelationship("resource", "foo", "owner", bobDID),
			Archived:     false,
			Actor:        bobDID,
		},
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)

	event := &types.EventObjectRegistered{
		Actor:          bobDID,
		PolicyId:       pol.Id,
		ObjectId:       "foo",
		ObjectResource: "resource",
	}
	testutil.AssertEventEmmited(t, ctx, event)
}

func TestRegisterObject_RegisteringObjectInAnUndefinedResourceErrors(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	creator := accKeep.FirstAcc()
	addr := creator.GetAddress().String()

	msg := types.NewMsgRegisterObject(addr, pol.Id, types.NewObject("unregistered-resource-abc", "foo"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, got)
	require.NotNil(t, err) // Error should be issue by Zanzi, the internal error codes aren't stable yet
}

func TestRegisterObject_CreatorWithInvalidBech32AccErrors(t *testing.T) {
	ctx, srv, pol, _ := registerObjectTestSetup(t)

	msg := types.NewMsgRegisterObject("some-non-bech23-addresss", pol.Id, types.NewObject("unregistered-resource-abc", "foo"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, got)
	require.NotNil(t, err)
}

func TestRegisterObject_RegisteringToUnknownPolicyReturnsError(t *testing.T) {
	ctx, srv, _, accKeep := registerObjectTestSetup(t)
	creator := accKeep.FirstAcc()
	addr := creator.GetAddress().String()

	msg := types.NewMsgRegisterObject(addr, "some-policy-id", types.NewObject("resource", "foo"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrPolicyNotFound)
}

func TestRegisterObject_UnknownCreatorErrors(t *testing.T) {
	ctx, srv, pol, _ := registerObjectTestSetup(t)
	unregisteredBech32 := "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj"

	msg := types.NewMsgRegisterObject(unregisteredBech32, pol.Id, types.NewObject("resource", "foo"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrAccNotFound)
}

func TestRegisterObject_BlankResourceErrors(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	creator := accKeep.FirstAcc()
	addr := creator.GetAddress().String()

	msg := types.NewMsgRegisterObject(addr, pol.Id, types.NewObject("", "obj"), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, got)
	require.NotNil(t, err)
}

func TestRegisterObject_BlankObjectIdErrors(t *testing.T) {
	ctx, srv, pol, accKeep := registerObjectTestSetup(t)
	creator := accKeep.FirstAcc()
	addr := creator.GetAddress().String()

	msg := types.NewMsgRegisterObject(addr, pol.Id, types.NewObject("resource", ""), timestamp)
	got, err := srv.RegisterObject(ctx, msg)

	require.Nil(t, got)
	require.NotNil(t, err)
}
