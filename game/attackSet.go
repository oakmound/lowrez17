package game

import (
	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/oak/alg"
	"time"
)

type AttackSet struct {
	NextAttack    time.Time
	WaitMillis    intrange.Range
	AttackWeights []float64
	Attacks       []*Action
}

func (a *AttackSet) attack(e HasE) {
	if time.Now().After(a.NextAttack) {
		a.Attacks[alg.ChooseX(a.AttackWeights, 1)[0]].Do(e)
		a.NextAttack = time.Now().Add(time.Duration(a.WaitMillis.Poll()) * time.Millisecond)
	}
}

func NewAttackSet(wait intrange.Range, weights []float64, attacks []*Action) AttackSet {
	return AttackSet{time.Now().Add(time.Duration(wait.Poll()) * time.Millisecond),
		wait,
		weights,
		attacks}
}
