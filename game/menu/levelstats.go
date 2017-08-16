package menu

type LevelStats struct {
	Score   int64   `json:"score"`
	Time    int64   `json:"time"`
	Cleared float64 `json:"cleared"`
	Level   int     `json:"level"`
}

func (ls *LevelStats) CalculateScore() {
	if ls.Cleared == 0 {
		return
	}
	// Fairly basic but sufficient, lower score is better
	ls.Score = int64(float64(ls.Time) * 1 / ls.Cleared)
}

type LevelStorage struct {
	Stats [5]LevelStats `json:"stats"`
}
