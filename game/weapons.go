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

type Weapon struct {
	left, right, space *Action
}

var (
	Sword = Weapon{
		left:  NewAction(SwordLeft(Ally), 250*time.Millisecond),
		right: NewAction(SwordRight(Ally), 250*time.Millisecond),
		space: NewAction(SwordDash(Ally), 750*time.Millisecond),
	}
	Whip = Weapon{
		left:  NewAction(WhipLeft(Ally), 300*time.Millisecond),
		right: NewAction(WhipRight(Ally), 300*time.Millisecond),
		space: NewAction(WhipTwirl(Ally), 1200*time.Millisecond),
	}
	Spear = Weapon{
		left:  NewAction(SpearJab(Ally), 350*time.Millisecond),
		right: NewAction(SpearThrust(Ally), 900*time.Millisecond),
		space: NewAction(SpearDash(Ally), 1000*time.Millisecond),
	}
	Net = Weapon{
		left:  NewAction(NetLeft(Ally), 100*time.Millisecond),
		right: NewAction(NetRight(Ally), 100*time.Millisecond),
		space: NewAction(NetTwirl(Ally), 600*time.Millisecond),
	}
)

const (
	Stun collision.Label = 100
)

func SpearJab(label collision.Label) func(*Entity) {
	return func(p *Entity) {
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

//Net Functions
func NetLeft(label collision.Label) func(*Entity) {
	return func(p *Entity) {
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
