package day5

import (
	"adventofcode/day5/crane"
	"adventofcode/day5/crate"
	_ "embed"
	"fmt"
	"regexp"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(rearrange(sampleInput))
	fmt.Println(rearrange(actualInput))
}

func rearrange(input string) string {
	inputParts := regexp.MustCompile(`\n\s*\n`).Split(input, -1)
	stacks := crate.CreateStacks(inputParts[0])
	procedure := crane.CreateProcedure(inputParts[1])
	processedStacks := crane.Start(stacks, procedure)

	return string(crate.GetTopCrateFromEachStack(processedStacks))
}
