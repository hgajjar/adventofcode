package strategy

import (
	"adventofcode/day2/shape"
	"fmt"
	"strings"
)

type Strategy struct {
	OpponentShape shape.Shape
	strategy      string
}

type Guide []Strategy

func NewGuide(input string) Guide {
	var guide Guide
	for _, i := range strings.Split(input, "\n") {
		letters := strings.Split(i, " ")
		if len(letters) != 2 {
			panic(fmt.Sprintf("invalid input found: '%s'", i))
		}

		s := Strategy{
			OpponentShape: mapStrategyLetterToShape(letters[0]),
			strategy:      letters[1],
		}

		guide = append(guide, s)
	}

	return guide
}

func (s Strategy) Lose() bool {
	return s.strategy == "X"
}

func (s Strategy) Draw() bool {
	return s.strategy == "Y"
}

func (s Strategy) Win() bool {
	return s.strategy == "Z"
}

func mapStrategyLetterToShape(letter string) shape.Shape {
	switch letter {
	case "A":
		return shape.NewRock()
	case "B":
		return shape.NewPaper()
	case "C":
		return shape.NewScissors()
	}

	panic(fmt.Sprintf("invalid letter found in the Strategy guide: '%s'", letter))
}
