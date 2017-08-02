package game

import (
	"image/color"

	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Body struct {
	overlay   render.Modifiable
	veinColor color.RGBA
	graph     []BodyNode
	adjacency [][]int
}

// Connect connects two bodyNodes on a body, and returns whether
// it succeeds. Failure indicates that the two nodes were already
// connected, or an input node did not exist
func (b *Body) Connect(a, c int) bool {
	if len(b.adjacency) <= a || len(b.adjacency) <= c {
		return false
	}
	for _, v := range b.adjacency[a] {
		if v == c {
			return false
		}
	}
	// All connections are dual connections
	b.adjacency[a] = append(b.adjacency[a], c)
	b.adjacency[c] = append(b.adjacency[c], a)
	return true
}

func (b *Body) AddNode(n BodyNode) {
	b.graph = append(b.graph, n)
	b.adjacency = append(b.adjacency, []int{})
}

type BodyNode interface {
	Vec() physics.Vector
	Organ() (Organ, bool)
}

func DemoBody() *Body {
	b := new(Body)
	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.AddNode(NewVeinNode(10, 10))
	b.AddNode(NewVeinNode(15, 20))
	b.AddNode(NewLiver(40, 5))
	b.Connect(0, 1)
	b.Connect(0, 2)
	b.Connect(1, 2)
	return b
}
