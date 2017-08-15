package menu

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
)

var (
	currentLevel  = 0
	sceneContinue = true
	nextScene     = "menu"
	levelData     = ""
	stats         *LevelStorage
)

func StartScene(_ string, levelData interface{}) {
	initLetters()
	if stats == nil {
		stats = &LevelStorage{}
		f, err := os.Open("save.json")
		if err == nil {
			defer f.Close()
			err = json.NewDecoder(f).Decode(stats)
			if err != nil {
				dlog.Error(err)
			}
		} else {
			dlog.Error(err)
		}
	}
	if levelData != nil {
		sData := levelData.(LevelStats)
		sData.CalculateScore()
		oldScore := stats.Stats[sData.Level].Score
		if oldScore <= 0 || oldScore > sData.Score {
			stats.Stats[sData.Level] = sData
		}
		currentLevel = (sData.Level + 1) % 6
	} else {
		// Set current level based on which levels have been completed already
		if stats.Stats[4].Score > 0 {
			currentLevel = 0 // Endurance?
		} else if stats.Stats[3].Score > 0 {
			currentLevel = 4
		} else if stats.Stats[2].Score > 0 {
			currentLevel = 3
		} else if stats.Stats[1].Score > 0 {
			currentLevel = 2
		} else if stats.Stats[0].Score > 0 {
			currentLevel = 1
		}
	}
	//currentLevel = 1
	jsonData, err := json.Marshal(stats)
	if err != nil {
		dlog.Error(err)
	} else {
		ioutil.WriteFile("save.json", jsonData, os.ModeType)
	}
	p := NewPlayer()
	p.SetPos(53, 34)
	nextScene = "menu"
	// Create blocking zones
	// walls
	collision.Add(collision.NewLabeledSpace(-2, -2, 2, 66, blocking))
	collision.Add(collision.NewLabeledSpace(-2, -2, 28, 19, blocking))
	collision.Add(collision.NewLabeledSpace(38, -2, 28, 19, blocking))
	collision.Add(collision.NewLabeledSpace(-2, 64, 64, 2, blocking))
	collision.Add(collision.NewLabeledSpace(64, -66, 2, 128, blocking))
	collision.Add(collision.NewLabeledSpace(64, 0, 64, 35, blocking))
	collision.Add(collision.NewLabeledSpace(126, 0, 2, 64, blocking))
	collision.Add(collision.NewLabeledSpace(64, 60, 30, 2, blocking))
	collision.Add(collision.NewLabeledSpace(102, 60, 30, 2, blocking))
	// not walls
	collision.Add(collision.NewLabeledSpace(0, 17, 14, 22, blocking))
	collision.Add(collision.NewLabeledSpace(43, 17, 20, 17, blocking))
	// Create zones that lead to levels, menu
	collision.Add(collision.NewLabeledSpace(50, 35, 10, 2, wasd))
	// Next level zone
	collision.Add(collision.NewLabeledSpace(19, 35, 27, 10, nextLevel))
	collision.Add(collision.NewLabeledSpace(21, 37, 23, 6, blocking))

	// Add level specific zones if those levels have been cleared before
	// todo: also add coloration indicating the morgue slots are open
	if stats.Stats[0].Score > 0 {
		collision.Add(collision.NewLabeledSpace(67, 40, 8, 2, level1))
	}
	if stats.Stats[1].Score > 0 {
		collision.Add(collision.NewLabeledSpace(79, 40, 8, 2, level2))
	}
	if stats.Stats[2].Score > 0 {
		collision.Add(collision.NewLabeledSpace(91, 40, 8, 2, level3))
	}
	if stats.Stats[3].Score > 0 {
		collision.Add(collision.NewLabeledSpace(103, 40, 8, 2, level4))
	}
	if stats.Stats[4].Score > 0 {
		collision.Add(collision.NewLabeledSpace(115, 40, 8, 2, level5))
	}
	// door to morgue
	collision.Add(collision.NewLabeledSpace(26, 17, 12, 2, door))
	// door back
	collision.Add(collision.NewLabeledSpace(26+64, 60, 10, 2, doorBack))
	// Background
	render.Draw(render.LoadSprite(filepath.Join("raw", "toplayer.png")), backgroundLayer)
	morgue := render.LoadSprite(filepath.Join("raw", "morgue.png"))
	morgue.SetPos(64, 0)
	render.Draw(morgue, backgroundLayer)
}

func LoopScene() bool {
	return sceneContinue
}

func EndScene() (string, *oak.SceneResult) {
	sceneContinue = true
	return nextScene, &oak.SceneResult{
		NextSceneInput: levelData,
		Transition:     oak.TransitionZoom(.51, .67, 50, .009),
	}
}

var letters = map[rune]render.Modifiable{}

func initLetters() {
	frameRate := 1.5
	sh := render.GetSheet(filepath.Join("5x7", "letters.png"))
	runeMap := map[rune]int{
		'e': 0,
		'w': 1,
		'a': 2,
		's': 3,
		'd': 4,
	}
	for r, col := range runeMap {
		letters[r] = render.NewSequence([]render.Modifiable{sh[0][col], sh[1][col]}, frameRate)
	}
}

