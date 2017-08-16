package game

import (
	"image/color"

	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/lowrez17/game/sfx"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

var (
	envFriction = 0.3
	traveler    *BodyTraveler
	thisBody    *Body
	player      *Player
)

//LevelInit sets up a level
func LevelInit(prevScene string, inLevel interface{}) {

	Init()
	sfx.Audios["fantastic"].Play()
	b := GetBody(inLevel.(string))
	player = nil

	thisBody = b
	var firstVein = true
	var playerStart int
	render.Draw(b.overlay, layers.BodyOverlayLayer)
	for i, n := range b.graph {
		pos := n.Vec()
		var r render.Modifiable
		if o, ok := n.Organ(); ok {
			r = o.R()
			w, h := r.GetDims()
			pos = pos.Copy().Sub(physics.NewVector(float64(w)/2, float64(h)/2))
			r.SetPos(pos.X(), pos.Y())
			n.SetPos(pos)
			render.Draw(r, layers.OrganLayer)
		} else {
			if firstVein {
				playerStart = i
			}
			firstVein = false
			// We could modify color at nodes
			r := n.R()
			pos.Sub(physics.NewVector(1, 1))
			r.SetPos(pos.X(), pos.Y())
			n.SetPos(pos)
			render.Draw(r, layers.VeinLayer)
		}
	}
	for i, list := range b.adjacency {
		for _, j := range list {
			if i > j {
				v := NewVein(b.graph[i], b.graph[j], b)
				render.Draw(v, layers.VeinLayer)
				b.veins[i][j] = v
				b.veins[j][i] = v
			}
		}
	}
	diseasedPalette = []color.Color{}
	rd, gd, bd, _ := thisBody.veinColor2.RGBA()
	for i := uint8(0); i < 127; i++ {
		diseasedPalette = append(diseasedPalette, color.RGBA{i * (uint8(rd)), i * (uint8(gd)), i * (uint8(bd)), 255})
	}

	// Place player
	pos := NodeCenter(b.graph[playerStart])
	traveler = NewBodyTraveler(pos.X(), pos.Y())
	// Bindings ...
	event.GlobalBind(enterOrgan, "HitNode")
	event.GlobalBind(spreadInfection, "EnterFrame")
	event.GlobalBind(heartBeat, "EnterFrame")
}

func heartBeat(no int, frame interface{}) int {
	f := frame.(int)
	// Todo: increase heart rate with poor body health
	if f%120 == 0 {
		event.Trigger("Heartbeat", nil)
	}
	return 0
}

func enterOrgan(no int, nothing interface{}) int {
	i := thisBody.VecIndex(traveler.Vector)
	if o, ok := thisBody.graph[i].Organ(); ok {
		if o.DiseaseLevel() > 0 && o.DiseaseLevel() < 1 {
			oak.SetScreenFilter(FadeBy(o.DiseaseLevel()))
			traveler.active = false
			NewPlayer()
			o.Place()
			startupWalls()
			startupFans()
			startupPlayer()
		}
	}
	return 0
}

func LevelLoop() bool {
	return !thisBody.complete
}

func LevelEnd() (nextScene string, result *oak.SceneResult) {
	sfx.Audios["fantastic"].Stop()
	oak.ClearScreenFilter()
	return "menu", &oak.SceneResult{
		NextSceneInput: thisBody.Stats(),
		Transition:     oak.TransitionFade(.03, 100),
	}
}
