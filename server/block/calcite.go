package block

// Calcite is a carbonate mineral found as part of amethyst geodes.
type Calcite struct {
	solid
	bassDrum
}

// BreakInfo ...
func (c Calcite) BreakInfo() BreakInfo {
	return NewBreakInfo(0.75, PickaxeHarvestable, PickaxeEffective, OneOf(c))
}

// EncodeItem ...
func (c Calcite) EncodeItem() (name string, meta int16) {
	return "minecraft:calcite", 0
}

// EncodeBlock ...
func (c Calcite) EncodeBlock() (string, map[string]any) {
	return "minecraft:calcite", nil
}
