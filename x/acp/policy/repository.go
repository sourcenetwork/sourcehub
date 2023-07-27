package policy

import (
	"context"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	zanziapi "github.com/sourcenetwork/zanzi/pkg/api"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"github.com/sourcenetwork/sourcehub/x/acp/zanzi"
)

type Repository struct {
	key    storetypes.StoreKey
	mapper mapper
}

func NewRepository(key storetypes.StoreKey) Repository {
	return Repository{
		key:    key,
		mapper: mapper{},
	}
}

func (r *Repository) Set(ctx context.Context, pol *types.Policy) error {
	service := r.getService(ctx)

	zanziPol, err := r.mapper.ToZanzi(pol)
	if err != nil {
		// wrap in invalid policy err
		return err
	}

	req := zanziapi.SetPolicyRequest{
		Policy: zanziPol,
	}
	_, err = service.Set(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, id string) (*types.Policy, error) {
	service := r.getService(ctx)

	req := zanziapi.GetPolicyRequest{
		Id: id,
	}
	res, err := service.Get(ctx, &req)
	if err != nil {
		return nil, err
	}
	if res.Policy == nil {
		return nil, nil
	}

	mapped, err := r.mapper.FromZanzi(res.Policy)
	if err != nil {
		return nil, err
	}

	return mapped, nil
}

func (r *Repository) getService(goCtx context.Context) zanziapi.PolicyServiceServer {
	ctx := sdk.UnwrapSDKContext(goCtx)
	kv := ctx.KVStore(r.key)
	zanzi := zanzi.New(kv)
	return zanzi.GetPolicyService()
}
