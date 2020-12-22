package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/hgajjar/adventofcode/puzzle/day5"
	"github.com/hgajjar/adventofcode/puzzle/day6"

	"github.com/hgajjar/adventofcode/puzzle/day1"
	"github.com/hgajjar/adventofcode/puzzle/day2"
	"github.com/hgajjar/adventofcode/puzzle/day3"
	"github.com/hgajjar/adventofcode/puzzle/day4"
)

func main() {
	day1input := convertStringsToInt(readFileIntoStringsArray("data/day1.txt"))

	day1part1result := day1.ReportRepairPart1(day1input)
	fmt.Println(fmt.Sprintf("Day1 Part1 Result: %d", day1part1result))
	day1part2result := day1.ReportRepairPart2(day1input)
	fmt.Println(fmt.Sprintf("Day1 Part2 Result: %d", day1part2result))

	day2part1result := day2.ValidatePasswordsByStrategy(readFileIntoStringsArray("data/day2.txt"), day2.OldStrategy)
	fmt.Println(fmt.Sprintf("Day2 Part1 Result: %d", day2part1result))
	day2part2result := day2.ValidatePasswordsByStrategy(readFileIntoStringsArray("data/day2.txt"), day2.NewStrategy)
	fmt.Println(fmt.Sprintf("Day2 Part2 Result: %d", day2part2result))

	day3input := day3.TreeMap{
		Coords: convertStringsToArrayOfChar(readFileIntoStringsArray("data/day3.txt")),
	}
	day3part1result := day3input.FindNumberOfTreesAtPath(day3.SlopePath{Right: 1, Down: 2})
	fmt.Println(fmt.Sprintf("Day3 Part1 Result: %d", day3part1result))
	day3part2result := day3input.FindProductOfNumberOfTreesAtPaths([]day3.SlopePath{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	})
	fmt.Println(fmt.Sprintf("Day3 Part2 Result: %d", day3part2result))

	day4part1result := day4.CountValidPassports(day4.ParseInputToPassports("data/day4.txt"), day4.ValidationStrategyEmptyness)
	fmt.Println(fmt.Sprintf("Day4 Part1 Result: %d", day4part1result))
	day4part2result := day4.CountValidPassports(day4.ParseInputToPassports("data/day4.txt"), day4.ValidationStrategyStrict)
	fmt.Println(fmt.Sprintf("Day4 Part2 Result: %d", day4part2result))

	day5input := day5.ParseInputToBoardingPasses("data/day5.txt")
	day5part1result := day5.GetHighestSeatID(day5input)
	fmt.Println(fmt.Sprintf("Day5 Part1 Result: %d", day5part1result))
	day5part2result := day5.FindRemainingSeatID(day5input)
	fmt.Println(fmt.Sprintf("Day5 Part2 Result: %d", day5part2result))

	day6part1result := day6.CountAnswersWhereAnyoneInGroupSaidYes(day6.ParseInput("data/day6.txt"))
	fmt.Println(fmt.Sprintf("Day6 Part1 Result: %d", day6part1result))
	day6part2result := day6.CountAnswersWhereEveryoneInGroupSaidYes(day6.ParseInput("data/day6.txt"))
	fmt.Println(fmt.Sprintf("Day6 Part2 Result: %d", day6part2result))
}

func readFileIntoStringsArray(fileName string) []string {
	content, _ := ioutil.ReadFile(fileName)
	return strings.Split(string(content), "\n")
}

func convertStringsToInt(array []string) []int {
	var result = []int{}

	for _, str := range array {
		convertedStr, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}

		result = append(result, convertedStr)
	}

	return result
}

func convertStringsToArrayOfChar(inputArray []string) [][]string {
	var result = make([][]string, len(inputArray))

	for i, str := range inputArray {
		charArray := []string{}
		for _, s := range str {
			charArray = append(charArray, fmt.Sprintf("%c", s))
		}
		result[i] = charArray
	}

	return result
}
