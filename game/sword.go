package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

func SwordLeft(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		center := p.CenterPos()
		rot := p.Dir.Copy().Rotate(-55)
		pos := center.Copy().Add(rot.Scale(7))
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(-90).Normalize(), 5)
		basePos := pos.Copy()

		sword := render.NewReverting(images["sword"].Copy())
		render.Draw(sword, layers.DebugLayer)

		// Might want to tweak this one
		go func(pos physics.Vector) {
			rot2 := rot.Copy()
			for i := -55; i <= 45.0; i += 10 {
				rot2.Rotate(10)
				SwordRotateAbout(sword, pos, center, rot2.Angle())
				time.Sleep(10 * time.Millisecond)
			}
			sword.UnDraw()
		}(pos)

		for j := -55.0; j <= 45.0; j += 10.0 {
			yDelta := p.Dir.Copy().Rotate(j).Scale(4)
			pos = basePos.Copy()
			for i := 0; i < 4; i++ {
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv, false)
				pos.Add(yDelta)
			}
		}
	}
}
func SwordRight(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		center := p.CenterPos()
		rot := p.Dir.Copy().Rotate(55)
		pos := center.Add(rot.Copy().Scale(7))

		sword := render.NewReverting(images["sword"].Copy())
		render.Draw(sword, layers.DebugLayer)

		go func(pos physics.Vector) {
			rot2 := rot.Copy()
			for i := 55; i >= -45.0; i -= 10 {
				rot2.Rotate(-10)
				SwordRotateAbout(sword, pos, center, rot2.Angle())
				time.Sleep(10 * time.Millisecond)
			}
			sword.UnDraw()
		}(pos)

		fv := physics.NewForceVector(p.Dir.Copy().Rotate(90).Normalize(), 5)
		basePos := pos.Copy()
		for j := 55.0; j >= -45.0; j -= 10.0 {
			yDelta := p.Dir.Copy().Rotate(j).Scale(4)
			pos = basePos.Copy()
			for i := 0; i < 4; i++ {
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv, false)
				pos.Add(yDelta)
			}
		}
	}
}

func SwordRotateAbout(r *render.Reverting, pos, center physics.Vector, angle float64) {
	r.RevertAndModify(1, render.Rotate(int(-angle)))
	pos2 := pos.Copy().Add(physics.AngleVector(angle).Scale(3))
	r.SetPos(pos2.X(), pos2.Y())
	w, h := r.GetDims()
	if pos2.X() < center.X()-1 {
		r.ShiftX(float64(-w))
	}
	if pos2.Y() < center.Y()-1 {
		r.ShiftY(float64(-h))
	}
}

func SwordDash(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		p.Delta.Add(p.Dir.Copy().Scale(24 * p.Speed.Y()))
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(180).Normalize(), 10)
		delta := p.Dir.Copy().Scale(3)
		perpendicular := delta.Copy().Rotate(90)
		pos := p.CenterPos().Add(delta, perpendicular, perpendicular)

		sword := render.NewReverting(images["sword"].Copy())
		render.Draw(sword, layers.DebugLayer)
		SwordRotateAbout(sword, pos.Copy().Sub(perpendicular, perpendicular), p.CenterPos(), p.Dir.Angle())

		go func() {
			for i := 0; i < 8; i++ {
				sword.ShiftX(delta.X())
				sword.ShiftY(delta.Y())
				time.Sleep(15 * time.Millisecond)
			}
			sword.UnDraw()
		}()

		perpendicular.Scale(-1)
		basePos := pos.Copy()
		for i := 0; i < 3; i++ {
			pos = basePos.Add(perpendicular).Copy()
			for j := 0; j < 12; j++ {
				pos.Add(delta)
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 75*time.Millisecond, label, fv, false)
			}
		}
	}
}
