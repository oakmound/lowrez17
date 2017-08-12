package game

import "github.com/oakmound/oak/render"

func NewLiver(x, y float64) Organ {
	r := images["midliver"].Copy()
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Liver)
}

func NewHeart(x, y float64) Organ {
	r := images["midheart"].Copy()
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Heart)
}

func NewLung(x, y float64) Organ {
	r := images["midlung"].Copy()
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Lung)
}
func NewRLung(x, y float64) Organ {
	r := images["midlung"].Copy()
	r = r.Modify(render.FlipX).Copy()
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Lung)
}

func NewStomach(x, y float64) Organ {
	r := images["midstomach"].Copy()
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Stomach)
}

func NewBrain(x, y float64) Organ {
	r := images["midbrain"].Copy()
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Brain)
}
