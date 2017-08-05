package game

import (
	"github.com/oakmound/oak/collision"
	"time"
)

type Weapon map[string]*Action

var (
	Sword = Weapon(map[string]*Action{
		"left":  NewAction(SwordLeft(Ally), 1*time.Second),
		"right": NewAction(SwordRight(Ally), 1*time.Second),
		"dash":  NewAction(SwordDash(Ally), 1*time.Second),
	})
)

func SwordLeft(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		pos := p.CenterPos().Add(p.Dir.Copy().Rotate(-55).Scale(7))
		basePos := pos.Copy()
		for j := -55.0; j <= 45.0; j += 10.0 {
			yDelta := p.Dir.Copy().Rotate(j).Scale(4)
			pos = basePos.Copy()
			for i := 0; i < 4; i++ {
				NewHurtBox(pos.X(), pos.Y(), 3, 3, 50*time.Millisecond, label)
				pos.Add(yDelta)
			}
		}
	}
}
func SwordRight(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		pos := p.CenterPos().Add(p.Dir.Copy().Rotate(55).Scale(7))
		basePos := pos.Copy()
		for j := 55.0; j >= -45.0; j -= 10.0 {
			yDelta := p.Dir.Copy().Rotate(j).Scale(4)
			pos = basePos.Copy()
			for i := 0; i < 4; i++ {
				NewHurtBox(pos.X(), pos.Y(), 3, 3, 50*time.Millisecond, label)
				pos.Add(yDelta)
			}
		}
	}
}

func SwordDash(label collision.Label) func(p *Entity) {
	return func(p *Entity) {
		p.Delta.Add(p.Dir.Copy().Scale(24 * p.Speed.Y()))
		delta := p.Dir.Copy().Scale(3)
		perpendicular := delta.Copy().Rotate(90)
		pos := p.CenterPos().Add(delta, perpendicular, perpendicular)
		perpendicular.Scale(-1)
		basePos := pos.Copy()
		for i := 0; i < 3; i++ {
			pos = basePos.Add(perpendicular).Copy()
			for j := 0; j < 12; j++ {
				pos.Add(delta)
				NewHurtBox(pos.X(), pos.Y(), 3, 3, 50*time.Millisecond, label)
			}
		}
	}
}
