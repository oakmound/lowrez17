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

type TileType int

const (
	LiverTile TileType = iota
	LungTile
	HeartTile
	StomachTile
	BrainTile
)

const (
	Open Tile = iota
	Blocked
	Exit
	// ...
)

var (
	tileColors = map[TileType]map[Tile]colorrange.Range{
		LiverTile: {
			Open:    colorrange.NewLinear(color.RGBA{230, 50, 5, 254}, color.RGBA{254, 60, 140, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 60, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
		BrainTile: {
			Open:    colorrange.NewLinear(color.RGBA{100, 120, 125, 254}, color.RGBA{190, 200, 210, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{30, 35, 40, 254}, color.RGBA{40, 60, 90, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
		HeartTile: {
			Open:    colorrange.NewLinear(color.RGBA{230, 10, 5, 254}, color.RGBA{254, 20, 60, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{110, 10, 5, 254}, color.RGBA{140, 20, 30, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
		LungTile: {
			Open:    colorrange.NewLinear(color.RGBA{230, 230, 5, 254}, color.RGBA{254, 254, 140, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{110, 110, 5, 254}, color.RGBA{140, 140, 60, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
		StomachTile: {
			Open:    colorrange.NewLinear(color.RGBA{50, 230, 5, 254}, color.RGBA{60, 254, 140, 254}),
			Blocked: colorrange.NewLinear(color.RGBA{10, 110, 5, 254}, color.RGBA{20, 110, 60, 254}),
			Exit:    colorrange.NewLinear(color.RGBA{230, 100, 5, 254}, color.RGBA{254, 254, 140, 254}),
		},
	}
	tileInit = map[Tile]func(x, y int){
		Open:    func(int, int) {},
		Blocked: addTileSpace(collision.Label(Blocked)),
		Exit:    addTileSpace(collision.Label(Exit)),
	}
	tileRs     = []render.Renderable{}
	tileSpaces = []*collision.Space{}
)

func (t Tile) Place(x, y int, typ TileType) {
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

func CleanupTiles() {
	for _, r := range tileRs {
		r.UnDraw()
	}
	collision.Remove(tileSpaces...)
	tileRs = []render.Renderable{}
	tileSpaces = []*collision.Space{}

}
