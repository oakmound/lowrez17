package game

import (
	"image/color"

	"github.com/200sc/go-dist/colorrange"
	"github.com/oakmound/lowrez17/game/layers"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

const (
	tileDim            = 4
	tileDimf64 float64 = 4.0
)

type Tile int

type OrganType int

const (
	Liver OrganType = iota
	Lung
	Heart
	Stomach
	Brain
)

const (
	Open Tile = iota
	Blocked
	Exit
	PlayerStart
	Anchor
	PressureFan
	Ventricle
	Acid
	LowDamage
	HighDamage
	// ...
)

var (
	tileColors = map[OrganType]map[Tile]colorrange.Range{
		Liver: {
			Open:    colorrange.NewLinear(color.RGBA{60, 50, 160, 254}, color.RGBA{90, 70, 200, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{30, 25, 80, 254}, color.RGBA{45, 35, 100, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
		Brain: {
			Open:       colorrange.NewLinear(color.RGBA{230, 50, 5, 254}, color.RGBA{254, 60, 140, 254}),
			Blocked:    colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 60, 254}),
			HighDamage: colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 60, 254}),
			Exit:       colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
		Heart: {
			Open:      colorrange.NewLinear(color.RGBA{230, 10, 5, 254}, color.RGBA{254, 20, 60, 254}),
			Blocked:   colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 30, 254}),
			Exit:      colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			Ventricle: colorrange.NewLinear(color.RGBA{70, 10, 10, 254}, color.RGBA{80, 15, 15, 255}),
			LowDamage: colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 30, 254}),
		},
		Lung: {
			Open:        colorrange.NewLinear(color.RGBA{20, 140, 100, 254}, color.RGBA{40, 230, 180, 254}),
			Blocked:     colorrange.NewLinear(color.RGBA{30, 35, 40, 254}, color.RGBA{40, 60, 90, 254}),
			Exit:        colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			PressureFan: colorrange.NewLinear(color.RGBA{220, 220, 220, 254}, color.RGBA{230, 254, 240, 254}),
		},
		Stomach: {
			Open:    colorrange.NewLinear(color.RGBA{230, 230, 5, 254}, color.RGBA{254, 254, 140, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{110, 110, 5, 254}, color.RGBA{140, 140, 60, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			Acid:    colorrange.NewLinear(color.RGBA{0, 200, 0, 255}, color.RGBA{40, 255, 40, 255}),
		},
	}
	tileInit = map[Tile]func(x, y int, r render.Renderable){
		Open: func(int, int, render.Renderable) {},
		PlayerStart: func(x, y int, r render.Renderable) {
			player.SetPos(float64(x)*tileDimf64, float64(y)*tileDimf64)
		},
		Blocked:     addTo(&walls),
		LowDamage:   addTo(&lowDamageWalls),
		HighDamage:  addTo(&highDamageWalls),
		Exit:        addTileSpace(collision.Label(Exit)),
		Anchor:      addTo(&anchors),
		PressureFan: addTo(&fans),
		Acid:        addTileSpace(collision.Label(Acid)),
		Ventricle:   NewVent,
	}
	tileUnDraw = map[Tile]bool{
		Ventricle: true,
	}
	tileRs     = []render.Renderable{}
	tileSpaces = []*collision.Space{}
)

func (t Tile) Place(x, y int, typ OrganType) {
	var c colorrange.Range
	var ok bool
	if c, ok = tileColors[typ][t]; !ok {
		c = tileColors[typ][Open]
	}
	cb := render.NewColorBox(tileDim, tileDim, c.Poll())
	if !tileUnDraw[t] {
		cb.SetPos(float64(x)*tileDimf64, float64(y)*tileDimf64)
		render.Draw(cb, layers.TileLayer)
		tileRs = append(tileRs, cb)
	}
	tileInit[t](x, y, cb)
}

func addTileSpace(l collision.Label) func(x, y int, r render.Renderable) {
	return func(x, y int, r render.Renderable) {
		s := collision.NewLabeledSpace(float64(x)*tileDimf64, float64(y)*tileDimf64, tileDimf64, tileDimf64, l)
		collision.Add(s)
		tileSpaces = append(tileSpaces, s)
	}
}

//CleanupTiles removes and undraws all tiles
func CleanupTiles() {
	for _, r := range tileRs {
		r.UnDraw()
	}
	collision.Remove(tileSpaces...)
	tileRs = []render.Renderable{}
	tileSpaces = []*collision.Space{}
	anchors = []physics.Vector{}
	walls = []physics.Vector{}
}
