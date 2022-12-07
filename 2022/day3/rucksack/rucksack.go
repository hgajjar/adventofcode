package rucksack

import "adventofcode/day3/item"

type Rucksack struct {
	Compartments []Compartment
}

type Compartment []item.Item

func New(items string) Rucksack {
	return Rucksack{
		Compartments: []Compartment{
			Compartment(items[:len(items)/2]),
			Compartment(items[len(items)/2:]),
		},
	}
}
