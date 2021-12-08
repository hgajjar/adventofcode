package day4_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day4"
)

func TestParseInputToPassports(t *testing.T) {
	passports := day4.ParseInputToPassports("../../data/day4_test.txt")

	expectedPassports := []day4.Passport{
		day4.Passport{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"},
		day4.Passport{Iyr: "2013", Ecl: "amb", Cid: "350", Eyr: "2023", Pid: "028048884", Hcl: "#cfa07d", Byr: "1929"},
		day4.Passport{Hcl: "#ae17e1", Iyr: "2013", Eyr: "2024", Ecl: "brn", Pid: "760753108", Byr: "1931", Hgt: "179cm"},
		day4.Passport{Hcl: "#cfa07d", Eyr: "2025", Pid: "166559648", Iyr: "2011", Ecl: "brn", Hgt: "59in"},
	}

	if !comparePasswords(expectedPassports, passports) {
		t.Errorf("Expected passports: %#v, found %#v", expectedPassports, passports)
	}
}

func comparePasswords(slice1 []day4.Passport, slice2 []day4.Passport) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
