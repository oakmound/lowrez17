package game

import (
	"image/color"

	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Body struct {
	overlay               render.Modifiable
	veinColor, veinColor2 color.RGBA
	graph                 []BodyNode
	adjacency             [][]int
	veins                 [][]*Vein
	infected              []int
	cleansed              map[int]bool
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

//AddNodes adds a set of body nodes to the body, placing them in the graph
func (b *Body) AddNodes(ns ...BodyNode) {
	for _, n := range ns {
		b.graph = append(b.graph, n)
		b.adjacency = append(b.adjacency, []int{})
	}
}

//IsAdjacent checks to see whether a node has a second node in its adjacency list
func (b *Body) IsAdjacent(i, j int) bool {
	for _, k := range b.adjacency[i] {
		if k == j {
			return true
		}
	}
	return false
}

func (b *Body) InitVeins() {
	w := len(b.graph)
	b.veins = make([][]*Vein, w)
	for i := range b.veins {
		b.veins[i] = make([]*Vein, w)
	}
}

// Infect infects organs that have not previously been cleansed.
// If an organ is not already infected it is then added to the body's diseased organs list.
func (b *Body) Infect(i int) {
	if _, ok := b.cleansed[i]; ok {
		return
	}
	if b.graph[i].Infect(.3) {
		b.infected = append(b.infected, i)
	}

}

//BodyNode is a node on the body that can be traveled to
type BodyNode interface {
	Vec() physics.Vector
	Dims() (int, int)
	SetPos(physics.Vector)
	Organ() (Organ, bool)
	Infect(...float64) bool
	DiseaseLevel() float64
	R() render.Modifiable
}

func DemoBody() *Body {
	b := new(Body)
	b.infected = []int{}
	b.cleansed = make(map[int]bool)
	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewVeinNode(10, 10, b.veinColor),
		NewVeinNode(15, 20, b.veinColor),
		NewLiver(40, 5),
		NewHeart(50, 40),
		NewBrain(50, 10),
		NewLung(30, 30),
		NewStomach(10, 40))
	b.Connect(0, 1)
	b.Connect(0, 2)
	b.Connect(1, 2)
	b.Connect(1, 3)
	b.Connect(1, 4)
	b.Connect(1, 5)
	b.Connect(1, 6)

	//b.Infect(0)
	//b.Infect(1)
	//b.Infect(2)
	b.Infect(3)
	//b.Infect(4)
	//b.Infect(5)
	//b.Infect(6)
	b.InitVeins()
	return b
}

func (b *Body) VecIndex(v physics.Vector) int {
	for i, n := range b.graph {
		dist := NodeCenter(n).Distance(v)
		if dist < 3 {
			return i
		}
	}
	return -1
}

func spreadInfection(id int, nothing interface{}) int {
	if traveler.active {
		for i, n := range thisBody.graph {
			if n.DiseaseLevel() > 0 {
				n.Infect()
				for j, v := range thisBody.veins[i] {
					if v != nil {
						v.Refresh(n, thisBody.graph[j], thisBody)
					}
				}
			}
		}
	}
	return 0
}

// Random body ideas:
// given a shape from an image,
// we can put in random locations for everything, where everything is a random
// set of organs / veins
