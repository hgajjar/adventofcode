package crane

import (
	"adventofcode/day5/crate"
)

func Start(stacks []crate.Stack, procedure Procedure) []crate.Stack {
	for _, p := range procedure {
		poppedCrates := stacks[p.from-1].Pop(p.move)
		stacks[p.to-1].Push(poppedCrates)
	}

	return stacks
}
