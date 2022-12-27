package day10

import (
	"adventofcode/day10/cpu"
	"adventofcode/day10/instruction"
	_ "embed"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	calculateTotalSignalStrength(sampleInput)
	calculateTotalSignalStrength(actualInput)
}

func calculateTotalSignalStrength(input string) {
	instructions := lo.Map(strings.Split(input, "\n"), func(item string, _ int) instruction.Instruction {
		parts := strings.Split(item, " ")
		var val int
		if len(parts) > 1 {
			val, _ = strconv.Atoi(parts[1])
		}
		return instruction.New(parts[0], val)
	})

	cpu.New().Execute(instructions)
}
