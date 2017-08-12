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

	organSheet := render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))

	images["midliver"] = organSheet[0][0].Copy()
	images["midliver"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midlung"] = organSheet[1][0].Copy()
	images["midlung"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midstomach"] = organSheet[2][0].Copy()
	images["midstomach"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midheart"] = organSheet[3][0].Copy()
	images["midheart"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	images["midbrain"] = organSheet[4][0].Copy()
	images["midbrain"].Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))

	enemySheet1 := render.GetSheet(filepath.Join("8x8", "genericfoes.png"))
	enemySheet2 := render.GetSheet(filepath.Join("16x16", "specialfoes.png"))

	images["meleeFoe"] = render.NewReverting(enemySheet1[0][0].Copy())
	images["rangedFoe"] = render.NewReverting(enemySheet1[0][1].Copy())
	images["stomachFoe"] = render.NewReverting(render.NewCompound("base", map[string]render.Modifiable{
		"base": enemySheet2[0][0].Copy().Modify(render.Rotate(180)),
		"attacking": render.NewSequence([]render.Modifiable{
			enemySheet2[0][0].Copy().Modify(render.Rotate(180)),
			enemySheet2[1][0].Copy().Modify(render.Rotate(180)),
			enemySheet2[2][0].Copy().Modify(render.Rotate(180))}, 2),
	}))
	images["heartFoe"] = render.NewReverting(render.NewSequence([]render.Modifiable{
		enemySheet2[0][1].Copy(),
		enemySheet2[0][1].Copy(),
		enemySheet2[0][1].Copy(),
		enemySheet2[1][1].Copy(),
		enemySheet2[2][1].Copy(),
		enemySheet2[1][1].Copy(),
	}, 4))
	images["liverFoe"] = render.NewReverting(render.NewCompound("base", map[string]render.Modifiable{
		"base": enemySheet2[0][2].Copy().Modify(render.Rotate(90)),
		"attacking": render.NewSequence([]render.Modifiable{
			enemySheet2[0][2].Copy().Modify(render.Rotate(90)),
			enemySheet2[1][2].Copy().Modify(render.Rotate(90)),
			enemySheet2[2][2].Copy().Modify(render.Rotate(90))}, 2),
	}))
	images["lungFoe"] = render.NewReverting(enemySheet2[0][4].Copy())
	images["brainFoe"] = render.NewReverting(
		render.NewCompound("base", map[string]render.Modifiable{
			"base": enemySheet2[0][3].Copy(),
			"teleLeft": render.NewSequence(
				[]render.Modifiable{
					enemySheet2[0][3].Copy(),
					enemySheet2[1][3].Copy(),
					enemySheet2[2][3].Copy()}, 2,
			),
			"teleRight": render.NewSequence(
				[]render.Modifiable{
					enemySheet2[0][3].Copy().Modify(render.Rotate(180)),
					enemySheet2[1][3].Copy().Modify(render.Rotate(180)),
					enemySheet2[2][3].Copy().Modify(render.Rotate(180))}, 2,
			),
			"teleBack": render.NewSequence(
				[]render.Modifiable{
					enemySheet2[0][3].Copy().Modify(render.Rotate(-90)),
					enemySheet2[1][3].Copy().Modify(render.Rotate(-90)),
					enemySheet2[2][3].Copy().Modify(render.Rotate(-90))}, 2,
			),
			"teleForward": render.NewSequence(
				[]render.Modifiable{
					enemySheet2[0][3].Copy().Modify(render.Rotate(90)),
					enemySheet2[1][3].Copy().Modify(render.Rotate(90)),
					enemySheet2[2][3].Copy().Modify(render.Rotate(90))}, 2,
			),
		}))

	levelTypes := map[string]OrganType{
		"liver":   Liver,
		"lung":    Lung,
		"heart":   Heart,
		"stomach": Stomach,
		"brain":   Brain,
	}
	for k, v := range levelTypes {
		for i := 1; i < 6; i++ {
			s := k + strconv.Itoa(i) + ".png"
			sp := render.LoadSprite(filepath.Join("raw", s))
			ts := ImageTiles(sp.GetRGBA())
			levels[v] = append(levels[v], ts)
		}
	}

	levelBodies = map[string]*Body{
		"level1": Body1(),
		"level2": DemoBody(),
		"level3": DemoBody(),
		"level4": DemoBody(),
		"level5": DemoBody(),
	}
}
