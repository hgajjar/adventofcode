package day12

type waypoint struct {
	coords coordinates
}

func newWaypoint() *waypoint {
	return &waypoint{coords: coordinates{long: 10, lat: -1}}
}

func (waypoint *waypoint) moveNorth(val int) {
	waypoint.coords.lat -= val
}

func (waypoint *waypoint) moveSouth(val int) {
	waypoint.coords.lat += val
}

func (waypoint *waypoint) moveEast(val int) {
	waypoint.coords.long += val
}

func (waypoint *waypoint) moveWest(val int) {
	waypoint.coords.long -= val
}

func (waypoint *waypoint) rotateLeft(angle int) {
	for x := angle / 90; x > 0; x-- {
		waypoint.rotateCounterClockwise()
	}
}

func (waypoint *waypoint) rotateCounterClockwise() {
	long := waypoint.coords.lat
	lat := -waypoint.coords.long
	waypoint.coords = coordinates{long: long, lat: lat}
}

func (waypoint *waypoint) rotateRight(angle int) {
	for x := angle / 90; x > 0; x-- {
		waypoint.rotateClockwise()
	}
}

func (waypoint *waypoint) rotateClockwise() {
	long := -waypoint.coords.lat
	lat := waypoint.coords.long
	waypoint.coords = coordinates{long: long, lat: lat}
}