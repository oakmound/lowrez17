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
	spearWidth  = 9
	spearDamage = 7
)

func SpearJab(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("SpearLight", p.X(), p.Y())

		fv := physics.NewForceVector(physics.NewVector(0, 0), 0)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(20))
		for i := 0; i < spearDamage; i++ {
			forceSpace.NewHurtBox(pos.X(), pos.Y(), spearWidth, spearWidth, 200*time.Millisecond, label, fv)
		}

		spear := images["spear"].Copy()
		RotateAbout(spear, p.CenterPos().Add(p.Dir.Copy().Scale(4)), p.CenterPos(), p.Dir.Angle())
		render.Draw(spear, layers.DebugLayer)

		stick := collision.NewLabeledSpace(pos.X(), pos.Y(), spearWidth, spearWidth, Stun)
		collision.Add(stick)
		go timing.DoAfter(200*time.Millisecond, func() {
			collision.Remove(stick)
			spear.Undraw()
		})
	}
}

func SpearThrust(label collision.Label) func(*Entity) {
	return func(p *Entity) {

		PlayAt("SpearLight", p.X(), p.Y())

		fv := physics.NewForceVector(physics.NewVector(0, 0), 0)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(23))
		for i := 0; i < spearDamage; i++ {
			forceSpace.NewHurtBox(pos.X(), pos.Y(), spearWidth, spearWidth, 700*time.Millisecond, label, fv)
		}

		spear := images["spear"].Copy()
		RotateAbout(spear, p.CenterPos().Add(p.Dir.Copy().Scale(5)), p.CenterPos(), p.Dir.Angle())
		render.Draw(spear, layers.DebugLayer)

		stick := collision.NewLabeledSpace(pos.X(), pos.Y(), spearWidth, spearWidth, Stun)
		collision.Add(stick)
		go timing.DoAfter(500*time.Millisecond, func() {
			collision.Remove(stick)
			spear.Undraw()
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
		render.DrawForTime(spear, 75*time.Millisecond)

		// More boxes -- more damage
		for i := 0; i < spearDamage*3; i++ {
			forceSpace.NewHurtBox(pos.X(), pos.Y(), spearWidth, spearWidth, 400*time.Millisecond, label, fv)
		}
		go timing.DoAfter(SpearDashCooldown, func() {
			PlayAt("SpearReady", p.X(), p.Y())
		})
	}
}
