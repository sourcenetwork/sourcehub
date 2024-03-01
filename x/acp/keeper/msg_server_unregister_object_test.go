package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func unregisterObjectTestSetup(t *testing.T) (sdk.Context, types.MsgServer, *types.Policy, *testutil.AccountKeeperStub) {
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

func TestUnregisterObject_RegisteredObjectCanBeUnregisteredByAuthor(t *testing.T)               {}
func TestUnregisterObject_UnregisteringObjectRemovesRelationshipsLeavingTheObject(t *testing.T) {}

func TestUnregisterObject_ActorCannotUnregisterObjectTheyDoNotOwn(t *testing.T)                  {}
func TestUnregisterObject_UnregisteringAnObjectThatDoesNotExistReturnsUnauthorized(t *testing.T) {}
func TestUnregisterObject_UnregisteringAnAlreadyArchivedObjectIsANoop(t *testing.T)              {}
func TestUnregisterObject_InvalidCreatorAddrIsReject(t *testing.T)                               {}
func TestUnregisterObject_AddressToNonExistingAccIsRejected(t *testing.T)                        {}
func TestUnregisterObject_SendingInvalidPolicyIdErrors(t *testing.T)                             {}
