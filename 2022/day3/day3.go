package day3

import (
	"adventofcode/day3/item"
	"adventofcode/day3/rucksack"
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(countTotalPrioritiesOfAllRucksacks(sampleInput))
	fmt.Println(countTotalPrioritiesOfAllRucksacks(actualInput))
}

func countTotalPrioritiesOfAllRucksacks(input string) int {
	var sharedItemsSumByRucksack []item.Priority
	for _, items := range strings.Split(input, "\n") {
		r := rucksack.New(items)
		sharedItems := lo.Uniq(lo.Intersect[item.Item](r.Compartments[0], r.Compartments[1]))

		sum := lo.SumBy(sharedItems, func(i item.Item) item.Priority {
			return i.Priority()
		})

		sharedItemsSumByRucksack = append(sharedItemsSumByRucksack, sum)
	}

	return int(lo.Sum(sharedItemsSumByRucksack))
}
