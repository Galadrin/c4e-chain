package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeKey     sdk.StoreKey
		memKey       sdk.StoreKey
		paramstore   paramtypes.Subspace
		bank         types.BankKeeper
		staking      types.StakingKeeper
		account      types.AccountKeeper
		distribution types.DistributionKeeper
		gov          types.GovKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bank types.BankKeeper,
	staking types.StakingKeeper,
	account types.AccountKeeper,
	distribution types.DistributionKeeper,
	gov types.GovKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		paramstore:   ps,
		bank:         bank,
		staking:      staking,
		account:      account,
		distribution: distribution,
		gov:          gov,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
