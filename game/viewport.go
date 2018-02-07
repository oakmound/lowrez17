package game

import (
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/alg"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/mouse"
	"github.com/oakmound/oak/physics"
)

func viewportFollow(id int, frame interface{}) int {
	e := event.GetEntity(id).(HasE).E()
	v := physics.NewVector(mouse.LastEvent.X(), mouse.LastEvent.Y())
	viewportGoalPos := e.CenterPos().Copy().Add(v.Sub(physics.NewVector(32, 32)))
	viewportGoalPos = viewportGoalPos.Sub(physics.NewVector(float64(oak.ScreenWidth/2), float64(oak.ScreenHeight/2)))
	delta := viewportGoalPos.Sub(physics.NewVector(float64(oak.ViewPos.X), float64(oak.ViewPos.Y)))

	if delta.Magnitude() < 2 {
		return 0
	}

	oak.SetScreen(
		alg.RoundF64(float64(oak.ViewPos.X)+delta.X()/9),
		alg.RoundF64(float64(oak.ViewPos.Y)+delta.Y()/9))

	return 0
}
