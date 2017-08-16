package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/oakmound/lowrez17/game"
	"github.com/oakmound/lowrez17/game/menu"
	"github.com/oakmound/oak"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}
	oak.AddCommand("stopProf", func([]string) {
		pprof.StopCPUProfile()
	})
	err := oak.LoadConf("oak.config")
	if err != nil {
		log.Fatal(err)
	}
	oak.LoadConf("oak.config")
	oak.AddScene("menu",
		menu.StartScene,
		menu.LoopScene,
		menu.EndScene)
	oak.AddScene("level",
		game.LevelInit,
		game.LevelLoop,
		game.LevelEnd)
	// grayScale := []color.Color{}
	// for i := uint8(0); i < 127; i++ {
	// 	grayScale = append(grayScale, color.RGBA{i * 2, i * 2, i * 2, 255})
	// }
	// oak.SetPalette(grayScale)
	oak.Init("menu")
}
