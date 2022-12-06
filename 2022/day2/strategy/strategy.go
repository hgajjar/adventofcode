package strategy

import (
	"adventofcode/day2/shape"
	"fmt"
	"strings"
)

type Strategy struct {
	OpponentShape shape.Shape
	OwnShape      shape.Shape
}

type Guide []Strategy

func ParseGuide(input string) Guide {
	var guide Guide
	for _, i := range strings.Split(input, "\n") {
		letters := strings.Split(i, " ")
		if len(letters) != 2 {
			panic(fmt.Sprintf("invalid input found: '%s'", i))
		}

		s := Strategy{
			OpponentShape: mapStrategyLetterToShape(letters[0]),
			OwnShape:      mapStrategyLetterToShape(letters[1]),
		}

		guide = append(guide, s)
	}

	return guide
}

func mapStrategyLetterToShape(letter string) shape.Shape {
	switch letter {
	case "A", "X":
		return shape.NewRock()
	case "B", "Y":
		return shape.NewPaper()
	case "C", "Z":
		return shape.NewScissors()
	}

	panic(fmt.Sprintf("invalid letter found in the Strategy guide: '%s'", letter))
}
