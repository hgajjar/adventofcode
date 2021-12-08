package day6

import (
	"io/ioutil"
	"strings"
)

// ParseInput parses input text file into structure of groups, persons and their filled questions
func ParseInput(fileName string) []Group {
	content, _ := ioutil.ReadFile(fileName)
	inputGroups := []Group{}

	for _, groupStr := range extractGroups(content) {
		group := Group{}
		for _, personStr := range extractPersons(groupStr) {
			qa := extractQuestionAnswers(personStr)
			group.Persons = append(group.Persons, Person{CustomsForm{qa}})
		}
		inputGroups = append(inputGroups, group)
	}

	return inputGroups
}

func extractGroups(content []byte) []string {
	return strings.Split(string(content), "\n\n")
}

func extractPersons(groupAnswer string) []string {
	return strings.Split(groupAnswer, "\n")
}

func extractQuestionAnswers(input string) map[string]bool {
	qa := make(map[string]bool)
	for _, q := range input {
		qa[string(q)] = true
	}

	return qa
}
