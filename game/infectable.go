package game

import "github.com/oakmound/oak/render"

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
	i.Disease += infection
	if i.Disease > 1 {
		i.Disease = 1
	}

	// Update renderable
	//i.r.(*render.Reverting).RevertAll()
	//i.r.(*render.Reverting).Modify(render.Fade(int(-i.Disease)))
	return out
}

func (i *Infectable) R() render.Modifiable {
	return i.r
}

func (i *Infectable) DiseaseLevel() float64 {
	return i.Disease
}
