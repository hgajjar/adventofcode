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

func Execute() {
	fmt.Println(countTotalSizeofSmallDirs(sampleInput))
	fmt.Println(countTotalSizeofSmallDirs(actualInput))
}

func countTotalSizeofSmallDirs(input string) int64 {
	filesys := filesystem.Create(input)
	filesys.ComputeSize()
	dirs := filesys.FindDirsWithSizeLessThan(100000)
	return lo.SumBy(dirs, func(d *filesystem.Dir) int64 {
		return d.Size()
	})
}
