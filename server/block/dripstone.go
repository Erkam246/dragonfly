package block

// Dripstone is a rock block that allows pointed dripstone to grow beneath it.
type Dripstone struct {
	solid
	bassDrum
}

// BreakInfo ...
func (d Dripstone) BreakInfo() BreakInfo {
	return NewBreakInfo(1.5, PickaxeHarvestable, PickaxeEffective, OneOf(d)).withBlastResistance(5)
}

// EncodeItem ...
func (d Dripstone) EncodeItem() (name string, meta int16) {
	return "minecraft:dripstone_block", 0
}

// EncodeBlock ...
func (d Dripstone) EncodeBlock() (string, map[string]any) {
	return "minecraft:dripstone_block", nil
}
