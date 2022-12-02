package day1

import (
	"adventofcode/lib"
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(lib.SliceIntSum(getSortedTotalCaloriesByElves(sampleInput)[:3]))
	fmt.Println(lib.SliceIntSum(getSortedTotalCaloriesByElves(actualInput)[:3]))

}

func getSortedTotalCaloriesByElves(input string) []int {
	caloriesByElves := regexp.MustCompile(`\n\s*\n`).Split(input, -1)
	var totalCaloriesByElves []int
	for _, perElfCalories := range caloriesByElves {
		totalCaloriesByElves = append(totalCaloriesByElves, countTotalCaloriesOfElf(perElfCalories))
	}

	sort.Ints(totalCaloriesByElves)
	sort.Sort(sort.Reverse(sort.IntSlice(totalCaloriesByElves)))

	return totalCaloriesByElves
}

func countTotalCaloriesOfElf(elfCalories string) int {
	calories, err := lib.SliceAtoi(strings.Split(elfCalories, "\n"))
	if err != nil {
		panic(err)
	}

	return lib.SliceIntSum(calories)
}
