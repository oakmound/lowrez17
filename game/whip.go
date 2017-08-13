package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
)

func WhipLeft(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		out := p.Dir.Copy().Rotate(30)
		pos := p.CenterPos().Add(out.Copy().Scale(6))
		fv := physics.NewForceVector(out.Copy(), 20)
		for i := 0; i < 25; i++ {
			pos.Add(out)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv)
		}
	}
}
func WhipRight(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		out := p.Dir.Copy().Rotate(-30)
		pos := p.CenterPos().Add(out.Copy().Scale(6))
		fv := physics.NewForceVector(out.Copy(), 20)
		for i := 0; i < 25; i++ {
			pos.Add(out)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv)
		}

	}
}

func WhipTwirl(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		rot := p.Dir.Copy().Scale(16)
		basePos := p.CenterPos()
		for angle := 0; angle < 360; angle += 10 {
			pos := basePos.Copy().Add(rot.Rotate(10))
			fv := physics.NewForceVector(rot.Copy(), 20)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv)
		}
	}
}
