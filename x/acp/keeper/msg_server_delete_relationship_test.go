package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func testDeleteRelationshipSetup(t *testing.T) (sdk.Context, types.MsgServer, *types.Policy, *testutil.AccountKeeperStub) {
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

	_, err = srv.SetRelationship(ctx, types.NewMsgSetRelationship(
		acc.GetAddress().String(),
		resp.Policy.Id,
		types.NewActorRelationship("file", "foo", "writer", did),
		timestamp,
	))
	require.Nil(t, err)

	return ctx, srv, resp.Policy, accKeep
}

func TestDeleteRelationship_ObjectOwnerCanRemoveRelationship(t *testing.T) {
	ctx, srv, pol, accKeep := testDeleteRelationshipSetup(t)

	acc := accKeep.FirstAcc()
	did, _ := did.IssueDID(acc)
	resp, err := srv.DeleteRelationship(ctx, types.NewMsgDeleteRelationship(
		acc.GetAddress().String(),
		pol.Id,
		types.NewActorRelationship("file", "foo", "reader", did),
	))

	want := &types.MsgDeleteRelationshipResponse{
		RecordFound: true,
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}

func TestDeleteRelationship_ObjectManagerCanRemoveRelationshipsForRelationTheyManage(t *testing.T) {
	// Given Bob as Manager
	ctx, srv, pol, accKeep := testDeleteRelationshipSetup(t)
	acc := accKeep.FirstAcc()
	bobAcc := accKeep.GenAccount()
	bob, _ := did.IssueDID(bobAcc)
	_, err := srv.SetRelationship(ctx, types.NewMsgSetRelationship(
		acc.GetAddress().String(),
		pol.Id,
		types.NewActorRelationship("file", "foo", "admin", bob),
		timestamp,
	))
	require.Nil(t, err)

	resp, err := srv.DeleteRelationship(ctx, types.NewMsgDeleteRelationship(
		bobAcc.GetAddress().String(),
		pol.Id,
		types.NewActorRelationship("file", "foo", "reader", bob),
	))

	want := &types.MsgDeleteRelationshipResponse{
		RecordFound: false,
	}
	require.Equal(t, want, resp)
	require.Nil(t, err)
}
func TestDeleteRelationship_ObjectManagerCannotRemoveRelationshipForRelationTheyDontManage(t *testing.T) {
	// Given Bob as Manager
	ctx, srv, pol, accKeep := testDeleteRelationshipSetup(t)
	acc := accKeep.FirstAcc()
	bobAcc := accKeep.GenAccount()
	bob, _ := did.IssueDID(bobAcc)
	_, err := srv.SetRelationship(ctx, types.NewMsgSetRelationship(
		acc.GetAddress().String(),
		pol.Id,
		types.NewActorRelationship("file", "foo", "admin", bob),
		timestamp,
	))
	require.Nil(t, err)

	resp, err := srv.DeleteRelationship(ctx, types.NewMsgDeleteRelationship(
		bobAcc.GetAddress().String(),
		pol.Id,
		types.NewActorRelationship("file", "foo", "writer", bob),
	))

	require.Nil(t, resp)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}
