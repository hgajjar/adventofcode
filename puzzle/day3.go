package puzzle

// TreeMap structure
type TreeMap struct {
	Coords [][]string
}

// SlopePath structure is used for navigatinng the TreeMap structure
type SlopePath struct {
	X int
	Y int
}

// FindNumberOfTreesAtPath finds total number of trees using given SlopePath
func (treeMap *TreeMap) FindNumberOfTreesAtPath(slopePath SlopePath) int {

	currentPos := slopePath.X
	totalTreesAtPath := 0

	for i := 1; i < len(treeMap.Coords); i += slopePath.Y {
		if treeMap.isTreeAtLocation(i, currentPos) {
			totalTreesAtPath++
		}

		currentPos = currentPos + slopePath.X
	}

	return totalTreesAtPath
}

func (treeMap *TreeMap) isTreeAtLocation(x int, y int) bool {
	if len(treeMap.Coords[x]) < y {
		y = y % len(treeMap.Coords[x])
	}
	return treeMap.Coords[x][y] == "#"
}
