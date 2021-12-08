package day7_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day7"
)

var searchBag = day7.Bag{Color: "shiny gold", Qty: 1}
var realInput = day7.ParseInputToBags("../../data/day7.txt")
var testInput = day7.ParseInputToBags("../../data/test/day7.txt")

func TestCountOuterBags(t *testing.T) {
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
	// with concurrency: BenchmarkCountOuterBags-4            480           2649880 ns/op (56% faster)
	for i := 0; i < b.N; i++ {
		searchBag.CountOuterBags(realInput)
	}
}

func TestCountInnerBags(t *testing.T) {
	testOutput := searchBag.CountInnerBags(testInput)

	if testOutput != 32 {
		t.Errorf("Expected inner bags: %d, got: %d", 32, testOutput)
	}

	realOutput := searchBag.CountInnerBags(realInput)

	if realOutput != 108636 {
		t.Errorf("Expected inner bags: %d, got: %d", 108636, realOutput)
	}
}

func BenchmarkCountInnerBags(b *testing.B) {
	// without concurrency: BenchmarkCountInnerBags-4           7701            147729 ns/op
	// with concurrency: BenchmarkCountInnerBags-4          11367            107975 ns/op (27% faster)
	for i := 0; i < b.N; i++ {
		searchBag.CountInnerBags(realInput)
	}
}
