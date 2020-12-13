package day5_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day5"
)

func TestFindSeatID(t *testing.T) {
	flight := day5.Flight{day5.SeatRange{0, 127}, day5.SeatRange{0, 7}}

	boardingPass1 := day5.BoardingPass{PassNumber: "BFFFBBFRRR"}
	seatID1 := flight.FindSeatID(&boardingPass1)

	if seatID1 != 567 {
		t.Errorf("Expected seat ID: %d, found: %d", 567, seatID1)
	}

	boardingPass2 := day5.BoardingPass{PassNumber: "FFFBBBFRRR"}
	seatID2 := flight.FindSeatID(&boardingPass2)

	if seatID2 != 119 {
		t.Errorf("Expected seat ID: %d, found: %d", 119, seatID2)
	}

	boardingPass3 := day5.BoardingPass{PassNumber: "BBFFBBFRLL"}
	seatID3 := flight.FindSeatID(&boardingPass3)

	if seatID3 != 820 {
		t.Errorf("Expected seat ID: %d, found: %d", 820, seatID3)
	}
}

func TestGetHighestSeatID(t *testing.T) {
	highestSeatID := day5.GetHighestSeatID([]day5.BoardingPass{
		day5.BoardingPass{PassNumber: "BFFFBBFRRR"},
		day5.BoardingPass{PassNumber: "FFFBBBFRRR"},
		day5.BoardingPass{PassNumber: "BBFFBBFRLL"},
	})

	if highestSeatID != 820 {
		t.Errorf("Expected highest seat ID: %d, found %d", 820, highestSeatID)
	}
}
