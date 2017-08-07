package main

import (
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
	oak.Init("menu")
}
