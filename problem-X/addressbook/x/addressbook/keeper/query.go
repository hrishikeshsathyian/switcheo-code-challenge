package keeper

import (
	"addressbook/x/addressbook/types"
)

var _ types.QueryServer = Keeper{}
