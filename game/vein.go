package game

import (
	"image/color"

	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type VeinNode struct {
	physics.Vector
}

func NewVeinNode(x, y float64) VeinNode {
	return VeinNode{physics.NewVector(x, y)}
}

func (v VeinNode) Organ() (Organ, bool) {
	return nil, false
}

type Vein struct {
	*render.Sprite
}

func NewVein(n1, n2 BodyNode, c color.Color) *Vein {
	p1 := n1.Vec()
	p2 := n2.Vec()

	l := render.NewLine(p1.X(), p1.Y(), p2.X(), p2.Y(), c)
	return &Vein{l}
}
