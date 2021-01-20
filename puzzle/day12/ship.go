package day12

import "fmt"

type ship struct {
	coords coordinates
	dir string
}

type coordinates struct {
	lat int // north/south
	long int // west/east
}

func newShip() *ship {
	return &ship{dir: "E"}
}

func (ship *ship) followInstruction(inst instruction) {
	switch inst.action {
	case "N":
		ship.moveNorth(inst.value)
	case "S":
		ship.moveSouth(inst.value)
	case "E":
		ship.moveEast(inst.value)
	case "W":
		ship.moveWest(inst.value)
	case "L":
		ship.rotateLeft(inst.value)
	case "R":
		ship.rotateRight(inst.value)
	case "F":
		ship.moveForward(inst.value)
	default:
		panic(fmt.Sprintf("Action %s is not supported", inst.action))
	}
}

func (ship *ship) moveNorth(val int) {
	ship.coords.lat -= val
}

func (ship *ship) moveSouth(val int) {
	ship.coords.lat += val
}

func (ship *ship) moveEast(val int) {
	ship.coords.long -= val
}

func (ship *ship) moveWest(val int) {
	ship.coords.long += val
}

func (ship *ship) rotateLeft(angle int) {
	for x := angle / 90; x > 0; x-- {
		ship.rotateCounterClockwise()
	}
}

func (ship *ship) rotateCounterClockwise() {
	switch ship.dir {
	case "E":
		ship.dir = "N"
	case "N":
		ship.dir = "W"
	case "W":
		ship.dir = "S"
	case "S":
		ship.dir = "E"
	}
}

func (ship *ship) rotateRight(angle int) {
	for x := angle / 90; x > 0; x-- {
		ship.rotateClockwise()
	}
}

func (ship *ship) rotateClockwise() {
	switch ship.dir {
	case "E":
		ship.dir = "S"
	case "S":
		ship.dir = "W"
	case "W":
		ship.dir = "N"
	case "N":
		ship.dir = "E"
	}
}

func (ship *ship) moveForward(val int) {
	switch ship.dir {
	case "N":
		ship.moveNorth(val)
	case "S":
		ship.moveSouth(val)
	case "E":
		ship.moveEast(val)
	case "W":
		ship.moveWest(val)
	}
}