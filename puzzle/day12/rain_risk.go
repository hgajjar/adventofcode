package day12

import "math"

type instructionSet []instruction

type instruction struct {
	action action
	value int
}

type action string

func findManhattanDistance(instSet instructionSet) int {
	ship := newShip()

	for _, inst := range instSet {
		ship.followInstruction(inst)
	}

	return int(math.Abs(float64(ship.coords.lat)) + math.Abs(float64(ship.coords.long)))
}

