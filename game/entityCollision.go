package game

import "github.com/oakmound/oak/collision"

const (
	// Shift to avoid same values as tiles
	_                        = iota
	Ally     collision.Label = iota << 4 // 00001 -> 010000
	Opposing                 = iota << 4 // 00010 -> 100000
	Stun                     = iota << 4 // 00011 -> 110000
)
