package game

import (
	"github.com/oakmound/lowrez17/game/forceSpace"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

// We keep track of all anchors and walls, and run an initialization on said anchors and walls
// once they are all tracked

var (
	anchors = []physics.Vector{}
	walls   = []physics.Vector{}
	fans    = []physics.Vector{}
)

func addTo(vs *[]physics.Vector) func(x, y int, r render.Renderable) {
	return func(x, y int, r render.Renderable) {
		*vs = append(*vs, physics.NewVector(float64(x)*tileDimf64, float64(y)*tileDimf64))
	}
}

func startupWalls() {
	for _, w := range walls {
		// Find the minimum distance anchor to this wall
		minDist := w.Distance(anchors[0])
		minV := anchors[0]
		for i := 1; i < len(anchors); i++ {
			dist := w.Distance(anchors[i])
			if minDist > dist {
				minDist = dist
				minV = anchors[i]
			}
		}
		// Initialize a directional collision space pointing toward the nearby anchor
		ds := forceSpace.NewDirectionSpace(collision.NewLabeledSpace(w.X(), w.Y(), tileDimf64, tileDimf64, collision.Label(Blocked)),
			physics.NewForceVector(w.Sub(minV).Normalize(), 10))
		collision.Add(ds.Space)
		tileSpaces = append(tileSpaces, ds.Space)
	}
}

func startupFans() {
	if len(fans) > 1 {
		f := fans[0]
		fans = fans[1:]
		for len(fans) > 0 {
			// Find a close by second fan
			dist := f.Distance(fans[0])
			minFan := 0
			for i := 1; i < len(fans); i++ {
				d := f.Distance(fans[i])
				if d < dist {
					minFan = i
					dist = d
				}
			}
			// set the direction of the fan to be towards that close second fan
			ds := forceSpace.NewDirectionSpace(collision.NewLabeledSpace(f.X()-tileDimf64, f.Y()-tileDimf64, tileDimf64*3, tileDimf64*3, collision.Label(PressureFan)),
				physics.NewForceVector(f.Sub(fans[minFan]).Normalize(), 1))
			collision.Add(ds.Space)
			tileSpaces = append(tileSpaces, ds.Space)

			// update current fan, reduce length of list
			fans[0], fans[minFan] = fans[minFan], fans[0]
			f = fans[0]
			fans = fans[1:]
		}
	}
}
