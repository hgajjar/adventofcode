package day11_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day11"
)

var (
	testInput = day11.ReadSeatLayout("../../data/test/day11.txt")
	realInput = day11.ReadSeatLayout("../../data/day11.txt")
)

func TestSimulateAndReturnOccupiedSeatsCount(t *testing.T) {
	testResult := testInput.SimulateAndReturnOccupiedSeatsCount()
	if testResult != 37 {
		t.Errorf("Expected count %d, got %d", 37, testResult)
	}

	realResult := realInput.SimulateAndReturnOccupiedSeatsCount()
	if realResult != 2494 {
		t.Errorf("Expected count %d, got %d", 2494, realResult)
	}
}

func TestSimulateAndReturnOccupiedSeatsCountPart2(t *testing.T) {
	testResult := testInput.SimulateAndReturnOccupiedSeatsCountPart2()
	if testResult != 26 {
		t.Errorf("Expected count %d, got %d", 26, testResult)
	}

	realResult := realInput.SimulateAndReturnOccupiedSeatsCountPart2()
	if realResult != 2306 {
		t.Errorf("Expected count %d, got %d", 2306, realResult)
	}
}
