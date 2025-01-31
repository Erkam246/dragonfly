package block

// ReinforcedDeepslate is a tough decorative block that spawns in ancient cities.
type ReinforcedDeepslate struct {
	solid
	bassDrum
}

// BreakInfo ...
func (r ReinforcedDeepslate) BreakInfo() BreakInfo {
	return NewBreakInfo(55, AlwaysHarvestable, NothingEffective, OneOf(r)).withBlastResistance(3600)
}

// EncodeItem ...
func (ReinforcedDeepslate) EncodeItem() (name string, meta int16) {
	return "minecraft:reinforced_deepslate", 0
}

// EncodeBlock ...
func (ReinforcedDeepslate) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:reinforced_deepslate", nil
}
