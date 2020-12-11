package day2_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day2"
)

func TestValidatePasswordsByStrategy(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	resultByOldStrategy := day2.ValidatePasswordsByStrategy(input, day2.OldStrategy)

	if resultByOldStrategy != 2 {
		t.Errorf("Expected %d, got %d", 2, resultByOldStrategy)
	}

	resultByNewStrategy := day2.ValidatePasswordsByStrategy(input, day2.NewStrategy)

	if resultByNewStrategy != 1 {
		t.Errorf("Expected %d, got %d", 1, resultByNewStrategy)
	}
}
