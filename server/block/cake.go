package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// Cake is an edible block.
type Cake struct {
	transparent
	sourceWaterDisplacer

	// Bites is the amount of bites taken out of the cake.
	Bites int
}

// SideClosed ...
func (c Cake) SideClosed(cube.Pos, cube.Pos, *world.World) bool {
	return false
}

// UseOnBlock ...
func (c Cake) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	pos, _, used := FirstReplaceable(w, pos, face, c)
	if !used {
		return false
	}

	if _, air := w.Block(pos.Side(cube.FaceDown)).(Air); air {
		return false
	}

	Place(w, pos, c, user, ctx)
	return Placed(ctx)
}

// NeighbourUpdateTick ...
func (c Cake) NeighbourUpdateTick(pos, _ cube.Pos, w *world.World) {
	if _, air := w.Block(pos.Side(cube.FaceDown)).(Air); air {
		w.SetBlock(pos, nil, nil)
		w.AddParticle(pos.Vec3Centre(), particle.BlockBreak{Block: c})
	}
}

// Activate ...
func (c Cake) Activate(pos cube.Pos, _ cube.Face, w *world.World, u item.User, _ *item.UseContext) bool {
	if i, ok := u.(interface {
		Saturate(food int, saturation float64)
	}); ok {
		i.Saturate(2, 0.4)
		w.PlaySound(u.Position().Add(mgl64.Vec3{0, 1.5}), sound.Burp{})
		c.Bites++
		if c.Bites > 6 {
			w.SetBlock(pos, nil, nil)
			return true
		}
		w.SetBlock(pos, c, nil)
		return true
	}
	return false
}

// BreakInfo ...
func (c Cake) BreakInfo() BreakInfo {
	return NewBreakInfo(0.5, NeverHarvestable, NothingEffective, SimpleDrops())
}

// EncodeItem ...
func (c Cake) EncodeItem() (name string, meta int16) {
	return "minecraft:cake", 0
}

// EncodeBlock ...
func (c Cake) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:cake", map[string]any{"bite_counter": int32(c.Bites)}
}

// Model ...
func (c Cake) Model() world.BlockModel {
	return model.Cake{Bites: c.Bites}
}

// allCake ...
func allCake() (cake []world.Block) {
	for bites := 0; bites < 7; bites++ {
		cake = append(cake, Cake{Bites: bites})
	}
	return
}
