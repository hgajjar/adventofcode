package day2

import (
	"adventofcode/day2/game"
	"adventofcode/day2/strategy"
	_ "embed"
	"fmt"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(game.Play(strategy.NewGuide(sampleInput)))
	fmt.Println(game.Play(strategy.NewGuide(actualInput)))
}
