package day9_test

import (
	"reflect"
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day9"
)

func TestReadXmasEncryption(t *testing.T) {
	expectedXmasSeries := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	testXmasSeries := day9.ReadXmasEncryption("../../data/test/day9.txt")

	if !reflect.DeepEqual(expectedXmasSeries, testXmasSeries) {
		t.Errorf("Invalid Xmas series")
	}
}
