package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/VelaChain/orion/testutil/keeper"
	"github.com/VelaChain/orion/x/amm/keeper"
	"github.com/VelaChain/orion/x/amm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AmmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
