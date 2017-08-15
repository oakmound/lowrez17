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
	// Todo: Distinguish these two, when we start tracking hits on walls
	e.RSpace.Add(collision.Label(Blocked), bounceEntity)
	e.RSpace.Add(collision.Label(LowDamage), infectBounce(0.001))
	e.RSpace.Add(collision.Label(HighDamage), infectBounce(0.005))
	e.RSpace.Add(collision.Label(PressureFan), nudgeEntity)
	e.RSpace.Add(Stun, stopEntity)
	return e
}

func (e *Entity) CenterPos() physics.Vector {
	return e.Vector.Copy().Add(physics.NewVector(e.W/2, e.H/2))
}

func (e *Entity) E() *Entity {
	return e
}

func (e *Entity) Cleanup() {
	e.UnbindAll()
	collision.Remove(e.RSpace.Space)
	e.R.UnDraw()
	event.DestroyEntity(int(e.CID))
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

func infectBounce(rate float64) func(s1, s2 *collision.Space) {
	return func(s1, s2 *collision.Space) {
		i := thisBody.VecIndex(traveler.Vector)
		o := thisBody.graph[i]
		//oak.SetPalette(grayScale)
		o.Infect(rate)
		bounceEntity(s1, s2)
	}
}

func bounceEntity(s1, s2 *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if hase, ok := ent.(HasE); ok {
		e := hase.E()
		e.collided++
		if psh, ok := event.GetEntity(int(s2.CID)).(physics.Pushes); ok {
			physics.Push(psh, e)
		} else {
			e.Delta.Add(s1.OverlapVector(s2).Scale(.5))
		}
	}
}
func nudgeEntity(s1, s2 *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if hase, ok := ent.(HasE); ok {
		e := hase.E()
		//e.collided++
		if psh, ok := event.GetEntity(int(s2.CID)).(physics.Pushes); ok {
			physics.Push(psh, e)
		} else {
			e.Delta.Add(s1.OverlapVector(s2).Scale(.2))
		}
	}
}

func stopEntity(s1, _ *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if hase, ok := ent.(HasE); ok {
		e := hase.E()
		e.Delta.SetPos(0, 0)
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

func (e *Entity) teleportForward(distance float64) {
	e.Vector.Add(e.Dir.Copy().Scale(distance))
}

func (e *Entity) teleportBack(distance float64) {
	e.Vector.Add(e.Dir.Copy().Scale(-distance))
}

func (e *Entity) teleportRight(distance float64) {
	e.Vector.Add(e.Dir.Copy().Rotate(90).Scale(distance))
}

func (e *Entity) teleportLeft(distance float64) {
	e.Vector.Add(e.Dir.Copy().Rotate(90).Scale(-distance))
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
