package block

import (
	"github.com/df-mc/dragonfly/server/item"
)

// IronOre is a mineral block found underground.
type IronOre struct {
	solid
	bassDrum

	// Type is the type of iron ore.
	Type OreType
}

// BreakInfo ...
func (i IronOre) BreakInfo() BreakInfo {
	b := NewBreakInfo(i.Type.Hardness(), func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, PickaxeEffective, SilkTouchOneOf(item.RawIron{}, i))
	if i.Type == DeepslateOre() {
		b = b.withBlastResistance(9)
	}
	return b
}

// SmeltInfo ...
func (IronOre) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.IronIngot{}, 1), 0.7)
}

// EncodeItem ...
func (i IronOre) EncodeItem() (name string, meta int16) {
	return "minecraft:" + i.Type.Prefix() + "iron_ore", 0
}

// EncodeBlock ...
func (i IronOre) EncodeBlock() (string, map[string]any) {
	return "minecraft:" + i.Type.Prefix() + "iron_ore", nil
}
