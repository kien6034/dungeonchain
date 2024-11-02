package v1

import (
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"cosmossdk.io/math"
	"github.com/Crypto-Dungeon/dungeonchain/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RunForkLogic(ctx sdk.Context, ak *upgrades.AppKeepers) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.Logger().Info("Starting fork logic...")

	// Mint 500M DGN to first address
	addr1, err := sdk.AccAddressFromBech32("dungeon1ctzzm9rvdtzqd9j46tj9ezrfyzfe2vr7fqyv8j")
	if err != nil {
		panic(err)
	}
	coins1 := sdk.NewCoins(sdk.NewCoin("udgn", math.NewInt(500_000_000_000_000))) // 500M * 1_000_000 (for 6 decimals)
	if err := ak.BankKeeper.MintCoins(sdkCtx, minttypes.ModuleName, coins1); err != nil {
		panic(err)
	}
	if err := ak.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, minttypes.ModuleName, addr1, coins1); err != nil {
		panic(err)
	}

	// Mint 150M DGN to second address
	addr2, err := sdk.AccAddressFromBech32("dungeon1aj5jlmvqp8dd26rsec6624szthlazdn2vhxak9")
	if err != nil {
		panic(err)
	}
	coins2 := sdk.NewCoins(sdk.NewCoin("udgn", math.NewInt(150_000_000_000_000))) // 150M * 1_000_000 (for 6 decimals)
	if err := ak.BankKeeper.MintCoins(sdkCtx, minttypes.ModuleName, coins2); err != nil {
		panic(err)
	}
	if err := ak.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, minttypes.ModuleName, addr2, coins2); err != nil {
		panic(err)
	}
}
