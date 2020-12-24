package day5

import (
	"fmt"
	"sync"
)

// Flight defines the seating arrangement structure
type Flight struct {
	SeatRows    SeatRange
	SeatColumns SeatRange
}

// SeatRange defines the lower and upper bounds of seat rows and columns
type SeatRange struct {
	Lower int
	Upper int
}

// BoardingPass defines the binary space partitioning number and its associated seat ID
type BoardingPass struct {
	PassNumber string
}

type seat struct {
	row    int
	column int
	seatID int
}

// FindSeatID finds the seat ID of the given boarding pass
func (flight Flight) FindSeatID(pass *BoardingPass) int {
	seat := flight.findSeat(pass)
	return seat.seatID
}

// GetHighestSeatID finds seat IDs of each boarding pass and returns
// the one with highest seat ID
func GetHighestSeatID(passes []BoardingPass) uint32 {
	flight := Flight{SeatRange{0, 127}, SeatRange{0, 7}}
	var wg sync.WaitGroup
	wg.Add(len(passes))

	in := make(chan uint32, len(passes))
	out := make(chan uint32)

	for _, pass := range passes {
		go func(p BoardingPass) {
			defer wg.Done()
			in <- uint32(flight.FindSeatID(&p))
		}(pass)
	}

	go findHighest(in, out)

	wg.Wait()
	close(in)

	return <-out
}

func findHighest(in chan uint32, out chan uint32) {
	var highest uint32 = 0
	for id := range in {
		if id > highest {
			highest = id
		}
	}
	out <- highest
}

// FindRemainingSeatID finds the remaining seat ID by eliminating the occupied seat IDs
// it ignores the non-occupied seats which does not have an occupied seat IDs around it
func FindRemainingSeatID(passes []BoardingPass) int {
	flight := Flight{SeatRange{0, 127}, SeatRange{0, 7}}
	occupiedSeatMatrix := make([][]int, flight.SeatRows.Upper+1)
	occupiedSeatIds := map[int]bool{}

	for i := range occupiedSeatMatrix {
		occupiedSeatMatrix[i] = make([]int, flight.SeatColumns.Upper+1)
	}

	for _, pass := range passes {
		seat := flight.findSeat(&pass)
		occupiedSeatMatrix[seat.row][seat.column] = seat.seatID
		occupiedSeatIds[seat.seatID] = true
	}

	for row := flight.SeatRows.Lower; row <= flight.SeatRows.Upper; row++ {
		for column := flight.SeatColumns.Lower; column <= flight.SeatColumns.Upper; column++ {
			if occupiedSeatMatrix[row][column] == 0 {
				currentSeatID := calculateSeatID(row, column)
				if occupiedSeatIds[currentSeatID-1] && occupiedSeatIds[currentSeatID+1] {
					return currentSeatID
				}
			}
		}
	}

	panic("can not find remaining seat ID")
}

func (flight Flight) findSeat(pass *BoardingPass) seat {
	passNumberArr := convertStringToArrayOfChar(pass.PassNumber)

	for _, binaryPartitionNumber := range passNumberArr {
		switch binaryPartitionNumber {
		case "F":
			flight.SeatRows.Upper = flight.SeatRows.Lower + splitRangeInHalf(flight.SeatRows)
		case "B":
			flight.SeatRows.Lower = flight.SeatRows.Lower + splitRangeInHalf(flight.SeatRows) + 1
		case "L":
			flight.SeatColumns.Upper = flight.SeatColumns.Lower + splitRangeInHalf(flight.SeatColumns)
		case "R":
			flight.SeatColumns.Lower = flight.SeatColumns.Lower + splitRangeInHalf(flight.SeatColumns) + 1
		default:
			panic(fmt.Sprintf("%s is not a valid binary partition number", binaryPartitionNumber))
		}
	}

	if flight.SeatRows.Lower != flight.SeatRows.Upper {
		panic("Seat row not found")
	}
	if flight.SeatColumns.Lower != flight.SeatColumns.Upper {
		panic("Seat column not found")
	}

	s := seat{row: flight.SeatRows.Lower, column: flight.SeatColumns.Lower}
	s.seatID = calculateSeatID(s.row, s.column)

	return s
}

func calculateSeatID(row int, column int) int {
	return row*8 + column
}

func convertStringToArrayOfChar(inputString string) []string {
	charArray := []string{}
	for _, s := range inputString {
		charArray = append(charArray, fmt.Sprintf("%c", s))
	}
	return charArray
}

func splitRangeInHalf(seatRange SeatRange) int {
	return (seatRange.Upper - seatRange.Lower - 1) / 2
}
