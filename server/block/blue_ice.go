package block

// BlueIce is a solid block similar to packed ice.
type BlueIce struct {
	solid
}

// BreakInfo ...
func (b BlueIce) BreakInfo() BreakInfo {
	return NewBreakInfo(2.8, AlwaysHarvestable, PickaxeEffective, SilkTouchOnlyDrop(b))
}

// Friction ...
func (b BlueIce) Friction() float64 {
	return 0.989
}

// EncodeItem ...
func (BlueIce) EncodeItem() (name string, meta int16) {
	return "minecraft:blue_ice", 0
}

// EncodeBlock ...
func (BlueIce) EncodeBlock() (string, map[string]any) {
	return "minecraft:blue_ice", nil
}
