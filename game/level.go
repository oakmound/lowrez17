package game

import (
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

var (
	envFriction = 0.7
)

func LevelInit(prevScene string, body interface{}) {
	// b := body.(Body)
	// Will remove this once we actually get bodies into scenes
	b := DemoBody()
	var firstVein bool
	var playerStart int
	render.Draw(b.overlay, bodyOverlayLayer)
	for i, n := range b.graph {
		pos := n.Vec()
		var r render.Modifiable
		if o, ok := n.Organ(); ok {
			r = o.R()
			w, h := r.GetDims()
			pos = pos.Copy().Sub(physics.NewVector(float64(w)/2, float64(h)/2))
			r.SetPos(pos.X(), pos.Y())
			n.SetPos(pos)
			render.Draw(r, organLayer)
		} else {
			if firstVein {
				playerStart = i
			}
			firstVein = false
			// We could modify color at nodes
			r := render.NewColorBox(veinNodeWidth, veinNodeWidth, b.veinColor)
			pos.Sub(physics.NewVector(1, 1))
			r.SetPos(pos.X(), pos.Y())
			n.SetPos(pos)
			render.Draw(r, veinLayer)

		}
	}
	// This will draw all veins twice
	for i, list := range b.adjacency {
		for _, j := range list {
			v := NewVein(b.graph[i], b.graph[j], b.veinColor)
			render.Draw(v, veinLayer)
		}
	}
	// Place player
	pos := b.graph[playerStart].Vec()
	NewBodyTraveler(pos.X(), pos.Y())
	// Bindings ...
	//BindBodyTravel()
}

func LevelLoop() bool {
	return true
}

func LevelEnd() (nextScene string, result *oak.SceneResult) {
	return "firstScene", nil
}
