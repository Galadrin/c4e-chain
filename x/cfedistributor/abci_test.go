package cfedistributor_test

import (
	testapp "github.com/chain4energy/c4e-chain/app"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestBurningDistributorMainCollectorDes(t *testing.T) {
	BurningDistributorTest(t, testutils.MainCollector)
}

func TestBurningDistributorModuleAccountDest(t *testing.T) {
	BurningDistributorTest(t, testutils.ModuleAccount)
}

func TestBurningDistributorInternalAccountDest(t *testing.T) {
	BurningDistributorTest(t, testutils.InternalAccount)
}

func TestBurningDistributorBaseAccountDest(t *testing.T) {
	BurningDistributorTest(t, testutils.BaseAccount)
}

func BurningDistributorTest(t *testing.T, destinationType testutils.DestinationType) {

	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	denom := "uc4e"
	testapp.AddMaccPerms(collector, perms)
	testapp.AddMaccPerms("c4e_distributor", nil)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	// app.AccountKeeper.GetModuleAccount(ctx, "c4e_distributor");
	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
	app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))
	require.EqualValues(t, cointToMint, app.BankKeeper.GetSupply(ctx, denom))
	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, testutils.PrepareBurningDistributor(destinationType))

	app.CfedistributorKeeper.SetParams(ctx, types.NewParams(subdistributors))
	ctx = ctx.WithBlockHeight(int64(2))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})

	//coin on "burnState" should be equal 498, remains: 1 and 0.33 on remains
	burnState, _ := app.CfedistributorKeeper.GetState(ctx, "burn_state_key")
	ctx.Logger().Error(burnState.String())
	//burnState, _ := app.CfedistributorKeeper.GetAllStates()
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf("uc4e"))

	if destinationType == testutils.MainCollector {
		mainCollectorCoins :=
			app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		require.EqualValues(t, 1, len(app.CfedistributorKeeper.GetAllStates(ctx)))
		require.EqualValues(t, sdk.NewInt(499), mainCollectorCoins.AmountOf(denom))
	} else if destinationType == testutils.ModuleAccount {
		mainCollectorCoins :=
			app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		c4eModulAccountCoins :=
			app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "c4e_distributor")
		require.EqualValues(t, 2, len(app.CfedistributorKeeper.GetAllStates(ctx)))
		require.EqualValues(t, sdk.NewInt(498), c4eModulAccountCoins.AmountOf(denom))
		require.EqualValues(t, sdk.NewInt(1), mainCollectorCoins.AmountOf(denom))
		c4eDistrState, _ := app.CfedistributorKeeper.GetState(ctx, "c4e_distributor")
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains.AmountOf("uc4e"))

	} else if destinationType == testutils.InternalAccount {
		mainCollectorCoins :=
			app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		require.EqualValues(t, 2, len(app.CfedistributorKeeper.GetAllStates(ctx)))
		require.EqualValues(t, sdk.NewInt(499), mainCollectorCoins.AmountOf(denom))
		c4eDistrState, _ := app.CfedistributorKeeper.GetState(ctx, "c4e_distributor")
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("498.33"), coinRemains.AmountOf("uc4e"))
	} else {
		address, _ := sdk.AccAddressFromBech32("cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
		mainCollectorCoins :=
			app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)

		accountCoins :=
			app.CfedistributorKeeper.GetAccountCoins(ctx, address)

		require.EqualValues(t, 2, len(app.CfedistributorKeeper.GetAllStates(ctx)))
		ctx.Logger().Error(accountCoins.AmountOf(denom).String())
		println("Amount: " + accountCoins.AmountOf(denom).String())

		require.EqualValues(t, sdk.NewInt(498), accountCoins.AmountOf(denom))
		require.EqualValues(t, sdk.NewInt(1), mainCollectorCoins.AmountOf(denom))

		c4eDistrState, _ := app.CfedistributorKeeper.GetState(ctx, "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains.AmountOf("uc4e"))
	}
	require.EqualValues(t, sdk.NewCoin(denom, sdk.NewInt(499)), app.BankKeeper.GetSupply(ctx, denom))
}

