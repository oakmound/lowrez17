package game

import (
	"github.com/oakmound/lowrez17/game/sfx"
	"github.com/oakmound/oak/audio"
	"github.com/oakmound/oak/dlog"
)

func PlayAt(s string, x, y float64) {
	aud, err := sfx.Audios[s].Copy()
	if err != nil {
		dlog.Error(err)
		return
	}
	a := aud.(*audio.Audio)
	a.X = &x
	a.Y = &y
	a.Play()
}
