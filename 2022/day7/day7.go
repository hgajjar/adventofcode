package day7

import (
	"adventofcode/day7/filesystem"
	_ "embed"
	"fmt"
	"github.com/samber/lo"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

const (
	totalSpace        = 70000000
	freeSpaceRequired = 30000000
)

func Execute() {
	fmt.Println(findTotalSizeOfSmallestDirToFreeEnoughSpace(sampleInput))
	fmt.Println(findTotalSizeOfSmallestDirToFreeEnoughSpace(actualInput))
}

func findTotalSizeOfSmallestDirToFreeEnoughSpace(input string) int64 {
	filesys := filesystem.Create(input)
	filesys.ComputeSize()
	freeSpace := totalSpace - filesys.Size()
	deleteSpaceRequired := freeSpaceRequired - freeSpace

	smallestDir := lo.MinBy(filesys.FindDirsWithSizeGreaterThan(deleteSpaceRequired), func(dir *filesystem.Dir, min *filesystem.Dir) bool {
		return dir.Size() < min.Size()
	})

	return smallestDir.Size()
}
