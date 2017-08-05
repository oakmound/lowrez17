package game

import (
	"image/color"

	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

var (
	veinNodeWidth    = 2
	veinNodeWidthf64 = float64(veinNodeWidth)
)

//VeinNode is a simple node that is NOT an Organ
type VeinNode struct {
	physics.Vector
	*BodyButton
}

//NewVeinNode creates a vein node
func NewVeinNode(x, y float64) *VeinNode {
	vn := &VeinNode{Vector: physics.NewVector(x, y)}
	vn.BodyButton = NewBodyButton(veinNodeWidthf64, veinNodeWidthf64)
	return vn
}

//SetPos sets the position of the vein
func (vn *VeinNode) SetPos(v physics.Vector) {
	vn.Vector.SetPos(v.X(), v.Y())
	vn.BodyButton.SetPos(v)
}

//Dims returns a static size for all veins
func (vn *VeinNode) Dims() (int, int) {
	return 3, 3
}

//Organ returns that vein is not an organ
func (vn *VeinNode) Organ() (Organ, bool) {
	return nil, false
}

//A Vein is a graphical connection between nodes on the body
type Vein struct {
	*render.Sprite
}

//NewVein creates a Vein
func NewVein(n1, n2 BodyNode, c color.Color) *Vein {
	p1 := NodeCenter(n1)
	p2 := NodeCenter(n2)

	l := render.NewLine(p1.X(), p1.Y(), p2.X(), p2.Y(), c)
	return &Vein{l}
}

//TODO: Is this function in the right file?
//NodeCenter returns the center of a body node
func NodeCenter(bn BodyNode) physics.Vector {
	pos := bn.Vec()
	w, h := bn.Dims()
	return physics.NewVector(pos.X()+float64(w)/2, pos.Y()+float64(h)/2)
}
