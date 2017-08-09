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

func (dir Direction) Teleport(distance float64) *Action {
	switch dir {
	case Forward:
		return TeleportForward(distance)
	case Backward:
		return TeleportBackward(distance)
	case Left:
		return TeleportLeft(distance)
	case Right:
		return TeleportRight(distance)
	}
	return nil
}

func Teleport(dir Direction, distance float64) []*Action {
	// Todo: preparatory / ending actions
	return []*Action{
		dir.Teleport(distance),
	}
}

