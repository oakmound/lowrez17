package game

import (
	"image/color"
	"time"

	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/timing"
)

type HurtBox struct {
	*DirectionSpace
}

func NewHurtBox(x, y, w, h float64, duration time.Duration, l collision.Label, fv physics.ForceVector) {
	hb := new(HurtBox)
	hb.DirectionSpace = NewDirectionSpace(collision.NewLabeledSpace(x, y, w, h, l), fv)
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
