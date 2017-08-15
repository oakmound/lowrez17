package game

import (
	"fmt"
	"time"

	"github.com/200sc/go-dist/intrange"
	"github.com/oakmound/lowrez17/game/sfx"
)

type Wave struct {
	EnemyDist
	Difficulty float64
	Timelimit  time.Duration
}

var (
	enemyCh    = make(chan bool)
	waveExitCh = make(chan bool)
)

func handleWaves(waves []Wave, tiles [][]Tile, typ OrganType) {
	wrange := intrange.NewLinear(0, len(tiles)-1)
	hrange := intrange.NewLinear(0, len(tiles[0])-1)
	i := 0
	enemiesLeft := 0
	for {
		es := waves[i].Poll()
		enemiesLeft += len(es)
		fmt.Println("wave", i, "enemies left", enemiesLeft)
		for _, t := range es {
			x := wrange.Poll()
			y := hrange.Poll()
			for tiles[x][y] != Open {
				x = wrange.Poll()
				y = hrange.Poll()
			}
			enemies = append(enemies, enemyFns[t][typ](x, y, waves[i].Difficulty, false))
		}
	handleOuter:
		for {
			select {
			case <-enemyCh:
				enemiesLeft--
				fmt.Println("Enemies Left", enemiesLeft)
				if enemiesLeft <= 0 {
					break handleOuter
				}
			case <-waveExitCh:
				return
			case <-time.After(waves[i].Timelimit):
				break handleOuter
			}
		}
		<-time.After(2 * time.Second)
		i++
		if i >= len(waves) {
			break
		}
	}
	for enemiesLeft > 0 {
		select {
		case <-enemyCh:
			fmt.Println("Enemies Left", enemiesLeft)
			enemiesLeft--
		case <-waveExitCh:
			return
		}
	}
	// Beat the organ
	fmt.Println("Organ beat")
	sfx.Audios["ClearOrgan"].Play()
	<-time.After(2 * time.Second)
	CleanupActiveOrgan(true)
}
