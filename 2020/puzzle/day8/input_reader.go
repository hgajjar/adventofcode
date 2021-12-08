package day8

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadBootCode(fileName string) []Instruction {
	content, _ := ioutil.ReadFile(fileName)
	instuctionsText := strings.Split(string(content), "\n")

	instructions := []Instruction{}

	for _, i := range instuctionsText {
		parts := strings.Split(i, " ")
		instArgument, _ := strconv.Atoi(parts[1])
		instruction := Instruction{Operation: parts[0], Argument: int32(instArgument)}
		instructions = append(instructions, instruction)
	}

	return instructions
}
