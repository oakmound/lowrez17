package game

import (
	"image/color"
	"path/filepath"
	"strconv"

	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/render/mod"
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

	shtt, _ := render.GetSheet(filepath.Join("64x64", "midbodies.png"))
	bodySheet := shtt.ToSprites()
	images["body1"] = bodySheet[0][0]
	images["body2"] = bodySheet[0][1]
	images["body3"] = bodySheet[0][2]
	images["body4"] = bodySheet[1][0]
	images["body5"] = bodySheet[1][1]
	shtt, _ = render.GetSheet(filepath.Join("16x16", "midlevelorgans.png"))
	organSheet := shtt.ToSprites()
	images["midliver"] = organSheet[0][0].Copy()
	images["midlung"] = organSheet[1][0].Copy()
	images["midstomach"] = organSheet[2][0].Copy()
	images["midheart"] = organSheet[3][0].Copy()
	images["midbrain"] = organSheet[4][0].Copy()

	shtt, _ = render.GetSheet(filepath.Join("8x8", "genericfoes.png"))
	enemySheet1 := shtt.ToSprites()
	shtt, _ = render.GetSheet(filepath.Join("16x16", "specialfoes.png"))
	enemySheet2 := shtt.ToSprites()
	images["meleeFoe"] = render.NewReverting(enemySheet1[0][0].Copy())
	images["rangedFoe"] = render.NewReverting(enemySheet1[0][1].Copy())
	images["stomachFoe"] = render.NewReverting(render.NewSwitch("base", map[string]render.Modifiable{
		"base": enemySheet2[0][0].Copy(),
		"attacking": render.NewSequence(2,
			enemySheet2[0][0].Copy(),
			enemySheet2[1][0].Copy(),
			enemySheet2[2][0].Copy(),
			enemySheet2[2][0].Copy(),
			enemySheet2[2][0].Copy(),
			enemySheet2[1][0].Copy(),
			enemySheet2[1][0].Copy()),
	}))
	sh, err := render.LoadSheet("images", filepath.Join("16x16", "specialfoes.png"), 16, 16, 0)
	if err != nil {
		dlog.Error(err)
	}
	an, err := render.NewSheetSequence(sh, 2, 0, 1, 0, 1, 0, 1, 1, 1, 2, 1, 1, 1)
	if err != nil {
		dlog.Error(err)
	}
	images["heartFoe"] = render.NewReverting(an)
	images["liverFoe"] = render.NewReverting(render.NewSwitch("base", map[string]render.Modifiable{
		"base": enemySheet2[0][2].Copy(),
		"attacking": render.NewSequence(1,
			enemySheet2[0][2].Copy(),
			enemySheet2[1][2].Copy(),
			enemySheet2[2][2].Copy(),
			enemySheet2[2][2].Copy(),
			enemySheet2[2][2].Copy(),
			enemySheet2[1][2].Copy()),
	}))
	images["lungFoe"] = render.NewReverting(enemySheet2[0][4].Copy())
	images["brainFoe"] = render.NewReverting(
		render.NewSwitch("base", map[string]render.Modifiable{
			"base": enemySheet2[0][3].Copy(),
			"teleLeft": render.NewSequence(2,
				enemySheet2[0][3].Copy(),
				enemySheet2[1][3].Copy(),
				enemySheet2[2][3].Copy(),
			),
			"teleRight": render.NewSequence(2,
				enemySheet2[0][3].Copy().Modify(mod.Rotate(180)),
				enemySheet2[1][3].Copy().Modify(mod.Rotate(180)),
				enemySheet2[2][3].Copy().Modify(mod.Rotate(180)),
			),
			"teleBack": render.NewSequence(2,
				enemySheet2[0][3].Copy().Modify(mod.Rotate(-90)),
				enemySheet2[1][3].Copy().Modify(mod.Rotate(-90)),
				enemySheet2[2][3].Copy().Modify(mod.Rotate(-90)),
			),
			"teleForward": render.NewSequence(2,
				enemySheet2[0][3].Copy().Modify(mod.Rotate(90)),
				enemySheet2[1][3].Copy().Modify(mod.Rotate(90)),
				enemySheet2[2][3].Copy().Modify(mod.Rotate(90)),
			),
		}))

	images["whip"], _ = render.GetSprite(filepath.Join("raw", "whip.png"))
	images["sword"], _ = render.GetSprite(filepath.Join("raw", "sword.png"))
	images["net"], _ = render.GetSprite(filepath.Join("raw", "net.png"))
	images["spear"], _ = render.GetSprite(filepath.Join("raw", "spear.png"))

	for k, v := range images {
		images[k] = v.Modify(mod.TrimColor(color.RGBA{1, 1, 1, 1}))
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
			sp, _ := render.GetSprite(filepath.Join("raw", s))
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
