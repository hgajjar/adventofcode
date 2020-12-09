package puzzle_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle"
)

func TestValidatePasswordsByStrategy(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	resultByOldStrategy := puzzle.ValidatePasswordsByStrategy(input, puzzle.OldStrategy)

	if resultByOldStrategy != 2 {
		t.Errorf("Expected %d, got %d", 2, resultByOldStrategy)
	}

	resultByNewStrategy := puzzle.ValidatePasswordsByStrategy(input, puzzle.NewStrategy)

	if resultByNewStrategy != 1 {
		t.Errorf("Expected %d, got %d", 1, resultByNewStrategy)
	}
}
