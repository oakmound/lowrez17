package game

import (
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"math"
)

//A Vein is a graphical connection between nodes on the body
type Vein struct {
	*render.Sprite
}

//NewVein creates a Vein
func NewVein(n1, n2 BodyNode, b *Body) *Vein {
	v := new(Vein)
	p1 := NodeCenter(n1)
	p2 := NodeCenter(n2)
	c1 := render.GradientColorAt(b.veinColor, b.veinColor2, math.Min(1.0, (3*n1.DiseaseLevel()+n2.DiseaseLevel())/3))
	c2 := render.GradientColorAt(b.veinColor, b.veinColor2, math.Min(1.0, (3*n2.DiseaseLevel()+n1.DiseaseLevel())/3))
	v.Sprite = render.NewGradientLine(p1.X(), p1.Y(), p2.X(), p2.Y(), c2, c1, 0)
	return v
}

func (v *Vein) Refresh(n1, n2 BodyNode, b *Body) {
	p1 := NodeCenter(n1)
	p2 := NodeCenter(n2)
	p1.Sub(v.Vec())
	p2.Sub(v.Vec())

	c1 := render.GradientColorAt(b.veinColor, b.veinColor2, math.Min(1.0, (3*n1.DiseaseLevel()+n2.DiseaseLevel())/3))
	c2 := render.GradientColorAt(b.veinColor, b.veinColor2, math.Min(1.0, (3*n2.DiseaseLevel()+n1.DiseaseLevel())/3))

	rgba := v.GetRGBA()
	render.DrawGradientLine(rgba, int(p1.X()), int(p1.Y()), int(p2.X()), int(p2.Y()), c2, c1, 0)
}

//TODO: Is this function in the right file?
//NodeCenter returns the center of a body node
func NodeCenter(bn BodyNode) physics.Vector {
	pos := bn.Vec()
	w, h := bn.Dims()
	return physics.NewVector(pos.X()+float64(w)/2, pos.Y()+float64(h)/2)
}
