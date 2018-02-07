package menu

import (
	"path/filepath"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/render/mod"
)

type Player struct {
	*entities.Reactive
	collision.Phase
	stop       bool
	interactFn func()
	interactR  render.Renderable
}

func (p *Player) Init() event.CID {
	return event.NextID(p)
}

func NewPlayer() *Player {
	p := new(Player)
	shtt, _ := render.GetSheet(filepath.Join("4x16", "topplayer.png"))
	sh := shtt.ToSprites()
	p.Reactive = entities.NewReactive(5, 5, 4, 16, render.NewSwitch(
		"forward",
		map[string]render.Modifiable{
			"forward": sh[0][0].Copy(),
			"right":   sh[1][0].Copy(),
			"left":    sh[1][0].Copy().Modify(mod.FlipX),
			"back":    sh[2][0].Copy(),
		}), nil, p.Init())
	var err error
	if err != nil {
		dlog.Error(err)
	}
	collision.Add(p.RSpace.Space)
	render.Draw(p.R, entityLayer)
	p.RSpace.Add(blocking, playerStop)
	collision.PhaseCollision(p.RSpace.Space)
	p.Bind(triggerInteractive, "CollisionStart")
	p.Bind(unbindInteractive, "CollisionStop")
	p.Bind(playerWalk, "EnterFrame")
	p.Bind(playerInteract, "KeyUpE")
	return p
}

func playerStop(s1, s2 *collision.Space) {
	p := s1.CID.E().(*Player)
	p.stop = true
}

func playerInteract(id int, nothing interface{}) int {
	p := event.CID(id).E().(*Player)
	if p.interactFn != nil {
		p.interactFn()
	}
	return 0
}

func playerWalk(id int, nothing interface{}) int {
	p := event.CID(id).E().(*Player)
	shiftX := 0.0
	shiftY := 0.0
	if oak.IsDown("W") {
		p.R.(*render.Switch).Set("back")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftY--
	} else if oak.IsDown("S") {
		p.R.(*render.Switch).Set("forward")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftY++
	} else if oak.IsDown("A") {
		p.R.(*render.Switch).Set("left")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftX--
	} else if oak.IsDown("D") {
		p.R.(*render.Switch).Set("right")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftX++
	}
	if p.interactR != nil {
		p.interactR.SetPos(p.X()-1, p.Y()-8)
	}
	collision.ShiftSpace(shiftX, shiftY, p.RSpace.Space)
	<-p.RSpace.CallOnHits()
	// If near interactable, show potential interaction
	// Stop if hit something
	collision.ShiftSpace(-shiftX, -shiftY, p.RSpace.Space)
	if p.stop {
		p.stop = false
	} else {
		p.ShiftPos(shiftX, shiftY)
	}
	return 0
}
