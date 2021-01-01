package day8_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day8"
)

var testInput = day8.ReadBootCode("../../data/test/day8.txt")
var realInput = day8.ReadBootCode("../../data/day8.txt")

func TestFindAccWhenRecursionDetected(t *testing.T) {
	testAccVal := day8.FindAccWhenRecursionDetected(testInput)

	if testAccVal != 5 {
		t.Errorf("Expected value %d, got %d", 5, testAccVal)
	}

	accVal := day8.FindAccWhenRecursionDetected(realInput)

	if accVal != 1614 {
		t.Errorf("Expected value %d, got %d", 1614, accVal)
	}
}

func TestFindAccAfterInstructionsFixed(t *testing.T) {
	testResult := day8.FindAccAfterInstructionsFixed(testInput)

	if testResult != 8 {
		t.Errorf("Expected value %d, got %d", 8, testResult)
	}

	realResult := day8.FindAccAfterInstructionsFixed(realInput)

	if realResult != 1260 {
		t.Errorf("Expected value %d, got %d", 1260, realResult)
	}
}
