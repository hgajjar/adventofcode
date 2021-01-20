package day12

import "testing"

var (
	testInput = readNavigationInstructions("../../data/test/day12.txt")
	realInput = readNavigationInstructions("../../data/day12.txt")
)

func TestFindManhattanDistance(t *testing.T) {
	testResult := findManhattanDistance(testInput)

	if testResult != 25 {
		t.Errorf("Expected Manhattan distance: %d, found: %d", 25, testResult)
	}

	realResult := findManhattanDistance(realInput)

	if realResult != 1148 {
		t.Errorf("Expected Manhattan distance: %d, found: %d", 1148, realResult)
	}
}
