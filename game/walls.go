package game

import (
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
)

// We keep track of all anchors and walls, and run an initialization on said anchors and walls
// once they are all tracked

var (
	anchors = []physics.Vector{}
	walls   = []physics.Vector{}
)

func addAnchor(x, y int) {
	anchors = append(anchors, physics.NewVector(float64(x)*tileDimf64, float64(y)*tileDimf64))
}

func addWall(x, y int) {
	walls = append(walls, physics.NewVector(float64(x)*tileDimf64, float64(y)*tileDimf64))
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
		ds := NewDirectionSpace(collision.NewLabeledSpace(w.X(), w.Y(), tileDimf64, tileDimf64, collision.Label(Blocked)),
			physics.NewForceVector(w.Sub(minV).Normalize(), 20))
		collision.Add(ds.Space)
	}
}
