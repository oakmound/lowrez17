package game

import (
	"time"

	"github.com/oakmound/oak/collision"
)

type Weapon struct {
	left, right, space *Action
}

var (
	Sword = Weapon{
		left:  NewAction(SwordLeft(Ally), 100*time.Millisecond),
		right: NewAction(SwordRight(Ally), 100*time.Millisecond),
		space: NewAction(SwordDash(Ally), 500*time.Millisecond),
	}
	Whip = Weapon{
		left:  NewAction(WhipLeft(Ally), 200*time.Millisecond),
		right: NewAction(WhipRight(Ally), 200*time.Millisecond),
		space: NewAction(WhipTwirl(Ally), 700*time.Millisecond),
	}
	Spear = Weapon{
		left:  NewAction(SpearJab(Ally), 150*time.Millisecond),
		right: NewAction(SpearJab(Ally), 150*time.Millisecond),
		space: NewAction(SpearDash(Ally), 1000*time.Millisecond),
	}
	Net = Weapon{
		left:  NewAction(NetLeft(Ally), 50*time.Millisecond),
		right: NewAction(NetRight(Ally), 50*time.Millisecond),
		space: NewAction(NetTwirl(Ally), 400*time.Millisecond),
	}
)

func SpearJab(label collision.Label) func(*Entity) {
	return func(p *Entity) {

	}
}

func SpearDash(label collision.Label) func(*Entity) {
	return func(p *Entity) {

	}
}

func NetLeft(label collision.Label) func(*Entity) {
	return func(p *Entity) {

	}
}

func NetRight(label collision.Label) func(*Entity) {
	return func(p *Entity) {

	}
}

func NetTwirl(label collision.Label) func(*Entity) {
	return func(p *Entity) {

	}
}
