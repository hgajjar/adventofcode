package day4

import (
	"reflect"
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

// CountValidPassports returns the total number of valid passports
// Validation Rule: all fields except cid are required
func CountValidPassports(passports []Passport) int {
	validPassports := 0

	for _, passport := range passports {
		if passport.isValid() {
			validPassports++
		}
	}

	return validPassports
}

func (passport *Passport) isValid() bool {
	v := reflect.Indirect(reflect.ValueOf(passport))
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == "Cid" {
			continue
		}

		if v.Field(i).Interface() == reflect.Zero(t.Field(i).Type).Interface() {
			return false
		}
	}

	return true
}
