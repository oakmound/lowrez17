package game

import (
	"image/color"

	"math/rand"
	"time"

	"github.com/oakmound/lowrez17/game/menu"
	"github.com/oakmound/lowrez17/game/sfx"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Body struct {
	overlay               render.Modifiable
	veinColor, veinColor2 color.RGBA
	graph                 []BodyNode
	adjacency             [][]int
	veins                 [][]*Vein

	infectionPattern [][]int
	infectionSet     int
	complete         bool

	level     int
	startTime time.Time
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
	b.graph[i].Infect(float64(5-b.level)*0.08 + rand.Float64()*0.1)
	//if _, ok := b.cleansed[i]; ok {
	//	return
	//}
	//if b.graph[i].Infect(.3 + rand.Float64()*0.1) {
	//	b.infected = append(b.infected, i)
	//}
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

//InfectionProgress is called when an organ is finished and is responsible for
// updating overall infection and level progress
func (b *Body) InfectionProgress() {
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		if b.graph[oNum].DiseaseLevel() != 0 && b.graph[oNum].DiseaseLevel() != 1 {
			return
		}
	}
	b.infectionSet++
	if b.infectionSet >= len(b.infectionPattern) {
		b.infectionSet--

		<-time.After(2 * time.Second)
		thisBody.complete = true
		return
	}
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}
}

func spreadInfection(id int, frame interface{}) int {

	if traveler.active && frame.(int)%20 == 0 {
		for i, n := range thisBody.graph {
			if n.DiseaseLevel() > 0 {
				if n.Infect() == true {
					thisBody.InfectionProgress()
				}
				for j, v := range thisBody.veins[i] {
					if v != nil {
						v.Refresh(n, thisBody.graph[j], thisBody)
					}
				}
			}
		}
	} else if !traveler.active {
		o := thisBody.graph[thisBody.VecIndex(traveler.Vector)]
		if frame.(int)%200 == 0 {
			oak.SetScreenFilter(FadeBy(o.DiseaseLevel()))
			if o.Infect() {
				sfx.Audios["FailOrgan"].Play()
				CleanupActiveOrgan(false)
			}
		}
	}

	return 0
}

func FadeBy(f float64) render.InPlaceMod {
	return render.InPlace(render.Saturate(-float32(f) * 99))
}

// Random body ideas:
// given a shape from an image,
// we can put in random locations for everything, where everything is a random
// set of organs / veins

func (b *Body) Stats() menu.LevelStats {
	cleared := 0.0
	total := 0.0
	for _, col := range b.infectionPattern {
		for _, o := range col {
			if b.graph[o].DiseaseLevel() <= 0 {
				cleared++
			}
			total++
		}
	}
	cleared /= total
	return menu.LevelStats{
		Score:   -1,
		Time:    time.Now().Sub(b.startTime).Nanoseconds(),
		Cleared: cleared,
		Level:   b.level,
	}
}
