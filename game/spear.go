package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/timing"
)

const (
	Stun collision.Label = 100
)

func SpearJab(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("SpearLight", p.X(), p.Y())

		fv := physics.NewForceVector(physics.NewVector(0, 0), 0)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(20))
		forceSpace.NewHurtBox(pos.X(), pos.Y(), 7, 7, 75*time.Millisecond, label, fv)

		spear := images["spear"].Copy()
		RotateAbout(spear, p.CenterPos().Add(p.Dir.Copy().Scale(4)), p.CenterPos(), p.Dir.Angle())
		render.Draw(spear, layers.DebugLayer)

		stick := collision.NewLabeledSpace(pos.X(), pos.Y(), 7, 7, Stun)
		collision.Add(stick)
		go timing.DoAfter(75*time.Millisecond, func() {
			collision.Remove(stick)
			spear.UnDraw()
		})
	}
}

func SpearThrust(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("SpearLight", p.X(), p.Y())

		fv := physics.NewForceVector(physics.NewVector(0, 0), 0)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(23))
		forceSpace.NewHurtBox(pos.X(), pos.Y(), 7, 7, 500*time.Millisecond, label, fv)

		spear := images["spear"].Copy()
		RotateAbout(spear, p.CenterPos().Add(p.Dir.Copy().Scale(5)), p.CenterPos(), p.Dir.Angle())
		render.Draw(spear, layers.DebugLayer)

		stick := collision.NewLabeledSpace(pos.X(), pos.Y(), 7, 7, Stun)
		collision.Add(stick)
		go timing.DoAfter(500*time.Millisecond, func() {
			collision.Remove(stick)
			spear.UnDraw()
		})
	}
}

func SpearDash(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("SpearHeavy", p.X(), p.Y())

		p.Delta.Add(p.Dir.Copy().Scale(24 * p.Speed.Y()))
		fv := physics.NewForceVector(p.Dir.Copy(), 30)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(23))

		spear := images["spear"].Copy()
		RotateAbout(spear, p.CenterPos().Add(p.Dir.Copy().Scale(5)), p.CenterPos(), p.Dir.Angle())
		render.Draw(spear, layers.DebugLayer)
		render.UndrawAfter(spear, 75*time.Millisecond)

		forceSpace.NewHurtBox(pos.X(), pos.Y(), 7, 7, 75*time.Millisecond, label, fv)

	}
}
