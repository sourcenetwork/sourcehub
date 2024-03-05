package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setRelationshipTestSetup(t *testing.T) (sdk.Context, types.MsgServer, *types.Policy, *testutil.AccountKeeperStub) {
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
	acc := accKeep.FirstAcc()

	resp, err := srv.CreatePolicy(ctx, &types.MsgCreatePolicy{
		Creator:      acc.GetAddress().String(),
		CreationTime: timestamp,
		Policy:       policy,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
	})
	require.Nil(t, err)

	_, err = srv.RegisterObject(ctx, types.NewMsgRegisterObject(acc.GetAddress().String(), resp.Policy.Id, types.NewObject("file", "foo"), timestamp))
	require.Nil(t, err)

	return ctx, srv, resp.Policy, accKeep
}

func TestSetRelationship_OwnerCanSetRelationshipForObjectTheyOwn(t *testing.T) {
	ctx, srv, pol, accKeep := setRelationshipTestSetup(t)

	creator := accKeep.FirstAcc().GetAddress().String()
	creatorDID, _ := did.IssueDID(accKeep.FirstAcc())
	bobAcc := accKeep.GenAccount()
	bobDID, _ := did.IssueDID(bobAcc)
	msg := types.NewMsgSetRelationship(creator, pol.Id, types.NewActorRelationship("file", "foo", "reader", bobDID), timestamp)
	got, err := srv.SetRelationship(ctx, msg)

	want := &types.MsgSetRelationshipResponse{
		RecordExisted: false,
		Record: &types.RelationshipRecord{
			Actor:        creatorDID,
			CreationTime: timestamp,
			Creator:      creator,
			PolicyId:     pol.Id,
			Relationship: types.NewActorRelationship("file", "foo", "reader", bobDID),
			Archived:     false,
		},
	}
	require.Equal(t, want, got)
	require.Nil(t, err)
}
func TestSetRelationship_ActorCannotSetRelationshipForUnregisteredObject(t *testing.T) {
	ctx, srv, pol, accKeep := setRelationshipTestSetup(t)

	creator := accKeep.FirstAcc().GetAddress().String()
	creatorDID, _ := did.IssueDID(accKeep.FirstAcc())
	msg := types.NewMsgSetRelationship(creator, pol.Id, types.NewActorRelationship("file", "does-not-exist", "reader", creatorDID), timestamp)
	got, err := srv.SetRelationship(ctx, msg)

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrObjectNotFound)
}
func TestSetRelationship_ActorCannotSetRelationshipForObjectTheyDoNotOwn(t *testing.T) {
	ctx, srv, pol, accKeep := setRelationshipTestSetup(t)

	bob := accKeep.GenAccount()
	bobDID, _ := did.IssueDID(bob)
	msg := types.NewMsgSetRelationship(bob.GetAddress().String(), pol.Id, types.NewActorRelationship("file", "foo", "reader", bobDID), timestamp)
	got, err := srv.SetRelationship(ctx, msg)

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}
func TestSetRelationship_ManagerActorCanSetRelationshipForARelationTheyManage(t *testing.T) {
	// Given object foo and Bob as a manager
	ctx, srv, pol, accKeep := setRelationshipTestSetup(t)
	creator := accKeep.FirstAcc()
	bob := accKeep.GenAccount()
	bobDID, _ := did.IssueDID(bob)
	rel := types.NewActorRelationship("file", "foo", "admin", bobDID)
	srv.SetRelationship(ctx, types.NewMsgSetRelationshipNow(creator.GetAddress().String(), pol.Id, rel))

	msg := types.NewMsgSetRelationship(bob.GetAddress().String(), pol.Id, types.NewActorRelationship("file", "foo", "reader", bobDID), timestamp)
	got, err := srv.SetRelationship(ctx, msg)

	want := &types.MsgSetRelationshipResponse{
		RecordExisted: false,
		Record: &types.RelationshipRecord{
			Actor:        bobDID,
			CreationTime: timestamp,
			Creator:      bob.GetAddress().String(),
			PolicyId:     pol.Id,
			Relationship: types.NewActorRelationship("file", "foo", "reader", bobDID),
			Archived:     false,
		},
	}
	t.Logf("err %v", err)
	require.Equal(t, want, got)
	require.Nil(t, err)
}

func TestSetRelationship_ManagerActorCannotSetRelationshipToRelationshipsTheyDoNotManage(t *testing.T) {
	// Given object foo and Bob as a manager
	ctx, srv, pol, accKeep := setRelationshipTestSetup(t)
	creator := accKeep.FirstAcc()
	bob := accKeep.GenAccount()
	bobDID, _ := did.IssueDID(bob)
	rel := types.NewActorRelationship("file", "foo", "admin", bobDID)
	srv.SetRelationship(ctx, types.NewMsgSetRelationshipNow(creator.GetAddress().String(), pol.Id, rel))

	msg := types.NewMsgSetRelationship(bob.GetAddress().String(), pol.Id, types.NewActorRelationship("file", "foo", "admin", bobDID), timestamp)
	got, err := srv.SetRelationship(ctx, msg)

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}
func TestSetRelationship_ActorIsNotAllowedToSetAnOwnerRelationshipDirectly(t *testing.T) {
	// Given object foo and Bob as a manager
	ctx, srv, pol, accKeep := setRelationshipTestSetup(t)
	creator := accKeep.FirstAcc()
	bob := accKeep.GenAccount()
	bobDID, _ := did.IssueDID(bob)
	rel := types.NewActorRelationship("file", "foo", "admin", bobDID)
	srv.SetRelationship(ctx, types.NewMsgSetRelationshipNow(creator.GetAddress().String(), pol.Id, rel))

	msg := types.NewMsgSetRelationship(bob.GetAddress().String(), pol.Id, types.NewActorRelationship("file", "foo", "owner", bobDID), timestamp)
	got, err := srv.SetRelationship(ctx, msg)

	require.Nil(t, got)
	require.ErrorIs(t, err, types.ErrAcpProtocolViolation)
}
