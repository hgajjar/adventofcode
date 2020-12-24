package day7

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func ParseInputToBags(fileName string) []Bag {
	content, _ := ioutil.ReadFile(fileName)
	//dim teal bags contain 2 posh tan bags, 1 striped cyan bag, 2 muted lavender bags, 5 wavy lime bags.
	regex := regexp.MustCompile(" (\\d) ([a-z]+ [a-z]+) bags?.?")
	bags := []Bag{}

	for _, bagStr := range extractBag(content) {
		parts := strings.Split(bagStr, " bags contain")
		bag := Bag{Color: parts[0]}

		childBagsParts := strings.Split(parts[1], ",")
		for _, childBagPart := range childBagsParts {
			if !regex.MatchString(childBagPart) {
				continue
			}

			childBagMatch := regex.FindStringSubmatch(childBagPart)
			bag.Bags = append(bag.Bags, Bag{Color: childBagMatch[2]})
		}

		bags = append(bags, bag)
	}

	return bags
}

func extractBag(content []byte) []string {
	return strings.Split(string(content), "\n")
}
