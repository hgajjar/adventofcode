package day8_test

import (
	"reflect"
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day8"
)

var expectedInstructions []day8.Instruction = []day8.Instruction{
	{"nop", 0},
	{"acc", 1},
	{"jmp", 4},
	{"acc", 3},
	{"jmp", -3},
	{"acc", -99},
	{"acc", 1},
	{"jmp", -4},
	{"acc", 6},
}

func TestReadBootCode(t *testing.T) {
	testResult := day8.ReadBootCode("../../data/test/day8.txt")

	if !reflect.DeepEqual(testResult, expectedInstructions) {
		t.Errorf("Expected %#v \n Got %#v", expectedInstructions, testResult)
	}
}
