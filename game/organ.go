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

func NewLiver(x, y float64) Organ {
	bo := &basicOrgan{}
	bo.Vector = physics.NewVector(x, y)
	bo.r = render.NewColorBox(6, 4, color.RGBA{240, 170, 230, 255})
	// get some liver map
	// for now this is a test map
	// bo.tiles = [][]Tile{
	// 	{Exit, Exit, Exit, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked},
	// 	{Exit, Exit, Exit, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked},
	// 	{Exit, Exit, Exit, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Open, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Open, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked},
	// 	{Exit, Exit, Exit, Open, Open, Open, Open, Open, Open, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked, Blocked},
	// }
	bo.tiles = ShapeTiles(shape.Heart, 64, 64)
	bo.tiles[32][50] = Exit
	bo.typ = LiverTile
	bo.BodyButton = NewBodyButton(6, 4)
	return bo
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
