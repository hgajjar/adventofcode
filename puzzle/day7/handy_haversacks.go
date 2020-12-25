package day7

import (
	"errors"
	"sync"
	"time"
)

type Bag struct {
	Color string
	Bags  []Bag
}

func (searchBag *Bag) CountOuterBags(allBags []Bag) int {
	bagChan := make(chan map[string]Bag, len(allBags))
	countChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(allBags))

	searchBag.findOuterBags(allBags, bagChan, &wg)
	go countBags(bagChan, countChan)

	wg.Wait()
	close(bagChan)

	return <-countChan
}

func (searchBag *Bag) findOuterBags(allBags []Bag, c chan map[string]Bag, wg *sync.WaitGroup) {
	for _, bag := range allBags {
		if wg != nil {
			go func(bag Bag, c chan map[string]Bag, wg *sync.WaitGroup) {
				defer wg.Done()

				if _, err := searchBag.isWithinBags(bag.Bags); err == nil {
					c <- map[string]Bag{bag.Color: bag}
					bag.findOuterBags(allBags, c, nil)
				}
			}(bag, c, wg)
		} else {
			if _, err := searchBag.isWithinBags(bag.Bags); err == nil {
				c <- map[string]Bag{bag.Color: bag}
				bag.findOuterBags(allBags, c, nil)
			}
		}
	}
}

func countBags(bagChan chan map[string]Bag, countChan chan int) {
	uniqueBags := make(map[string]Bag)
	for bagMap := range bagChan {
		for k, v := range bagMap {
			uniqueBags[k] = v
		}
	}

	countChan <- len(uniqueBags)
}

func (searchBag *Bag) isWithinBags(bags []Bag) (Bag, error) {
	time.Sleep(time.Nanosecond * 5)
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
