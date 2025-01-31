package block

import (
	"github.com/df-mc/dragonfly/server/item"
)

// DiamondOre is a rare ore that generates underground.
type DiamondOre struct {
	solid
	bassDrum

	// Type is the type of diamond ore.
	Type OreType
}

// BreakInfo ...
func (d DiamondOre) BreakInfo() BreakInfo {
	i := NewBreakInfo(d.Type.Hardness(), func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierIron.HarvestLevel
	}, PickaxeEffective, SilkTouchOneOf(item.Diamond{}, d)).withXPDropRange(3, 7)
	if d.Type == DeepslateOre() {
		i = i.withBlastResistance(9)
	}
	return i
}

// SmeltInfo ...
func (DiamondOre) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.Diamond{}, 1), 1)
}

// EncodeItem ...
func (d DiamondOre) EncodeItem() (name string, meta int16) {
	return "minecraft:" + d.Type.Prefix() + "diamond_ore", 0
}

// EncodeBlock ...
func (d DiamondOre) EncodeBlock() (string, map[string]any) {
	return "minecraft:" + d.Type.Prefix() + "diamond_ore", nil
}
