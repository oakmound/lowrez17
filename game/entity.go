package game

import (
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
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

func (e *Entity) E() *Entity {
	return e
}

type HasE interface {
	E() *Entity
}
