package cpu

import (
	"adventofcode/day10/crt"
	"adventofcode/day10/instruction"
)

var interestingCycles = []int{20, 60, 100, 140, 180, 220}

type Cpu struct {
	RegisterX int
	Crt       crt.CRT
}

func New() Cpu {
	return Cpu{RegisterX: 1, Crt: crt.New()}
}

func (c Cpu) Execute(instructions []instruction.Instruction) {
	cycle := 1

	for _, inst := range instructions {
		for i := 1; i <= inst.Cycles; i++ {
			if i == inst.Cycles {
				c.RegisterX += inst.Value
			}

			c.Crt.DrawPixel(cycle, c.RegisterX)

			cycle++
		}
	}

	c.Crt.Print()
}
