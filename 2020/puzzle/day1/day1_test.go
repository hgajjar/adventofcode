package day1_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day1"
)

func TestReportRepair(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	part1Result := day1.ReportRepairPart1(input)
	if part1Result != 514579 {
		t.Errorf("Expected %d, got %d", 514579, part1Result)
	}

	part2Result := day1.ReportRepairPart2(input)
	if part2Result != 241861950 {
		t.Errorf("Expected %d, got %d", 241861950, part2Result)
	}
}
