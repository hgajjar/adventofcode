package day11

import (
	"errors"
)

type SeatLayout [][]Seat

func (sl SeatLayout) SimulateAndReturnOccupiedSeatsCount() int {
	return sl.simulate(true, 4).countOccupiedSeats()
}

func (sl SeatLayout) SimulateAndReturnOccupiedSeatsCountPart2() int {
	return sl.simulate(false, 5).countOccupiedSeats()
}

func (sl SeatLayout) simulate(onlyAdjacent bool, occupiedCount int) SeatLayout {
	slc := sl.copy()

	for x, row := range sl {
		for y, seat := range row {
			count := 0
			if onlyAdjacent {
				count = sl.countOccupiedAdjacentSeats(x, y)
			} else {
				count = sl.countFirstOccupiedSeatsInAllDirectionsFrom(x, y)
			}
			if count == 0 && seat.IsEmpty() {
				slc[x][y] = seat.Occupy()
			} else if count >= occupiedCount && seat.IsOccupied() {
				slc[x][y] = seat.Empty()
			}
		}
	}

	if slc.IsEqualTo(sl) {
		return slc
	}

	return slc.simulate(onlyAdjacent, occupiedCount)
}

func (sl SeatLayout) countOccupiedAdjacentSeats(x, y int) int {
	c := 0
	adjacentPositions := [8][2]int{
		{x, y - 1},
		{x, y + 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y + 1},
	}

	for _, pos := range adjacentPositions {
		s, err := sl.getSeatAt(pos[0], pos[1])

		if err != nil {
			continue
		}

		if s.IsOccupied() {
			c++
		}
	}

	return c
}

func (sl SeatLayout) countFirstOccupiedSeatsInAllDirectionsFrom(x, y int) int {
	c := 0
	c += sl.countLeft(x, y)
	c += sl.countRight(x, y)
	c += sl.countTop(x, y)
	c += sl.countBottom(x, y)
	c += sl.countTopLeft(x, y)
	c += sl.countTopRight(x, y)
	c += sl.countBottomLeft(x, y)
	c += sl.countBottomRight(x, y)

	return c
}

func (sl SeatLayout) countLeft(x, y int) int {
	for i := y - 1; i >= 0; i-- {
		s, _ := sl.getSeatAt(x, i)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}
	return 0
}

func (sl SeatLayout) countRight(x, y int) int {
	for i := y + 1; i < len(sl[x]); i++ {
		s, _ := sl.getSeatAt(x, i)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}
	return 0
}

func (sl SeatLayout) countTop(x, y int) int {
	for i := x - 1; i >= 0; i-- {
		s, _ := sl.getSeatAt(i, y)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}

	return 0
}

func (sl SeatLayout) countBottom(x, y int) int {
	for i := x + 1; i < len(sl); i++ {
		s, _ := sl.getSeatAt(i, y)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}
	return 0
}

func (sl SeatLayout) countTopLeft(x, y int) int {
	for x = x - 1; x >= 0; x-- {
		if y < 0 {
			break
		}
		y--
		s, _ := sl.getSeatAt(x, y)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}

	return 0
}

func (sl SeatLayout) countTopRight(x, y int) int {
	for x := x - 1; x >= 0; x-- {
		if y == len(sl[x]) {
			break
		}
		y++
		s, _ := sl.getSeatAt(x, y)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}

	return 0
}

func (sl SeatLayout) countBottomLeft(x, y int) int {
	for x = x + 1; x < len(sl); x++ {
		if y < 0 {
			break
		}
		y--
		s, _ := sl.getSeatAt(x, y)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}

	}

	return 0
}

func (sl SeatLayout) countBottomRight(x, y int) int {
	for x = x + 1; x < len(sl); x++ {
		if y == len(sl[x]) {
			break
		}
		y++
		s, _ := sl.getSeatAt(x, y)
		if s.IsOccupied() {
			return 1
		}
		if s.IsEmpty() {
			return 0
		}
	}

	return 0
}

func (sl SeatLayout) countOccupiedSeats() int {
	c := 0
	for _, r := range sl {
		for _, s := range r {
			if s.IsOccupied() {
				c++
			}
		}
	}
	return c
}

func (sl SeatLayout) getSeatAt(x, y int) (Seat, error) {
	if x < 0 || len(sl) <= x {
		return "", errors.New("Index out of range")
	}
	if y < 0 || len(sl[x]) <= y {
		return "", errors.New("Index out of range")
	}
	return sl[x][y], nil
}

func (original SeatLayout) copy() SeatLayout {
	copy := SeatLayout{}
	for _, r := range original {
		row := []Seat{}
		for _, c := range r {
			row = append(row, c)
		}
		copy = append(copy, row)
	}

	return copy
}

func (sl1 SeatLayout) IsEqualTo(sl2 SeatLayout) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for i, r := range sl1 {
		if len(sl1[i]) != len(sl2[i]) {
			return false
		}
		for j := range r {
			if sl1[i][j] != sl2[i][j] {
				return false
			}
		}
	}

	return true
}
