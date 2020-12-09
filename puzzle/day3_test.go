package puzzle_test

import (
	"fmt"
	"testing"

	"github.com/hgajjar/adventofcode/puzzle"
)

func TestFindNumberOfTreesAtPath(t *testing.T) {
	treeMap := puzzle.TreeMap{
		Coords: [][]string{
			[]string{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
			[]string{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
			[]string{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
			[]string{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
			[]string{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
			[]string{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
			[]string{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
			[]string{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			[]string{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
			[]string{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
			[]string{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
		},
	}

	findPath := puzzle.SlopePath{3, 1}

	totalTreesAtPath := treeMap.FindNumberOfTreesAtPath(findPath)

	if totalTreesAtPath != 7 {
		t.Errorf(fmt.Sprintf("Expected %d trees, found %d", 7, totalTreesAtPath))
	}
}
