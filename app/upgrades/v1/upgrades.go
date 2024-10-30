package v1

import (
	"context"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/Crypto-Dungeon/dungeonchain/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func CreateUpgradeHandler(mm upgrades.ModuleManager,
	configurator module.Configurator,
	ak *upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		sdkCtx.Logger().Info("Starting module migrations...")

		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		// Mint 500M DGN to first address
		addr1, err := sdk.AccAddressFromBech32("dungeon1ctzzm9rvdtzqd9j46tj9ezrfyzfe2vr7fqyv8j")
		if err != nil {
			return vm, err
		}
		coins1 := sdk.NewCoins(sdk.NewCoin("udgn", math.NewInt(500_000_000_000_000))) // 500M * 1_000_000 (for 6 decimals)
		if err := ak.BankKeeper.MintCoins(sdkCtx, minttypes.ModuleName, coins1); err != nil {
			return vm, err
		}
		if err := ak.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, minttypes.ModuleName, addr1, coins1); err != nil {
			return vm, err
		}

		// Mint 150M DGN to second address
		addr2, err := sdk.AccAddressFromBech32("dungeon1aj5jlmvqp8dd26rsec6624szthlazdn2vhxak9")
		if err != nil {
			return vm, err
		}
		coins2 := sdk.NewCoins(sdk.NewCoin("udgn", math.NewInt(150_000_000_000_000))) // 150M * 1_000_000 (for 6 decimals)
		if err := ak.BankKeeper.MintCoins(sdkCtx, minttypes.ModuleName, coins2); err != nil {
			return vm, err
		}
		if err := ak.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, minttypes.ModuleName, addr2, coins2); err != nil {
			return vm, err
		}

		return vm, nil
	}
}
