package game

import (
	"image/color"
	"time"

	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Shot struct {
	Entity
	accel     float64
	deathTime time.Time
}

func (s *Shot) Init() event.CID {
	s.CID = event.NextID(s)
	return s.CID
}

func MakeShot(pos physics.Vector, dir physics.Vector, speed, accel float64, w int, c color.Color,
	label collision.Label, dur time.Duration, friction, mass float64) *Shot {
	s := new(Shot)
	s.Entity = *NewEntity(pos.X(), pos.Y(), float64(w), float64(w), render.NewColorBox(w, w, c), s.Init(), mass, friction)
	s.Dir = dir.Copy()
	s.accel = accel
	s.Speed = dir.Copy().Scale(speed)
	s.speedMax = 10
	s.RSpace.Space.UpdateLabel(label)
	render.Draw(s.R, entityLayer)
	s.deathTime = time.Now().Add(dur)
	s.Bind(shotEnter, "EnterFrame")
	s.RSpace.Add(collision.Label(Blocked), shotBlocked)
	s.RSpace.Add(collision.Label(Ally), shotReflect)

	return s
}

func shotEnter(id int, nothing interface{}) int {
	s := event.GetEntity(id).(*Shot)
	s.moveForward()
	s.Speed.Scale(s.accel)
	if time.Now().After(s.deathTime) {
		s.Cleanup()
		return 0
	}
	s.applyMovement()
	return 0
}

func shotBlocked(s1, s2 *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if sh, ok := ent.(*Shot); ok && sh != nil {
		sh.Cleanup()
	}
	// Todo: hurt organ
}

func shotReflect(s1, s2 *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if sh, ok := ent.(*Shot); ok {
		s1.UpdateLabel(Ally)
		if psh, ok := event.GetEntity(int(s2.CID)).(physics.Pushes); ok {
			physics.Push(psh, sh)
			sh.Speed = sh.Delta.Copy().Normalize()
			sh.Dir = sh.Speed.Copy()
		}
	}
}

func Shoot(speed, accel float64, w int, c color.Color, label collision.Label, dur time.Duration, friction, mass float64) func(e *Entity) {
	return func(e *Entity) {
		v := e.Vector.Copy().Add(e.Dir.Copy().Scale(4).Rotate(180))
		MakeShot(v, e.Dir.Copy(), speed, accel, w, c, label, dur, friction, mass)
	}
}
