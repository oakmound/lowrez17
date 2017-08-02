package game

import (
	"fmt"
	"image/color"

	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type BodyTraveler struct {
	entities.Doodad
	moving      bool
	targetPos   physics.Vector
	travelSpeed physics.Vector
}

func (bt *BodyTraveler) Init() event.CID {
	bt.CID = event.NextID(bt)
	return bt.CID
}

func NewBodyTraveler(x, y float64) *BodyTraveler {
	bt := new(BodyTraveler)
	bt.Doodad = entities.NewDoodad(x, y, render.NewColorBox(3, 3, color.RGBA{0, 0, 255, 255}), bt.Init())
	bt.Bind(startTravelerMove, "MoveTraveler")
	bt.Bind(moveTraveler, "EnterFrame")
	render.Draw(bt.R, travelerLayer)
	return bt
}

func startTravelerMove(id int, pos interface{}) int {
	bt := event.GetEntity(id).(*BodyTraveler)
	bt.targetPos = pos.(physics.Vector)
	bt.moving = true
	return 0
}

func moveTraveler(id int, nothing interface{}) int {
	bt := event.GetEntity(id).(*BodyTraveler)
	if bt.moving {
		delta := bt.targetPos.Copy().Sub(bt.Vector)
		fmt.Println("Moving", delta)
		if delta.Magnitude() < 2 {
			bt.moving = false
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
