package game

import (
	"path/filepath"
	"time"

	"github.com/disintegration/gift"
	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/lowrez17/game/sfx"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/render/mod"
	"github.com/oakmound/oak/timing"
)

type Player struct {
	Entity
	Weapon
	leaveTime time.Time
	hitExit   bool
}

func (p *Player) Init() event.CID {
	return event.NextID(p)
}

func NewPlayer() *Player {
	if player == nil {
		e := new(Player)
		shtt, _ := render.GetSheet(filepath.Join("8x8", "lowerlevelplayer.png"))
		sh := shtt.ToSprites()
		s := sh[1][0].Copy().Modify(mod.FlipX)
		r := render.NewReverting(s)
		e.Entity = *NewEntity(0, 0, 8, 8, r, e.Init(), .8, 20)
		e.Speed = physics.NewVector(.3, .3)
		e.Dir = physics.NewVector(1, 0)
		e.RSpace.Add(collision.Label(Exit), leaveOrgan)

		e.RSpace.Add(collision.Label(Opposing), bounceEntity)
		e.RSpace.Add(collision.Label(Opposing), playerHurt)
		e.speedMax = 6.5
		collision.Add(e.RSpace.Space)
		player = e
		player.Weapon = Sword
		sfx.UpdateEars(player)
	}
	return player
}

func startupPlayer() {
	render.Draw(player.R, layers.EntityLayer)
	collision.Add(player.RSpace.Space)

	player.Bind(playerMove, "EnterFrame")
	player.Bind(viewportFollow, "EnterFrame")
	player.Bind(playerAttack, "MouseRelease")
	player.Bind(equip(Sword), "KeyUp1")
	player.Bind(equip(Whip), "KeyUp2")
	player.Bind(equip(Net), "KeyUp3")
	player.Bind(equip(Spear), "KeyUp4")
}

func equip(w Weapon) func(id int, nothing interface{}) int {
	return func(id int, nothing interface{}) int {
		p := event.GetEntity(id).(*Player)
		// These globals will keep their cooldowns
		p.Weapon = w
		return 0
	}
}

func stopPlayer() {
	player.SetPos(-1000, -1000)
	collision.Remove(player.RSpace.Space)
	player.UnbindAll()
	player.R.Undraw()
}

func leaveOrgan(_, _ *collision.Space) {

	if player.hitExit {
		if time.Now().After(player.leaveTime) {
			if player.X() > -1000 {
				CleanupActiveOrgan(false)
			}
		}
	} else {
		player.leaveTime = time.Now().Add(3000 * time.Millisecond)
		go timing.DoAfter(3100*time.Millisecond, func() {
			player.hitExit = false
		})
	}
	player.hitExit = true
}

func playerAttack(id int, mouseEvent interface{}) int {
	p := event.GetEntity(id).(*Player)
	me := mouseEvent.(mouse.Event)
	switch me.Button {
	case "LeftMouse":
		p.Weapon.left.Do(p)
	case "RightMouse":
		p.Weapon.right.Do(p)
	}
	return 0
}

func playerMove(id int, frame interface{}) int {
	p := event.GetEntity(id).(*Player)
	// Calculate direction based on mouse position
	me := mouse.LastEvent
	// Oak viewPos would be great as a vector
	center := p.CenterPos().Sub(oak.ViewVector())
	p.Dir = physics.NewVector(me.X(), me.Y()).Sub(center).Normalize()
	p.R.(*render.Reverting).RevertAndModify(1,
		mod.RotateInterpolated(float32(-p.Dir.Angle()), gift.NearestNeighborInterpolation))

	if oak.IsDown("W") {
		p.moveForward()
	}
	if oak.IsDown("S") {
		p.moveBack()
	}
	if oak.IsDown("A") {
		p.moveLeft()
	}
	if oak.IsDown("D") {
		p.moveRight()
	}
	p.scaleDiagonal()

	if oak.IsDown("Spacebar") {
		p.Weapon.space.Do(p)
	}

	p.applyMovement()
	return 0
}
func playerHurt(s1, s2 *collision.Space) {
	//oak.SetPalette(hurtPalette)
	bounceEntity(s1, s2)
}
