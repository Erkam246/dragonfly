package block

// NetherWartBlock is a decorative block found in crimson forests and crafted using Nether wart.
type NetherWartBlock struct {
	solid

	// Warped is the turquoise variant found in warped forests, but cannot be crafted unlike Nether wart block.
	Warped bool
}

// BreakInfo ...
func (n NetherWartBlock) BreakInfo() BreakInfo {
	return NewBreakInfo(1, AlwaysHarvestable, HoeEffective, OneOf(n))
}

// CompostChance ...
func (NetherWartBlock) CompostChance() float64 {
	return 0.85
}

// EncodeItem ...
func (n NetherWartBlock) EncodeItem() (name string, meta int16) {
	if n.Warped {
		return "minecraft:warped_wart_block", 0
	}
	return "minecraft:nether_wart_block", 0
}

// EncodeBlock ...
func (n NetherWartBlock) EncodeBlock() (name string, properties map[string]interface{}) {
	if n.Warped {
		return "minecraft:warped_wart_block", nil
	}
	return "minecraft:nether_wart_block", nil
}
