package game

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/shape"
	"github.com/oakmound/oak/timing"
)

type Organ interface {
	BodyNode
	Place()
}

type basicOrgan struct {
	physics.Vector
	*BodyButton
	Infectable
	tiles [][]Tile
	typ   OrganType
	waves []Wave
	w, h  float64
}

func (b *basicOrgan) Place() {
	oak.SetViewportBounds(0, 0, len(b.tiles)*tileDim, len(b.tiles[0])*tileDim)
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			t.Place(x, y, b.typ)
		}
	}
	go timing.DoAfter(time.Second, func() {
		go handleWaves(b.waves, b.tiles, b.typ)
	})
}

func (b *basicOrgan) SetPos(v physics.Vector) {
	b.Vector.SetPos(v.X(), v.Y())
	b.BodyButton.SetPos(v)
}

func (b *basicOrgan) Dims() (int, int) {
	return int(b.w), int(b.h)
}

func (b *basicOrgan) Organ() (Organ, bool) {
	return b, true
}

//NewBasicOrgan creates a new default organ
func NewBasicOrgan(x, y float64, w, h float64, r render.Modifiable, typ OrganType) *basicOrgan {
	bo := &basicOrgan{}
	bo.Vector = physics.NewVector(x, y)
	bo.r = render.NewReverting(r)
	// for now this is a test map, each NewXXX function will populate this themsleves
	//bo.tiles = ShapeTiles(shape.Heart, 64, 64)
	//bo.tiles[32][50] = Exit
	bo.tiles = levels[typ][rand.Intn(5)]
	bo.typ = typ
	bo.BodyButton = NewBodyButton(float64(w), float64(h))
	//bo.waves = []Wave{
	//	Wave{SmallRangedDist, 1.0, 10 * time.Second},
	//	Wave{SmallMeleeDist, 1.0, 30 * time.Second}}
	bo.waves = []Wave{
		Wave{SingleMelee, 1.0, time.Second},
	}
	bo.w = w
	bo.h = h
	bo.diseaseRate = .001
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
			case color.RGBA{255, 0, 110, 255}:
				out[x][y] = Anchor
			case color.RGBA{0, 148, 255, 255}:
				out[x][y] = PressureFan
			case color.RGBA{255, 106, 0, 255}:
				out[x][y] = Acid
			case color.RGBA{255, 0, 0, 255}:
				out[x][y] = Ventricle
			}
		}
	}
	return out
}

//TODO: Refactor this name
//CleanupActiveOrgan cleans up when leaving an organ to return to body map
func CleanupActiveOrgan(cleared bool) {
	stopPlayer()
	CleanupTiles()
	CleanupEnemies()
	oak.SetScreen(0, 0)
	oak.ClearPalette()

	i := thisBody.VecIndex(traveler.Vector)
	o := thisBody.graph[i]
	if cleared {
		o.Cleanse()
		for j, v := range thisBody.veins[i] {
			if v != nil {
				v.Refresh(o, thisBody.graph[j], thisBody)
			}
		}
	}
	if o.DiseaseLevel() == 0 || o.DiseaseLevel() == 1 {
		thisBody.InfectionProgress()
	}
	traveler.active = true
	select {
	case waveExitCh <- true:
	default:
	}
}
