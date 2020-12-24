package day6

import (
	"sync"
	"sync/atomic"
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
func CountAnswersWhereAnyoneInGroupSaidYes(groups []Group) uint32 {
	var c uint32
	var wg sync.WaitGroup
	wg.Add(len(groups))

	for _, g := range groups {
		go func(g Group) {
			defer wg.Done()
			atomic.AddUint32(&c, g.countUniqueAnswers())
		}(g)
	}

	wg.Wait()

	return c
}

// CountAnswersWhereEveryoneInGroupSaidYes counts answers per group where everyone said Yes and returns sum of them
func CountAnswersWhereEveryoneInGroupSaidYes(groups []Group) uint32 {
	var c uint32
	var wg sync.WaitGroup
	wg.Add(len(groups))

	for _, g := range groups {
		go func(g Group) {
			defer wg.Done()
			atomic.AddUint32(&c, g.countCommonAnswers())
		}(g)
	}

	wg.Wait()

	return c
}

func (g *Group) countUniqueAnswers() uint32 {
	uniqueAnswers := make(map[string]bool)
	for _, p := range g.Persons {
		for k, v := range p.FilledForm.QuestionAnswers {
			uniqueAnswers[k] = v
		}
	}

	return uint32(len(uniqueAnswers))
}

func (g *Group) countCommonAnswers() uint32 {
	answers := make(map[string]int)
	for _, p := range g.Persons {
		for k := range p.FilledForm.QuestionAnswers {
			if _, ok := answers[k]; !ok {
				answers[k] = 0
			}
			answers[k] = answers[k] + 1
		}
	}

	var commonAnswersCount uint32 = 0
	for _, v := range answers {
		if v >= len(g.Persons) {
			commonAnswersCount++
		}
	}

	return commonAnswersCount
}
