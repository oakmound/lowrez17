package game

import (
	"fmt"
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
	fmt.Println(s)
	s.Entity = *NewEntity(pos.X(), pos.Y(), float64(w), float64(w), render.NewColorBox(w, w, c), s.Init(), mass, friction)
	s.Dir = dir.Copy()
	s.Speed = dir.Copy().Scale(speed)
	s.speedMax = 10
	s.RSpace.Space.UpdateLabel(label)
	render.Draw(s.R, entityLayer)
	s.deathTime = time.Now().Add(dur)
	s.Bind(shotEnter, "EnterFrame")

	return s
}

func shotEnter(id int, nothing interface{}) int {
	s := event.GetEntity(id).(*Shot)
	s.Delta.Add(s.Speed)
	s.Speed.Scale(s.accel)
	if time.Now().After(s.deathTime) {
		s.Cleanup()
		return 0
	}
	s.applyMovement()
	return 0
}

func Shoot(speed, accel float64, w int, c color.Color, label collision.Label, dur time.Duration, friction, mass float64) func(e *Entity) {
	return func(e *Entity) {
		MakeShot(e.Vector.Copy().Add(e.Dir.Copy().Scale(4)), e.Dir.Copy(), speed, accel, w, c, label, dur, friction, mass)
	}
}
