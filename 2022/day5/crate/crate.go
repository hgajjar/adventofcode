package crate

import (
	"github.com/samber/lo"
	"strings"
)

type crate byte

type Stack []crate

func (s *Stack) Push(crates []crate) {
	for _, c := range crates {
		*s = append(*s, c)
	}
}

func (s *Stack) Pop(n int) []crate {
	offset := len(*s) - n
	ss := *s
	popped := ss[offset:]
	*s = ss[:offset]

	return popped
}

func CreateStacks(input string) []Stack {
	var stacks []Stack

	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		i := 1
		j := 0
		for {
			if i > len(line) {
				break
			}
			if len(stacks) <= j {
				stacks = append(stacks, Stack{})
			}
			crateChar := line[i]
			if len(strings.TrimSpace(string(crateChar))) > 0 {
				stacks[j] = append(stacks[j], crate(crateChar))
			}
			i += 4
			j++
		}
	}

	var reverseStacks []Stack
	for _, s := range stacks {
		reverseStacks = append(reverseStacks, lo.Reverse(s))
	}

	return reverseStacks
}

func GetTopCrateFromEachStack(stacks []Stack) []crate {
	var topCrates []crate
	for _, s := range stacks {
		topCrates = append(topCrates, s[len(s)-1])
	}

	return topCrates
}
