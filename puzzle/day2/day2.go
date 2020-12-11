package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// OldStrategy checks if the password contains give char between given limits
// Example:
// input: "1-3 a: abcde"
// 1 is min limit, 3 is max limit, "a" is match character, and "abcde" is password
// it checks if the password "abcde" contains "a" min 1 time or max 3 times
const OldStrategy = "OldStrategy"

// NewStrategy checks if the password contains give char at either of specific positions
// Example:
// input: "1-3 a: abcde"
// 1 is first position, 3 is second position, "a" is match character, and "abcde" is password
// it checks if the password "abcde" contains "a" at either first or second position
const NewStrategy = "NewStrategy"

// ValidatePasswordsByStrategy validates given password by strategy
// and returns the total number of passwords which are valid as per their policy
func ValidatePasswordsByStrategy(passwordsWithPolicy []string, strategy string) int {
	regex := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)")
	validPasswordsCount := 0

	for _, passwordWithPolicy := range passwordsWithPolicy {
		if validate(passwordWithPolicy, regex, strategy) {
			validPasswordsCount++
		}
	}

	return validPasswordsCount
}

func validate(passwordWithPolicy string, regex *regexp.Regexp, strategy string) bool {
	if strategy == OldStrategy {
		return validateByOldStrategy(passwordWithPolicy, regex)
	}

	return validateByNewStrategy(passwordWithPolicy, regex)
}

func validateByOldStrategy(passwordWithPolicy string, regex *regexp.Regexp) bool {
	min, max, matchChar, password := extract(passwordWithPolicy, regex)

	matchCount := strings.Count(password, matchChar)

	if matchCount >= min && matchCount <= max {
		return true
	}

	return false
}

func validateByNewStrategy(passwordWithPolicy string, regex *regexp.Regexp) bool {
	position1, position2, matchChar, password := extract(passwordWithPolicy, regex)

	return (string(password[position1-1]) == matchChar) != (string(password[position2-1]) == matchChar)
}

func extract(passwordWithPolicy string, regex *regexp.Regexp) (min int, max int, matchChar string, password string) {
	if !regex.MatchString(passwordWithPolicy) {
		panic(fmt.Sprintf("%s does not match with regex", passwordWithPolicy))
	}

	parts := regex.FindStringSubmatch(passwordWithPolicy)
	min, _ = strconv.Atoi(parts[1])
	max, _ = strconv.Atoi(parts[2])
	matchChar = parts[3]
	password = parts[4]
	return
}
