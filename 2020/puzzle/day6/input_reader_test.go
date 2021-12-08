package day6_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day6"
)

var expectedGroups = []day6.Group{
	{
		[]day6.Person{
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
						"b": true,
						"c": true,
					},
				},
			},
		},
	},
	{
		[]day6.Person{
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
					},
				},
			},
			{
				day6.CustomsForm{
					map[string]bool{
						"b": true,
					},
				},
			},
			{
				day6.CustomsForm{
					map[string]bool{
						"c": true,
					},
				},
			},
		},
	},
	{
		[]day6.Person{
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
						"b": true,
					},
				},
			},
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
						"c": true,
					},
				},
			},
		},
	},
	{
		[]day6.Person{
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
					},
				},
			},
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
					},
				},
			},
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
					},
				},
			},
			{
				day6.CustomsForm{
					map[string]bool{
						"a": true,
					},
				},
			},
		},
	},
	{
		[]day6.Person{
			{
				day6.CustomsForm{
					map[string]bool{
						"b": true,
					},
				},
			},
		},
	},
}

func TestParseInput(t *testing.T) {
	groups := day6.ParseInput("../../data/test/day6.txt")

	if len(groups) != len(expectedGroups) {
		t.Errorf("Expected number of groups: %d, got: %d", len(expectedGroups), len(groups))
	}

	for i, expectedGroup := range expectedGroups {
		if len(expectedGroup.Persons) != len(groups[i].Persons) {
			t.Errorf("Expected number of persons in group %d: %d, got: %d", i, len(expectedGroup.Persons), len(groups[i].Persons))
		}

		for j, expectedPerson := range expectedGroup.Persons {
			if len(expectedPerson.FilledForm.QuestionAnswers) != len(groups[i].Persons[j].FilledForm.QuestionAnswers) {
				t.Errorf("Expected number of answers in group %d by person %d: %d, got: %d", i, j, len(expectedPerson.FilledForm.QuestionAnswers), len(groups[i].Persons[j].FilledForm.QuestionAnswers))
			}
		}
	}
}
