package day10

import (
	"sort"
)

type Adapter struct {
	joltage int
}

func NewAdapter(joltage int) Adapter {
	return Adapter{joltage: joltage}
}

func FindProductOfJoltDifferences(adapters []Adapter) int {
	sortAdapters(adapters)
	preparedAdapters := addBuiltInAdapters(adapters)
	oneJoltDiff, threeJoltDiff := findJoltDifferences(preparedAdapters)

	return oneJoltDiff * threeJoltDiff
}

func CountCombinationsOfAdapterConnections(adapters []Adapter) int {
	sortAdapters(adapters)
	preparedAdapters := addBuiltInAdapters(adapters)

	adaptersChunks := splitByThreeJoltDiff(preparedAdapters)

	totalCombinations := 1
	for _, ads := range adaptersChunks {
		switch len(ads) {
		case 1:
			totalCombinations *= 2
		case 2:
			totalCombinations *= 4
		default:
			totalCombinations *= (4 + len(ads))
		}
	}

	return totalCombinations
}

func sortAdapters(adapters []Adapter) []Adapter {
	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i].joltage < adapters[j].joltage
	})
	return adapters
}

func addBuiltInAdapters(adapters []Adapter) []Adapter {
	adaptersWithBuiltIn := []Adapter{NewAdapter(0)}
	adaptersWithBuiltIn = append(adaptersWithBuiltIn, adapters...)
	adaptersWithBuiltIn = append(adaptersWithBuiltIn, NewAdapter(adaptersWithBuiltIn[len(adaptersWithBuiltIn)-1].joltage+3))

	return adaptersWithBuiltIn
}

func findJoltDifferences(adapters []Adapter) (int, int) {
	oneJoltDiff := 0
	threeJoltDiff := 0

	for i := range adapters {
		switch joltDiff := getJoltDiffWithNextAdapter(adapters, i); joltDiff {
		case 1:
			oneJoltDiff++
		case 3:
			threeJoltDiff++
		}
	}

	return oneJoltDiff, threeJoltDiff
}

func splitByThreeJoltDiff(adapters []Adapter) [][]Adapter {
	splitedAdapters := [][]Adapter{}
	lastSplitIndex := 1
	for i := range adapters {
		if getJoltDiffWithNextAdapter(adapters, i) == 3 {
			if lastSplitIndex < i {
				splitedAdapters = append(splitedAdapters, adapters[lastSplitIndex:i])
			}
			lastSplitIndex = i + 2
		}
	}

	return splitedAdapters
}

func getJoltDiffWithNextAdapter(adapters []Adapter, index int) int {
	var joltDiff int
	if index == len(adapters)-1 {
		joltDiff = 0
	} else {
		joltDiff = adapters[index+1].joltage - adapters[index].joltage
	}

	return joltDiff
}