type Player struct {
	entities.Reactive
	collision.Phase
	stop       bool
	interactFn func()
	interactR  render.Renderable
}

func (p *Player) Init() event.CID {
	p.CID = event.NextID(p)
	return p.CID
}

func NewPlayer() *Player {
	p := new(Player)
	sh := render.GetSheet(filepath.Join("4x16", "topplayer.png"))
	p.Reactive = entities.NewReactive(5, 5, 4, 16, render.NewCompound(
		"forward",
		map[string]render.Modifiable{
			"forward": sh[0][0].Copy(),
			"right":   sh[1][0].Copy(),
			"left":    sh[1][0].Copy().Modify(render.FlipX),
			"back":    sh[2][0].Copy(),
		}), p.Init())
	collision.Add(p.RSpace.Space)
	render.Draw(p.R, entityLayer)
	p.RSpace.Add(blocking, playerStop)
	collision.PhaseCollision(p.RSpace.Space)
	p.Bind(triggerInteractive, "CollisionStart")
	p.Bind(unbindInteractive, "CollisionStop")
	p.Bind(playerWalk, "EnterFrame")
	p.Bind(playerInteract, "KeyUpE")
	return p
}

const (
	blocking collision.Label = iota
	nextLevel
	// morgue labels
	level1
	level2
	level3
	level4
	level5
	endurance
	// other places
	settings
	wasd
	door
	doorBack
	/// ...
)

func playerStop(s1, s2 *collision.Space) {
	p := s1.CID.E().(*Player)
	p.stop = true
}

func triggerInteractive(id int, label interface{}) int {
	p := event.CID(id).E().(*Player)
	switch label.(collision.Label) {
	case level1:
		levelData = "level1"
		setLevelInteracts(p)
	case level2:
		levelData = "level2"
		setLevelInteracts(p)
	case level3:
		levelData = "level3"
		setLevelInteracts(p)
	case level4:
		levelData = "level4"
		setLevelInteracts(p)
	case level5:
		levelData = "level5"
		setLevelInteracts(p)
	case endurance:
		// Todo: this isn't placed, could make it "level6"
		levelData = "endurance"
	case nextLevel:
		levelData = "level" + strconv.Itoa((currentLevel+1)%6)
		// Todo: don't try to go to level6, reset to 0
		setLevelInteracts(p)
	case wasd:
		w, a, s, d := letters['w'], letters['a'], letters['s'], letters['d']
		w.SetPos(0, -7)
		a.SetPos(-5, 0)
		d.SetPos(5, 0)
		p.interactR = render.NewComposite([]render.Modifiable{w, a, s, d})
		p.interactR.SetPos(p.X()-1, p.Y()-8)
		render.Draw(p.interactR, uiLayer)
	case door:
		p.SetPos(96, 40)
		oak.SetScreen(64, 0)
	case doorBack:
		p.SetPos(32, 20)
		oak.SetScreen(0, 0)
	}
	return 0
}

func setLevelInteracts(p *Player) {
	var err error
	p.interactR = letters['e']
	// I tried using attachment here, it was bugged?
	p.interactR.SetPos(p.X()-1, p.Y()-8)
	if err != nil {
		dlog.Error(err)
	}
	render.Draw(p.interactR, uiLayer)
	p.interactFn = func() {
		nextScene = "level"
		sceneContinue = false
	}
}

func unbindInteractive(id int, label interface{}) int {
	p := event.CID(id).E().(*Player)
	if label.(collision.Label) != blocking {
		p.Vector = p.Vector.Detach()
		p.interactFn = nil
		p.interactR.UnDraw()
		levelData = ""
	}
	return 0
}

func playerInteract(id int, nothing interface{}) int {
	p := event.CID(id).E().(*Player)
	if p.interactFn != nil {
		p.interactFn()
	}
	return 0
}

func playerWalk(id int, nothing interface{}) int {
	p := event.CID(id).E().(*Player)
	shiftX := 0.0
	shiftY := 0.0
	if oak.IsDown("W") {
		p.R.(*render.Compound).Set("back")
		shiftY--
	}
	if oak.IsDown("S") {
		p.R.(*render.Compound).Set("forward")
		shiftY++
	}
	if oak.IsDown("A") {
		p.R.(*render.Compound).Set("left")
		shiftX--
	}
	if oak.IsDown("D") {
		p.R.(*render.Compound).Set("right")
		shiftX++
	}
	if p.interactR != nil {
		p.interactR.SetPos(p.X()-1, p.Y()-8)
	}
	collision.ShiftSpace(shiftX, shiftY, p.RSpace.Space)
	<-p.RSpace.CallOnHits()
	// If near interactable, show potential interaction
	// Stop if hit something
	collision.ShiftSpace(-shiftX, -shiftY, p.RSpace.Space)
	if p.stop {
		p.stop = false
	} else {
		p.ShiftPos(shiftX, shiftY)
	}
	return 0
}
