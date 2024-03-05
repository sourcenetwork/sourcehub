package keeper

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setupTestVerifyAccessRequest(t *testing.T) (context.Context, Keeper, *types.Policy) {
	obj := types.NewObject("file", "1")
	creator := "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj"

	policyStr := `
name: policy
resources:
  file:
    relations: 
      owner:
        types:
          - actor
      rm-root:
    permissions: 
      read: 
        expr: owner
      write: 
        expr: owner
`

	msg := types.MsgCreatePolicy{
		Creator:      creator,
		Policy:       policyStr,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
		CreationTime: timestamp,
	}

	keeper, ctx := setupKeeper(t)
	msgServer := NewMsgServerImpl(keeper)

	resp, err := msgServer.CreatePolicy(ctx, &msg)
	require.Nil(t, err)

	_, err = msgServer.RegisterObject(ctx, &types.MsgRegisterObject{
		Creator:      creator,
		PolicyId:     resp.Policy.Id,
		Object:       obj,
		CreationTime: timestamp,
	})
	require.Nil(t, err)

	return ctx, keeper, resp.Policy
}

func TestVerifyAccessRequest_QueryingObjectsTheActorHasAccessToReturnsTrue(t *testing.T) {
	ctx, keeper, pol := setupTestVerifyAccessRequest(t)

	req := &types.QueryVerifyAccessRequestRequest{
		PolicyId: pol.Id,
		AccessRequest: &types.AccessRequest{
			Operations: []*types.Operation{
				{
					Object:     types.NewObject("file", "1"),
					Permission: "read",
				},
				{
					Object:     types.NewObject("file", "1"),
					Permission: "write",
				},
			},
			Actor: &types.Actor{
				Id: "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj",
			},
		},
	}
	result, err := keeper.VerifyAccessRequest(ctx, req)

	want := &types.QueryVerifyAccessRequestResponse{
		Valid: true,
	}
	require.Equal(t, want, result)
	require.Nil(t, err)
}

func TestVerifyAccessRequest_QueryingOperationActorIsNotAuthorizedReturnNotValid(t *testing.T) {
	ctx, keeper, pol := setupTestVerifyAccessRequest(t)

	req := &types.QueryVerifyAccessRequestRequest{
		PolicyId: pol.Id,
		AccessRequest: &types.AccessRequest{
			Operations: []*types.Operation{
				{
					Object:     types.NewObject("file", "1"),
					Permission: "rm-root",
				},
			},
			Actor: &types.Actor{
				Id: "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj",
			},
		},
	}
	result, err := keeper.VerifyAccessRequest(ctx, req)

	want := &types.QueryVerifyAccessRequestResponse{
		Valid: false,
	}
	require.Equal(t, want, result)
	require.Nil(t, err)
}

func TestVerifyAccessRequest_QueryingObjectThatDoesNotExistReturnValidFalse(t *testing.T) {
	ctx, keeper, pol := setupTestVerifyAccessRequest(t)

	req := &types.QueryVerifyAccessRequestRequest{
		PolicyId: pol.Id,
		AccessRequest: &types.AccessRequest{
			Operations: []*types.Operation{
				{
					Object:     types.NewObject("file", "file-that-is-not-registered"),
					Permission: "read",
				},
			},
			Actor: &types.Actor{
				Id: "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj",
			},
		},
	}
	result, err := keeper.VerifyAccessRequest(ctx, req)

	want := &types.QueryVerifyAccessRequestResponse{
		Valid: false,
	}
	require.Equal(t, want, result)
	require.Nil(t, err)
}
