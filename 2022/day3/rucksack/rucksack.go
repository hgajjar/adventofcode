package rucksack

import (
	"adventofcode/day3/item"
	"github.com/samber/lo"
	"strings"
)

type Rucksack []item.Item

func Groups(input string) [][]Rucksack {
	var rucksacks []Rucksack
	for _, items := range strings.Split(input, "\n") {
		rucksacks = append(rucksacks, Rucksack(items))
	}

	return lo.Chunk[Rucksack](rucksacks, 3)
}
