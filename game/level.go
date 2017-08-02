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
	render.Draw(b.overlay, bodyOverlayLayer)
	for _, n := range b.graph {
		pos := n.Vec()
		var r render.Modifiable
		if o, ok := n.Organ(); ok {
			r = o.R()
			w, h := r.GetDims()
			pos = pos.Copy().Sub(physics.NewVector(float64(w)/2, float64(h)/2))
			r.SetPos(pos.X(), pos.Y())
			render.Draw(r, organLayer)
		} else {
			// We could modify color at nodes
			r := render.NewColorBox(2, 2, b.veinColor)
			pos.Sub(physics.NewVector(1, 1))
			r.SetPos(pos.X(), pos.Y())
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
	// Bindings ...
}

func LevelLoop() bool {
	return true
}

func LevelEnd() (nextScene string, result *oak.SceneResult) {
	return "firstScene", nil
}
