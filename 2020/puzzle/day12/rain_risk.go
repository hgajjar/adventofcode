package day12

import (
	"fmt"
	"math"
)

type instructionSet []instruction

type instruction struct {
	action action
	value  int
}

type action string

type coordinates struct {
	lat  int // north/south
	long int // west/east
}

const (
	NORTH   = "N"
	SOUTH   = "S"
	EAST    = "E"
	WEST    = "W"
	LEFT    = "L"
	RIGHT   = "R"
	FORWARD = "F"
)

func findManhattanDistance(instSet instructionSet) int {
	ship := newShip()
	waypoint := newWaypoint()

	for _, inst := range instSet {
		ship.followInstructions(waypoint, inst)
	}

	return int(math.Abs(float64(ship.coords.lat)) + math.Abs(float64(ship.coords.long)))
}

func (ship *ship) followInstructions(waypoint *waypoint, inst instruction) {
	switch inst.action {
	case NORTH:
		waypoint.moveNorth(inst.value)
	case SOUTH:
		waypoint.moveSouth(inst.value)
	case EAST:
		waypoint.moveEast(inst.value)
	case WEST:
		waypoint.moveWest(inst.value)
	case LEFT:
		waypoint.rotateLeft(inst.value)
	case RIGHT:
		waypoint.rotateRight(inst.value)
	case FORWARD:
		ship.moveForward(waypoint, inst.value)
	default:
		panic(fmt.Sprintf("Action %s is not supported", inst.action))
	}
}
