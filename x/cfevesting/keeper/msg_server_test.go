package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	// mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	// minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/chain4energy/c4e-chain/app"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"


	"github.com/stretchr/testify/require"
)
const helperModuleAccount = "heleprTestAcc"

func addHelperModuleAccountPerms() {
	perms := []string{authtypes.Minter}
	app.AddMaccPerms(helperModuleAccount, perms)
}

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CfevestingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func addCoinsToAccount(vested uint64, mintTo string, ctx sdk.Context, bank bankkeeper.Keeper, toAddr sdk.AccAddress) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	bank.MintCoins(ctx, mintTo, mintedCoins)
	bank.SendCoinsFromModuleToAccount(ctx, mintTo, toAddr, mintedCoins)
	return denom
}

func createAccountVestings(addr string, vt1 string, vested uint64, withdrawn uint64) (types.AccountVestings, *types.Vesting) {
	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting1 := types.Vesting{vt1, 1000, 10000, 110000, vested, 0, 0, 10, 0, false, withdrawn}
	vestingsArray := []*types.Vesting{&vesting1}
	accountVestings.Vestings = vestingsArray
	return accountVestings, &vesting1
}

func addCoinsToModuleByName(vested uint64, modulaName string, mintTo string, ctx sdk.Context, bank bankkeeper.Keeper) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	bank.MintCoins(ctx, mintTo, mintedCoins)
	bank.SendCoinsFromModuleToModule(ctx, mintTo, modulaName, mintedCoins)
	return denom
}

func verifyAccountBalance(t *testing.T, bank bankkeeper.Keeper, ctx sdk.Context, accAddr sdk.AccAddress, denom string, expectedAmount uint64) {
	balance := bank.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, sdk.NewIntFromUint64(expectedAmount), balance.Amount)
}

func verifyModuleAccountByName(accName string, auth authkeeper.AccountKeeper, ctx sdk.Context, bank bankkeeper.Keeper, denom string, t *testing.T, expected uint64) {
	moduleAccAddr := auth.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := bank.GetBalance(ctx, moduleAccAddr, denom)
	require.EqualValues(t, sdk.NewIntFromUint64(expected), moduleBalance.Amount)
}

func verifyModuleAccount(auth authkeeper.AccountKeeper, ctx sdk.Context, bank bankkeeper.Keeper, denom string, t *testing.T, expected uint64) {
	verifyModuleAccountByName(types.ModuleName, auth, ctx, bank, denom, t, expected)
}

func createValidator(t *testing.T, ctx sdk.Context, sk stakingkeeper.Keeper, addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	msg, err := stakingtypes.NewMsgCreateValidator(addr, pk, coin, stakingtypes.Description{}, commisions, sdk.OneInt())
	msgSrvr := stakingkeeper.NewMsgServerImpl(sk)
	require.NoError(t, err)
	res, err := msgSrvr.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, res)

}
