package keeper

import (
	"addressBook/x/addressbook/types"
)

var _ types.QueryServer = Keeper{}
