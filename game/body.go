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

func (b *Body) AddNodes(ns ...BodyNode) {
	for _, n := range ns {
		b.graph = append(b.graph, n)
		b.adjacency = append(b.adjacency, []int{})
	}
}

func (b *Body) IsAdjacent(i, j int) bool {
	for _, k := range b.adjacency[i] {
		if k == j {
			return true
		}
	}
	return false
}

type BodyNode interface {
	Vec() physics.Vector
	SetPos(physics.Vector)
	Organ() (Organ, bool)
}

func DemoBody() *Body {
	b := new(Body)
	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.AddNodes(NewVeinNode(10, 10),
		NewVeinNode(15, 20),
		NewLiver(40, 5),
		NewHeart(50, 40),
		NewBrain(50, 10),
		NewLung(30, 30),
		NewStomach(4, 40))
	b.Connect(0, 1)
	b.Connect(0, 2)
	b.Connect(1, 2)
	b.Connect(1, 3)
	b.Connect(1, 4)
	b.Connect(1, 5)
	b.Connect(1, 6)
	return b
}

func (b *Body) VecIndex(v physics.Vector) int {
	for i, n := range b.graph {
		dist := n.Vec().Distance(v)
		if dist < 1 {
			return i
		}
	}
	return -1
}

// Random body ideas:
// given a shape from an image,
// we can put in random locations for everything, where everything is a random
// set of organs / veins
