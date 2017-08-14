package game

import (
	"github.com/200sc/go-dist/floatrange"
	"github.com/oakmound/oak/alg"
)

type EnemyDist struct {
	N       int
	Chances []float64
}

// We're assuming here that the difficulty of the enemies chosen (recoloringm etc)
// will be decided separate from the type of enemy chosen.

func (ed EnemyDist) Poll() []EnemyType {
	es := make([]EnemyType, ed.N)
	chosen := alg.ChooseX(ed.Chances, ed.N)
	for i, v := range chosen {
		es[i] = EnemyType(v)
	}
	return es
}

var (
	SingleMelee           = EnemyDist{1, []float64{1.0, 0.0, 0.0}}
	SmallMeleeDist        = EnemyDist{8, []float64{1.0, 0.0, 0.0}}
	ModerateMeleeDist     = EnemyDist{16, []float64{1.0, 0.0, 0.0}}
	LargeMeleeDist        = EnemyDist{40, []float64{1.0, 0.0, 0.0}}
	SmallRangedDist       = EnemyDist{8, []float64{0.0, 1.0, 0.0}}
	ModerateRangedDist    = EnemyDist{16, []float64{0.0, 1.0, 0.0}}
	LargeRangedDist       = EnemyDist{40, []float64{0.0, 1.0, 0.0}}
	SmallSpecialDist      = EnemyDist{8, []float64{0.0, 0.0, 1.0}}
	ModerateSpecialDist   = EnemyDist{16, []float64{0.0, 0.0, 1.0}}
	LargeSpecialDist      = EnemyDist{40, []float64{0.0, 0.0, 1.0}}
	SmallBalancedDist     = EnemyDist{8, []float64{1.0, 1.0, 1.0}}
	ModerateBalancedDist  = EnemyDist{16, []float64{1.0, 1.0, 1.0}}
	LargeBalancedDist     = EnemyDist{40, []float64{1.0, 1.0, 1.0}}
	SmallNoMeleeDist      = EnemyDist{8, []float64{0.0, 1.0, 1.0}}
	ModerateNoMeleeDist   = EnemyDist{16, []float64{0.0, 1.0, 1.0}}
	LargeNoMeleeDist      = EnemyDist{40, []float64{0.0, 1.0, 1.0}}
	SmallNoSpecialDist    = EnemyDist{8, []float64{1.0, 1.0, 0.0}}
	ModerateNoSpecialDist = EnemyDist{16, []float64{1.0, 1.0, 0.0}}
	LargeNoSpecialDist    = EnemyDist{40, []float64{1.0, 1.0, 0.0}}
	SmallNoRangedDist     = EnemyDist{8, []float64{1.0, 0.0, 1.0}}
	ModerateNoRangedDist  = EnemyDist{16, []float64{1.0, 0.0, 1.0}}
	LargeNoRangedDist     = EnemyDist{40, []float64{1.0, 0.0, 1.0}}
)

// Short names

var (
	SMD  = SmallMeleeDist
	MMD  = ModerateMeleeDist
	LMD  = LargeMeleeDist
	SRD  = SmallRangedDist
	MRD  = ModerateRangedDist
	LRD  = LargeRangedDist
	SSD  = SmallSpecialDist
	MSD  = ModerateSpecialDist
	LSD  = LargeSpecialDist
	SBD  = SmallBalancedDist
	MBD  = ModerateBalancedDist
	LBD  = LargeBalancedDist
	SNMD = SmallNoMeleeDist
	MNMD = ModerateNoMeleeDist
	LNMD = LargeNoMeleeDist
	SNSD = SmallNoSpecialDist
	MNSD = ModerateNoSpecialDist
	LNSD = LargeNoSpecialDist
	SNRD = SmallNoRangedDist
	MNRD = ModerateNoRangedDist
	LNRD = LargeNoRangedDist
)

func RandomDist(enemyCount int) EnemyDist {
	fr := floatrange.NewLinear(0.1, 1.0)
	ed := EnemyDist{
		enemyCount,
		[]float64{fr.Poll(), fr.Poll(), fr.Poll()},
	}
	return ed
}
