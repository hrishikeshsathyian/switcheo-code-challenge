package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "addressBook/testutil/keeper"
	"addressBook/x/addressbook/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AddressbookKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
