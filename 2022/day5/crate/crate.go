package crate

import (
	"github.com/samber/lo"
	"strings"
)

type crate byte

type Stack []crate

func (s *Stack) Push(c crate) {
	*s = append(*s, c)
}

func (s *Stack) Pop() crate {
	n := len(*s) - 1
	ss := *s
	c := ss[n]
	*s = ss[:n]

	return c
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
		topCrates = append(topCrates, s.Pop())
	}

	return topCrates
}
