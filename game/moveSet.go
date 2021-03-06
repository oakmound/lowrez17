package game

import (
	"math/rand"

	"github.com/oakmound/oak/alg"
)

type MoveSet struct {
	moves       [][]*Action
	moveWeights []float64
	currentSet  int
}

func (m *MoveSet) move(frame int, e HasE) {
	if len(m.moves) == 0 {
		return
	}
	aIndex := frame % len(m.moves[m.currentSet])
	//Check if the move set should change
	if aIndex == 0 {
		m.currentSet = alg.ChooseX(m.moveWeights, 1)[0]
	}
	//Get the actual move action to use
	m.moves[m.currentSet][aIndex].MustDo(e)
}

func NewMoveSet(weights []float64, actions ...[]*Action) MoveSet {
	return MoveSet{actions, weights, rand.Intn(len(actions))}
}
