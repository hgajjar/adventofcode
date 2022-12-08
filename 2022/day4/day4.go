package day4

import (
	"adventofcode/day4/elf"
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(countTotalOverlappingPairs(sampleInput))
	fmt.Println(countTotalOverlappingPairs(actualInput))
}

func countTotalOverlappingPairs(input string) int {
	var total int
	for _, pair := range strings.Split(input, "\n") {
		assignments := strings.Split(pair, ",")
		elf1 := elf.NewWithAssignmentRange(assignments[0])
		elf2 := elf.NewWithAssignmentRange(assignments[1])

		if lo.Some(elf1.Sections, elf2.Sections) || lo.Some(elf2.Sections, elf1.Sections) {
			total++
		}
	}

	return total
}
