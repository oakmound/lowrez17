package forceSpace

import (
	"image/color"
	"time"

	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/timing"
)

type HurtBox struct {
	*DirectionSpace
}

func NewHurtBox(x, y, w, h float64, duration time.Duration, l collision.Label, fv physics.ForceVector, debug ...bool) {
	hb := new(HurtBox)
	hb.DirectionSpace = NewDirectionSpace(collision.NewLabeledSpace(x, y, w, h, l), fv)
	collision.Add(hb.Space)
	// Debug renderable to see the hurtbox
	if len(debug) == 0 {
		cb := render.NewColorBox(int(w), int(h), color.RGBA{100, 100, 100, 100})
		cb.SetPos(x, y)
		render.Draw(cb, layers.DebugLayer)
		go timing.DoAfter(duration, func() {
			collision.Remove(hb.Space)
			cb.UnDraw()
		})
	} else {
		go timing.DoAfter(duration, func() {
			collision.Remove(hb.Space)
		})
	}
}
