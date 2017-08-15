package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

func NetLeft(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("NetLight", p.X(), p.Y())

		fv := physics.NewForceVector(p.Dir.Copy().Rotate(180), 3)
		basePos := p.CenterPos()
		rot := p.Dir.Copy().Rotate(-130)

		net := render.NewReverting(images["net"].Copy().Modify(render.FlipY))
		render.Draw(net, layers.DebugLayer)
		go func(rot physics.Vector) {
			for a := 0; a < 90; a += 10 {
				pos := basePos.Copy()
				rot.Rotate(10)
				NetRotateAbout(net, pos, basePos, rot.Angle())
				time.Sleep(20 * time.Millisecond)
			}
			net.UnDraw()
		}(rot.Copy())

		for a := 0; a < 90; a += 10 {
			pos := basePos.Copy().Add(rot.Copy().Scale(6))
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 5, 5, 75*time.Millisecond, label, fv, false)
			rot.Rotate(10)
		}
	}
}

func NetRight(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("NetLight", p.X(), p.Y())

		fv := physics.NewForceVector(p.Dir.Copy().Rotate(180), 3)
		basePos := p.CenterPos()
		rot := p.Dir.Copy().Rotate(130)

		net := render.NewReverting(images["net"].Copy())
		render.Draw(net, layers.DebugLayer)
		go func(rot physics.Vector) {
			for a := 0; a < 90; a += 10 {
				pos := basePos.Copy()
				rot.Rotate(-10)
				NetRotateAbout(net, pos, basePos, rot.Angle())
				time.Sleep(20 * time.Millisecond)
			}
			net.UnDraw()
		}(rot.Copy())

		for a := 0; a < 90; a += 10 {
			pos := basePos.Copy().Add(rot.Copy().Scale(6))
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 5, 5, 75*time.Millisecond, label, fv, false)
			rot.Rotate(-10)
		}
	}
}
func NetRotateAbout(r *render.Reverting, pos, center physics.Vector, angle float64) {
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

func NetTwirl(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("NetHeavy", p.X(), p.Y())

		basePos := p.CenterPos()
		rot := p.Dir.Copy().Rotate(-10)
		go func(basePos, rot physics.Vector) {
			net := render.NewReverting(images["net"].Copy())
			render.Draw(net, layers.DebugLayer)
			for a := 0; a < 260; a += 10 {
				pos := basePos.Copy()
				rot.Rotate(-10)
				NetRotateAbout(net, pos, basePos, rot.Angle())
				time.Sleep(20 * time.Millisecond)
			}
			net.UnDraw()
		}(basePos.Copy(), rot.Copy())

		go func(basePos, rot physics.Vector) {
			for a := 0; a < 260; a += 10 {
				pos := basePos.Copy().Add(rot.Copy().Scale(6))
				fv := physics.NewForceVector(rot.Copy().Rotate(90), 3)
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 5, 5, 75*time.Millisecond, label, fv, false)
				rot.Rotate(-10)
				time.Sleep(5 * time.Millisecond)
			}
		}(basePos, rot)
	}
}
