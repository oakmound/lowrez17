package game

func GetBody(level string) *Body {
	if level == "endurance" {
		// randomize
		return nil
	}
	return levelBodies[level]
}
