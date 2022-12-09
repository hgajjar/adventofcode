package crane

import (
	"adventofcode/day5/crate"
)

func Start(stacks []crate.Stack, procedure Procedure) []crate.Stack {
	for _, p := range procedure {
		for i := 0; i < p.move; i++ {
			pop := stacks[p.from-1].Pop()
			stacks[p.to-1].Push(pop)
		}
	}

	return stacks
}
