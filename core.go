package main

import (
	"image/color"

	"github.com/oakmound/lowrez17/game"
	"github.com/oakmound/lowrez17/game/menu"
	"github.com/oakmound/oak"
)

func main() {
	oak.LoadConf("oak.config")
	oak.AddScene("menu",
		menu.StartScene,
		menu.LoopScene,
		menu.EndScene)
	oak.AddScene("level",
		game.LevelInit,
		game.LevelLoop,
		game.LevelEnd)
	grayScale := []color.Color{}
	for i := uint8(0); i < 127; i++ {
		grayScale = append(grayScale, color.RGBA{i * 2, i * 2, i * 2, 255})
	}
	oak.SetPalette(grayScale)
	oak.Init("menu")
}
