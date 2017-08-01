package game

import "github.com/oakmound/oak"

var (
	envFriction = 0.7
)

func LevelInit(prevScene string, inData interface{}) {
	NewEntity(20, 20)
}

func LevelLoop() bool {
	return true
}

func LevelEnd() (nextScene string, result *oak.SceneResult) {
	return "firstScene", nil
}
