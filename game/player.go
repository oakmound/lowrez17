package game

import (
	"image/color"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Player struct {
	Entity
	Weapon
}

func (p *Player) Init() event.CID {
	p.CID = event.NextID(p)
	return p.CID
}

func NewPlayer() *Player {
	if player == nil {
		e := new(Player)
		r := render.NewReverting(render.NewColorBox(8, 8, color.RGBA{0, 0, 255, 255}))
		e.Entity = *NewEntity(0, 0, 8, 8, r, e.Init(), .8, 10)
		e.Speed = physics.NewVector(.3, .3)
		e.Dir = physics.NewVector(1, 0)
		e.RSpace.Add(collision.Label(Exit), leaveOrgan)
		e.RSpace.Add(collision.Label(Blocked), bounceEntity)
		collision.Add(e.RSpace.Space)
		player = e
	}
	player.Weapon = Sword
	return player
}

func startupPlayer() {
	render.Draw(player.R, entityLayer)

	player.Bind(playerMove, "EnterFrame")
	player.Bind(viewportFollow, "EnterFrame")
	player.Bind(playerAttack, "MouseRelease")
}

func stopPlayer() {
	player.SetPos(-1000, -1000)
	player.UnbindAll()
	player.R.UnDraw()
}

func leaveOrgan(_, _ *collision.Space) {
	CleanupTiles()
	stopPlayer()
	oak.SetScreen(0, 0)
	traveler.active = true
}

func playerAttack(id int, mouseEvent interface{}) int {
	p := event.GetEntity(id).(*Player)
	me := mouseEvent.(mouse.Event)
	switch me.Button {
	case "LeftMouse":
		p.Weapon["left"].Do(p)
	case "RightMouse":
		p.Weapon["right"].Do(p)
	}
	return 0
}

func playerMove(id int, frame interface{}) int {
	p := event.GetEntity(id).(*Player)

	p.ApplyFriction(envFriction)

	// Calculate direction based on mouse position
	me := mouse.LastMouseEvent
	// Oak viewPos would be great as a vector
	center := p.CenterPos().Sub(physics.NewVector(float64(oak.ViewPos.X), float64(oak.ViewPos.Y)))
	p.Dir = physics.NewVector(float64(me.X), float64(me.Y)).Sub(center).Normalize()
	p.R.(*render.Reverting).RevertAndModify(1, render.Rotate(int(-p.Dir.Angle())))
	var vertDown, horzDown bool
	if oak.IsDown("W") {
		p.Delta.Add(p.Dir.Copy().Scale(p.Speed.Y()))
		vertDown = true
	}
	if oak.IsDown("S") {
		p.Delta.Add(p.Dir.Copy().Scale(-p.Speed.Y()))
		vertDown = true
	}
	if oak.IsDown("A") {
		p.Delta.Add(p.Dir.Copy().Rotate(90).Scale(p.Speed.X()))
		horzDown = true
	}
	if oak.IsDown("D") {
		p.Delta.Add(p.Dir.Copy().Rotate(90).Scale(-p.Speed.X()))
		horzDown = true
	}
	if horzDown && vertDown {
		p.Delta.Scale(.8)
	}

	if oak.IsDown("Spacebar") {
		p.Weapon["dash"].Do(p)
	}
	p.ShiftPos(p.Delta.X(), p.Delta.Y())
	<-p.RSpace.CallOnHits()
	if p.Delta.X() > 10 {
		p.Delta.SetX(10)
	} else if p.Delta.X() < -10 {
		p.Delta.SetX(-10)
	}
	if p.Delta.Y() > 10 {
		p.Delta.SetY(10)
	} else if p.Delta.Y() < -10 {
		p.Delta.SetY(-10)
	}

	return 0
}
