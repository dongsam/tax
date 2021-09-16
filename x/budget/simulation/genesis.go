package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/tendermint/budget/x/budget/types"
)

// DONTCOVER

// Simulation parameter constants
const (
	EpochBlocks = "epoch_blocks"
	Budgets     = "budgets"
)

// GenEpochBlocks return default DefaultEpochBlocks
func GenEpochBlocks(r *rand.Rand) uint32 {
	// TODO: randomize
	return types.DefaultEpochBlocks
}

// GenGenBudgets return randomized budgets
func GenBudgets(r *rand.Rand) []types.Budget {
	// TODO: randomize
	budgetSamples := []types.Budget{
		{
			Name:                "budget1",
			Rate:                sdk.MustNewDecFromStr("0.5"),
			BudgetSourceAddress: "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta", // Address corresponding to fee_collector module account in cosmoshub case
			CollectionAddress:   sdk.AccAddress(address.Module(types.ModuleName, []byte("GravityDEXFarmingBudget"))).String(),
			StartTime:           types.ParseTime("2000-01-01T00:00:00Z"),
			EndTime:             types.ParseTime("9999-12-31T00:00:00Z"),
		},
	}
	return budgetSamples
}

// RandomizedGenState generates a random GenesisState for budget
func RandomizedGenState(simState *module.SimulationState) {
	var epochBlocks uint32
	var budgets []types.Budget
	simState.AppParams.GetOrGenerate(
		simState.Cdc, EpochBlocks, &epochBlocks, simState.Rand,
		func(r *rand.Rand) { epochBlocks = GenEpochBlocks(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, Budgets, &budgets, simState.Rand,
		func(r *rand.Rand) { budgets = GenBudgets(r) },
	)

	budgetGenesis := types.GenesisState{
		Params: types.Params{
			EpochBlocks: epochBlocks,
			Budgets:     budgets,
		},
	}

	bz, _ := json.MarshalIndent(&budgetGenesis, "", " ")
	fmt.Printf("Selected randomly generated budget parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&budgetGenesis)
}