func TestBurningWithInflationDistributorPassThroughMainCollector(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.MainCollector, true)
}

func TestBurningWithInflationDistributorPassThroughModuleAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.ModuleAccount, true)
}

func TestBurningWithInflationDistributorPassInternalAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.InternalAccount, true)
}

func TestBurningWithInflationDistributorPassBaseAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.BaseAccount, true)
}

func TestBurningWithInflationDistributorPassThroughMainCollectorNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.MainCollector, false)
}

func TestBurningWithInflationDistributorPassThroughModuleAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.ModuleAccount, false)
}

func TestBurningWithInflationDistributorPassInternalAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.InternalAccount, false)
}

func TestBurningWithInflationDistributorPassBaseAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, testutils.BaseAccount, false)
}

func BurningWithInflationDistributorTest(t *testing.T, passThroughAccoutType testutils.DestinationType, toValidators bool) {
	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	testapp.AddMaccPerms("c4e_distributor", nil)
	testapp.AddMaccPerms("no_validators", nil)
	denom := "uc4e"
	testapp.AddMaccPerms(collector, perms)
	testapp.AddMaccPerms(types.DistributorMainAccount, perms)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
	app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))

	cointToMintFromInflation := sdk.NewCoin(denom, sdk.NewInt(5044))
	app.BankKeeper.MintCoins(ctx, types.DistributorMainAccount, sdk.NewCoins(cointToMintFromInflation))
	require.EqualValues(t, sdk.NewCoin(denom, sdk.NewInt(1017+5044)), app.BankKeeper.GetSupply(ctx, denom))

	var subDistributors []types.SubDistributor

	subDistributors = append(subDistributors, testutils.PrepareBurningDistributor(testutils.MainCollector))
	if passThroughAccoutType != testutils.MainCollector {
		subDistributors = append(subDistributors, testutils.PrepareInflationToPassAcoutSubDistr(passThroughAccoutType))
	}
	subDistributors = append(subDistributors, testutils.PrepareInflationSubDistributor(passThroughAccoutType, toValidators))

	app.CfedistributorKeeper.SetParams(ctx, types.NewParams(subDistributors))
	ctx = ctx.WithBlockHeight(int64(2))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})

	if passThroughAccoutType == testutils.MainCollector {
		require.EqualValues(t, 3, len(app.CfedistributorKeeper.GetAllStates(ctx)))
	} else if passThroughAccoutType == testutils.ModuleAccount {
		require.EqualValues(t, 4, len(app.CfedistributorKeeper.GetAllStates(ctx)))
	} else if passThroughAccoutType == testutils.InternalAccount {
		require.EqualValues(t, 4, len(app.CfedistributorKeeper.GetAllStates(ctx)))
	} else {
		require.EqualValues(t, 4, len(app.CfedistributorKeeper.GetAllStates(ctx)))
	}

	// coins flow:
	// fee 1017*51% = 518.67 to burn, so 518 burned - and burn remains 0.67

	require.EqualValues(t, sdk.NewCoin(denom, sdk.NewInt(1017+5044-518)), app.BankKeeper.GetSupply(ctx, denom))
	burnState, _ := app.CfedistributorKeeper.GetBurnState(ctx)
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf("uc4e"))

	// added 499 to main collector
	// main collector state = 499 + 5044 = 5543, but 5543 - 0,67 = 5542.33 to distribute

	if passThroughAccoutType == testutils.ModuleAccount || passThroughAccoutType == testutils.InternalAccount {
		// 5542.33 moved to c4e_distributor module or internal account
		// and all is distributed further, and 0 in remains
		c4eDIstrCoins := app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "c4e_distributor")
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), c4eDIstrCoins.AmountOf(denom).ToDec())

		remains, _ := app.CfedistributorKeeper.GetState(ctx, "c4e_distributor")
		//require.EqualValues(t, passThroughAccoutType == ModuleAccount, remains.Account.IsModuleAccount)
		//require.EqualValues(t, passThroughAccoutType == InternalAccount, remains.Account.IsInternalAccount)
		//require.EqualValues(t, false, remains.Account.IsMainCollector)

		coinRemainsDevelopmentFund := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinRemainsDevelopmentFund.AmountOf("uc4e"))
	} else if passThroughAccoutType == testutils.BaseAccount {
		// 5542.33 moved to cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck account
		// and all is distributed further, and 0 in remains
		address, _ := sdk.AccAddressFromBech32("cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")

		c4eDIstrCoins := app.CfedistributorKeeper.GetAccountCoins(ctx, address)
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), c4eDIstrCoins.AmountOf(denom).ToDec())

		remains, _ := app.CfedistributorKeeper.GetState(ctx, "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
		//require.EqualValues(t, passThroughAccoutType == ModuleAccount, remains.Account.IsModuleAccount)
		//require.EqualValues(t, passThroughAccoutType == InternalAccount, remains.Account.IsInternalAccount)
		//require.EqualValues(t, false, remains.Account.IsMainCollector)

		coinRemainsDevelopmentFund := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinRemainsDevelopmentFund.AmountOf("uc4e"))
	}

	// 5542.33*10.345% = 573.3540385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 573 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.3540385 on its cfedistributor state

	acc, _ := sdk.AccAddressFromBech32("cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	developmentFundAccount := app.CfedistributorKeeper.GetAccountCoins(ctx, acc)
	require.EqualValues(t, sdk.MustNewDecFromStr("573"), developmentFundAccount.AmountOf(denom).ToDec())

	remains, _ := app.CfedistributorKeeper.GetState(ctx, "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	coinRemainsDevelopmentFund := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.3540385"), coinRemainsDevelopmentFund.AmountOf("uc4e"))

	// 5542.33 - 573.3540385 = 4968.9759615 to validators_rewards_collector, so
	// 4968 on validators_rewards_collector or no_validators module account and 0.9759615 on its cfedistributor state

	if toValidators {
		// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

		validatorRewardCollectorAccountCoin := app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.ValidatorsRewardsCollector)
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), validatorRewardCollectorAccountCoin.AmountOf(denom).ToDec())
		// still 0.9759615 on its cfedistributor state remains
		remains, _ = app.CfedistributorKeeper.GetState(ctx, types.ValidatorsRewardsCollector)
		coinRemainsValidatorsReward := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.9759615"), coinRemainsValidatorsReward.AmountOf("uc4e"))
		// and 4968 to validators rewards
		distrCoins := app.CfedistributorKeeper.GetAccountCoins(ctx, app.DistrKeeper.GetDistributionAccount(ctx).GetAddress())
		require.EqualValues(t, sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(4968))), distrCoins)
	} else {
		// no_validators module account coins amount is 4968,
		// and remains 0.9759615 on its cfedistributor state

		NoValidatorsCoin := app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "no_validators")
		require.EqualValues(t, sdk.MustNewDecFromStr("4968"), NoValidatorsCoin.AmountOf(denom).ToDec())

		remains, _ = app.CfedistributorKeeper.GetState(ctx, "no_validators")
		coinRemainsValidatorsReward := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.9759615"), coinRemainsValidatorsReward.AmountOf("uc4e"))
	}

	// 5543 - 573 - 4968 = 2 (its ramains 0,67 + 0.3540385 + 0.9759615 = 2) on main collector
	coinOnDistributorAccount :=
		app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
	require.EqualValues(t, sdk.MustNewDecFromStr("2"), coinOnDistributorAccount.AmountOf(denom).ToDec())

}

