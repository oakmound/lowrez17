package game

import "time"

type Weapon struct {
	left, right, space *Action
}

var (
	Sword = Weapon{
		left:  NewAction(SwordLeft(Ally), 250*time.Millisecond),
		right: NewAction(SwordRight(Ally), 250*time.Millisecond),
		space: NewAction(SwordDash(Ally), 5*time.Second),
	}
	Whip = Weapon{
		left:  NewAction(WhipLeft(Ally), 300*time.Millisecond),
		right: NewAction(WhipRight(Ally), 300*time.Millisecond),
		space: NewAction(WhipTwirl(Ally), 8*time.Second),
	}
	Spear = Weapon{
		left:  NewAction(SpearJab(Ally), 400*time.Millisecond),
		right: NewAction(SpearThrust(Ally), 900*time.Millisecond),
		space: NewAction(SpearDash(Ally), 3*time.Second),
	}
	Net = Weapon{
		left:  NewAction(NetLeft(Ally), 100*time.Millisecond),
		right: NewAction(NetRight(Ally), 100*time.Millisecond),
		space: NewAction(NetTwirl(Ally), 1*time.Second),
	}
)
