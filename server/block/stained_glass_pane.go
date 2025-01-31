package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// StainedGlassPane is a transparent block that can be used as a more efficient alternative to glass blocks.
type StainedGlassPane struct {
	transparent
	thin
	clicksAndSticks
	sourceWaterDisplacer

	// Colour specifies the colour of the block.
	Colour item.Colour
}

// SideClosed ...
func (p StainedGlassPane) SideClosed(cube.Pos, cube.Pos, *world.World) bool {
	return false
}

// BreakInfo ...
func (p StainedGlassPane) BreakInfo() BreakInfo {
	return NewBreakInfo(0.3, AlwaysHarvestable, NothingEffective, SilkTouchOnlyDrop(p))
}

// EncodeItem ...
func (p StainedGlassPane) EncodeItem() (name string, meta int16) {
	return "minecraft:" + p.Colour.String() + "_stained_glass_pane", 0
}

// EncodeBlock ...
func (p StainedGlassPane) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + p.Colour.String() + "_stained_glass_pane", nil
}

// allStainedGlassPane returns stained-glass panes with all possible colours.
func allStainedGlassPane() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, c := range item.Colours() {
		b = append(b, StainedGlassPane{Colour: c})
	}
	return b
}
