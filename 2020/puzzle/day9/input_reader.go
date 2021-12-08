package day9

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadXmasEncryption(fileName string) []int {
	content, _ := ioutil.ReadFile(fileName)
	xmasSeries := []int{}

	for _, num := range strings.Split(string(content), "\n") {
		intNum, _ := strconv.Atoi(num)
		xmasSeries = append(xmasSeries, intNum)
	}

	return xmasSeries
}
