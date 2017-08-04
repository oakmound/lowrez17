package game

import "time"

type Wave struct {
	EnemyDist
	Difficulty float64
	Timelimit  time.Duration
}
