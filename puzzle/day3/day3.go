package day3

// TreeMap structure
type TreeMap struct {
	Coords [][]string
}

// SlopePath structure is used for navigatinng the TreeMap structure
type SlopePath struct {
	Right int
	Down  int
}

// FindNumberOfTreesAtPath finds total number of trees using given SlopePath
func (treeMap *TreeMap) FindNumberOfTreesAtPath(slopePath SlopePath) int {
	x := 0
	y := 0
	totalTreesAtPath := 0

	for x < len(treeMap.Coords)-1 {
		x += slopePath.Down
		y += slopePath.Right

		if treeMap.isTreeAtLocation(x, y) {
			totalTreesAtPath++
		}
	}

	return totalTreesAtPath
}

// FindProductOfNumberOfTreesAtPaths finds total number of trees using given slope paths and returns product of them
func (treeMap *TreeMap) FindProductOfNumberOfTreesAtPaths(slopePaths []SlopePath) int {
	product := 1
	for _, slopePath := range slopePaths {
		product *= treeMap.FindNumberOfTreesAtPath(slopePath)
	}

	return product
}

func (treeMap *TreeMap) isTreeAtLocation(x int, y int) bool {
	if y >= len(treeMap.Coords[x]) {
		y %= len(treeMap.Coords[x])
	}

	return treeMap.Coords[x][y] == "#"
}
