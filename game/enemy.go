package game

import (
	"image/color"

	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
)

type EnemyType int

const (
	Melee EnemyType = iota
	Ranged
	Special
)

type Enemy struct {
	Entity
	Health int
}

func (e *Enemy) Init() event.CID {
	e.CID = event.NextID(e)
	return e.CID
}

type EnemyFn func(x, y int, difficulty float64) *Enemy

var (
	enemyFns = map[EnemyType]map[OrganType]EnemyFn{
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
	e := new(Enemy)
	r := render.NewColorBox(8, 8, color.RGBA{120, 120, 120, 255})
	e.Entity = *NewEntity(float64(x*tileDim), float64(y*tileDim), 8, 8, r, e.Init(), 0.5, 5)
	render.Draw(e.R, entityLayer)
	collision.Add(e.RSpace.Space)
	return e
}

func NewRanged(x, y int, diff float64) *Enemy {
	return nil
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
