package game

import "time"

type Action struct {
	Cooldown  time.Duration
	ReadyTime time.Time
	action    func(*Entity)
}

func (a *Action) Do(e HasE) {
	if time.Now().After(a.ReadyTime) {
		a.action(e.E())
		a.ReadyTime = time.Now().Add(a.Cooldown)
	}
}

func NewAction(a func(*Entity), cooldown time.Duration) *Action {
	return &Action{Cooldown: cooldown, action: a}
}

var (
	MoveForward  = NewAction((*Entity).moveForward, 0)
	MoveBackward = NewAction((*Entity).moveBack, 0)
	MoveRight    = NewAction((*Entity).moveRight, 0)
	MoveLeft     = NewAction((*Entity).moveLeft, 0)
)

func (a *Action) MustDo(e HasE) {
	a.action(e.E())
}
