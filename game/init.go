package game

import (
	"fmt"
	"image/color"
	"path/filepath"
	"strconv"

	"github.com/oakmound/oak/render"
)

var (
	inited bool
	images = map[string]render.Modifiable{}
	levels = map[OrganType][][][]Tile{
		Liver:   [][][]Tile{},
		Stomach: [][][]Tile{},
		Lung:    [][][]Tile{},
		Brain:   [][][]Tile{},
		Heart:   [][][]Tile{},
	}
)

func Init() {
	if inited {
		return
	}
	inited = true
	images["midliver"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[0][0]
	images["midliver"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midlung"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[1][0]
	images["midlung"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midstomach"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[2][0]
	images["midstomach"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midheart"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[3][0]
	images["midheart"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midbrain"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[4][0]
	images["midbrain"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	for i := 1; i < 6; i++ {
		s := "liver" + strconv.Itoa(i) + ".png"
		sp := render.LoadSprite(filepath.Join("raw", s))
		fmt.Println(s, sp)
		ts := ImageTiles(sp.GetRGBA())
		levels[Liver] = append(levels[Liver], ts)
	}
	// Placeholder
	levels[Stomach] = levels[Liver]
	levels[Lung] = levels[Liver]
	levels[Brain] = levels[Liver]
	levels[Heart] = levels[Liver]
}