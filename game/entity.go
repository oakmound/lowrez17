package game

import (
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Entity struct {
	entities.Interactive
	physics.Mass
	Dir                 physics.Vector
	speedMax            float64
	collided            int
	moveVert, moveHoriz bool
}

func (e *Entity) Init() event.CID {
	e.CID = event.NextID(e)
	return e.CID
}

func NewEntity(x, y, w, h float64, r render.Renderable, id event.CID,
	friction, mass float64) *Entity {
	e := new(Entity)
	e.SetMass(mass)
	e.Interactive = entities.NewInteractive(x, y, w, h, r, id.Parse(e), friction)
	e.RSpace.Add(collision.Label(Blocked), bounceEntity)
	return e
}

func (e *Entity) CenterPos() physics.Vector {
	return e.Vector.Copy().Add(physics.NewVector(e.W/2, e.H/2))
}

func (e *Entity) E() *Entity {
	return e
}

type HasE interface {
	E() *Entity
}

func (e *Entity) applyMovement() {
	//Movement logic
	e.enforceSpeedMax()
	e.ShiftPos(e.Delta.X(), e.Delta.Y())
	<-e.RSpace.CallOnHits()
	e.enforceSpeedMax()
	if e.collided > 0 {
		e.collided = 0
		e.Delta.Scale(-1)
		e.ShiftPos(e.Delta.X(), e.Delta.Y())
	}
	e.ApplyFriction(envFriction)
}

func bounceEntity(s1, s2 *collision.Space) {
	e := event.GetEntity(int(s1.CID)).(HasE).E()
	e.collided++
	if ds, ok := event.GetEntity(int(s2.CID)).(*DirectionSpace); ok {
		physics.Push(ds, e)
	} else {
		e.Delta.Add(s1.OverlapVector(s2).Scale(.5))
	}
}

func (e *Entity) moveForward() {
	e.Delta.Add(e.Dir.Copy().Scale(e.Speed.Y()))
	e.moveVert = true
}
func (e *Entity) moveBack() {
	e.Delta.Add(e.Dir.Copy().Scale(-e.Speed.Y()))
	e.moveVert = true
}
func (e *Entity) moveRight() {
	e.Delta.Add(e.Dir.Copy().Rotate(90).Scale(e.Speed.X()))
	e.moveHoriz = true
}
func (e *Entity) moveLeft() {
	e.Delta.Add(e.Dir.Copy().Rotate(90).Scale(-e.Speed.X()))
	e.moveHoriz = true
}
func (e *Entity) scaleDiagonal() {
	if e.moveHoriz && e.moveVert {
		e.Delta.Scale(.8)
	}
	e.moveHoriz = false
	e.moveVert = false
}
func (e *Entity) enforceSpeedMax() {
	if e.Delta.Magnitude() > e.speedMax {
		e.Delta.Scale(e.speedMax / e.Delta.Magnitude())
	}
}
