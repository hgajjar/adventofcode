package day6_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day6"
)

func TestCountTotalUniqueGroupAnswers(t *testing.T) {
	count := day6.CountAnswersWhereAnyoneInGroupSaidYes(day6.ParseInput("../../data/test/day6.txt"))

	if count != 11 {
		t.Errorf("Expected count: %d, got: %d", 11, count)
	}
}

func TestCountAnswersWhereEveryoneInGroupSaidYes(t *testing.T) {
	count := day6.CountAnswersWhereEveryoneInGroupSaidYes(day6.ParseInput("../../data/test/day6.txt"))

	if count != 6 {
		t.Errorf("Expected count: %d, got: %d", 6, count)
	}
}

func BenchmarkCountTotalUniqueGroupAnswers(t *testing.B) {
	day6.CountAnswersWhereAnyoneInGroupSaidYes(day6.ParseInput("../../data/day6.txt"))
}

func BenchmarkCountAnswersWhereEveryoneInGroupSaidYes(t *testing.B) {
	day6.CountAnswersWhereEveryoneInGroupSaidYes(day6.ParseInput("../../data/day6.txt"))
}
