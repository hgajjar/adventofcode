package day9_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day9"
)

var testInput = day9.ReadXmasEncryption("../../data/test/day9.txt")
var realInput = day9.ReadXmasEncryption("../../data/day9.txt")

func TestFindEncodingError(t *testing.T) {
	testOutput := day9.FindEncodingError(testInput, day9.NewPreambleRange(5))

	if testOutput != 127 {
		t.Errorf("Expected %d, got %d", 127, testOutput)
	}

	realOutput := day9.FindEncodingError(realInput, day9.NewPreambleRange(25))

	if realOutput != 1038347917 {
		t.Errorf("Expected %d, got %d", 1038347917, realOutput)
	}
}

func TestFindSumOfEncryptionWeakness(t *testing.T) {
	testOutput := day9.FindSumOfEncryptionWeakness(testInput, day9.NewPreambleRange(5))

	if testOutput != 62 {
		t.Errorf("Expected %d, got %d", 62, testOutput)
	}

	realOutput := day9.FindSumOfEncryptionWeakness(realInput, day9.NewPreambleRange(25))

	if realOutput != 62 {
		t.Errorf("Expected %d, got %d", 62, realOutput)
	}
}
