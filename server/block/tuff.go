package block

// Tuff is an ornamental rock formed from volcanic ash, occurring in underground blobs below Y=16.
type Tuff struct {
	solid
	bassDrum
}

// BreakInfo ...
func (t Tuff) BreakInfo() BreakInfo {
	return NewBreakInfo(1.5, PickaxeHarvestable, PickaxeEffective, OneOf(t)).withBlastResistance(30)
}

// EncodeItem ...
func (t Tuff) EncodeItem() (name string, meta int16) {
	return "minecraft:tuff", 0
}

// EncodeBlock ...
func (t Tuff) EncodeBlock() (string, map[string]any) {
	return "minecraft:tuff", nil
}
