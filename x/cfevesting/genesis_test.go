package cfevesting_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"


	"github.com/chain4energy/c4e-chain/app"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestGenesisWholeApp(t *testing.T) {

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		// this line is used by starport scaffolding # genesis/test/state
		VestingTypes: []types.GenesisVestingType{},
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	cfevesting.InitGenesis(ctx, app.CfevestingKeeper, genesisState, app.AccountKeeper)
	got := cfevesting.ExportGenesis(ctx, app.CfevestingKeeper)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisVestingTypes(t *testing.T) {
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params:       types.NewParams("test_denom"),
		VestingTypes: vestingTypesArray,
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	cfevesting.InitGenesis(ctx, k, genesisState, ak)
	got := cfevesting.ExportGenesis(ctx, k)

	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func TestGenesisVestingTypesUnitsSecondsToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*60*24, keeper.Second, keeper.Day)
}

func TestGenesisVestingTypesUnitsSecondsToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*60, keeper.Second, keeper.Hour)
}

func TestGenesisVestingTypesUnitsSecondsToMinutes(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60, keeper.Second, keeper.Minute)
}

func TestGenesisVestingTypesUnitsSecondsToSeconds(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Second, keeper.Second)
}


func TestGenesisVestingTypesUnitsMinutesToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*24, keeper.Minute, keeper.Day)
}

func TestGenesisVestingTypesUnitsMinutesToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60, keeper.Minute, keeper.Hour)
}

func TestGenesisVestingTypesUnitsMinutesToMinutes(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Minute, keeper.Minute)
}

func TestGenesisVestingTypesUnitsHoursToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 24, keeper.Hour, keeper.Day)
}

func TestGenesisVestingTypesUnitsHoursToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Hour, keeper.Hour)
}

func TestGenesisVestingTypesUnitsDaysToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Day, keeper.Day)
}

func genesisVestingTypesUnitsTest(t *testing.T, multiplier int64, srcUnits string, dstUnits string) {
	vestingTypesArray := generateGenesisVestingTypes(1, 1)
	vestingTypesArray[0].LockupPeriod = 234 * multiplier
	vestingTypesArray[0].LockupPeriodUnit = srcUnits

	vestingTypesArray[0].VestingPeriod = 345 * multiplier
	vestingTypesArray[0].VestingPeriodUnit = srcUnits

	vestingTypesArray[0].TokenReleasingPeriod = 23 * multiplier
	vestingTypesArray[0].TokenReleasingPeriodUnit = srcUnits

	genesisState := types.GenesisState{
		Params:       types.NewParams("test_denom"),
		VestingTypes: vestingTypesArray,
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	cfevesting.InitGenesis(ctx, k, genesisState, ak)

	vestingTypesArray[0].LockupPeriod = 234
	vestingTypesArray[0].LockupPeriodUnit = dstUnits

	vestingTypesArray[0].VestingPeriod = 345
	vestingTypesArray[0].VestingPeriodUnit = dstUnits

	vestingTypesArray[0].TokenReleasingPeriod = 23
	vestingTypesArray[0].TokenReleasingPeriodUnit = dstUnits

	got := cfevesting.ExportGenesis(ctx, k)

	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func TestGenesisAccountVestingsList(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	cfevesting.InitGenesis(ctx, k, genesisState, ak)
	got := cfevesting.ExportGenesis(ctx, k)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState.Params, got.GetParams())
	require.EqualValues(t, genesisState.VestingTypes, (*got).VestingTypes)
	require.EqualValues(t, len(accountVestingsListArray), len((*got).AccountVestingsList.Vestings))

	testutils.AssertAccountVestingsArrays(t, accountVestingsListArray, (*got).AccountVestingsList.Vestings)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}

func TestDurationFromUnits(t *testing.T) {
	amount := int64(456)
	require.EqualValues(t, amount*int64(time.Second), keeper.DurationFromUnits(keeper.Second, amount))
	require.EqualValues(t, amount*int64(time.Minute), keeper.DurationFromUnits(keeper.Minute, amount))
	require.EqualValues(t, amount*int64(time.Hour), keeper.DurationFromUnits(keeper.Hour, amount))
	require.EqualValues(t, amount*int64(time.Hour*24), keeper.DurationFromUnits(keeper.Day, amount))

}

func TestDurationFromUnitsWrongUnit(t *testing.T) {
	require.PanicsWithError(t, "Unknown PeriodUnit: das: invalid type", func() { keeper.DurationFromUnits("das", 234) }, "")

}

func TestUnitsFromDuration(t *testing.T) {
	unit, amount := keeper.UnitsFromDuration(234 * time.Second)
	require.EqualValues(t, keeper.Second, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = keeper.UnitsFromDuration(234 * 60 * time.Second)
	require.EqualValues(t, keeper.Minute, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = keeper.UnitsFromDuration(234 * 60 * 60 * time.Second)
	require.EqualValues(t, keeper.Hour, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = keeper.UnitsFromDuration(234 * 60 * 60 * 24 * time.Second)
	require.EqualValues(t, keeper.Day, unit)
	require.EqualValues(t, 234, amount)
}

func generateGenesisVestingTypes(numberOfVestingTypes int, startId int) []types.GenesisVestingType {
	vts := testutils.GenerateVestingTypes(numberOfVestingTypes, startId)
	result := []types.GenesisVestingType{}
	for _, vt := range vts {

		gvt := types.GenesisVestingType{
			Name:                     vt.Name,
			LockupPeriod:             vt.LockupPeriod.Nanoseconds() / int64(time.Hour),
			LockupPeriodUnit:         keeper.Day,
			VestingPeriod:            vt.VestingPeriod.Nanoseconds() / int64(time.Hour),
			VestingPeriodUnit:        keeper.Day,
			TokenReleasingPeriod:     vt.TokenReleasingPeriod.Nanoseconds() / int64(time.Hour),
			TokenReleasingPeriodUnit: keeper.Day,
			DelegationsAllowed:       vt.DelegationsAllowed,
		}
		result = append(result, gvt)
	}
	return result
}
