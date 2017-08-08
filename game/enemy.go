package game

import (
	"fmt"

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
}

func (e *Enemy) Init() event.CID {
	e.CID = event.NextID(e)
	return e.CID
}

func (e *Enemy) Destroy() {
	enemyCh <- true
	e.Cleanup()
}

type EnemyCreation func(x, y int, difficulty float64) *Enemy

func NewEnemy(x, y, w, h float64, r render.Renderable, friction, mass, speed, maxSpeed float64) (e *Enemy) {
	e = new(Enemy)
	e.Entity = *NewEntity(x, y, w, h, r, e.Init(), friction, mass)
	render.Draw(e.R, entityLayer)
	collision.Add(e.RSpace.Space)
	e.Dir = physics.NewVector(1, 0)
	e.Speed = physics.NewVector(speed, speed)
	e.RSpace.Add(collision.Label(Ally), hitEnemy)

	e.speedMax = maxSpeed

	e.Bind(enemyEnter, "EnterFrame")
	return e
}

func enemyEnter(id int, frame interface{}) int {
	e := event.GetEntity(id).(*Enemy)
	if e.Health < 1 {
		fmt.Println("An enemy is dead!")
		e.Destroy()
	}
	e.Dir = player.Vec().Copy().Sub(e.CenterPos()).Normalize()
	e.attack(e)
	e.move(frame.(int), e)
	e.applyMovement()
	return 0
}

func hitEnemy(s1, s2 *collision.Space) {
	ent := event.GetEntity(int(s1.CID))
	if e, ok := ent.(*Enemy); ok {
		e.Health--
		bounceEntity(s1, s2)
	}
}

func CleanupEnemies() {
	for _, e := range enemies {
		e.Cleanup()
	}
	enemies = []*Enemy{}
	fmt.Println("Enemies cleaned up")
}
