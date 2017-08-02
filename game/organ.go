package game

import (
	"image/color"

	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

type Organ interface {
	BodyNode
	Place()
	R() render.Modifiable
}

type basicOrgan struct {
	physics.Vector
	r     *render.Sprite
	tiles [][]Tile
	typ   TileType
}

func (b basicOrgan) R() render.Modifiable {
	return b.r
}

func (b basicOrgan) Place() {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			t.Place(x, y, b.typ)
		}
	}
}

func (b basicOrgan) Organ() (Organ, bool) {
	return b, true
}

func NewLiver(x, y float64) Organ {
	bo := basicOrgan{}
	bo.Vector = physics.NewVector(x, y)
	bo.r = render.NewColorBox(6, 4, color.RGBA{240, 170, 230, 255})
	// get some liver map
	bo.typ = LiverTile
	return bo
}
