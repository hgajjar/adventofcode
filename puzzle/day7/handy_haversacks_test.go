package day7_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day7"
)

var searchBag = day7.Bag{Color: "shiny gold"}
var realInput = day7.ParseInputToBags("../../data/day7.txt")

func TestCountOuterBags(t *testing.T) {

	testInput := day7.ParseInputToBags("../../data/test/day7.txt")
	testOutput := searchBag.CountOuterBags(testInput)

	if testOutput != 4 {
		t.Errorf("Expected outer bags: %d, got: %d", 4, testOutput)
	}

	realOutput := searchBag.CountOuterBags(realInput)

	if realOutput != 101 {
		t.Errorf("Expected outer bags: %d, got: %d", 101, realOutput)
	}
}

func BenchmarkCountOuterBags(b *testing.B) {
	// without concurrency: BenchmarkCountOuterBags-4            276           4140232 ns/op
	for i := 0; i < b.N; i++ {
		searchBag.CountOuterBags(realInput)
	}
}
