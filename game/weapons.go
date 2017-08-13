package game

import (
	"time"

	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/timing"
)

type Weapon struct {
	left, right, space *Action
}

var (
	Sword = Weapon{
		left:  NewAction(SwordLeft(Ally), 100*time.Millisecond),
		right: NewAction(SwordRight(Ally), 100*time.Millisecond),
		space: NewAction(SwordDash(Ally), 500*time.Millisecond),
	}
	Whip = Weapon{
		left:  NewAction(WhipLeft(Ally), 200*time.Millisecond),
		right: NewAction(WhipRight(Ally), 200*time.Millisecond),
		space: NewAction(WhipTwirl(Ally), 700*time.Millisecond),
	}
	Spear = Weapon{
		left:  NewAction(SpearJab(Ally), 150*time.Millisecond),
		right: NewAction(SpearThrust(Ally), 900*time.Millisecond),
		space: NewAction(SpearDash(Ally), 1000*time.Millisecond),
	}
	Net = Weapon{
		left:  NewAction(NetLeft(Ally), 50*time.Millisecond),
		right: NewAction(NetRight(Ally), 50*time.Millisecond),
		space: NewAction(NetTwirl(Ally), 400*time.Millisecond),
	}
)

const (
	Stun collision.Label = 100
)

func SpearJab(label collision.Label) func(*Entity) {
	return func(p *Entity) {
		fv := physics.NewForceVector(physics.NewVector(0, 0), 0)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(15))
		forceSpace.NewHurtBox(pos.X(), pos.Y(), 7, 7, 75*time.Millisecond, label, fv)
		stick := collision.NewLabeledSpace(pos.X(), pos.Y(), 7, 7, Stun)
		collision.Add(stick)
		go timing.DoAfter(75*time.Millisecond, func() {
			collision.Remove(stick)
		})
	}
}

func SpearThrust(label collision.Label) func(*Entity) {
	return func(p *Entity) {
		fv := physics.NewForceVector(physics.NewVector(0, 0), 0)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(20))
		forceSpace.NewHurtBox(pos.X(), pos.Y(), 7, 7, 500*time.Millisecond, label, fv)
		stick := collision.NewLabeledSpace(pos.X(), pos.Y(), 7, 7, Stun)
		collision.Add(stick)
		go timing.DoAfter(500*time.Millisecond, func() {
			collision.Remove(stick)
		})
	}
}

func SpearDash(label collision.Label) func(*Entity) {
	return func(p *Entity) {
		p.Delta.Add(p.Dir.Copy().Scale(24 * p.Speed.Y()))
		fv := physics.NewForceVector(p.Dir.Copy(), 30)
		pos := p.CenterPos().Add(p.Dir.Copy().Scale(15))
		forceSpace.NewHurtBox(pos.X(), pos.Y(), 7, 7, 75*time.Millisecond, label, fv)
	}
}

//Net Functions
func NetLeft(label collision.Label) func(*Entity) {
	return func(p *Entity) {
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(180), 3)
		basePos := p.CenterPos()
		rot := p.Dir.Copy().Rotate(-130)
		for a := 0; a < 90; a += 10 {
			pos := basePos.Copy().Add(rot.Copy().Scale(6))
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 5, 5, 75*time.Millisecond, label, fv)
			rot.Rotate(10)
		}
	}
}

func NetRight(label collision.Label) func(*Entity) {
	return func(p *Entity) {
		fv := physics.NewForceVector(p.Dir.Copy().Rotate(180), 3)
		basePos := p.CenterPos()
		rot := p.Dir.Copy().Rotate(130)
		for a := 0; a < 90; a += 10 {
			pos := basePos.Copy().Add(rot.Copy().Scale(6))
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 5, 5, 75*time.Millisecond, label, fv)
			rot.Rotate(-10)
		}
	}
}

func NetTwirl(label collision.Label) func(*Entity) {
	return func(p *Entity) {
		go func() {
			basePos := p.CenterPos()
			rot := p.Dir.Copy().Rotate(-10)
			for a := 0; a < 260; a += 10 {
				pos := basePos.Copy().Add(rot.Copy().Scale(6))
				fv := physics.NewForceVector(rot.Copy().Rotate(90), 3)
				forceSpace.NewHurtBox(pos.X(), pos.Y(), 5, 5, 75*time.Millisecond, label, fv)
				rot.Rotate(-10)
				time.Sleep(5 * time.Millisecond)
			}
		}()
	}
}
