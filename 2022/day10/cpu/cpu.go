package cpu

import (
	"adventofcode/day10/instruction"
	"github.com/samber/lo"
)

var interestingCycles = []int{20, 60, 100, 140, 180, 220}

type Cpu struct {
	RegisterX int
}

func New() Cpu {
	return Cpu{RegisterX: 1}
}

func (c Cpu) Execute(instructions []instruction.Instruction) (totalSignalStrength int) {
	cycle := 1

	for _, inst := range instructions {
		for i := 1; i <= inst.Cycles; i++ {
			cycle++

			if i == inst.Cycles {
				c.RegisterX += inst.Value
			}

			if lo.Contains(interestingCycles, cycle) {
				totalSignalStrength += cycle * c.RegisterX
				//fmt.Printf("cycle: %d, X: %d\n", cycle, c.RegisterX)
			}
			//fmt.Printf("cycle: %d, X: %d\n", cycle, c.RegisterX)
		}
	}

	return
}
