package day10

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadAdapters(fileName string) []Adapter {
	content, _ := ioutil.ReadFile(fileName)

	adapters := []Adapter{}
	for _, joltage := range strings.Split(string(content), "\n") {
		intJoltage, _ := strconv.Atoi(joltage)
		adapters = append(adapters, NewAdapter(intJoltage))
	}

	return adapters
}
