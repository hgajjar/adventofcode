package lib

import "strconv"

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func SliceIntSum(s []int) int {
	var sum int
	for _, e := range s {
		sum += e
	}

	return sum
}