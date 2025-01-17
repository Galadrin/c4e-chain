package v101

import (
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getAllV100SubDistributorStatesAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (list []v100cfedistributor.State, err error) {
	prefixStore := prefix.NewStore(store, v100cfedistributor.RemainsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v100cfedistributor.State
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewSubDistributorStates(store sdk.KVStore, cdc codec.BinaryCodec, oldStates []v100cfedistributor.State) error {
	prefixStore := prefix.NewStore(store, types.StateKeyPrefix)

	for _, oldState := range oldStates {
		var newAccount *types.Account
		if oldState.Burn == true {
			newAccount = nil
		} else {
			newAccount = &types.Account{
				Id:   oldState.Account.Id,
				Type: oldState.Account.Type,
			}
		}

		newState := types.State{
			Account: newAccount,
			Burn:    oldState.Burn,
			Remains: oldState.CoinsStates,
		}
		err := newState.Validate()
		if err != nil {
			return err
		}
		av, err := cdc.Marshal(&newState)
		if err != nil {
			return err
		}
		prefixStore.Set([]byte(newState.GetStateKey()), av)
	}
	return nil
}

func migrateSubdistributorStates(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldDistributorStates, err := getAllV100SubDistributorStatesAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewSubDistributorStates(store, cdc, oldDistributorStates)
}

//MigrateStore performs in-place store migrations from v1.0.0 to v1.0.1. The
//migration includes:
//
//- SubDistributor State rename CoinStates to Remains.
//- If burn is set to true state account must be nil
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	return migrateSubdistributorStates(store, cdc)
}
