package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetGroupChoices(t *testing.T) {
	testcases := []struct {
		Description string
		Input       string
		Expected    map[rune]struct{}
	}{
		{
			Description: "One person, three questions",
			Input:       `abc`,
			Expected:    map[rune]struct{}{'a': {}, 'b': {}, 'c': {}},
		},
		{
			Description: "three people, three questions",
			Input: `a
b
c`,
			Expected: map[rune]struct{}{'a': {}, 'b': {}, 'c': {}},
		},
		{
			Description: "two people, three questions",
			Input: `ab
ac`,
			Expected: map[rune]struct{}{'a': {}, 'b': {}, 'c': {}},
		},
		{
			Description: "four people, one question",
			Input: `a
a
a
a`,
			Expected: map[rune]struct{}{'a': {}},
		},

		{
			Description: "one person, one question",
			Input:       `b`,
			Expected:    map[rune]struct{}{'b': {}},
		},
	}

	for _, test := range testcases {
		t.Run(test.Description, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.Expected, extractCommonYesUnion(test.Input))
		})
	}
}

func TestExtractCommonYesUnion(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	assert.Equal(t, 11, getResult(input, extractCommonYesUnion))
}

func TestGetQuestionPartTwo(t *testing.T) {
	testcases := []struct {
		Description string
		Input       string
		Expected    map[rune]struct{}
	}{
		{
			Description: "One person, three questions",
			Input:       `abc`,
			Expected:    map[rune]struct{}{'a': {}, 'b': {}, 'c': {}},
		},
		{
			Description: "three people, three questions",
			Input: `a
b
c`,
			Expected: map[rune]struct{}{},
		},
		{
			Description: "two people, three questions",
			Input: `ab
ac`,
			Expected: map[rune]struct{}{'a': {}},
		},
		{
			Description: "four people, one question",
			Input: `a
a
a
a`,
			Expected: map[rune]struct{}{'a': {}},
		},

		{
			Description: "one person, one question",
			Input:       `b`,
			Expected:    map[rune]struct{}{'b': {}},
		},
	}

	for _, test := range testcases {
		t.Run(test.Description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, extractCommonYesIntersection(test.Input))
		})
	}
}

func TestExtractCommonYesIntersection(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	assert.Equal(t, 6, getResult(input, extractCommonYesIntersection))
}
