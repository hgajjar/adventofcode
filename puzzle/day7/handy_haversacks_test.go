package day7_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day7"
)

var searchBag = day7.Bag{Color: "shiny gold"}
var realInput = day7.ParseInputToBags("../../data/day7.txt")

func TestCountOuterBags(t *testing.T) {

	// testInput := day7.ParseInputToBags("../../data/test/day7.txt")
	// testOutput := searchBag.CountOuterBags(testInput)

	// if testOutput != 4 {
	// 	t.Errorf("Expected outer bags: %d, got: %d", 4, testOutput)
	// }

	realOutput := searchBag.CountOuterBags(realInput)

	if realOutput != 101 {
		t.Errorf("Expected outer bags: %d, got: %d", 101, realOutput)
	}
}

func BenchmarkCountOuterBags(b *testing.B) {
	// without concurrency: BenchmarkCountOuterBags-4            276           4140232 ns/op
	// with concurrency: BenchmarkCountOuterBags-4            480           2649880 ns/op (56% faster)

	// with 5 ns delay
	// without BenchmarkCountOuterBags-3             25          49470532 ns/op
	// with BenchmarkCountOuterBags-3             79          14439482 ns/op (71% faster)
	for i := 0; i < b.N; i++ {
		searchBag.CountOuterBags(realInput)
	}
}
