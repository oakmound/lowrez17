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
	complete              bool
	infectionPattern      [][]int
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

// IsAdjacent returns whether the nodes at indices i and j in this graph are adjacent
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

func (b *Body) InfectionPattern(pattern [][]int) {
	b.infectionPattern = pattern
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

// MonitorInfections Monitors ongoing cleared organs, advancing to the next set
// or beating the level when appropriate
func (b *Body) MonitorInfections() {
	go func() {
		// Todo
	}()
}

func spreadInfection(id int, frame interface{}) int {

	if traveler.active && frame.(int)%4 == 0 {
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
