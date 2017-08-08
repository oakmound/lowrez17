package game

import "time"

type Weapon map[string]*Action

var (
	Sword = Weapon(map[string]*Action{
		"left":  NewAction(SwordLeft(Ally), 100*time.Millisecond),
		"right": NewAction(SwordRight(Ally), 100*time.Millisecond),
		"space": NewAction(SwordDash(Ally), 500*time.Millisecond),
	})
	Whip = Weapon(map[string]*Action{
		"left":  NewAction(WhipLeft(Ally), 200*time.Millisecond),
		"right": NewAction(WhipRight(Ally), 200*time.Millisecond),
		"space": NewAction(WhipTwirl(Ally), 700*time.Millisecond),
	})
)
