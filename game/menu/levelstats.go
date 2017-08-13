package menu

type LevelStats struct {
	Score   int     `json:"score"`
	Time    int     `json:"time"`
	Cleared float64 `json:"cleared"`
	Level   int     `json:"level"`
}
