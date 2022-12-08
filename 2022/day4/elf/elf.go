package elf

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type Elf struct {
	Sections []sectionId
}

type sectionId int

// NewWithAssignmentRange creates new elf with given section range
// input format: 2-4
func NewWithAssignmentRange(sectionRange string) Elf {
	limits := strings.Split(sectionRange, "-")
	min, _ := strconv.Atoi(limits[0])
	max, _ := strconv.Atoi(limits[1])

	return Elf{
		Sections: lo.RangeWithSteps[sectionId](sectionId(min), sectionId(max+1), 1),
	}
}
