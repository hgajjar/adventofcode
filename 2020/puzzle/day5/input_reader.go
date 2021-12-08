package day5

import (
	"io/ioutil"
	"strings"
)

// ParseInputToBoardingPasses reads given file and parses it into array of BoardingPass
func ParseInputToBoardingPasses(fileName string) []BoardingPass {
	content, _ := ioutil.ReadFile(fileName)
	passes := []BoardingPass{}

	for _, line := range strings.Split(string(content), "\n") {
		pass := BoardingPass{PassNumber: string(line)}
		passes = append(passes, pass)
	}

	return passes
}
