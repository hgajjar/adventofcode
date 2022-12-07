package day3

import (
	"adventofcode/day3/item"
	"adventofcode/day3/rucksack"
	_ "embed"
	"fmt"
	"github.com/samber/lo"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(countTotalPriorities(sampleInput))
	fmt.Println(countTotalPriorities(actualInput))
}

func countTotalPriorities(input string) item.Priority {
	var total item.Priority
	for _, group := range rucksack.Groups(input) {
		sharedItems := lo.Uniq(lo.Intersect[item.Item](group[0], group[1]))
		sharedItems = lo.Uniq(lo.Intersect[item.Item](sharedItems, group[2]))

		for _, s := range sharedItems {
			total += s.Priority()
		}
	}

	return total
}
