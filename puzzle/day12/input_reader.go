package day12

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readNavigationInstructions(fileName string) instructionSet {
	content, _ := ioutil.ReadFile(fileName)
	set := instructionSet{}
	for  _, line := range strings.Split(string(content), "\n") {
		act := action(fmt.Sprintf("%c", line[0]))
		val, _ := strconv.Atoi(line[1:])
		inst := instruction{act, val}
		set = append(set, inst)
	}
	return set
}
