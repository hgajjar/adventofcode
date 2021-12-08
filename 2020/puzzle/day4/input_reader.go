package day4

import (
	"io/ioutil"
	"reflect"
	"strings"
)

// ParseInputToPassports reads the input file and parses it into array of Passport
func ParseInputToPassports(fileName string) []Passport {
	content, _ := ioutil.ReadFile(fileName)
	passportStrings := extractPasswordStrings(content)
	passports := []Passport{}

	for _, passportString := range passportStrings {
		passportFields := extractPassportFields(passportString)
		passport := createPassport(passportFields)
		passports = append(passports, passport)
	}

	return passports
}

func extractPasswordStrings(content []byte) []string {
	return strings.Split(string(content), "\n\n")
}

func extractPassportFields(passportString string) []string {
	passportFields := []string{}
	stringsWithSpaces := strings.Split(passportString, "\n")

	for _, stringWithSpaces := range stringsWithSpaces {
		passportFields = append(passportFields, strings.Split(stringWithSpaces, " ")...)
	}

	return passportFields
}

func createPassport(passportFields []string) Passport {
	passport := Passport{}
	v := reflect.ValueOf(&passport).Elem()

	for _, passportField := range passportFields {
		fieldArray := strings.Split(passportField, ":")
		field := v.FieldByName(strings.Title(fieldArray[0]))
		if field.IsValid() {
			field.SetString(fieldArray[1])
		}
	}

	return passport
}
