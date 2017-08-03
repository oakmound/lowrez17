package game

import (
	"image/color"

	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/shape"
)

type Organ interface {
	BodyNode
	Place()
	R() render.Modifiable
}

type basicOrgan struct {
	physics.Vector
	*BodyButton
	r     *render.Sprite
	tiles [][]Tile
	typ   TileType
}

func (b *basicOrgan) R() render.Modifiable {
	return b.r
}

func (b *basicOrgan) Place() {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			t.Place(x, y, b.typ)
		}
	}
}

func (b *basicOrgan) SetPos(v physics.Vector) {
	b.Vector.SetPos(v.X(), v.Y())
	b.BodyButton.SetPos(v)
}

func (b *basicOrgan) Organ() (Organ, bool) {
	return b, true
}

func NewBasicOrgan(x, y float64, w, h int, c color.Color, typ TileType) *basicOrgan {
	bo := &basicOrgan{}
	bo.Vector = physics.NewVector(x, y)
	// Eventually this will take in a renderable instead of a color
	bo.r = render.NewColorBox(w, h, c)
	// for now this is a test map, each NewXXX function will populate this themsleves
	bo.tiles = ShapeTiles(shape.Heart, 64, 64)
	bo.tiles[32][50] = Exit
	bo.typ = typ
	bo.BodyButton = NewBodyButton(float64(w), float64(h))
	return bo
}

func NewLiver(x, y float64) Organ {
	return NewBasicOrgan(x, y, 9, 8, color.RGBA{240, 170, 230, 255}, LiverTile)
}

func NewHeart(x, y float64) Organ {
	return NewBasicOrgan(x, y, 3, 3, color.RGBA{220, 30, 30, 255}, HeartTile)
}

func NewLung(x, y float64) Organ {
	return NewBasicOrgan(x, y, 3, 8, color.RGBA{240, 220, 80, 255}, LungTile)
}

func NewStomach(x, y float64) Organ {
	return NewBasicOrgan(x, y, 8, 6, color.RGBA{120, 210, 50, 255}, StomachTile)
}

func NewBrain(x, y float64) Organ {
	return NewBasicOrgan(x, y, 4, 4, color.RGBA{130, 160, 170, 255}, BrainTile)
}

func ShapeTiles(sh shape.Shape, w, h int) [][]Tile {
	out := make([][]Tile, w)
	for x := 0; x < len(out); x++ {
		out[x] = make([]Tile, h)
		for y := 0; y < len(out[x]); y++ {
			if sh.In(x, y, w, h) {
				out[x][y] = Open
			} else {
				out[x][y] = Blocked
			}
		}
	}
	return out
}
