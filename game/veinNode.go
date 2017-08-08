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
	Infectable
}

//NewVeinNode creates a vein node
func NewVeinNode(x, y float64, veinColor color.Color) *VeinNode {
	vn := &VeinNode{Vector: physics.NewVector(x, y)}
	vn.BodyButton = NewBodyButton(veinNodeWidthf64, veinNodeWidthf64)
	vn.diseaseRate = .0001
	vn.r = render.NewReverting(render.NewColorBox(veinNodeWidth, veinNodeWidth, veinColor))
	return vn
}

// NewVeinNodes returns a set of vein nodes, pairing together adjacent float inputs
// as x,y pairs
func NewVeinNodes(veinColor color.Color, positions ...float64) []*VeinNode {
	vns := make([]*VeinNode, len(positions)/2)
	if len(positions)%2 != 0 {
		return vns
	}
	for i := 0; i < len(positions); i += 2 {
		vns[i/2] = NewVeinNode(positions[i], positions[i+1], veinColor)
	}
	return vns
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
