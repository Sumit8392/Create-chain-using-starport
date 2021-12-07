package keeper

import (
	"github.com/demo/test/x/test/types"
)

var _ types.QueryServer = Keeper{}
