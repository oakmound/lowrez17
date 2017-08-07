package menu

import (
	"image/color"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
)

var (
	currentLevel  = 0
	sceneContinue = true
	nextScene     = "menu"
)

func StartScene(string, interface{}) {
	NewPlayer()
	nextScene = "menu"
	// Create blocking zones
	collision.Add(collision.NewLabeledSpace(-2, -2, 2, 66, blocking))
	collision.Add(collision.NewLabeledSpace(-2, -2, 40, 2, blocking))
	collision.Add(collision.NewLabeledSpace(-2, 64, 64, 2, blocking))
	collision.Add(collision.NewLabeledSpace(64, -66, 2, 128, blocking))
	// Create zones that lead to levels, menu
	collision.Add(collision.NewLabeledSpace(20, 20, 20, 10, nextLevel))
	collision.Add(collision.NewLabeledSpace(22, 22, 16, 6, blocking))
	table := render.NewColorBox(16, 6, color.RGBA{200, 200, 200, 255})
	table.SetPos(22, 22)
	render.Draw(table, 12)
}

func LoopScene() bool {
	return sceneContinue
}

func EndScene() (string, *oak.SceneResult) {
	sceneContinue = false
	return nextScene, nil
}

type Player struct {
	entities.Reactive
	collision.Phase
	stop       bool
	interactFn func()
}

func (p *Player) Init() event.CID {
	p.CID = event.NextID(p)
	return p.CID
}

func NewPlayer() {
	p := new(Player)
	r := render.NewColorBox(3, 7, color.RGBA{255, 255, 255, 255})
	p.Reactive = entities.NewReactive(5, 5, 3, 7, r, p.Init())
	collision.Add(p.RSpace.Space)
	render.Draw(p.R, 10)
	p.RSpace.Add(blocking, playerStop)
	collision.PhaseCollision(p.RSpace.Space)
	p.Bind(triggerInteractive, "CollisionStart")
	p.Bind(unbindInteractive, "CollisionStop")
	p.Bind(playerWalk, "EnterFrame")
	p.Bind(playerInteract, "KeyUpE")

}

const (
	blocking collision.Label = iota
	nextLevel
	// morgue labels
	level1
	level2
	level3
	level4
	level5
	endurance
	// other places
	settings
	/// ...
)

func playerStop(s1, s2 *collision.Space) {
	p := s1.CID.E().(*Player)
	p.stop = true
}

func triggerInteractive(id int, label interface{}) int {
	p := event.CID(id).E().(*Player)
	switch label.(collision.Label) {
	case nextLevel:
		p.interactFn = func() {
			nextScene = "level"
			sceneContinue = false
		}
	}
	return 0
}

func unbindInteractive(id int, label interface{}) int {
	p := event.CID(id).E().(*Player)
	if label.(collision.Label) != blocking {
		p.interactFn = nil
	}
	return 0
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
		shiftY--
	}
	if oak.IsDown("S") {
		shiftY++
	}
	if oak.IsDown("A") {
		shiftX--
	}
	if oak.IsDown("D") {
		shiftX++
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
