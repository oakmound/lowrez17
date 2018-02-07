package game

import (
	"image/color"

	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type BodyButton struct {
	mouse.CollisionPhase
	*collision.Space
	highlight render.Renderable
}

//Init aqcuires CID
func (bb *BodyButton) Init() event.CID {
	bb.CID = event.NextID(bb)
	return bb.CID
}

//SetPos sets the location for the BodyButton
func (bb *BodyButton) SetPos(v physics.Vector) {
	mouse.UpdateSpace(v.X(), v.Y(), bb.GetW(), bb.GetH(), bb.Space)
}

func (bb *BodyButton) IsTravelerAdjacent() bool {
	thisIndex := thisBody.VecIndex(bb.CenterPos())
	travelIndex := thisBody.VecIndex(traveler.CenterPos())
	if thisIndex == -1 || travelIndex == -1 {
		return false
	}
	return thisBody.IsAdjacent(thisIndex, travelIndex)
}

//CenterPos
func (bb *BodyButton) CenterPos() physics.Vector {
	return physics.NewVector(bb.Space.GetCenter())
}

func NewBodyButton(w, h float64) *BodyButton {
	bb := &BodyButton{}
	bb.Space = collision.NewSpace(0, 0, w, h, 0)
	bb.Space.CID = bb.Init()
	mouse.PhaseCollision(bb.Space)
	mouse.Add(bb.Space)
	bb.CID.Bind(tryHighlightBB, "HitNode")
	bb.CID.Bind(highlightBB, "MouseCollisionStart")
	bb.CID.Bind(unhighlightBB, "MouseCollisionStop")
	bb.CID.Bind(moveToBB, "MouseReleaseOn")
	bb.highlight = render.NewColorBox(int(w+2), int(h+2), color.RGBA{255, 255, 255, 255})
	return bb
}

func tryHighlightBB(id int, nothing interface{}) int {
	bb := event.GetEntity(id).(*BodyButton)
	me := mouse.LastEvent
	hits := mouse.Hits(me.ToSpace())
	for _, h := range hits {
		if h == bb.Space {
			return highlightBB(id, nil)
		}
	}
	return 0
}

func highlightBB(id int, nothing interface{}) int {
	bb := event.GetEntity(id).(*BodyButton)
	if bb.IsTravelerAdjacent() {
		bb.highlight.SetPos(bb.X()-1, bb.Y()-1)
		render.Draw(bb.highlight, layers.HighlightLayer)
	}
	return 0
}

func unhighlightBB(id int, nothing interface{}) int {
	bb := event.GetEntity(id).(*BodyButton)
	bb.highlight.Undraw()
	return 0
}

func moveToBB(id int, nothing interface{}) int {
	bb := event.GetEntity(id).(*BodyButton)
	if bb.IsTravelerAdjacent() {
		event.Trigger("MoveTraveler", bb.CenterPos())
	}
	return 0
}
