package day12

import (
	"reflect"
	"testing"
)

func TestReadNavigationInstructions(t *testing.T) {
	instSet := readNavigationInstructions("../../data/test/day12.txt")

	expectedSet := instructionSet{
		instruction{action("F"), 10},
		instruction{action("N"), 3},
		instruction{action("F"), 7},
		instruction{action("R"), 90},
		instruction{action("F"), 11},
	}

	if !reflect.DeepEqual(instSet, expectedSet) {
		t.Errorf("Expected instruction set: %#v, got %#v", expectedSet, instSet)
	}
}
