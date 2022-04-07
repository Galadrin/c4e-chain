package keeper_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/stretchr/testify/require"

)

func TestWithdrawAllAvailableOnVestingStart(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	addCoinsToCfevestingModule(vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), "", 1, vested, 0, false)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, 0, vested)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestingsOnVestingStart(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addCoinsToCfevestingModule(3*vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), "", 3, vested, 0, false)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, 3*vested, 0, 3*vested)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableSomeToWithdraw(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = 1000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	addCoinsToCfevestingModule(vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), "", 1, vested, 0, false)
	
	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, withdrawable, vested-withdrawable)
	accountVestings.Vestings[0].Withdrawn = sdk.NewInt(withdrawable)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(withdrawable)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdraw(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = 1000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addCoinsToCfevestingModule(3*vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), "", 3, vested, 0, false)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, 3*vested, 3*withdrawable, 3*vested-3*withdrawable)
	for _, vesting := range accountVestings.Vestings {
		vesting.Withdrawn = sdk.NewInt(withdrawable)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawable)
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = 1000
	const withdrawn = 300

	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	addCoinsToCfevestingModule(vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), "", 1, vested, withdrawn, false)
	
	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, withdrawable-withdrawn, vested-withdrawable+withdrawn)
	accountVestings.Vestings[0].Withdrawn = sdk.NewInt(withdrawable)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(withdrawable)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)

}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = 1000
	const withdrawn = 300

	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addCoinsToCfevestingModule(3*vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), "", 3, vested, withdrawn, false)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, 3*vested, 3*withdrawable-3*withdrawn, 3*vested-3*withdrawable+3*withdrawn)
	for _, vesting := range accountVestings.Vestings {
		vesting.Withdrawn = sdk.NewInt(withdrawable)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawable)
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]
	addCoinsToAccount(vested, ctx, app, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 1, 1, false, 1)
	
	makeVesting(t, ctx, app, accAddr, false, true, false, false, *vestingTypes.VestingTypes[0], vested, vested, 0, 0, 0, 0, vested)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, 0, vested)

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*vestingTypes.VestingTypes[0]}, []int64{vested}, []int64{0})

	oldCtx := ctx
	ctx = ctx.WithBlockHeight(int64(10100)).WithBlockTime(testutils.CreateTimeFromNumOfHours(10100))

	const withdrawn = 1000
	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, withdrawn, vested-withdrawn)

	verifyAccountVestings(t, oldCtx, app, accAddr, []types.VestingType{*vestingTypes.VestingTypes[0]}, []int64{vested}, []int64{withdrawn})

}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegable(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	addCoinsToAccount(3*vested, ctx, app, delegableAccAddr)
	
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	const withdrawn = 1000
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, 3*vested, 0, 3*withdrawn, 3*vested-3*withdrawn, 0)

	for _, vesting := range accountVestings.Vestings {
		vesting.Withdrawn = sdk.NewInt(withdrawn)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllSomeDelegable(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	addCoinsToCfevestingModule(2*vested, ctx, app)
	addCoinsToAccount(vested, ctx, app, delegableAccAddr)

	modifyVesting := func(v *types.Vesting) {
		if v.Id == 3 {
			v.DelegationAllowed = true
		}
	}
	accountVestings := setupAccountsVestingsWithModification(ctx, app, modifyVesting, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, false)
	const withdrawn = 1000
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, vested, 2*vested, 3*withdrawn, vested-withdrawn, 2*vested-2*withdrawn)

	for _, vesting := range accountVestings.Vestings {
		vesting.Withdrawn = sdk.NewInt(withdrawn)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}


func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount1(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(2*withdrawn, ctx, app, delegableAccAddr)
	
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, 2*withdrawn, 0, 2*withdrawn, 0, 0)

	for i, vesting := range accountVestings.Vestings {
		if i == 2 {
			vesting.Withdrawn = sdk.ZeroInt()
			vesting.LastModificationWithdrawn = sdk.ZeroInt()
		} else {
			vesting.Withdrawn = sdk.NewInt(withdrawn)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
		}
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount2(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(2*withdrawn+withdrawn/4, ctx, app, delegableAccAddr)
	
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, 2*withdrawn+withdrawn/4, 0, 2*withdrawn +withdrawn/4, 0, 0)

	for i, vesting := range accountVestings.Vestings {
		if i == 2 {
			vesting.Withdrawn = sdk.NewInt(withdrawn/4)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn/4)
		} else {
			vesting.Withdrawn = sdk.NewInt(withdrawn)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
		}
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount3(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(withdrawn/2, ctx, app, delegableAccAddr)
	
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, withdrawn/2, 0, withdrawn/2, 0, 0)

	for i, vesting := range accountVestings.Vestings {
		if i == 0 {
			vesting.Withdrawn = sdk.NewInt(withdrawn/2)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn/2)
		} else {
			vesting.Withdrawn = sdk.ZeroInt()
			vesting.LastModificationWithdrawn = sdk.ZeroInt()
		}
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount4(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(withdrawn + withdrawn/2, ctx, app, delegableAccAddr)

	modification := func(vesting *types.Vesting) {
		vesting.Withdrawn = sdk.NewInt(withdrawn)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
	}
	
	accountVestings := setupAccountsVestingsWithModification(ctx, app, modification, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, withdrawn + withdrawn/2, 0, 0, withdrawn + withdrawn/2, 0)

	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}


func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount5(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(withdrawn, ctx, app, delegableAccAddr)

	modification := func(vesting *types.Vesting) {
		vesting.Withdrawn = sdk.NewInt(withdrawn/2)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn/2)
	}
	
	accountVestings := setupAccountsVestingsWithModification(ctx, app, modification, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, withdrawn, 0, withdrawn, 0, 0)

	for i, vesting := range accountVestings.Vestings {
		if i != 2 {
			vesting.Withdrawn = sdk.NewInt(withdrawn)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
		}
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount6(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(0, ctx, app, delegableAccAddr)

	modification := func(vesting *types.Vesting) {
		vesting.Withdrawn = sdk.NewInt(withdrawn/2)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn/2)
	}
	
	accountVestings := setupAccountsVestingsWithModification(ctx, app, modification, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, 0, 0, 0, 0, 0)

	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegableNotEnoughOnDelegableAccount7(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	const withdrawn = 1000
	addCoinsToAccount(withdrawn + withdrawn/4, ctx, app, delegableAccAddr)

	modification := func(vesting *types.Vesting) {
		vesting.Withdrawn = sdk.NewInt(withdrawn/2)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn/2)
	}
	
	accountVestings := setupAccountsVestingsWithModification(ctx, app, modification, accAddr.String(), delegableAccAddr.String(), 3, vested, 0, true)
	
	withdrawAllAvailableDelegable(t, ctx, app, accAddr, delegableAccAddr, 0, withdrawn + withdrawn/4, 0, withdrawn + withdrawn/4, 0, 0)

	for i, vesting := range accountVestings.Vestings {
		if i != 2 {
			vesting.Withdrawn = sdk.NewInt(withdrawn)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn)
		} else {
			vesting.Withdrawn = sdk.NewInt(withdrawn-withdrawn/4)
			vesting.LastModificationWithdrawn = sdk.NewInt(withdrawn-withdrawn/4)
		}
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableBadAddress(t *testing.T) {

	app, ctx := setupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: "badaddress"}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)

	require.EqualError(t, err,
		"decoding bech32 failed: invalid separator index -1")


}