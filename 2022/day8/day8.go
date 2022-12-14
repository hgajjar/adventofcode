package day8

import (
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(getMaxScenicScore(sampleInput))
	fmt.Println(getMaxScenicScore(actualInput))
}

func getMaxScenicScore(input string) (max int) {
	grid := createGrid(input)
	for i, x := range grid {
		for j, y := range x {
			if score := scenicScore(grid, y, i, j); score > max {
				max = score
			}
		}
	}

	return
}

func createGrid(input string) [][]int {
	return lo.Map(strings.Split(input, "\n"), func(row string, _ int) []int {
		return lo.Map[byte, int]([]byte(row), func(b byte, _ int) int {
			i, _ := strconv.Atoi(string(b))
			return i
		})
	})
}

func scenicScore(grid [][]int, currentElement, i, j int) int {
	if isOnEdge(grid, i, j) {
		return 0
	}

	return checkPrev(grid, currentElement, i, j) * checkNext(grid, currentElement, i, j) *
		checkTop(grid, currentElement, i, j) * checkBottom(grid, currentElement, i, j)
}

func isOnEdge(grid [][]int, i, j int) bool {
	return i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1
}

func checkPrev(grid [][]int, currentElement, i, j int) int {
	for k := j - 1; k >= 0; k-- {
		if currentElement <= grid[i][k] {
			return j - k
		}
	}

	return j
}

func checkNext(grid [][]int, currentElement, i, j int) int {
	for k := j + 1; k < len(grid[i]); k++ {
		if currentElement <= grid[i][k] {
			return k - j
		}
	}

	return len(grid[i]) - 1 - j
}

func checkTop(grid [][]int, currentElement, i, j int) int {
	for k := i - 1; k >= 0; k-- {
		if currentElement <= grid[k][j] {
			return i - k
		}
	}

	return i
}

func checkBottom(grid [][]int, currentElement, i, j int) int {
	for k := i + 1; k < len(grid); k++ {
		if currentElement <= grid[k][j] {
			return k - i
		}
	}

	return len(grid) - 1 - i
}
