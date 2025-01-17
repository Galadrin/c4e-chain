package cfedistributor

import (
	"time"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	subDistributors := k.GetParams(ctx).SubDistributors
	states := k.GetAllStates(ctx)

	for _, subDistributor := range subDistributors {
		allCoinsToDistribute := sdk.NewDecCoins()
		for _, source := range subDistributor.Sources {
			var coinsToDistribute sdk.DecCoins
			if source.Type == types.MAIN {
				coinsToDistribute = k.PrepareCoinToDistributeForMainAccount(ctx, states, subDistributor.Name)
			} else {
				coinsToDistribute = k.PrepareCoinToDistributeForNotMainAccount(ctx, *source, states, subDistributor.Name)
			}

			if len(coinsToDistribute) == 0 {
				continue
			}
			allCoinsToDistribute = allCoinsToDistribute.Add(coinsToDistribute...)
		}

		if allCoinsToDistribute.IsZero() {
			continue
		}

		newStates, distributions, burn := k.StartDistributionProcess(ctx, &states, allCoinsToDistribute, subDistributor)
		for _, distribution := range distributions {
			err := ctx.EventManager().EmitTypedEvent(distribution)
			if err != nil {
				k.Logger(ctx).Error("distributions emit event error", "distribution", distribution, "error", err.Error())
			}
		}
		if burn != nil {
			err := ctx.EventManager().EmitTypedEvent(burn)
			if err != nil {
				k.Logger(ctx).Error("distributions emit event error", "distribution_burn", burn, "error", err.Error())
			}
		}
		states = *newStates
	}



	k.SendCoinsFromStates(ctx, states)
}
