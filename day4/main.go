package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type validator func(string) bool

var (
	requiredKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	eclColor     = map[string]struct{}{"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn": {}, "hzl": {}, "oth": {}}
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Cannot read file: %v", err)
	}

	docs := strings.Split(string(data), "\n\n")

	validPass := getValidPassports(docs, isValidPartOne)
	fmt.Printf("Result: %d\n", len(validPass))

	result := validPassportPart2(validPass)
	fmt.Printf("Result %d\n", len(result))
}

func validPassportPart2(validPass []string) []string {
	var result []string
	validators := map[string]validator{
		"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
	}

	for _, passport := range validPass {
		fields := getFields(passport)

		if areFieldsValid(fields, validators) {
			result = append(result, passport)
		}
	}

	return result
}

func getFields(s string) map[string]string {
	fields := strings.Split(s, " ")

	data := map[string]string{}
	for _, field := range fields {
		f := strings.Split(field, ":")
		if len(f) != 2 {
			continue
		}
		data[f[0]] = f[1]
	}

	return data
}

func getValidPassports(docs []string, isValid func(s string) bool) []string {
	var validPass []string

	for _, doc := range docs {
		content := strings.ReplaceAll(doc, "\n", " ")
		if isValid(content) {
			validPass = append(validPass, content)
		}
	}

	return validPass
}

func isValidPartOne(s string) bool {
	for _, key := range requiredKeys {
		if !strings.Contains(s, key) {
			return false
		}
	}

	return true
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func byr(s string) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return n >= 1920 && n <= 2002
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func iyr(s string) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return n >= 2010 && n <= 2020
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030
func eyr(s string) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return n >= 2020 && n <= 2030
}

// hgt (Height) - a number followed by either cm or in:
//    If cm, the number must be at least 150 and at most 193.
//    If in, the number must be at least 59 and at most 76.
func hgt(s string) bool {
	unit := s[len(s)-2:]
	if unit != "cm" && unit != "in" {
		return false
	}

	n, err := strconv.Atoi(s[:len(s)-2])
	if err != nil {
		return false
	}

	if unit == "cm" {
		return n >= 150 && n <= 193
	}

	return n >= 59 && n <= 76
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func hcl(s string) bool {
	match, err := regexp.MatchString(`^#[0-9a-f]{6}$`, s)
	if err != nil {
		return false
	}

	return match
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func ecl(s string) bool {
	_, ok := eclColor[s]
	return ok
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func pid(s string) bool {
	match, err := regexp.MatchString(`^\d{9}$`, s)
	if err != nil {
		return false
	}

	return match
}

func areFieldsValid(fields map[string]string, validators map[string]validator) bool {
	for key, isValid := range validators {
		value, ok := fields[key]
		if !ok || !isValid(value) {
			return false
		}
	}

	return true
}
