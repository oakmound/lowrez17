package game

import (
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

//BodyNode is a node on the body that can be traveled to
type BodyNode interface {
	Vec() physics.Vector
	Dims() (int, int)
	SetPos(physics.Vector)
	Organ() (Organ, bool)
	Infect(...float64) bool
	DiseaseLevel() float64
	Cleanse()
	R() render.Modifiable
}
