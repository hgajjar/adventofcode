package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/hgajjar/adventofcode/puzzle"
)

func main() {
	day1input := convertStringsToInt(readFileIntoStringsArray("data/day1.txt"))

	day1part1result := puzzle.ReportRepairPart1(day1input)
	fmt.Println(fmt.Sprintf("Day1 Part1 Result: %d", day1part1result))
	day1part2result := puzzle.ReportRepairPart2(day1input)
	fmt.Println(fmt.Sprintf("Day1 Part2 Result: %d", day1part2result))

	day2part1result := puzzle.ValidatePasswordsByStrategy(readFileIntoStringsArray("data/day2.txt"), puzzle.OldStrategy)
	fmt.Println(fmt.Sprintf("Day2 Part1 Result: %d", day2part1result))
	day2part2result := puzzle.ValidatePasswordsByStrategy(readFileIntoStringsArray("data/day2.txt"), puzzle.NewStrategy)
	fmt.Println(fmt.Sprintf("Day2 Part2 Result: %d", day2part2result))

	day3input := puzzle.TreeMap{
		Coords: convertStringsToArrayOfChar(readFileIntoStringsArray("data/day3.txt")),
	}
	day3part1result := day3input.FindNumberOfTreesAtPath(puzzle.SlopePath{Right: 1, Down: 2})
	fmt.Println(fmt.Sprintf("Day3 Part1 Result: %d", day3part1result))
	day3part2result := day3input.FindProductOfNumberOfTreesAtPaths([]puzzle.SlopePath{
		puzzle.SlopePath{Right: 1, Down: 1},
		puzzle.SlopePath{Right: 3, Down: 1},
		puzzle.SlopePath{Right: 5, Down: 1},
		puzzle.SlopePath{Right: 7, Down: 1},
		puzzle.SlopePath{Right: 1, Down: 2},
	})
	fmt.Println(fmt.Sprintf("Day3 Part2 Result: %d", day3part2result))
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
