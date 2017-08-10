package game

import (
	"image/color"

	"github.com/200sc/go-dist/colorrange"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
)

type Vent struct {
	cmp  *render.Compound
	s    *collision.Space
	open bool
	event.CID
}

func (v *Vent) Init() event.CID {
	v.CID = event.NextID(v)
	return v.CID
}

var (
	openVentColor = colorrange.NewLinear(color.RGBA{200, 120, 120, 255}, color.RGBA{210, 130, 130, 255})
)

func NewVent(x, y int, r render.Renderable) {
	v := new(Vent)
	m1 := r.(render.Modifiable)
	m2 := render.NewColorBox(tileDim, tileDim, openVentColor.Poll())
	v.cmp = render.NewCompound("closed", map[string]render.Modifiable{
		"closed": m1,
		"open":   m2,
	})
	v.cmp.SetPos(float64(x)*tileDimf64, float64(y)*tileDimf64)
	render.Draw(v.cmp, tileLayer)
	v.s = collision.NewLabeledSpace(float64(x)*tileDimf64, float64(y)*tileDimf64,
		tileDimf64, tileDimf64, collision.Label(Blocked))
	collision.Add(v.s)
	v.Bind(toggleVent, "Heartbeat")
}

func toggleVent(id int, nothing interface{}) int {
	v := event.GetEntity(id).(*Vent)
	if v.open {
		v.open = false
		collision.Add(v.s)
		v.cmp.Set("closed")
	} else {
		v.open = true
		collision.Remove(v.s)
		v.cmp.Set("open")
	}
	return 0
}
