package keeper

import (
	"context"
	"testing"

	prototypes "github.com/cosmos/gogoproto/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

type queryObjectOwnerSuite struct {
	suite.Suite

	obj     *types.Object
	creator string
}

func TestObjectOwner(t *testing.T) {
	suite.Run(t, &queryObjectOwnerSuite{})
}

func (s *queryObjectOwnerSuite) setup(t *testing.T) (context.Context, Keeper, string) {
	s.obj = types.NewObject("file", "1")
	s.creator = "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj"

	policyStr := `
name: policy
description: ok
resources:
  file:
    relations: 
      owner:
        doc: owner owns
        types:
          - actor-resource
      reader:
      admin:
        manages:
          - reader
    permissions: 
      own:
        expr: owner
        doc: own doc
      read: 
        expr: owner + reader
actor:
  name: actor-resource
  doc: my actor
          `

	var timestamp = prototypes.TimestampNow()

	msg := types.MsgCreatePolicy{
		Creator:      s.creator,
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
		Object:       s.obj,
		CreationTime: timestamp,
	})
	require.Nil(t, err)

	return ctx, keeper, resp.Policy.Id
}

func (s *queryObjectOwnerSuite) TestQueryReturnsObjectOwner() {
	ctx, keeper, policyId := s.setup(s.T())

	resp, err := keeper.ObjectOwner(ctx, &types.QueryObjectOwnerRequest{
		PolicyId: policyId,
		Object:   s.obj,
	})

	require.Nil(s.T(), err)
	require.Equal(s.T(), resp, &types.QueryObjectOwnerResponse{
		IsRegistered: true,
		OwnerId:      s.creator,
	})
}

func (s *queryObjectOwnerSuite) TestQueryingForUnregisteredObjectReturnsEmptyOwner() {
	ctx, keeper, policyId := s.setup(s.T())

	resp, err := keeper.ObjectOwner(ctx, &types.QueryObjectOwnerRequest{
		PolicyId: policyId,
		Object:   types.NewObject("file", "404"),
	})

	require.Nil(s.T(), err)
	require.Equal(s.T(), resp, &types.QueryObjectOwnerResponse{
		IsRegistered: false,
		OwnerId:      "",
	})
}

func (s *queryObjectOwnerSuite) TestQueryingPolicyThatDoesNotExistReturnError() {
	ctx, keeper, _ := s.setup(s.T())

	resp, err := keeper.ObjectOwner(ctx, &types.QueryObjectOwnerRequest{
		PolicyId: "some-policy",
		Object:   s.obj,
	})

	require.ErrorIs(s.T(), err, types.ErrPolicyNotFound)
	require.Nil(s.T(), resp)
}

func (s *queryObjectOwnerSuite) TestQueryingForObjectInNonExistingPolicyReturnsError() {
	ctx, keeper, policyId := s.setup(s.T())

	resp, err := keeper.ObjectOwner(ctx, &types.QueryObjectOwnerRequest{
		PolicyId: policyId,
		Object:   types.NewObject("missing-resource", "abc"),
	})

	require.Nil(s.T(), resp)
	require.NotNil(s.T(), err)
}
