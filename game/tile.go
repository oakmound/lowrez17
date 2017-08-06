package game

import (
	"image/color"

	"github.com/200sc/go-dist/colorrange"
	"github.com/oakmound/oak/collision"
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
	// ...
)

var (
	tileColors = map[OrganType]map[Tile]colorrange.Range{
		Liver: {
			Open:        colorrange.NewLinear(color.RGBA{60, 50, 160, 254}, color.RGBA{90, 70, 200, 254}),
			Blocked:     colorrange.NewLinear(color.RGBA{30, 25, 80, 254}, color.RGBA{45, 35, 100, 254}),
			Exit:        colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			PlayerStart: colorrange.NewLinear(color.RGBA{60, 50, 160, 254}, color.RGBA{90, 70, 200, 254}),
		},
		Brain: {
			Open:        colorrange.NewLinear(color.RGBA{230, 50, 5, 254}, color.RGBA{254, 60, 140, 254}),
			Blocked:     colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 60, 254}),
			Exit:        colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			PlayerStart: colorrange.NewLinear(color.RGBA{230, 50, 5, 254}, color.RGBA{254, 60, 140, 254}),
		},
		Heart: {
			Open:        colorrange.NewLinear(color.RGBA{230, 10, 5, 254}, color.RGBA{254, 20, 60, 254}),
			Blocked:     colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 30, 254}),
			Exit:        colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			PlayerStart: colorrange.NewLinear(color.RGBA{230, 10, 5, 254}, color.RGBA{254, 20, 60, 254}),
		},
		Lung: {
			Open:        colorrange.NewLinear(color.RGBA{100, 120, 125, 254}, color.RGBA{190, 200, 210, 254}),
			Blocked:     colorrange.NewLinear(color.RGBA{30, 35, 40, 254}, color.RGBA{40, 60, 90, 254}),
			Exit:        colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			PlayerStart: colorrange.NewLinear(color.RGBA{100, 120, 125, 254}, color.RGBA{190, 200, 210, 254}),
		},
		Stomach: {
			Open:        colorrange.NewLinear(color.RGBA{230, 230, 5, 254}, color.RGBA{254, 254, 140, 254}),
			Blocked:     colorrange.NewLinear(color.RGBA{110, 110, 5, 254}, color.RGBA{140, 140, 60, 254}),
			Exit:        colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
			PlayerStart: colorrange.NewLinear(color.RGBA{230, 230, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
	}
	tileInit = map[Tile]func(x, y int){
		Open: func(int, int) {},
		PlayerStart: func(x, y int) {
			player.SetPos(float64(x)*tileDimf64, float64(y)*tileDimf64)
		},
		Blocked: addTileSpace(collision.Label(Blocked)),
		Exit:    addTileSpace(collision.Label(Exit)),
	}
	tileRs     = []render.Renderable{}
	tileSpaces = []*collision.Space{}
)

func (t Tile) Place(x, y int, typ OrganType) {
	cb := render.NewColorBox(tileDim, tileDim, tileColors[typ][t].Poll())
	cb.SetPos(float64(x)*tileDimf64, float64(y)*tileDimf64)
	render.Draw(cb, tileLayer)
	tileInit[t](x, y)
	tileRs = append(tileRs, cb)
}

func addTileSpace(l collision.Label) func(x, y int) {
	return func(x, y int) {
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

}
