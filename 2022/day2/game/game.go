package game

import (
	"adventofcode/day2/shape"
	"adventofcode/day2/strategy"
)

const (
	winScore  = 6
	drawScore = 3
	loseScore = 0
)

func Play(strategyGuide strategy.Guide) int {
	var totalScore int
	for _, s := range strategyGuide {
		totalScore += playRound(s)
	}

	return totalScore
}

func playRound(s strategy.Strategy) int {
	if s.Draw() {
		var chosenShape shape.Shape

		if s.OpponentShape.IsRock() {
			chosenShape = shape.NewRock()
		} else if s.OpponentShape.IsPaper() {
			chosenShape = shape.NewPaper()
		} else if s.OpponentShape.IsScissors() {
			chosenShape = shape.NewScissors()
		}

		return drawScore + chosenShape.Score()
	} else if s.Lose() {
		var chosenShape shape.Shape

		if s.OpponentShape.IsRock() {
			chosenShape = shape.NewScissors()
		} else if s.OpponentShape.IsPaper() {
			chosenShape = shape.NewRock()
		} else if s.OpponentShape.IsScissors() {
			chosenShape = shape.NewPaper()
		}

		return loseScore + chosenShape.Score()
	} else {
		var chosenShape shape.Shape

		if s.OpponentShape.IsScissors() {
			chosenShape = shape.NewRock()
		} else if s.OpponentShape.IsRock() {
			chosenShape = shape.NewPaper()
		} else if s.OpponentShape.IsPaper() {
			chosenShape = shape.NewScissors()
		}

		return winScore + chosenShape.Score()
	}
}
