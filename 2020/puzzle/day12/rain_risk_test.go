package day12

import "testing"

var (
	testInput = readNavigationInstructions("../../data/test/day12.txt")
	realInput = readNavigationInstructions("../../data/day12.txt")
)

func TestFindManhattanDistance(t *testing.T) {
	testResult := findManhattanDistance(testInput)

	if testResult != 286 {
		t.Errorf("Expected Manhattan distance: %d, found: %d", 286, testResult)
	}

	realResult := findManhattanDistance(realInput)

	if realResult != 52203 {
		t.Errorf("Expected Manhattan distance: %d, found: %d", 52203, realResult)
	}
}
