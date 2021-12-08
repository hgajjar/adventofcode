package day7_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day7"
)

var expectedBags = []day7.Bag{
	{Color: "light red", Bags: []day7.Bag{{Color: "bright white", Qty: 1}, {Color: "muted yellow", Qty: 2}}},
	{Color: "dark orange", Bags: []day7.Bag{{Color: "bright white", Qty: 3}, {Color: "muted yellow", Qty: 4}}},
	{Color: "bright white", Bags: []day7.Bag{{Color: "shiny gold", Qty: 1}}},
	{Color: "muted yellow", Bags: []day7.Bag{{Color: "shiny gold", Qty: 2}, {Color: "faded blue", Qty: 9}}},
	{Color: "shiny gold", Bags: []day7.Bag{{Color: "dark olive", Qty: 1}, {Color: "vibrant plum", Qty: 2}}},
	{Color: "dark olive", Bags: []day7.Bag{{Color: "faded blue", Qty: 3}, {Color: "dotted black", Qty: 4}}},
	{Color: "vibrant plum", Bags: []day7.Bag{{Color: "faded blue", Qty: 5}, {Color: "dotted black", Qty: 6}}},
	{Color: "faded blue"},
	{Color: "dotted black"},
}

func TestParseInputToBags(t *testing.T) {
	result := day7.ParseInputToBags("../../data/test/day7.txt")

	for i, bag := range result {
		if bag.Color != expectedBags[i].Color {
			t.Errorf("Expected bag color %s, got %s", expectedBags[i].Color, bag.Color)
		}

		if len(bag.Bags) != len(expectedBags[i].Bags) {
			t.Errorf("Bag %s should contain %d child bags, got %d", bag.Color, len(expectedBags[i].Color), len(bag.Bags))
		}

		for j, childBag := range bag.Bags {
			if childBag.Color != expectedBags[i].Bags[j].Color {
				t.Errorf("Expected bag color %s, got %s", expectedBags[i].Bags[j].Color, childBag.Color)
			}
		}
	}
}
