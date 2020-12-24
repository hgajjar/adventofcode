package day7

import (
	"errors"
)

type Bag struct {
	Color string
	Bags  []Bag
}

func (searchBag *Bag) CountOuterBags(allBags []Bag) int {
	return len(searchBag.findOuterBags(allBags))
}

func (searchBag *Bag) findOuterBags(allBags []Bag) map[string]Bag {
	fbs := map[string]Bag{}
	for _, bag := range allBags {
		if _, err := searchBag.isWithinBags(bag.Bags); err == nil {
			fbs[bag.Color] = bag
			for color, fcb := range bag.findOuterBags(allBags) {
				fbs[color] = fcb
			}
		}
	}

	return fbs
}

func (searchBag *Bag) isWithinBags(bags []Bag) (Bag, error) {
	for _, b := range bags {
		if b.Color == searchBag.Color {
			return Bag{}, nil
		}

		if b.Bags != nil {
			if _, err := searchBag.isWithinBags(b.Bags); err == nil {
				return b, nil
			}
		}
	}

	return Bag{}, errors.New("Bag not found withing given bags")
}
