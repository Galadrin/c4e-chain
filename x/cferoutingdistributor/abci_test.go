package cferoutingdistributor_test

import (
	testapp "github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func prepareBurningDistributor() types.RoutingDistributor {
	destAccount := types.Account{
		Address:         "c4e_distributor",
		IsModuleAccount: true,
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("51"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: burnShare,
	}

	distributor1 := types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []string{"fee_collector"},
		Destination: destination,
		Order:       0,
	}

	routingDistributor := types.RoutingDistributor{
		SubDistributor:           []types.SubDistributor{distributor1},
		ModuleAccounts:           nil,
		RemainsCoinModuleAccount: "remains",
	}

	return routingDistributor
}

func prepareInflationSubDistributor() types.SubDistributor {

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	destAccount := types.Account{
		Address:         "validators_rewards_collector",
		IsModuleAccount: true,
	}

	shareDevelopmentFundAccount := types.Account{
		Address:         "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag",
		IsModuleAccount: false,
	}

	shareDevelopmentFund := types.Share{
		Name:    "development_fund",
		Percent: sdk.MustNewDecFromStr("10.345"),
		Account: shareDevelopmentFundAccount,
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     []types.Share{shareDevelopmentFund},
		BurnShare: burnShare,
	}

	return types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []string{"c4e_distributor"},
		Destination: destination,
		Order:       0,
	}
}

func TestBurningDistributor(t *testing.T) {

	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	denom := "uc4e"
	testapp.AddMaccPerms(collector, perms)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
	app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))

	app.CferoutingdistributorKeeper.SetRoutingDistributor(ctx, prepareBurningDistributor())
	ctx = ctx.WithBlockHeight(int64(2))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})

	//coin on "c4e_distributor" should be equal 498, remains: 1 and 0.33 on remains
	coinRemains := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["c4e_distributor"].LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains)

	coinOnRemainAccount := app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "remains")
	require.EqualValues(t, sdk.NewInt(1), coinOnRemainAccount.AmountOf(denom))

	coinAfterDistribution :=
		app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "c4e_distributor")

	require.EqualValues(t, sdk.NewInt(498), coinAfterDistribution.AmountOf(denom))
}

func TestBurningWithInflationDistributor(t *testing.T) {
	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	denom := "uc4e"
	inflationCollector := "c4e_distributor"
	testapp.AddMaccPerms(collector, perms)
	testapp.AddMaccPerms(inflationCollector, perms)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
	app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))

	//prepare coin minted from inflation 899
	cointToMintFromInflation := sdk.NewCoin(denom, sdk.NewInt(5044))
	app.BankKeeper.MintCoins(ctx, inflationCollector, sdk.NewCoins(cointToMintFromInflation))

	routingDistibutor := prepareBurningDistributor()
	subDistributors := append(routingDistibutor.SubDistributor, prepareInflationSubDistributor())
	routingDistibutor.SubDistributor = subDistributors

	app.CferoutingdistributorKeeper.SetRoutingDistributor(ctx, routingDistibutor)
	ctx = ctx.WithBlockHeight(int64(2))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})

	//c4e_distributor should be empty
	coinOnDistributorAccount :=
		app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "c4e_distributor")
	require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinOnDistributorAccount.AmountOf(denom).ToDec())

	//coin on tx_fee_distributor distributor should have 0.33 remains left
	coinRemains := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["c4e_distributor"].LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains)

	//development_fund account should have 573
	acc, _ := sdk.AccAddressFromBech32("cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	developmentFundAccount := app.CferoutingdistributorKeeper.GetAccountCoins(ctx, acc)
	require.EqualValues(t, sdk.MustNewDecFromStr("573"), developmentFundAccount.AmountOf(denom).ToDec())

	//development_fund account  remains should have 0.00955
	coinRemainsDevelopmentFund := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag"].LeftoverCoin

	require.EqualValues(t, sdk.MustNewDecFromStr("0.3199"), coinRemainsDevelopmentFund)

	//validators_rewards_collector should have be 0, distributor getting the coin
	validatorRewardCollectorAccountCoin := app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "validators_rewards_collector")
	require.EqualValues(t, sdk.MustNewDecFromStr("0"), validatorRewardCollectorAccountCoin.AmountOf(denom).ToDec())

	coinRemainsValidatorsReward := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["validators_rewards_collector"].LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.6801"), coinRemainsValidatorsReward)

	//coins on remains module account should be equal 2
	coinOnRemainAccount := app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "remains")
	require.EqualValues(t, sdk.NewInt(2), coinOnRemainAccount.AmountOf(denom))
}

func TestBurningWithInflationDistributorAfter3001Blocks(t *testing.T) {
	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	denom := "uc4e"
	inflationCollector := "c4e_distributor"
	testapp.AddMaccPerms(collector, perms)
	testapp.AddMaccPerms(inflationCollector, perms)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	routingDistibutor := prepareBurningDistributor()
	subDistributors := append(routingDistibutor.SubDistributor, prepareInflationSubDistributor())
	routingDistibutor.SubDistributor = subDistributors

	app.CferoutingdistributorKeeper.SetRoutingDistributor(ctx, routingDistibutor)

	//coin from fee 1017 * 3000
	for i := 1; i <= 3001; i++ {
		//prepare module account with coin to distribute fee_collector 1017

		cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
		app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))

		//prepare coin minted from inflation 899
		cointToMintFromInflation := sdk.NewCoin(denom, sdk.NewInt(5044))
		app.BankKeeper.MintCoins(ctx, inflationCollector, sdk.NewCoins(cointToMintFromInflation))
		ctx = ctx.WithBlockHeight(int64(i))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})
	}

	ctx = ctx.WithBlockHeight(int64(3002))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})
	app.EndBlocker(ctx, abci.RequestEndBlock{})

	coinRemains := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["c4e_distributor"].LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains)

	burnRemains := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["burn"].LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.670000000000000000"), burnRemains)

	acc, _ := sdk.AccAddressFromBech32("cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	developmentFundAccount := app.CferoutingdistributorKeeper.GetAccountCoins(ctx, acc)
	require.EqualValues(t, sdk.MustNewDecFromStr("1720635"), developmentFundAccount.AmountOf(denom).ToDec())

	coinRemainsDevelopmentFund := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag"].LeftoverCoin

	require.EqualValues(t, sdk.MustNewDecFromStr("0.4354"), coinRemainsDevelopmentFund)

	validatorRewardCollectorAccountCoin := app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "validators_rewards_collector")
	require.EqualValues(t, sdk.MustNewDecFromStr("0"), validatorRewardCollectorAccountCoin.AmountOf(denom).ToDec())

	coinRemainsValidatorsReward := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).RemainsMap["validators_rewards_collector"].LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.5646"), coinRemainsValidatorsReward)

	coinOnRemainAccount := app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "remains")
	require.EqualValues(t, sdk.NewInt(2), coinOnRemainAccount.AmountOf(denom))
}
