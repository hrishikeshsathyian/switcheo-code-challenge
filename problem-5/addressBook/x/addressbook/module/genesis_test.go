package addressbook_test

import (
	"testing"

	keepertest "addressbook/testutil/keeper"
	"addressbook/testutil/nullify"
	addressbook "addressbook/x/addressbook/module"
	"addressbook/x/addressbook/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AddressbookKeeper(t)
	addressbook.InitGenesis(ctx, k, genesisState)
	got := addressbook.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
