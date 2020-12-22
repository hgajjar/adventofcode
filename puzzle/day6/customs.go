package day6

import (
	"sync"
)

type Group struct {
	Persons []Person
}

type Person struct {
	FilledForm CustomsForm
}

type CustomsForm struct {
	// key: question, value: answer. Ex: 'a' => true
	QuestionAnswers map[string]bool
}

// CountAnswersWhereAnyoneInGroupSaidYes counts unique answers per group and returns sum of them
func CountAnswersWhereAnyoneInGroupSaidYes(groups []Group) int {
	groupCount := make(chan int)
	total := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(groups))

	for _, g := range groups {
		go func(g Group) {
			defer wg.Done()
			groupCount <- g.countUniqueAnswers()
		}(g)
	}

	go sumCounts(groupCount, total, len(groups))

	wg.Wait()

	return <-total
}

// CountAnswersWhereEveryoneInGroupSaidYes counts answers per group where everyone said Yes and returns sum of them
func CountAnswersWhereEveryoneInGroupSaidYes(groups []Group) int {
	groupCount := make(chan int)
	total := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(groups))

	for _, g := range groups {
		go func(g Group) {
			wg.Done()
			groupCount <- g.countCommonAnswers()
		}(g)
	}

	go sumCounts(groupCount, total, len(groups))

	wg.Wait()

	return <-total
}

func (g *Group) countUniqueAnswers() int {
	uniqueAnswers := make(map[string]bool)
	for _, p := range g.Persons {
		for k, v := range p.FilledForm.QuestionAnswers {
			uniqueAnswers[k] = v
		}
	}

	return len(uniqueAnswers)
}

func (g *Group) countCommonAnswers() int {
	answers := make(map[string]int)
	for _, p := range g.Persons {
		for k := range p.FilledForm.QuestionAnswers {
			if _, ok := answers[k]; !ok {
				answers[k] = 0
			}
			answers[k] = answers[k] + 1
		}
	}

	commonAnswersCount := 0
	for _, v := range answers {
		if v >= len(g.Persons) {
			commonAnswersCount++
		}
	}

	return commonAnswersCount
}

func sumCounts(groupCount chan int, sendTotal chan int, iterations int) {
	total := 0
	for i := 0; i < iterations; i++ {
		total += <-groupCount
	}
	sendTotal <- total
}
