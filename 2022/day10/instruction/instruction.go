package instruction

import "fmt"

const (
	Noop string = "noop"
	Addx string = "addx"
)

type Instruction struct {
	Value  int
	Cycles int
}

func New(name string, value int) Instruction {
	switch name {
	case Noop:
		return Instruction{Value: value, Cycles: 1}
	case Addx:
		return Instruction{Value: value, Cycles: 2}
	default:
		panic(fmt.Sprintf("invalid instruction name '%s'", name))
	}
}
