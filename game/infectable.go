package game

import (
	"github.com/oakmound/oak/render"

	"image/color"
)

type Infectable struct {
	Disease     float64
	diseaseRate float64
	r           render.Modifiable
}

func (i *Infectable) Infect(fs ...float64) bool {
	var infection float64
	if len(fs) == 0 {
		infection = i.diseaseRate
	} else {
		for _, f := range fs {
			infection += f
		}
	}
	out := i.Disease == 0
	//pastDisease := i.Disease
	i.Disease += infection
	if i.Disease > 1 {
		i.Disease = 1
	}
	if len(fs) != 0 {
		//Infect with fs is currently used only for setup
		//If modifications are applied while in setup (predraw) it can cause the image to disappear on revert.
		return true
	}

	// Update renderable
	i.r.(*render.Reverting).RevertAll()

	i.r.(*render.Reverting).Modify(render.Brighten(float32(i.Disease) * 10))
	if int(i.Disease/i.diseaseRate)%40 < 10 {
		i.r.(*render.Reverting).Modify(render.Fade(int(-i.Disease)))
		i.r.(*render.Reverting).Modify(render.ApplyColor(color.RGBA{0, 0, 255, 255}))
	}

	return out
}

func (i *Infectable) R() render.Modifiable {
	return i.r
}

func (i *Infectable) DiseaseLevel() float64 {
	return i.Disease
}
