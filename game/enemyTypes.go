package game

import (
	"image/color"
	"time"

	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/oak/render"
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
	e.AttackSet = NewAttackSet(intrange.NewLinear(5000, 15000), []float64{1.0}, []*Action{NewAction(SwordDash(Opposing), 0)})
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
	e.AttackSet = NewAttackSet(intrange.NewLinear(1000, 2000), []float64{1.0},
		[]*Action{NewAction(Shoot(1, 1, 2, color.RGBA{255, 255, 255, 255}, Opposing, 3*time.Second, .5, 10), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, .2, 0.2, 1.0, 2.0},
		Move(Left, 5),
		Move(Forward, 5),
		Move(Backward, 5),
		Move(Right, 5),
		Move(Wait, 30))
	return e
}

func NewBoomer(x, y int, diff float64) *Enemy {
	return nil
}

func NewWizard(x, y int, diff float64) *Enemy {
	return nil
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
