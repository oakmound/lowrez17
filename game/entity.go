package game

import (
	"image/color"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/alg"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Entity struct {
	entities.Interactive
	physics.Mass
	Dir physics.Vector
}

func (e *Entity) Init() event.CID {
	e.CID = event.NextID(e)
	return e.CID
}

func (e *Entity) CenterPos() physics.Vector {
	return e.Vector.Copy().Add(physics.NewVector(e.W/2, e.H/2))
}

func NewEntity(x, y float64) {
	e := new(Entity)
	e.SetMass(10)
	r := render.NewReverting(render.NewColorBox(8, 8, color.RGBA{255, 0, 0, 255}))
	e.Interactive = entities.NewInteractive(x, y, 8, 8, r, e.Init(), .7)
	e.Speed = physics.NewVector(.5, .5)
	e.Dir = physics.NewVector(1, 0)
	e.Bind(playerMove, "EnterFrame")
	e.Bind(viewportFollow, "EnterFrame")
	render.Draw(e.R, entityLayer)
}

func playerMove(id int, frame interface{}) int {
	p := event.GetEntity(id).(*Entity)
	p.ApplyFriction(envFriction)
	// Calculate direction based on mouse position
	me := mouse.LastMouseEvent
	// Oak viewPos would be great as a vector
	center := p.CenterPos().Sub(physics.NewVector(float64(oak.ViewPos.X), float64(oak.ViewPos.Y)))
	p.Dir = physics.NewVector(float64(me.X), float64(me.Y)).Sub(center).Normalize()
	p.R.(*render.Reverting).RevertAndModify(1, render.Rotate(int(-p.Dir.Angle())))
	if oak.IsDown("W") {
		p.Delta.Add(p.Dir.Copy().Scale(p.Speed.Y()))
	}
	if oak.IsDown("S") {
		p.Delta.Add(p.Dir.Copy().Scale(-p.Speed.Y()))
	}
	if oak.IsDown("A") {
		p.Delta.Add(p.Dir.Copy().Rotate(90).Scale(p.Speed.X()))
	}
	if oak.IsDown("D") {
		p.Delta.Add(p.Dir.Copy().Rotate(90).Scale(-p.Speed.X()))
	}
	p.ShiftPos(p.Delta.X(), p.Delta.Y())
	return 0
}

func viewportFollow(id int, frame interface{}) int {
	p := event.GetEntity(id).(*Entity)
	viewportGoalPos := p.CenterPos().Sub(physics.NewVector(float64(oak.ScreenWidth/2), float64(oak.ScreenHeight/2)))
	delta := viewportGoalPos.Sub(physics.NewVector(float64(oak.ViewPos.X), float64(oak.ViewPos.Y)))

	if delta.Magnitude() < 2 {
		return 0
	}

	oak.SetScreen(
		alg.RoundF64(float64(oak.ViewPos.X)+delta.X()/16),
		alg.RoundF64(float64(oak.ViewPos.Y)+delta.Y()/16))

	return 0
}
