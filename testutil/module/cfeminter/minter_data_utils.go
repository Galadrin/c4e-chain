package cfeminter

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
)

func CompareMinters(t *testing.T, m1 types.Minter, m2 types.Minter) {
	require.True(t, m1.Start.Equal(m2.Start))
	for i, p1 := range m1.Periods {
		p2 := m2.Periods[i]
		if p1.PeriodEnd == nil {
			require.Nil(t, p2.PeriodEnd)
		} else {
			require.True(t, p1.PeriodEnd.Equal(*p2.PeriodEnd))
		}
		require.EqualValues(t, p1.Position, p2.Position)
		require.EqualValues(t, p1.TimeLinearMinter, p2.TimeLinearMinter)
		require.EqualValues(t, p1.Type, p2.Type)
	}
}

func CompareMinterStates(t *testing.T, expected types.MinterState, state types.MinterState) {
	require.EqualValues(t, expected.Position, state.Position)
	require.Truef(t, expected.AmountMinted.Equal(state.AmountMinted), "expected.AmountMinted %s <> state.AmountMinted %s", expected.AmountMinted, state.AmountMinted)
	require.Truef(t, expected.RemainderToMint.Equal(state.RemainderToMint), "expected.RemainderToMint %s <> state.RemainderToMint %s", expected.RemainderToMint, state.RemainderToMint)
	require.Truef(t, expected.LastMintBlockTime.Equal(state.LastMintBlockTime), "expected.LastMintBlockTime %s <> state.LastMintBlockTime %s", expected.LastMintBlockTime.Local(), state.LastMintBlockTime.Local())
	require.Truef(t, expected.RemainderFromPreviousPeriod.Equal(state.RemainderFromPreviousPeriod), "expected.RemainderFromPreviousPeriod %s <> state.RemainderFromPreviousPeriod %s", expected.RemainderFromPreviousPeriod, state.RemainderFromPreviousPeriod)
}
