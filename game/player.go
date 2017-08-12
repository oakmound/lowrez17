package game

import (
	"path/filepath"
	"time"

	"github.com/disintegration/gift"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/timing"
)

type Player struct {
	Entity
	Weapon
	leaveTime time.Time
	hitExit   bool
}

func (p *Player) Init() event.CID {
	p.CID = event.NextID(p)
	return p.CID
}

func NewPlayer() *Player {
	if player == nil {
		e := new(Player)
		s := render.GetSheet(filepath.Join("8x8", "lowerlevelplayer.png"))[1][0].Copy().Modify(render.FlipX)
		r := render.NewReverting(s)
		e.Entity = *NewEntity(0, 0, 8, 8, r, e.Init(), .8, 13)
		e.Speed = physics.NewVector(.3, .3)
		e.Dir = physics.NewVector(1, 0)
		e.RSpace.Add(collision.Label(Exit), leaveOrgan)
		e.RSpace.Add(collision.Label(Opposing), bounceEntity)
		e.speedMax = 7
		collision.Add(e.RSpace.Space)
		player = e
	}
	player.Weapon = Whip
	return player
}

func startupPlayer() {
	render.Draw(player.R, entityLayer)
	collision.Add(player.RSpace.Space)

	player.Bind(playerMove, "EnterFrame")
	player.Bind(viewportFollow, "EnterFrame")
	player.Bind(playerAttack, "MouseRelease")
}

func stopPlayer() {
	player.SetPos(-1000, -1000)
	collision.Remove(player.RSpace.Space)
	player.UnbindAll()
	player.R.UnDraw()
}

func leaveOrgan(_, _ *collision.Space) {

	if player.hitExit {
		if time.Now().After(player.leaveTime) {
			if player.X() > -1000 {
				stopPlayer()
				CleanupTiles()
				CleanupEnemies()
				oak.SetScreen(0, 0)
				traveler.active = true
				select {
				case waveExitCh <- true:
				default:
				}
			}
		}
	} else {
		player.leaveTime = time.Now().Add(1500 * time.Millisecond)
		go timing.DoAfter(1600*time.Millisecond, func() {
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
	me := mouse.LastMouseEvent
	// Oak viewPos would be great as a vector
	center := p.CenterPos().Sub(oak.ViewVector())
	p.Dir = physics.NewVector(float64(me.X), float64(me.Y)).Sub(center).Normalize()
	p.R.(*render.Reverting).RevertAndModify(1,
		render.RotateInterpolated(int(-p.Dir.Angle()), gift.NearestNeighborInterpolation))

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
