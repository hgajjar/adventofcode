package day3_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day3"
)

var treeMap = day3.TreeMap{
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
	findPath := day3.SlopePath{3, 1}

	totalTreesAtPath := treeMap.FindNumberOfTreesAtPath(findPath)

	if totalTreesAtPath != 7 {
		t.Errorf("Expected %d trees, found %d", 7, totalTreesAtPath)
	}
}

func TestFindProductOfNumberOfTreesAtPaths(t *testing.T) {
	slopePaths := []day3.SlopePath{
		day3.SlopePath{1, 1},
		day3.SlopePath{3, 1},
		day3.SlopePath{5, 1},
		day3.SlopePath{7, 1},
		day3.SlopePath{1, 2},
	}

	productOfTotalTreesAtPaths := treeMap.FindProductOfNumberOfTreesAtPaths(slopePaths)

	if productOfTotalTreesAtPaths != 336 {
		t.Errorf("Expected product of total trees as %d, found %d", 336, productOfTotalTreesAtPaths)
	}
}
