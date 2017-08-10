package game

import (
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
	levelBodies = map[string]*Body{}
)

func Init() {
	if inited {
		return
	}
	inited = true
	images["midliver"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[0][0].Copy()
	images["midliver"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midlung"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[1][0].Copy()
	images["midlung"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midstomach"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[2][0].Copy()
	images["midstomach"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midheart"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[3][0].Copy()
	images["midheart"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midbrain"] = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))[4][0].Copy()
	images["midbrain"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	levelTypes := map[string]OrganType{
		"liver":   Liver,
		"lung":    Lung,
		"heart":   Heart,
		"stomach": Stomach,
	}
	for k, v := range levelTypes {
		for i := 1; i < 6; i++ {
			s := k + strconv.Itoa(i) + ".png"
			sp := render.LoadSprite(filepath.Join("raw", s))
			ts := ImageTiles(sp.GetRGBA())
			levels[v] = append(levels[v], ts)
		}
	}
	// Placeholder
	levels[Brain] = levels[Liver]

	levelBodies = map[string]*Body{
		"level1": Body1(),
		"level2": DemoBody(),
		"level3": DemoBody(),
		"level4": DemoBody(),
		"level5": DemoBody(),
	}
}
