package block

// QuartzBricks is a mineral block used only for decoration.
type QuartzBricks struct {
	solid
	bassDrum
}

// BreakInfo ...
func (q QuartzBricks) BreakInfo() BreakInfo {
	return NewBreakInfo(0.8, PickaxeHarvestable, PickaxeEffective, OneOf(q))
}

// EncodeItem ...
func (QuartzBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:quartz_bricks", 0
}

// EncodeBlock ...
func (QuartzBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:quartz_bricks", nil
}
