package game

import (
	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/font"
	"github.com/oakmound/oak/audio"
	"github.com/oakmound/oak/dlog"
)

var (
	LoudSFX = font.New()
	SoftSFX = font.New()
	Music   = font.New()
	audios  = map[string]*audio.Audio{}
)

func InitAudio() {
	files := map[string]*font.Font{
		"BoomerAttack": SoftSFX,
		"ClearOrgan":   LoudSFX,
		"DasherAttack": SoftSFX,
		"FailOrgan":    LoudSFX,
		"Footstep":     SoftSFX,
		"InfectHeavy":  SoftSFX,
		"InfectLight":  SoftSFX,
		"InfectMedium": SoftSFX,
		"NetHeavy":     SoftSFX,
		"NetLight":     SoftSFX,
		"RangedAttack": SoftSFX,
		"Shrink":       LoudSFX,
		"Shrink2":      LoudSFX,
		"Shrink3":      LoudSFX,
		"SpearHeavy":   SoftSFX,
		"SpearLight":   SoftSFX,
		"SwordHeavy":   SoftSFX,
		"SwordLight":   SoftSFX,
		"Vacuum":       SoftSFX,
		"WhipHeavy":    SoftSFX,
		"WhipLight":    SoftSFX,
		"WizardAttack": SoftSFX,
	}
	for s, f := range files {
		a, err := audio.Get(s + ".wav")
		if err != nil {
			dlog.Error(err)
			return
		}
		audios[s] = audio.New(f, a)
	}
}

func PlayAt(s string, x, y float64) {
	aud, err := audios[s].Copy()
	if err != nil {
		dlog.Error(err)
		return
	}
	a := aud.(*audio.Audio)
	a.X = &x
	a.Y = &y
	a.Play()
}

func UpdateEars(player *Player) {

	ears := audio.NewEars(player.Xp(), player.Yp(), 40, 70)

	LoudSFX.Filter(filter.Volume(.5), audio.PosFilter(ears))
	SoftSFX.Filter(filter.Volume(.25), audio.PosFilter(ears))
	Music.Filter(filter.Volume(.5))
}
