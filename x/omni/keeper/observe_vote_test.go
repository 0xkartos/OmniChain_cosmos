package keeper_test

import (
	"strconv"
	"testing"

	keepertest "omni/testutil/keeper"
	"omni/testutil/nullify"
	"omni/x/omni/keeper"
	"omni/x/omni/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNObserveVote(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ObserveVote {
	items := make([]types.ObserveVote, n)
	for i := range items {
		items[i].Index = (uint64)(i)
		items[i].Voter = ""
		items[i].Round = 0
		items[i].Value = sdk.NewInt(1)

		keeper.SetObserveVote(ctx, items[i])
	}
	return items
}

func TestObserveVoteGet(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNObserveVote(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetObserveVote(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestObserveVoteRemove(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNObserveVote(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveObserveVote(ctx,
			item.Index,
		)
		_, found := keeper.GetObserveVote(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestObserveVoteGetAll(t *testing.T) {
	keeper, ctx := keepertest.OmniKeeper(t)
	items := createNObserveVote(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllObserveVote(ctx)),
	)
}
