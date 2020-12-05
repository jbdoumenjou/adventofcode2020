package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByr(t *testing.T) {
	testcases := []struct {
		Desc     string
		Input    string
		Expected bool
	}{
		{
			Desc:     "valid upper limit",
			Input:    "2002",
			Expected: true,
		},
		{
			Desc:     "valid lower limit",
			Input:    "1920",
			Expected: true,
		},
		{
			Desc:     "too high",
			Input:    "2003",
			Expected: false,
		},
		{
			Desc:     "too low",
			Input:    "1919",
			Expected: false,
		},
		{
			Desc:     "not even a number",
			Input:    "19b19",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.Desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, byr(test.Input))
		})
	}
}

func TestIyr(t *testing.T) {
	testcases := []struct {
		desc     string
		Input    string
		Expected bool
	}{
		{
			desc:     "valid upper limit",
			Input:    "2020",
			Expected: true,
		},
		{
			desc:     "valid lower limit",
			Input:    "2010",
			Expected: true,
		},
		{
			desc:     "too high",
			Input:    "2021",
			Expected: false,
		},
		{
			desc:     "too low",
			Input:    "2009",
			Expected: false,
		},
		{
			desc:     "not even a number",
			Input:    "19b19",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, iyr(test.Input))
		})
	}
}

func TestEyr(t *testing.T) {
	testcases := []struct {
		desc     string
		Input    string
		Expected bool
	}{
		{
			desc:     "valid upper limit",
			Input:    "2030",
			Expected: true,
		},
		{
			desc:     "valid lower limit",
			Input:    "2020",
			Expected: true,
		},
		{
			desc:     "too high",
			Input:    "2031",
			Expected: false,
		},
		{
			desc:     "too low",
			Input:    "2019",
			Expected: false,
		},
		{
			desc:     "not even a number",
			Input:    "19b19",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, eyr(test.Input))
		})
	}
}

func TestHgt(t *testing.T) {
	testcases := []struct {
		desc     string
		Input    string
		Expected bool
	}{
		{
			desc:     "valid cm upper limit",
			Input:    "193cm",
			Expected: true,
		},
		{
			desc:     "valid cm lower limit",
			Input:    "150cm",
			Expected: true,
		},
		{
			desc:     "valid in upper limit",
			Input:    "76in",
			Expected: true,
		},
		{
			desc:     "valid cm lower limit",
			Input:    "59in",
			Expected: true,
		},
		{
			desc:     "too high cm",
			Input:    "194cm",
			Expected: false,
		},
		{
			desc:     "too low cm",
			Input:    "149in",
			Expected: false,
		},
		{
			desc:     "too high in",
			Input:    "77in",
			Expected: false,
		},
		{
			desc:     "too low in",
			Input:    "58in",
			Expected: false,
		},
		{
			desc:     "not even a number",
			Input:    "19b19cm",
			Expected: false,
		},
		{
			desc:     "not even a valid unit",
			Input:    "19b19im",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, hgt(test.Input))
		})
	}
}

func TestHcl(t *testing.T) {
	testcases := []struct {
		desc     string
		Input    string
		Expected bool
	}{
		{
			desc:     "valid color",
			Input:    "#123abc",
			Expected: true,
		},
		{
			desc:     "too short",
			Input:    "#12b34",
			Expected: false,
		},
		{
			desc:     "too long",
			Input:    "#12b4567",
			Expected: false,
		},
		{
			desc:     "bad start",
			Input:    "1#12b4567",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, hcl(test.Input))
		})
	}
}

func TestEcl(t *testing.T) {
	testcases := []struct {
		desc     string
		Input    string
		Expected bool
	}{
		{
			desc:     "valid amb",
			Input:    "amb",
			Expected: true,
		},
		{
			desc:     "valid blu",
			Input:    "blu",
			Expected: true,
		},
		{
			desc:     "valid brn",
			Input:    "brn",
			Expected: true,
		},
		{
			desc:     "valid gry",
			Input:    "gry",
			Expected: true,
		},
		{
			desc:     "valid grn",
			Input:    "grn",
			Expected: true,
		},
		{
			desc:     "valid hzl",
			Input:    "hzl",
			Expected: true,
		},
		{
			desc:     "valid oth",
			Input:    "oth",
			Expected: true,
		},
		{
			desc:     "other",
			Input:    "other",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, ecl(test.Input))
		})
	}
}

func TestPid(t *testing.T) {
	testcases := []struct {
		desc     string
		Input    string
		Expected bool
	}{
		{
			desc:     "valid number",
			Input:    "123456789",
			Expected: true,
		},
		{
			desc:     "valid number with leading 0",
			Input:    "000000009",
			Expected: true,
		},
		{
			desc:     "too long number",
			Input:    "0000000091",
			Expected: false,
		},
		{
			desc:     "too short number",
			Input:    "0000000091",
			Expected: false,
		},
		{
			desc:     "not number",
			Input:    "00a000091",
			Expected: false,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, pid(test.Input))
		})
	}
}
