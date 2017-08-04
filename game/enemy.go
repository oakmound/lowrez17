package game

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

type EnemyFn func(x, y, difficulty float64) *Enemy

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

func NewMelee(x, y, diff float64) *Enemy {
	return nil
}

func NewRanged(x, y, diff float64) *Enemy {
	return nil
}

func NewBoomer(x, y, diff float64) *Enemy {
	return nil
}

func NewWizard(x, y, diff float64) *Enemy {
	return nil
}

func NewDasher(x, y, diff float64) *Enemy {
	return nil
}

func NewSummoner(x, y, diff float64) *Enemy {
	return nil
}

func NewVacuumer(x, y, diff float64) *Enemy {
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
