package game

import (
	"image/color"
	"time"

	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/lowrez17/game/forceSpace"
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

func NewMelee(x, y int, diff float64, summoned bool) *Enemy {
	r := images["meleeFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.2, 5, 0.1, 4, summoned)
	e.Health = 70
	e.AttackSet = NewAttackSet(intrange.NewLinear(2000, 4000),
		[]float64{1.0},
		[]*Action{NewAction(EnemySwordDash, 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 1.0, 0.1, 0.1},
		Move(Left, 2),
		Move(Forward, 1),
		Move(Forward, 10),
		Move(Right, 10))
	return e
}

func NewRanged(x, y int, diff float64, summoned bool) *Enemy {
	r := images["rangedFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.2, 5, 0.1, 4, summoned)
	e.Health = 50
	e.AttackSet = NewAttackSet(intrange.NewLinear(1000, 2000),
		[]float64{1.0},
		[]*Action{NewAction(Shoot(1, 1, 2, color.RGBA{255, 255, 255, 255}, Opposing, 3*time.Second, .5, 2, "RangedAttack"), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, .2, 0.2, 1.0, 2.0},
		Move(Left, 5),
		Move(Forward, 5),
		Move(Backward, 5),
		Move(Right, 5),
		Move(Wait, 30))
	return e
}

func NewBoomer(x, y int, diff float64, summoned bool) *Enemy {
	r := images["heartFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.2, 20, 0.05, 2, summoned)
	e.Health = 300
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
			PlayAt("BoomerAttack", e.X(), e.Y())
			for i := 0; i < 360; i += 30 {
				dir := physics.AngleVector(float64(i))
				MakeShot(e.Vector, dir, 8, .9, 5, color.RGBA{100, 10, 10, 255}, Opposing, 5*time.Second, .5, 40)
			}
			e.Destroy()
		})
	}
	return 0
}

func NewWizard(x, y int, diff float64, summoned bool) *Enemy {
	r := images["brainFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.2, 5, 0.05, 2, summoned)
	e.Health = 150
	e.AttackSet = NewAttackSet(intrange.NewLinear(1200, 1800),
		[]float64{1.0},
		[]*Action{NewAction(Shoot(0.2, 1.05, 6, color.RGBA{190, 20, 200, 190}, Opposing, 5*time.Second, .25, 60, "WizardAttack"), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 0.5, 0.5, 1.0, 1.0, 1.0, 10.0},
		Teleport(Left, 25),
		Teleport(Forward, 20),
		Teleport(Backward, 10),
		Teleport(Forward, 30),
		Teleport(Backward, 5),
		Teleport(Right, 25),
		Move(Wait, 90))
	return e
}

func NewDasher(x, y int, diff float64, summoned bool) *Enemy {
	r := images["lungFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.4, 4, 2, 6, summoned)
	e.Health = 100
	e.AttackSet = NewAttackSet(intrange.NewLinear(200, 1000),
		[]float64{1.0, 1.0},
		[]*Action{NewAction(Shoot(2, 1, 2, color.RGBA{50, 50, 255, 255}, Opposing, 6*time.Second, .5, 2, "DasherAttack"), 0),
			NewAction(EnemySwordDash, 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 1.0, .2, 1.0, .2},
		Move(Left, 5),
		Move(Forward, 5),
		Move(Backward, 5),
		Move(Right, 5),
		Move(Wait, 5))
	return e
}

func NewSummoner(x, y int, diff float64, summoned bool) *Enemy {
	r := images["liverFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.2, 10, 0.1, 4, summoned)
	e.Health = 250
	e.AttackSet = NewAttackSet(intrange.NewLinear(6000, 18000),
		[]float64{1.0, 1.0},
		[]*Action{NewAction(Summon(NewMelee), 0),
			NewAction(Summon(NewRanged), 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, 1.0, 1.0, 1.0, 3.0},
		Move(Left, 5),
		Move(Backward, 10),
		Move(Forward, 5),
		Move(Right, 5),
		Move(Wait, 30))
	e.attackAnim = true
	e.R.(render.Triggerable).SetTriggerID(e.CID)
	e.Bind(stopAttacking, "AnimationEnd")
	return e
}

func stopAttacking(id int, nothing interface{}) int {
	e := event.GetEntity(id).(*Enemy)
	e.R.(*render.Reverting).Set("base")
	return 0
}

func Summon(ec EnemyCreation) func(*Entity) {
	return func(e *Entity) {
		PlayAt("SummonAttack", e.X(), e.Y())
		en := ec(int(e.X()+e.Dir.X()*4)/tileDim, int(e.Y()+e.Dir.Y()*4)/tileDim, 1.0, true)
		enemies = append(enemies, en)
	}
}

func NewVacuumer(x, y int, diff float64, summoned bool) *Enemy {
	r := images["stomachFoe"].Copy()
	w, h := r.GetDims()
	e := NewEnemy(float64(x*tileDim), float64(y*tileDim), float64(w), float64(h), r, 0.2, 5, 0.8, 2.3, summoned)
	e.Health = 160
	e.AttackSet = NewAttackSet(intrange.NewLinear(2000, 3000),
		[]float64{1.0},
		[]*Action{NewAction(Vacuum, 0)})
	e.MoveSet = NewMoveSet([]float64{1.0, .2, 0.2, 1.0, 2.0},
		Move(Left, 5),
		Move(Forward, 5),
		Move(Backward, 5),
		Move(Right, 5),
		Move(Wait, 30))
	e.attackAnim = true
	e.R.(render.Triggerable).SetTriggerID(e.CID)
	e.Bind(stopAttacking, "AnimationEnd")
	return e
}

func Vacuum(p *Entity) {
	PlayAt("Vacuum", p.X(), p.Y())
	fv := physics.NewForceVector(p.Dir.Copy().Normalize(), 5)
	delta := p.Dir.Copy().Scale(3)
	perpendicular := delta.Copy().Rotate(90)
	pos := p.CenterPos().Add(delta, perpendicular, perpendicular, perpendicular)
	perpendicular.Scale(-1)
	basePos := pos.Copy()
	for i := 0; i < 5; i++ {
		pos = basePos.Add(perpendicular).Copy()
		for j := 0; j < 15; j++ {
			pos.Add(delta)
			forceSpace.NewHurtBox(pos.X(), pos.Y(), 3, 3, 100*time.Millisecond, Opposing, fv)
		}
	}
}

// Notes:
// Each organ has waves, each wave has random or perscribed enemies
// each organ has valid positions for enemies to be placed, by shaping
// flow:
// enter organ
// first wave spawns at set of valid positions
// either after time passes or wave defeated, next wave spawns
// when all waves defeated, organ is saved
