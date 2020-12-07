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
