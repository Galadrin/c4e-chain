package cfedistributor

import (
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	randDistinationType := testutils.RandomCollectorName(simState.Rand)
	subdistributors = append(subdistributors, testutils.PrepareBurningDistributor(randDistinationType))
	subdistributors = append(subdistributors, testutils.PrepareInflationSubDistributor(randDistinationType, true))
	subdistributors = append(subdistributors, testutils.PrepareInflationToPassAcoutSubDistr(randDistinationType))

	genesisState.Params.SubDistributors = subdistributors
	cfedistributorGenesis := types.GenesisState{
		Params: types.NewParams(subdistributors),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfedistributorGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	// No decoder
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
