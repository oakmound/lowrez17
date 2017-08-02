package game

import (
	"image/color"

	"github.com/200sc/go-dist/colorrange"
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
	}
)

func (t Tile) Place(x, y int, typ TileType) {
	cb := render.NewColorBox(tileDim, tileDim, tileColors[typ][t].Poll())
	cb.SetPos(float64(x)*tileDimf64, float64(y)*tileDimf64)
	render.Draw(cb, tileLayer)
}
