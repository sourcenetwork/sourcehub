package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func testUnregisterObjectSetup(t *testing.T) (sdk.Context, types.MsgServer, *types.Policy, *testutil.AccountKeeperStub) {
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
	acc := accKeep.FirstAcc()
	did, _ := did.IssueDID(acc)

	resp, err := srv.CreatePolicy(ctx, &types.MsgCreatePolicy{
		Creator:      acc.GetAddress().String(),
		CreationTime: timestamp,
		Policy:       policy,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
	})
	require.Nil(t, err)

	_, err = srv.RegisterObject(ctx, types.NewMsgRegisterObject(
		acc.GetAddress().String(),
		resp.Policy.Id,
		types.NewObject("file", "foo"),
		timestamp,
	))
	require.Nil(t, err)

	_, err = srv.SetRelationship(ctx, types.NewMsgSetRelationship(
		acc.GetAddress().String(),
		resp.Policy.Id,
		types.NewActorRelationship("file", "foo", "reader", did),
		timestamp,
	))
	require.Nil(t, err)

	return ctx, srv, resp.Policy, accKeep
}

func TestUnregisterObject_RegisteredObjectCanBeUnregisteredByAuthor(t *testing.T) {
	ctx, srv, pol, accKeep := testUnregisterObjectSetup(t)

	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		accKeep.FirstAcc().GetAddress().String(),
		pol.Id,
		types.NewObject("file", "foo"),
	))

	want := &types.MsgUnregisterObjectResponse{
		Found: true,
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}
func TestUnregisterObject_UnregisteringObjectRemovesRelationshipsLeavingTheObject(t *testing.T) {
	// TODO
}

func TestUnregisterObject_ActorCannotUnregisterObjectTheyDoNotOwn(t *testing.T) {
	ctx, srv, pol, accKeep := testUnregisterObjectSetup(t)

	randomActor := accKeep.GenAccount().GetAddress().String()
	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		randomActor,
		pol.Id,
		types.NewObject("file", "foo"),
	))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestUnregisterObject_UnregisteringAnObjectThatDoesNotExistReturnsUnauthorized(t *testing.T) {
	ctx, srv, pol, accKeep := testUnregisterObjectSetup(t)

	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		accKeep.FirstAcc().GetAddress().String(),
		pol.Id,
		types.NewObject("file", "file-that-isn't-registered"),
	))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestUnregisterObject_UnregisteringAnAlreadyArchivedObjectIsANoop(t *testing.T) {
	// Given an Archived Object
	ctx, srv, pol, accKeep := testUnregisterObjectSetup(t)
	_, _ = srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		accKeep.FirstAcc().GetAddress().String(),
		pol.Id,
		types.NewObject("file", "foo"),
	))

	// When the Creator Unregisters it
	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		accKeep.FirstAcc().GetAddress().String(),
		pol.Id,
		types.NewObject("file", "foo"),
	))

	want := &types.MsgUnregisterObjectResponse{
		Found: true,
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}

func TestUnregisterObject_InvalidCreatorAddrIsReject(t *testing.T) {
	ctx, srv, pol, _ := testUnregisterObjectSetup(t)

	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		"invalid-account-address",
		pol.Id,
		types.NewObject("file", "foo"),
	))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrAcpInput)
}

func TestUnregisterObject_AddressToNonExistingAccIsRejected(t *testing.T) {
	ctx, srv, pol, _ := testUnregisterObjectSetup(t)

	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		"cosmos1mdnz0q84lgz9rkc05ga5wh9jy9vgth7ej9efmk",
		pol.Id,
		types.NewObject("file", "foo"),
	))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrAccNotFound)
}

func TestUnregisterObject_SendingInvalidPolicyIdErrors(t *testing.T) {
	ctx, srv, _, accKeep := testUnregisterObjectSetup(t)

	resp, err := srv.UnregisterObject(ctx, types.NewMsgUnregisterObject(
		accKeep.FirstAcc().GetAddress().String(),
		"any random policy id",
		types.NewObject("file", "foo"),
	))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrPolicyNotFound)
}
