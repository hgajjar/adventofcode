package puzzle_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle"
)

var treeMap = puzzle.TreeMap{
	Coords: [][]string{
		[]string{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
		[]string{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
		[]string{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
		[]string{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
		[]string{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
		[]string{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
		[]string{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
		[]string{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		[]string{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
		[]string{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
		[]string{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
	},
}

func TestFindNumberOfTreesAtPath(t *testing.T) {
	findPath := puzzle.SlopePath{3, 1}

	totalTreesAtPath := treeMap.FindNumberOfTreesAtPath(findPath)

	if totalTreesAtPath != 7 {
		t.Errorf("Expected %d trees, found %d", 7, totalTreesAtPath)
	}
}

func TestFindProductOfNumberOfTreesAtPaths(t *testing.T) {
	slopePaths := []puzzle.SlopePath{
		puzzle.SlopePath{1, 1},
		puzzle.SlopePath{3, 1},
		puzzle.SlopePath{5, 1},
		puzzle.SlopePath{7, 1},
		puzzle.SlopePath{1, 2},
	}

	productOfTotalTreesAtPaths := treeMap.FindProductOfNumberOfTreesAtPaths(slopePaths)

	if productOfTotalTreesAtPaths != 336 {
		t.Errorf("Expected product of total trees as %d, found %d", 336, productOfTotalTreesAtPaths)
	}
}
