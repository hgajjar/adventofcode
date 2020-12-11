package day4

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Passport struct holds the information about a passport
type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

// ValidationStrategyEmptyness is used to only check for emptyness of passport fields
const ValidationStrategyEmptyness = "Emptyness"

// ValidationStrategyStrict is used to validate passport fields with strict validation rules
const ValidationStrategyStrict = "Strict"

// CountValidPassports returns the total number of valid passports
// Validation Rule: all fields except cid are required
func CountValidPassports(passports []Passport, validationStrategy string) int {
	validPassports := 0

	for _, passport := range passports {
		if passport.isValid(validationStrategy) {
			validPassports++
		}
	}

	return validPassports
}

func (passport *Passport) isValid(strategy string) bool {
	v := reflect.Indirect(reflect.ValueOf(passport))
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == "Cid" {
			continue
		}

		if v.Field(i).Interface() == reflect.Zero(t.Field(i).Type).Interface() {
			return false
		}

		if strategy == ValidationStrategyStrict {
			if !validateFieldValue(t.Field(i).Name, v.Field(i).Interface().(string)) {
				return false
			}
		}
	}

	return true
}

func validateFieldValue(name string, value string) bool {
	switch name {
	case "Byr":
		return isValidInt(value, 1920, 2002)
	case "Iyr":
		return isValidInt(value, 2010, 2020)
	case "Eyr":
		return isValidInt(value, 2020, 2030)
	case "Hgt":
		return isValidHeightByUnit(value, "cm", 150, 193) || isValidHeightByUnit(value, "in", 59, 76)
	case "Hcl":
		regex := regexp.MustCompile("^#[0-9a-f]{6}$")
		return regex.MatchString(value)
	case "Ecl":
		validColors := map[string]bool{
			"amb": true,
			"blu": true,
			"brn": true,
			"gry": true,
			"grn": true,
			"hzl": true,
			"oth": true,
		}
		return validColors[value]
	case "Pid":
		regex := regexp.MustCompile("^[0-9]{9}$")
		return regex.MatchString(value)
	case "Cid":
		return true
	}

	return false
}

func isValidInt(value string, min int, max int) bool {
	v, _ := strconv.Atoi(value)
	return v >= min && v <= max
}

func isValidHeightByUnit(height string, unit string, min int, max int) bool {
	parts := strings.Split(height, unit)
	if len(parts) != 2 {
		return false
	}

	heightInInt, err := strconv.Atoi(parts[0])

	if err != nil {
		return false
	}

	return heightInInt >= min && heightInInt <= max
}
