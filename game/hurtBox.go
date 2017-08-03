package game

import (
	"image/color"
	"time"

	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/timing"
)

type HurtBox struct {
	*collision.Space
}

func NewHurtBox(x, y, w, h float64, duration time.Duration, l collision.Label) {
	hb := new(HurtBox)
	hb.Space = collision.NewLabeledSpace(x, y, w, h, l)
	collision.Add(hb.Space)
	// Debug renderable to see the hurtbox
	cb := render.NewColorBox(int(w), int(h), color.RGBA{100, 100, 100, 100})
	cb.SetPos(x, y)
	render.Draw(cb, debugLayer)
	go timing.DoAfter(duration, func() {
		collision.Remove(hb.Space)
		cb.UnDraw()
	})
}
