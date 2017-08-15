package game

import (
	"fmt"

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
			fmt.Println(f)
			infection += f
		}
	}

	//pastDisease := i.Disease
	i.Disease += infection
	if i.Disease >= 1 {
		i.Disease = 1
	}
	if len(fs) != 0 {
		//Infect with fs is currently used only for setup
		//If modifications are applied while in setup (predraw) it can cause the image to disappear on revert.
		return i.Disease == 1
	}

	// Update renderable

	if i.Disease == 1 {
		i.r.(*render.Reverting).RevertAndModify(1, render.ApplyColor(color.RGBA{0, 0, 255, 255}))
	} else if int(i.Disease/i.diseaseRate)%20 < 8 {
		i.r.(*render.Reverting).RevertAndModify(1,
			render.And(render.Fade(int(-i.Disease)), render.ApplyColor(color.RGBA{0, 0, 255, 255})))
	} else {
		i.r.(*render.Reverting).RevertAndModify(1, render.Brighten(float32(i.Disease)*10))
	}
	return i.Disease == 1
}

func (i *Infectable) R() render.Modifiable {
	return i.r
}

func (i *Infectable) DiseaseLevel() float64 {
	return i.Disease
}

func (i *Infectable) Cleanse() {
	i.Disease = 0
	i.r.(*render.Reverting).RevertAll()
}
