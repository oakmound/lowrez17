package sfx

import (
	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/font"
	"github.com/oakmound/oak/audio"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/physics"
)

var (
	LoudSFX = font.New()
	SoftSFX = font.New()
	Music   = font.New()
	Audios  = map[string]*audio.Audio{}
	inited  bool
)

func InitAudio() {
	if inited {
		return
	}
	inited = true
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
		"SummonAttack": SoftSFX,
		"SwordReady":   SoftSFX,
		"WhipReady":    SoftSFX,
		"NetReady":     SoftSFX,
		"SpearReady":   SoftSFX,
	}
	for s, f := range files {
		a, err := audio.Get(s + ".wav")
		if err != nil {
			dlog.Error(err)
			return
		}
		Audios[s] = audio.New(f, a)
	}
}

func UpdateEars(player physics.Attachable) {

	v := player.Vec()
	ears := audio.NewEars(v.Xp(), v.Yp(), 40, 70)

	LoudSFX.Filter(filter.Volume(.5), audio.PosFilter(ears))
	SoftSFX.Filter(filter.Volume(.25), audio.PosFilter(ears))
	Music.Filter(filter.Volume(.5))
}
