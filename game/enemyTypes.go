package game

import (
	"image/color"
	"time"

	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/timing"
)

type EnemyType int

const (
	Melee EnemyType = iota
	Ranged
	Special
)

var (
	enemyFns = map[EnemyType]map[OrganType]EnemyCreation{
		Melee: {
			Brain:   NewMelee,
			Heart:   NewMelee,
			Lung:    NewMelee,
			Stomach: NewMelee,
			Liver:   NewMelee,
		},
		Ranged: {
			Brain:   NewRanged,
			Heart:   NewRanged,
			Lung:    NewRanged,
			Stomach: NewRanged,
			Liver:   NewRanged,
		},
		Special: {
			Brain:   NewWizard,
			Heart:   NewBoomer,
			Lung:    NewDasher,
			Stomach: NewVacuumer,
			Liver:   NewSummoner,
		},
	}
)

func NewMelee(x, y int, diff float64) *Enemy {
	r := render.NewColorBox(8, 8, color.RGBA{120, 120, 120, 255})
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), 8, 8, r, 0.2, 5, 0.1, 4)
	e.Health = 50
	e.AttackSet = NewAttackSet(intrange.NewLinear(500, 1500),
		[]float64{1.0},
		[]*Action{NewAction(SwordDash(Opposing), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 1.0, 0.1, 0.1},
		Move(Left, 2),
		Move(Forward, 1),
		Move(Forward, 10),
		Move(Right, 10))
	return e
}

func NewRanged(x, y int, diff float64) *Enemy {
	r := render.NewColorBox(8, 8, color.RGBA{170, 170, 170, 255})
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), 8, 8, r, 0.2, 5, 0.1, 4)
	e.Health = 50
	e.AttackSet = NewAttackSet(intrange.NewLinear(1000, 2000),
		[]float64{1.0},
		[]*Action{NewAction(Shoot(1, 1, 2, color.RGBA{255, 255, 255, 255}, Opposing, 3*time.Second, .5, 2), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, .2, 0.2, 1.0, 2.0},
		Move(Left, 5),
		Move(Forward, 5),
		Move(Backward, 5),
		Move(Right, 5),
		Move(Wait, 30))
	return e
}

func NewBoomer(x, y int, diff float64) *Enemy {
	r := render.NewColorBox(12, 12, color.RGBA{200, 120, 120, 255})
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), 12, 12, r, 0.2, 20, 0.05, 2)
	e.Health = 150
	e.AttackSet = NewAttackSet(intrange.NewLinear(1000, 3000),
		[]float64{1.0, 1.0},
		[]*Action{NewAction(SwordLeft(Opposing), 0),
			NewAction(SwordRight(Opposing), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 1.0, 1.0, 1.0, 1.0},
		Move(Left, 4),
		Move(Forward, 1),
		Move(Forward, 10),
		Move(Right, 4),
		Move(Wait, 15))
	e.Bind(explode, "EnterFrame")
	return e
}

func explode(id int, nothing interface{}) int {
	e := event.GetEntity(id).(*Enemy)
	if e.Health < 50 {
		e.Health = 1000
		go timing.DoAfter(500*time.Millisecond, func() {
			// todo: Animate explosion
			for i := 0; i < 360; i += 5 {
				dir := physics.AngleVector(float64(i))
				MakeShot(e.Vector, dir, 3, .9, 3, color.RGBA{100, 10, 10, 255}, Opposing, 3*time.Second, .5, 5)
			}
			e.Destroy()
		})
	}
	return 0
}

func NewWizard(x, y int, diff float64) *Enemy {
	r := render.NewColorBox(12, 12, color.RGBA{200, 120, 120, 255})
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), 12, 12, r, 0.2, 20, 0.05, 2)
	e.Health = 150
	e.AttackSet = NewAttackSet(intrange.NewLinear(200, 1000),
		[]float64{1.0, 1.0},
		[]*Action{NewAction(Shoot(1, 1.2, 4, color.RGBA{190, 20, 20, 190}, Opposing, 5*time.Second, .25, 1), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 1.0, 1.0, 1.0, 1.0},
		Teleport(Left, 15),
		Teleport(Forward, 5),
		Teleport(Backward, 5),
		Teleport(Right, 15),
		Move(Wait, 45))
	return e
}

func NewDasher(x, y int, diff float64) *Enemy {
	return nil
}

func NewSummoner(x, y int, diff float64) *Enemy {
	return nil
}

func NewVacuumer(x, y int, diff float64) *Enemy {
	return nil
}

// Notes:
// Each organ has waves, each wave has random or perscribed enemies
// each organ has valid positions for enemies to be placed, by shaping
// flow:
// enter organ
// first wave spawns at set of valid positions
// either after time passes or wave defeated, next wave spawns
// when all waves defeated, organ is saved
