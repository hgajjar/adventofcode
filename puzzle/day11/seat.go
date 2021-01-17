package day11

import "fmt"

type Seat string

func NewSeat(pos string) (Seat, error) {
	if pos != "L" && pos != "#" && pos != "." {
		return "", fmt.Errorf("Invalid seat position %s", pos)
	}
	return Seat(pos), nil
}

func (s Seat) IsEmpty() bool {
	return string(s) == "L"
}

func (s Seat) IsOccupied() bool {
	return string(s) == "#"
}

func (s Seat) IsFloor() bool {
	return string(s) == "."
}

func (s Seat) Occupy() Seat {
	s = "#"
	return s
}

func (s Seat) Empty() Seat {
	s = "L"
	return s
}
