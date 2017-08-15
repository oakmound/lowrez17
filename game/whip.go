package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

func WhipLeft(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		out := p.Dir.Copy().Rotate(30)
		center := p.CenterPos()
		pos := center.Copy().Add(out.Copy().Scale(6))
		fv := physics.NewForceVector(out.Copy(), 20)

		whip := images["whip"].Copy()
		whip.Modify(render.Rotate(int(-out.Angle())))
		whip.SetPos(pos.X(), pos.Y())
		w, h := whip.GetDims()
		if pos.X() < center.X()-1 {
			whip.ShiftX(float64(-w))
		}
		if pos.Y() < center.Y()-1 {
			whip.ShiftY(float64(-h))
		}

		render.Draw(whip, layers.DebugLayer)
		render.UndrawAfter(whip, 75*time.Millisecond)

		for i := 0; i < 25; i++ {
			pos.Add(out)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 4, 4, 75*time.Millisecond, label, fv)
		}
	}
}
func WhipRight(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		out := p.Dir.Copy().Rotate(-30)
		center := p.CenterPos()
		pos := center.Copy().Add(out.Copy().Scale(6))

		whip := images["whip"].Copy()
		whip.Modify(render.Rotate(int(-out.Angle())))
		whip.SetPos(pos.X(), pos.Y())
		w, h := whip.GetDims()
		if pos.X() < center.X()-1 {
			whip.ShiftX(float64(-w))
		}
		if pos.Y() < center.Y()-1 {
			whip.ShiftY(float64(-h))
		}

		render.Draw(whip, layers.DebugLayer)
		render.UndrawAfter(whip, 75*time.Millisecond)

		fv := physics.NewForceVector(out.Copy(), 20)
		for i := 0; i < 25; i++ {
			pos.Add(out)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 4, 4, 75*time.Millisecond, label, fv)
		}

	}
}

func WhipTwirl(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		rot := p.Dir.Copy().Scale(16)
		basePos := p.CenterPos()
		whip := render.NewReverting(images["whip"].Copy())
		whip.SetPos(basePos.X(), basePos.Y())
		render.Draw(whip, layers.DebugLayer)
		go func(whip *render.Reverting) {
			for i := 0; i < 360; i += 5 {
				whip.RevertAndModify(1, render.Rotate(-i))
				whip.SetPos(basePos.X(), basePos.Y())
				w, h := whip.GetDims()
				if i > 90 && i < 270 {
					whip.ShiftX(float64(-w))
				}
				if i > 180 {
					whip.ShiftY(float64(-h))
				}
				time.Sleep(5 * time.Millisecond)
			}
			whip.UnDraw()
		}(whip)
		for angle := 0; angle < 360; angle += 10 {
			pos := basePos.Copy().Add(rot.Rotate(10))
			fv := physics.NewForceVector(rot.Copy(), 20)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 150*time.Millisecond, label, fv)
		}
	}
}
