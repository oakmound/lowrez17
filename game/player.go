package game

import (
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
	dashCooldown  time.Time
	leftCooldown  time.Time
	rightCooldown time.Time
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
	player.Bind(playerAttack, "MouseRelease")
}

func playerAttack(id int, mouseEvent interface{}) int {
	p := event.GetEntity(id).(*Player)
	me := mouseEvent.(mouse.Event)
	switch me.Button {
	case "LeftMouse":
		if time.Now().After(p.leftCooldown) {
			// Todo: generalize what is essentially making a cone of hurt boxes as
			// a sword swing
			pos := p.CenterPos().Add(p.Dir.Copy().Rotate(-55).Scale(7))
			basePos := pos.Copy()
			for j := -55.0; j <= 45.0; j += 10.0 {
				yDelta := p.Dir.Copy().Rotate(j).Scale(4)
				pos = basePos.Copy()
				for i := 0; i < 4; i++ {
					NewHurtBox(pos.X(), pos.Y(), 3, 3, 50*time.Millisecond, Ally)
					pos.Add(yDelta)
				}
			}
			p.leftCooldown = time.Now().Add(1 * time.Second)
		}
	case "RightMouse":
		if time.Now().After(p.rightCooldown) {
			pos := p.CenterPos().Add(p.Dir.Copy().Rotate(55).Scale(7))
			basePos := pos.Copy()
			for j := 55.0; j >= -45.0; j -= 10.0 {
				yDelta := p.Dir.Copy().Rotate(j).Scale(4)
				pos = basePos.Copy()
				for i := 0; i < 4; i++ {
					NewHurtBox(pos.X(), pos.Y(), 3, 3, 50*time.Millisecond, Ally)
					pos.Add(yDelta)
				}
			}
			p.rightCooldown = time.Now().Add(1 * time.Second)
		}
	}
	return 0
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
		p.dashCooldown = time.Now().Add(1 * time.Second)
		delta := p.Dir.Copy().Scale(3)
		perpendicular := delta.Copy().Rotate(90)
		pos := p.CenterPos().Add(delta, perpendicular, perpendicular)
		perpendicular.Scale(-1)
		basePos := pos.Copy()
		for i := 0; i < 3; i++ {
			pos = basePos.Add(perpendicular).Copy()
			for j := 0; j < 12; j++ {
				pos.Add(delta)
				NewHurtBox(pos.X(), pos.Y(), 3, 3, 50*time.Millisecond, Ally)
			}
		}
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

func viewportFollow(id int, frame interface{}) int {
	p := event.GetEntity(id).(*Player)
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
