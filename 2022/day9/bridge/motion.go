package bridge

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type Motion struct {
	Direction string
	Steps     int
}

func NewMotionGuide(input string) []Motion {
	return lo.Map(strings.Split(input, "\n"), func(row string, _ int) Motion {
		parts := strings.Split(row, " ")
		steps, _ := strconv.Atoi(parts[1])

		return Motion{
			Direction: parts[0],
			Steps:     steps,
		}
	})
}
