package game

func NewLiver(x, y float64) Organ {
	r := images["midliver"]
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Liver)
}

func NewHeart(x, y float64) Organ {
	r := images["midheart"]
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Heart)
}

func NewLung(x, y float64) Organ {
	r := images["midlung"]
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Lung)
}

func NewStomach(x, y float64) Organ {
	r := images["midstomach"]
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Stomach)
}

func NewBrain(x, y float64) Organ {
	r := images["midbrain"]
	w, h := r.GetDims()
	return NewBasicOrgan(x, y, float64(w), float64(h), r, Brain)
}
