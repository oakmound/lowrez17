package game

type Direction int

const (
	Forward Direction = iota
	Backward
	Left
	Right
	Wait
)

func (dir Direction) Move() *Action {
	switch dir {
	case Forward:
		return MoveForward
	case Backward:
		return MoveBackward
	case Left:
		return MoveLeft
	case Right:
		return MoveRight
	case Wait:
		return NewAction(func(*Entity) {}, 0)
	}
	return nil
}

func Move(dir Direction, n int) []*Action {
	out := make([]*Action, n)
	for i := range out {
		out[i] = dir.Move()
	}
	return out
}
