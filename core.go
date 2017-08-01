package main

import (
	"github.com/oakmound/lowrez17/game"
	"github.com/oakmound/oak"
)

func main() {
	oak.LoadConf("oak.config")
	oak.AddScene("level",
		game.LevelInit,
		game.LevelLoop,
		game.LevelEnd)
	oak.Init("level")
}
