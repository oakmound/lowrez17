package game

import (
	"fmt"
	"image"
	"image/color"
	"path/filepath"
	"time"

	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/shape"
	"github.com/oakmound/oak/timing"
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
	typ   OrganType
	waves []Wave
}

func (b *basicOrgan) R() render.Modifiable {
	return b.r
}

func (b *basicOrgan) Place() {
	oak.SetViewportBounds(0, 0, len(b.tiles)*tileDim, len(b.tiles[0])*tileDim)
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			t.Place(x, y, b.typ)
		}
	}
	go timing.DoAfter(time.Second, func() { b.PlaceWave(0) })
}

func (b *basicOrgan) PlaceWave(index int) {
	fmt.Println("Placing Wave", index)
	// This assumes most territory is open
	wrange := intrange.NewLinear(0, len(b.tiles)-1)
	hrange := intrange.NewLinear(0, len(b.tiles[0])-1)
	es := b.waves[index].Poll()
	for _, t := range es {
		x := wrange.Poll()
		y := hrange.Poll()
		for b.tiles[x][y] != Open {
			fmt.Println(b.tiles[x][y], x, y)
			x = wrange.Poll()
			y = hrange.Poll()
		}
		e := enemyFns[t][b.typ](x, y, b.waves[index].Difficulty)
		fmt.Println(e)
	}
	// Todo: check what wave is active, time the next wave, clear organ when
	// last wave cleared
}

func (b *basicOrgan) SetPos(v physics.Vector) {
	b.Vector.SetPos(v.X(), v.Y())
	b.BodyButton.SetPos(v)
}

func (b *basicOrgan) Dims() (int, int) {
	return b.r.GetDims()
}

func (b *basicOrgan) Organ() (Organ, bool) {
	return b, true
}

func NewBasicOrgan(x, y float64, w, h int, c color.Color, typ OrganType) *basicOrgan {
	bo := &basicOrgan{}
	bo.Vector = physics.NewVector(x, y)
	// Eventually this will take in a renderable instead of a color
	bo.r = render.NewColorBox(w, h, c)
	// for now this is a test map, each NewXXX function will populate this themsleves
	//bo.tiles = ShapeTiles(shape.Heart, 64, 64)
	//bo.tiles[32][50] = Exit
	bo.tiles = ImageTiles(render.LoadSprite(filepath.Join("raw", "baseliver.png")).GetRGBA())
	bo.typ = typ
	bo.BodyButton = NewBodyButton(float64(w), float64(h))
	bo.waves = []Wave{
		Wave{SmallMeleeDist, 1.0, 10 * time.Second},
		Wave{LargeMeleeDist, 1.0, 30 * time.Second}}
	return bo
}

func NewLiver(x, y float64) Organ {
	return NewBasicOrgan(x, y, 9, 8, color.RGBA{240, 170, 230, 255}, Liver)
}

func NewHeart(x, y float64) Organ {
	return NewBasicOrgan(x, y, 3, 3, color.RGBA{220, 30, 30, 255}, Heart)
}

func NewLung(x, y float64) Organ {
	return NewBasicOrgan(x, y, 3, 8, color.RGBA{240, 220, 80, 255}, Lung)
}

func NewStomach(x, y float64) Organ {
	return NewBasicOrgan(x, y, 8, 6, color.RGBA{120, 210, 50, 255}, Stomach)
}

func NewBrain(x, y float64) Organ {
	return NewBasicOrgan(x, y, 4, 4, color.RGBA{130, 160, 170, 255}, Brain)
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

func ImageTiles(rgba *image.RGBA) [][]Tile {
	rect := rgba.Bounds()
	w, h := rect.Max.X, rect.Max.Y
	out := make([][]Tile, w)
	for x := 0; x < len(out); x++ {
		out[x] = make([]Tile, h)
		for y := 0; y < len(out[x]); y++ {
			c := rgba.At(x, y)
			// This could be more lenient
			// see raw/baseliver.png
			switch c {
			case color.RGBA{255, 255, 255, 255}:
				out[x][y] = Open
			case color.RGBA{0, 0, 0, 255}:
				out[x][y] = Blocked
			case color.RGBA{255, 216, 0, 255}:
				out[x][y] = Exit
			case color.RGBA{0, 255, 33, 255}:
				out[x][y] = PlayerStart
			}
		}
	}
	return out
}
