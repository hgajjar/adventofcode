package day10_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day10"
)

var (
	testIn1 = day10.ReadAdapters("../../data/test/day10.txt")
	testIn2 = day10.ReadAdapters("../../data/test/day10_2.txt")
	realIn  = day10.ReadAdapters("../../data/day10.txt")
)

func TestFindProductOfJoltDifferences(t *testing.T) {
	testOp1 := day10.FindProductOfJoltDifferences(testIn1)

	if testOp1 != 35 {
		t.Errorf("Expected product of jolt differences: %d, got %d", 35, testOp1)
	}

	testOp2 := day10.FindProductOfJoltDifferences(testIn2)

	if testOp2 != 220 {
		t.Errorf("Expected product of jolt differences: %d, got %d", 220, testOp2)
	}

	realOp1 := day10.FindProductOfJoltDifferences(realIn)

	if realOp1 != 2046 {
		t.Errorf("Expected product of jolt differences: %d, got %d", 2046, realOp1)
	}
}

func TestCountCombinationsOfAdapterConnections(t *testing.T) {
	testOp1 := day10.CountCombinationsOfAdapterConnections(testIn1)

	if testOp1 != 8 {
		t.Errorf("Expected adapter combinations: %d, got %d", 8, testOp1)
	}

	testOp2 := day10.CountCombinationsOfAdapterConnections(testIn2)

	if testOp2 != 19208 {
		t.Errorf("Expected adapter combinations: %d, got %d", 19208, testOp2)
	}

	realOp := day10.CountCombinationsOfAdapterConnections(realIn)

	if realOp != 1157018619904 {
		t.Errorf("Expected adapter combinations: %d, got %d", 1157018619904, realOp)
	}
}
