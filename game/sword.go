package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
)

func SwordLeft(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		pos := p.CenterPos().Add(p.Dir.Copy().Rotate(-55).Scale(7))
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(-90).Normalize(), 5)
		basePos := pos.Copy()
		for j := -55.0; j <= 45.0; j += 10.0 {
			yDelta := p.Dir.Copy().Rotate(j).Scale(4)
			pos = basePos.Copy()
			for i := 0; i < 4; i++ {
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv)
				pos.Add(yDelta)
			}
		}
	}
}
func SwordRight(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		pos := p.CenterPos().Add(p.Dir.Copy().Rotate(55).Scale(7))
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(90).Normalize(), 5)
		basePos := pos.Copy()
		for j := 55.0; j >= -45.0; j -= 10.0 {
			yDelta := p.Dir.Copy().Rotate(j).Scale(4)
			pos = basePos.Copy()
			for i := 0; i < 4; i++ {
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv)
				pos.Add(yDelta)
			}
		}
	}
}

func SwordDash(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		p.Delta.Add(p.Dir.Copy().Scale(24 * p.Speed.Y()))
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(180).Normalize(), 10)
		delta := p.Dir.Copy().Scale(3)
		perpendicular := delta.Copy().Rotate(90)
		pos := p.CenterPos().Add(delta, perpendicular, perpendicular)
		perpendicular.Scale(-1)
		basePos := pos.Copy()
		for i := 0; i < 3; i++ {
			pos = basePos.Add(perpendicular).Copy()
			for j := 0; j < 12; j++ {
				pos.Add(delta)
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv)
			}
		}
	}
}
