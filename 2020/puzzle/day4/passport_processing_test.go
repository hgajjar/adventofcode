package day4_test

import (
	"testing"

	"github.com/hgajjar/adventofcode/puzzle/day4"
)

func TestCountValidPassportsPart1(t *testing.T) {
	passports := []day4.Passport{
		day4.Passport{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"},
		day4.Passport{Iyr: "2013", Ecl: "amb", Cid: "350", Eyr: "2023", Pid: "028048884", Hcl: "#cfa07d", Byr: "1929"},
		day4.Passport{Hcl: "#ae17e1", Iyr: "2013", Eyr: "2024", Ecl: "brn", Pid: "760753108", Byr: "1931", Hgt: "179cm"},
		day4.Passport{Hcl: "#cfa07d", Eyr: "2025", Pid: "166559648", Iyr: "2011", Ecl: "brn", Hgt: "59in"},
	}

	if validPassports := day4.CountValidPassports(passports, day4.ValidationStrategyEmptyness); validPassports != 2 {
		t.Errorf("Expected valid passports are %d, found %d", 2, validPassports)
	}
}

func TestCountValidPassportsPart2(t *testing.T) {
	passports := []day4.Passport{
		day4.Passport{Eyr: "1972", Cid: "100", Hcl: "#18171d", Ecl: "amb", Hgt: "170", Pid: "186cm", Iyr: "2018", Byr: "1926"},
		day4.Passport{Iyr: "2019", Hcl: "#602927", Eyr: "1967", Hgt: "170cm", Ecl: "grn", Pid: "012533040", Byr: "1946"},
		day4.Passport{Hcl: "dab227", Iyr: "2012", Ecl: "brn", Hgt: "182cm", Pid: "021572410", Eyr: "2020", Byr: "1992", Cid: "277"},
		day4.Passport{Hgt: "59cm", Ecl: "zzz", Eyr: "2038", Hcl: "74454a", Iyr: "2023", Pid: "3556412378", Byr: "2007"},
		day4.Passport{Pid: "087499704", Hgt: "74in", Ecl: "grn", Iyr: "2012", Eyr: "2030", Byr: "1980", Hcl: "#623a2f"},
		day4.Passport{Eyr: "2029", Ecl: "blu", Cid: "129", Byr: "1989", Iyr: "2014", Pid: "896056539", Hcl: "#a97842", Hgt: "165cm"},
		day4.Passport{Hcl: "#888785", Hgt: "164cm", Byr: "2001", Iyr: "2015", Cid: "88", Pid: "545766238", Ecl: "hzl", Eyr: "2022"},
	}
	if validPassports := day4.CountValidPassports(passports, day4.ValidationStrategyStrict); validPassports != 3 {
		t.Errorf("Expected valid passports are %d, found %d", 3, validPassports)
	}
}
