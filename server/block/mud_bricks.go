package block

// MudBricks are a decorative block obtained by crafting 4 packed mud.
type MudBricks struct {
	solid
	bassDrum
}

// BreakInfo ...
func (m MudBricks) BreakInfo() BreakInfo {
	return NewBreakInfo(2, AlwaysHarvestable, NothingEffective, OneOf(m))
}

// EncodeItem ...
func (MudBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:mud_bricks", 0
}

// EncodeBlock ...
func (MudBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:mud_bricks", nil
}
