package game

import (
	"image/color"

	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/lowrez17/game/layers"
)

//A BodyTraveler is an entity that travels along the veins of a body on the body layer
type BodyTraveler struct {
	entities.Doodad
	targetPos      physics.Vector
	travelSpeed    physics.Vector
	moving, active bool
}

//Init obtains a CID for an entity
func (bt *BodyTraveler) Init() event.CID {
	bt.CID = event.NextID(bt)
	return bt.CID
}

//NewBodyTraveler creates a BodyTraveler, binding the appropriate functions to a Sprite
func NewBodyTraveler(x, y float64) *BodyTraveler {
	bt := new(BodyTraveler)
	bt.Doodad = entities.NewDoodad(x, y, render.NewColorBox(3, 3, color.RGBA{0, 0, 255, 255}), bt.Init())
	bt.Bind(startTravelerMove, "MoveTraveler")
	bt.Bind(moveTraveler, "EnterFrame")
	bt.active = true
	render.Draw(bt.R, layers.TravelerLayer)
	return bt
}

//CenterPos  returns a vector to the center of BodyTraveler
func (bt *BodyTraveler) CenterPos() physics.Vector {
	v := bt.Vec()
	v.Copy().Add(physics.NewVector(1.5, 1.5))
	return v
}

//startTravelerMove is bound to a bodytraveler kicking off its motion along viens
func startTravelerMove(id int, pos interface{}) int {
	bt := event.GetEntity(id).(*BodyTraveler)
	if !bt.moving && bt.active {
		bt.targetPos = pos.(physics.Vector)
		bt.moving = true
	}
	return 0
}

//moveTraveler dictates the motion for the bodytraveler
func moveTraveler(id int, nothing interface{}) int {
	bt := event.GetEntity(id).(*BodyTraveler)
	if bt.moving && bt.active {
		delta := bt.targetPos.Copy().Sub(bt.Vector)
		if delta.Magnitude() < 1 {
			bt.moving = false
			// When you click on an organ, this will enter the lowest level once you hit it
			event.Trigger("HitNode", nil)
			return 0
		}
		delta = delta.Normalize()
		bt.ShiftX(delta.X())
		bt.ShiftY(delta.Y())
		bt.R.ShiftX(delta.X())
		bt.R.ShiftY(delta.Y())
	}
	return 0
}
