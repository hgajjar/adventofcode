package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadSeatLayout(fileName string) SeatLayout {
	content, _ := ioutil.ReadFile(fileName)
	seatLayout := SeatLayout{}
	for _, line := range strings.Split(string(content), "\n") {
		seatRow := []Seat{}
		for _, pos := range line {
			s, _ := NewSeat(fmt.Sprintf("%c", pos))
			seatRow = append(seatRow, s)
		}
		seatLayout = append(seatLayout, seatRow)
	}

	return seatLayout
}
