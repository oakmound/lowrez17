package menu

import (
	"path/filepath"

	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/lowrez17/game/sfx"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/audio"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
)

type Player struct {
	entities.Reactive
	collision.Phase
	stop       bool
	interactFn func()
	interactR  render.Renderable
	footCh     chan audio.ChannelSignal
}

func (p *Player) Init() event.CID {
	p.CID = event.NextID(p)
	return p.CID
}

func NewPlayer() *Player {
	p := new(Player)
	sh := render.GetSheet(filepath.Join("4x16", "topplayer.png"))
	p.Reactive = entities.NewReactive(5, 5, 4, 16, render.NewCompound(
		"forward",
		map[string]render.Modifiable{
			"forward": sh[0][0].Copy(),
			"right":   sh[1][0].Copy(),
			"left":    sh[1][0].Copy().Modify(render.FlipX),
			"back":    sh[2][0].Copy(),
		}), p.Init())
	var err error
	p.footCh, err = audio.GetChannel(sfx.SoftSFX, intrange.NewLinear(140, 300), "Footstep.wav")
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
		p.R.(*render.Compound).Set("back")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftY--
	} else if oak.IsDown("S") {
		p.R.(*render.Compound).Set("forward")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftY++
	} else if oak.IsDown("A") {
		p.R.(*render.Compound).Set("left")
		//p.footCh <- audio.NewPosSignal(0, p.X(), p.Y())
		shiftX--
	} else if oak.IsDown("D") {
		p.R.(*render.Compound).Set("right")
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
