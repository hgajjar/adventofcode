package game

import "adventofcode/day2/strategy"

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
	if (s.OpponentShape.IsRock() && s.OwnShape.IsRock()) ||
		(s.OpponentShape.IsPaper() && s.OwnShape.IsPaper()) ||
		(s.OpponentShape.IsScissors() && s.OwnShape.IsScissors()) {
		return drawScore + s.OwnShape.Score()
	}

	if (s.OpponentShape.IsRock() && s.OwnShape.IsPaper()) ||
		(s.OpponentShape.IsPaper() && s.OwnShape.IsScissors()) ||
		(s.OpponentShape.IsScissors() && s.OwnShape.IsRock()) {
		return winScore + s.OwnShape.Score()
	}

	return loseScore + s.OwnShape.Score()
}