func TestBurningWithInflationDistributorAfter3001Blocks(t *testing.T) {
	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	denom := "uc4e"
	// inflationCollector := "c4e_distributor"
	testapp.AddMaccPerms(collector, perms)
	testapp.AddMaccPerms(types.DistributorMainAccount, perms)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, testutils.PrepareBurningDistributor(testutils.MainCollector))
	subdistributors = append(subdistributors, testutils.PrepareInflationSubDistributor(testutils.MainCollector, true))
	app.CfedistributorKeeper.SetParams(ctx, types.NewParams(subdistributors))

	for i := int64(1); i <= 3001; i++ {

		cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
		app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))

		cointToMintFromInflation := sdk.NewCoin(denom, sdk.NewInt(5044))
		app.BankKeeper.MintCoins(ctx, types.DistributorMainAccount, sdk.NewCoins(cointToMintFromInflation))
		ctx = ctx.WithBlockHeight(int64(i))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})
		burn, _ := sdk.NewDecFromStr("518.67")
		burn = burn.MulInt64(i)
		burn.GT(burn.TruncateDec())
		totalExpected := sdk.NewDec(i * (1017 + 5044)).Sub(burn)

		totalExpectedTruncated := totalExpected.TruncateInt()

		if burn.GT(burn.TruncateDec()) {
			totalExpectedTruncated = totalExpectedTruncated.AddRaw(1)
		}
		require.EqualValues(t, sdk.NewCoin(denom, totalExpectedTruncated).String(), app.BankKeeper.GetSupply(ctx, denom).String())

	}

	ctx = ctx.WithBlockHeight(int64(3002))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})
	app.EndBlocker(ctx, abci.RequestEndBlock{})

	// coins flow:
	// fee 3001*1017*51% = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	// fee 3001*(1017*51%) = 3001*518.67 = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67

	require.EqualValues(t, sdk.NewCoin(denom, sdk.NewInt(3001*(1017+5044)-1556528)), app.BankKeeper.GetSupply(ctx, denom))
	burnState, _ := app.CfedistributorKeeper.GetBurnState(ctx)
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf("uc4e"))

	// added 3001*1017 - 1556528 = 1495489 to main collector
	// main collector state = 1495489 + 3001*5044 = 16632533, but 16632533 - 0.67 (burning remains) = 16632532.33 to distribute

	// 16632532.33*10.345% = 1720635.4695385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 1720635 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.4695385 on its cfedistributor state

	acc, _ := sdk.AccAddressFromBech32("cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	developmentFundAccount := app.CfedistributorKeeper.GetAccountCoins(ctx, acc)
	require.EqualValues(t, sdk.MustNewDecFromStr("1720635"), developmentFundAccount.AmountOf(denom).ToDec())

	remains, _ := app.CfedistributorKeeper.GetState(ctx, "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	coinRemainsDevelopmentFund := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.4695385"), coinRemainsDevelopmentFund.AmountOf("uc4e"))

	// 16632532.33- 1720635.4695385 = 14911896.8604615 to validators_rewards_collector, so
	// 14911896 on validators_rewards_collector or no_validators module account and 0.8604615 on its cfedistributor state

	// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

	validatorRewardCollectorAccountCoin := app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.ValidatorsRewardsCollector)
	require.EqualValues(t, sdk.MustNewDecFromStr("0"), validatorRewardCollectorAccountCoin.AmountOf(denom).ToDec())
	// still 0.8845 on its cfedistributor state remains
	remains, _ = app.CfedistributorKeeper.GetState(ctx, types.ValidatorsRewardsCollector)
	coinRemainsValidatorsReward := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.8604615"), coinRemainsValidatorsReward.AmountOf("uc4e"))
	// and 14906927 to validators rewards
	distrCoins := app.CfedistributorKeeper.GetAccountCoins(ctx, app.DistrKeeper.GetDistributionAccount(ctx).GetAddress())
	require.EqualValues(t, sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(14911896))), distrCoins)

	// 16632533 - 1720635 - 14911896 = 1 (its ramains 0.67 + 0.4695385 + 0.8604615 = 2) on main collector
	coinOnDistributorAccount :=
		app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
	require.EqualValues(t, sdk.MustNewDecFromStr("2"), coinOnDistributorAccount.AmountOf(denom).ToDec())

}
