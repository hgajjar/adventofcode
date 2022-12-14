package filesystem

import (
	"fmt"
	"github.com/samber/lo"
)

type Dir struct {
	name   string
	dirs   []*Dir
	files  []file
	parent *Dir
	size   int64
}

func (d *Dir) findChildDirByName(name string) *Dir {
	for _, childDir := range d.dirs {
		if childDir.name == name {
			return childDir
		}
	}

	panic(fmt.Sprintf("could not find child dir by name '%s'", name))
}

func (d *Dir) Parent() *Dir {
	return d.parent
}

func (d *Dir) Size() int64 {
	return d.size
}

func (d *Dir) ComputeSize() int64 {
	total := lo.SumBy(d.files, func(f file) int64 {
		return f.size
	})
	if len(d.dirs) > 0 {
		for _, d := range d.dirs {
			total += d.ComputeSize()
		}
	}

	d.size = total

	return total
}

func (d *Dir) FindDirsWithSizeGreaterThan(limit int64) []*Dir {
	var foundDirs []*Dir

	if d.size > limit {
		foundDirs = append(foundDirs, d)
	}

	for _, childDir := range d.dirs {
		foundDirs = append(foundDirs, childDir.FindDirsWithSizeGreaterThan(limit)...)
	}

	return foundDirs
}
