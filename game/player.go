package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/alg"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Player struct {
	Entity
	dashCooldown time.Time
}

func (p *Player) Init() event.CID {
	p.CID = event.NextID(p)
	return p.CID
}

func NewPlayer(x, y float64) *Player {
	if player == nil {
		e := new(Player)
		e.SetMass(10)
		r := render.NewReverting(render.NewColorBox(8, 8, color.RGBA{0, 0, 255, 255}))
		e.Interactive = entities.NewInteractive(x, y, 8, 8, r, e.Init(), .8)
		e.Speed = physics.NewVector(.3, .3)
		e.Dir = physics.NewVector(1, 0)
		e.RSpace.Add(collision.Label(Exit), leaveOrgan)
		e.RSpace.Add(collision.Label(Blocked), bounceEntity)
		collision.Add(e.RSpace.Space)
		player = e
	}
	player.SetPos(x, y)
	return player
}

func startupPlayer() {
	render.Draw(player.R, entityLayer)

	player.Bind(playerMove, "EnterFrame")
	player.Bind(viewportFollow, "EnterFrame")
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

func bounceEntity(s1, s2 *collision.Space) {
	// This will need work
	e := event.GetEntity(int(s1.CID)).(HasE).E()
	e.Delta.Scale(-1.5)
	e.ShiftPos(e.Delta.X(), e.Delta.Y())
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

	if oak.IsDown("Spacebar") && time.Now().After(p.dashCooldown) {
		p.Delta.Add(p.Dir.Copy().Scale(24 * p.Speed.Y()))
		p.dashCooldown = time.Now().Add(3 * time.Second)
	}
	p.ShiftPos(p.Delta.X(), p.Delta.Y())
	<-p.RSpace.CallOnHits()
	if p.Delta.X() > 5 {
		p.Delta.SetX(5)
	} else if p.Delta.X() < -5 {
		p.Delta.SetX(-5)
	}
	if p.Delta.Y() > 5 {
		p.Delta.SetY(5)
	} else if p.Delta.Y() < -5 {
		p.Delta.SetY(-5)
	}

	return 0
}

func viewportFollow(id int, frame interface{}) int {
	p := event.GetEntity(id).(*Player)
	fmt.Println(mouse.LastMouseEvent)
	viewportGoalPos := p.CenterPos().Copy().Add(mouse.LastMouseEvent.ToVector().Sub(physics.NewVector(32, 32)))
	viewportGoalPos = viewportGoalPos.Sub(physics.NewVector(float64(oak.ScreenWidth/2), float64(oak.ScreenHeight/2)))
	delta := viewportGoalPos.Sub(physics.NewVector(float64(oak.ViewPos.X), float64(oak.ViewPos.Y)))

	if delta.Magnitude() < 2 {
		return 0
	}

	oak.SetScreen(
		alg.RoundF64(float64(oak.ViewPos.X)+delta.X()/9),
		alg.RoundF64(float64(oak.ViewPos.Y)+delta.Y()/9))

	return 0
}
