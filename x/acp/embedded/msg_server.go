package embedded

import (
	"context"

	"cosmossdk.io/store"
	"github.com/sourcenetwork/sourcehub/x/acp/keeper"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var _ (types.MsgServer) = (*msgServer)(nil)

// msgServer implements the ACP module MsgServer interface.
//
// Effectively msgServer wraps a keeper with the module's native implementation
// but additionally it calls the commit method in the commit multi store after each method call
type msgServer struct {
	types.UnimplementedMsgServer

	msgServer types.MsgServer
	cms       store.CommitMultiStore
}

// NewMsgSrever creates a message server for Embedded ACP
func NewMsgServer(k keeper.Keeper, cms store.CommitMultiStore) types.MsgServer {
	srv := keeper.NewMsgServerImpl(k)
	return &msgServer{
		msgServer: srv,
		cms:       cms,
	}
}

func (s *msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	resp, err := s.msgServer.UpdateParams(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}

func (s *msgServer) CreatePolicy(ctx context.Context, msg *types.MsgCreatePolicy) (*types.MsgCreatePolicyResponse, error) {
	resp, err := s.msgServer.CreatePolicy(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}

func (s *msgServer) SetRelationship(ctx context.Context, msg *types.MsgSetRelationship) (*types.MsgSetRelationshipResponse, error) {
	resp, err := s.msgServer.SetRelationship(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}
func (s *msgServer) DeleteRelationship(ctx context.Context, msg *types.MsgDeleteRelationship) (*types.MsgDeleteRelationshipResponse, error) {
	resp, err := s.msgServer.DeleteRelationship(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}
func (s *msgServer) RegisterObject(ctx context.Context, msg *types.MsgRegisterObject) (*types.MsgRegisterObjectResponse, error) {
	resp, err := s.msgServer.RegisterObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}

func (s *msgServer) UnregisterObject(ctx context.Context, msg *types.MsgUnregisterObject) (*types.MsgUnregisterObjectResponse, error) {
	resp, err := s.msgServer.UnregisterObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}
func (s *msgServer) CheckAccess(ctx context.Context, msg *types.MsgCheckAccess) (*types.MsgCheckAccessResponse, error) {
	resp, err := s.msgServer.CheckAccess(ctx, msg)
	if err != nil {
		return nil, err
	}

	s.cms.Commit()
	return resp, nil
}
