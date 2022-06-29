package keeper_test

import (
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	getParams := k.GetParams(ctx)
	require.EqualValues(t, params.MintDenom, getParams.MintDenom)
	testminter.CompareMinters(t, params.Minter, getParams.Minter)
}

func TestGetParams2(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	params.MintDenom = "dfda"
	params.Minter = createMinter(time.Now())
	k.SetParams(ctx, params)

	getParams := k.GetParams(ctx)
	require.EqualValues(t, params.MintDenom, getParams.MintDenom)
	testminter.CompareMinters(t, params.Minter, getParams.Minter)

}
