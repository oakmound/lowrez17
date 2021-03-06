package game

import (
	"image/color"
	"path/filepath"
	"strconv"

	"github.com/oakmound/oak/dlog"
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
	levelBodies     = map[string]*Body{}
	hurtPalette     = []color.Color{}
	diseasedPalette = []color.Color{}
	grayScale       = []color.Color{}
)

func Init() {

	if inited {
		levelBodies = map[string]*Body{
			"level1": Body1(),
			"level2": Body2(),
			"level3": Body3(),
			"level4": Body4(),
			"level5": Body5(),
		}
		return
	}
	inited = true

	bodySheet := render.GetSheet(filepath.Join("64x64", "midbodies.png"))
	images["body1"] = bodySheet[0][0]
	images["body2"] = bodySheet[0][1]
	images["body3"] = bodySheet[0][2]
	images["body4"] = bodySheet[1][0]
	images["body5"] = bodySheet[1][1]
	organSheet := render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))

	images["midliver"] = organSheet[0][0].Copy()
	images["midlung"] = organSheet[1][0].Copy()
	images["midstomach"] = organSheet[2][0].Copy()
	images["midheart"] = organSheet[3][0].Copy()
	images["midbrain"] = organSheet[4][0].Copy()

	enemySheet1 := render.GetSheet(filepath.Join("8x8", "genericfoes.png"))
	enemySheet2 := render.GetSheet(filepath.Join("16x16", "specialfoes.png"))

	images["meleeFoe"] = render.NewReverting(enemySheet1[0][0].Copy())
	images["rangedFoe"] = render.NewReverting(enemySheet1[0][1].Copy())
	images["stomachFoe"] = render.NewReverting(render.NewCompound("base", map[string]render.Modifiable{
		"base": enemySheet2[0][0].Copy(),
		"attacking": render.NewSequence([]render.Modifiable{
			enemySheet2[0][0].Copy(),
			enemySheet2[1][0].Copy(),
			enemySheet2[2][0].Copy(),
			enemySheet2[2][0].Copy(),
			enemySheet2[2][0].Copy(),
			enemySheet2[1][0].Copy(),
			enemySheet2[1][0].Copy()}, 2),
	}))
	sh, err := render.LoadSheet("images", filepath.Join("16x16", "specialfoes.png"), 16, 16, 0)
	if err != nil {
		dlog.Error(err)
	}
	an, err := render.NewAnimation(sh, 2, []int{0, 1, 0, 1, 0, 1, 1, 1, 2, 1, 1, 1})
	if err != nil {
		dlog.Error(err)
	}
	images["heartFoe"] = render.NewReverting(an)
	images["liverFoe"] = render.NewReverting(render.NewCompound("base", map[string]render.Modifiable{
		"base": enemySheet2[0][2].Copy(),
		"attacking": render.NewSequence([]render.Modifiable{
			enemySheet2[0][2].Copy(),
			enemySheet2[1][2].Copy(),
			enemySheet2[2][2].Copy(),
			enemySheet2[2][2].Copy(),
			enemySheet2[2][2].Copy(),
			enemySheet2[1][2].Copy()}, 1),
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

	images["whip"] = render.LoadSprite(filepath.Join("raw", "whip.png"))
	images["sword"] = render.LoadSprite(filepath.Join("raw", "sword.png"))
	images["net"] = render.LoadSprite(filepath.Join("raw", "net.png"))
	images["spear"] = render.LoadSprite(filepath.Join("raw", "spear.png"))

	for k, v := range images {
		images[k] = v.Modify(render.TrimColor(color.RGBA{1, 1, 1, 1}))
	}

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

	for i := uint8(0); i < 127; i++ {
		hurtPalette = append(hurtPalette, color.RGBA{255, i * 2, i * 2, 255})
	}
	for i := uint8(0); i < 127; i++ {
		grayScale = append(grayScale, color.RGBA{i * 2, i * 2, i * 2, 255})
	}

	//for i := uint8(0); i < 127; i++ {
	//	diseasedPalette = append(diseasedPalette, color.RGBA{(i - 127) * 2, (i - 127) * 2, (i - 127) * 2, 255})
	//}

	levelBodies = map[string]*Body{
		"level1": Body1(),
		"level2": Body2(),
		"level3": Body3(),
		"level4": Body4(),
		"level5": Body5(),
	}
}
