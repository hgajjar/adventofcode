package day11_test

import (
	"fmt"
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day11"
)

func TestReadSeatLayout(t *testing.T) {
	expectedSeatLayout := day11.SeatLayout{
		{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", "L", "L", "L", "L", "L", "L", ".", "L", "L"},
		{"L", ".", "L", ".", "L", ".", ".", "L", ".", "."},
		{"L", "L", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
		{".", ".", "L", ".", "L", ".", ".", ".", ".", "."},
		{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"},
		{"L", ".", "L", "L", "L", "L", "L", "L", ".", "L"},
		{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
	}
	seatLayout := day11.ReadSeatLayout("../../data/test/day11.txt")

	if !seatLayout.IsEqualTo(expectedSeatLayout) {
		t.Errorf("Seat layout does not match with expected layout")
		fmt.Println(seatLayout)
		fmt.Println(expectedSeatLayout)
	}
}
