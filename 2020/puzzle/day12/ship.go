package day12

type ship struct {
	coords coordinates
}

func newShip() *ship {
	return &ship{}
}

func (ship *ship) moveForward(waypoint *waypoint, val int) {
	ship.coords.long += waypoint.coords.long * val
	ship.coords.lat += waypoint.coords.lat * val
}