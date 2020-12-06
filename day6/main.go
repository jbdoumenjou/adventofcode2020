package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("ERROR: Cannot open file %v\n", err)
		return
	}
	count := getResult(string(data), extractCommonYesUnion)
	fmt.Printf("Part One Count : %d\n", count)

	count = getResult(string(data), extractCommonYesIntersection)
	fmt.Printf("Part Two Count : %d\n", count)
}

func extractCommonYesUnion(group string) map[rune]struct{} {
	result := map[rune]struct{}{}

	for _, r := range group {
		if r >= 'a' && r <= 'z' {
			result[r] = struct{}{}
		}
	}

	return result
}

func extractCommonYesIntersection(group string) map[rune]struct{} {
	result := map[rune]struct{}{}

	people := strings.Split(group, "\n")
	if len(people) == 0 {
		return result
	}

	// reference map
	for _, c := range people[0] {
		result[c] = struct{}{}
	}

	for _, person := range people {
		if len(person) == 0 {
			continue
		}

		for c, _ := range result {
			if !strings.ContainsRune(person, c) {
				delete(result, c)
			}
		}
	}

	return result
}

func getResult(input string, extract func(string) map[rune]struct{}) int {
	groups := strings.Split(input, "\n\n")
	count := 0

	for _, group := range groups {
		count += len(extract(group))
	}

	return count
}
