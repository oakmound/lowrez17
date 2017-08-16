package game

import (
	"image/color"
	"time"

	"math/rand"

	"math"

	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

var (
	enemies []*Enemy
)

type Enemy struct {
	Entity
	Health int
	AttackSet
	MoveSet
	summoned   bool
	minimapR   render.Renderable
	attackAnim bool
	flashing   bool
	flashStop  time.Time
}

func (e *Enemy) Init() event.CID {
	e.CID = event.NextID(e)
	return e.CID
}

func (e *Enemy) Destroy() {
	if !e.summoned {
		go func() { enemyCh <- true }()
	}
	e.minimapR.UnDraw()
	e.Cleanup()
}

func (e *Enemy) Cleanup() {
	e.minimapR.UnDraw()
	e.Entity.Cleanup()
}

type EnemyCreation func(x, y int, difficulty float64, summoned bool) *Enemy

func NewEnemy(x, y, w, h float64, r render.Renderable, friction, mass, speed, maxSpeed float64, summoned bool) (e *Enemy) {
	e = new(Enemy)
	e.summoned = summoned
	render.Draw(r, layers.EntityLayer)
	e.Entity = *NewEntity(x, y, w, h, r, e.Init(), friction, mass)
	collision.Add(e.RSpace.Space)
	e.Dir = physics.NewVector(1, 0)
	e.Speed = physics.NewVector(speed, speed)
	e.RSpace.Add(collision.Label(Ally), hitEnemy)
	e.RSpace.Add(collision.Label(Acid), hurtEnemy)

	e.speedMax = maxSpeed

	e.minimapR = render.NewColorBox(1, 1, color.RGBA{0, 0, 0, 128})
	render.Draw(e.minimapR, layers.DebugLayer)

	e.Bind(enemyEnter, "EnterFrame")
	return e
}

func enemyEnter(id int, frame interface{}) int {
	e := event.GetEntity(id).(*Enemy)
	if e.Health < 1 && e.Health > -math.MaxInt32/2 {
		e.Health = -math.MaxInt32 / 2
		e.Destroy()
	}
	e.Dir = player.Vec().Copy().Sub(e.CenterPos()).Normalize()

	// Minimap logic
	v := oak.ViewVector()
	delta := e.Vec().Copy().Sub(v)
	if delta.X() > 0 {
		if delta.Y() > 0 {
			if delta.X() < 64 {
				if delta.Y() < 64 {
					v.Sub(physics.NewVector(1, 1))
				} else {
					v.Add(physics.NewVector(delta.X(), 63))
				}
			} else {
				if delta.Y() < 64 {
					v.Add(physics.NewVector(63, delta.Y()))
				} else {
					v.Add(physics.NewVector(63, 63))
				}
			}
		} else {
			if delta.X() > 64 {
				v.ShiftX(63)
			} else {
				v.ShiftX(delta.X())
			}
		}
	} else if delta.Y() > 0 {
		if delta.Y() > 64 {
			v.ShiftY(63)
		} else {
			v.ShiftY(delta.Y())
		}
	}

	e.minimapR.SetPos(v.X(), v.Y())

	if e.attack(e) && e.attackAnim {
		e.R.(*render.Reverting).Set("attacking")
	}
	if e.flashing && time.Now().After(e.flashStop) {
		e.R.(*render.Reverting).Revert(1)
		e.flashing = false
	}
	e.move(frame.(int), e)
	e.applyMovement()
	return 0
}

func hurtEnemy(s1, _ *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if e, ok := ent.(*Enemy); ok {
		if rand.Float64() < 0.05 {
			flash(e)
			e.Health--
		}
	}
}

func hitEnemy(s1, s2 *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if e, ok := ent.(*Enemy); ok {
		flash(e)
		e.Health--
		bounceEntity(s1, s2)
	}
}

func flash(ent *Enemy) {
	if !ent.flashing {
		ent.flashing = true
		if r, ok := ent.R.(*render.Reverting); ok {
			r.Modify(render.Brighten(50))
			ent.flashStop = time.Now().Add(50 * time.Millisecond)
		}
	} else {
		ent.flashStop = time.Now().Add(50 * time.Millisecond)
	}
}

func CleanupEnemies() {
	for _, e := range enemies {
		e.Cleanup()
	}
	enemies = []*Enemy{}
}
