package crane

import (
	"regexp"
	"strconv"
	"strings"
)

type step struct {
	move int
	from int
	to   int
}

type Procedure []step

func CreateProcedure(input string) Procedure {
	r := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	var p Procedure
	for _, line := range strings.Split(input, "\n") {
		matches := r.FindStringSubmatch(line)
		s := step{
			move: toInt(matches[1]),
			from: toInt(matches[2]),
			to:   toInt(matches[3]),
		}
		p = append(p, s)
	}

	return p
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
