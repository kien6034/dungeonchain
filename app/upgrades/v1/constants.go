package v1

import (
	"github.com/Crypto-Dungeon/dungeonchain/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v1"
)

var Upgrade = upgrades.Fork{
	UpgradeName:    UpgradeName,
	UpgradeHeight:  1083549,
	BeginForkLogic: RunForkLogic,
}
