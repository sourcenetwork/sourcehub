package keeper

import (
	"github.com/sourcenetwork/sourcehub/x/sourcehub/types"
)

var _ types.QueryServer = Keeper{}
