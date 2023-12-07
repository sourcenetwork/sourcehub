package keeper

import (
	"github.com/sourcenetwork/sourcehub/x/bulletin/types"
)

var _ types.QueryServer = Keeper{}
