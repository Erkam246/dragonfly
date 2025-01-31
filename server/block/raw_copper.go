package block

import (
	"github.com/df-mc/dragonfly/server/item"
)

// RawCopper is a raw metal block equivalent to nine raw copper.
type RawCopper struct {
	solid
	bassDrum
}

// BreakInfo ...
func (r RawCopper) BreakInfo() BreakInfo {
	return NewBreakInfo(5, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, PickaxeEffective, OneOf(r)).withBlastResistance(30)
}

// EncodeItem ...
func (RawCopper) EncodeItem() (name string, meta int16) {
	return "minecraft:raw_copper_block", 0
}

// EncodeBlock ...
func (RawCopper) EncodeBlock() (string, map[string]any) {
	return "minecraft:raw_copper_block", nil
}
