package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeat(t *testing.T) {
	testcases := []struct {
		Desc         string
		Input        string
		ExpectedSeat Seat
		ExpectedId   int
	}{
		{
			Desc:         "Valid 1",
			Input:        "BFFFBBFRRR",
			ExpectedSeat: Seat{70, 7},
			ExpectedId:   567,
		},
		{
			Desc:         "Valid 2",
			Input:        "FFFBBBFRRR",
			ExpectedSeat: Seat{14, 7},
			ExpectedId:   119,
		},
		{
			Desc:         "Valid 3",
			Input:        "BBFFBBFRLL",
			ExpectedSeat: Seat{102, 4},
			ExpectedId:   820,
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.Desc, func(t *testing.T) {
			t.Parallel()
			seat := getSeat(test.Input)
			assert.Equal(t, test.ExpectedSeat, seat)
			assert.Equal(t, test.ExpectedId, seat.GetID())
		})
	}
}

func TestName(t *testing.T) {
	toto := "FBFBBFFRLR"
	fmt.Printf("%s:%s", toto[:7], toto[7:])
}
func TestGetRow(t *testing.T) {
	testcases := []struct {
		Desc     string
		Input    string
		Expected int
	}{
		{
			Desc:     "val1",
			Input:    "BFFFBBF",
			Expected: 70,
		},
		{
			Desc:     "val2",
			Input:    "FBFBBFF",
			Expected: 14,
		},

		{
			Desc:     "val3",
			Input:    "BBFFBBFRLL",
			Expected: 102,
		},
	}

	for _, test := range testcases {
		t.Run(test.Desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, getValue(test.Input, BinaryPartition{
				Min:   0,
				Max:   127,
				Lower: 'F',
				Upper: 'B',
			}))
		})
	}

	assert.Equal(t, 44, getValue("FBFBBFF", BinaryPartition{
		Min:   0,
		Max:   127,
		Lower: 'F',
		Upper: 'B',
	}))
}

func TestGetCol(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Desc     string
		Input    string
		Expected int
	}{
		{
			Desc:     "val1",
			Input:    "RRR",
			Expected: 7,
		},
		{
			Desc:     "val2",
			Input:    "RLL",
			Expected: 4,
		},
	}

	for _, test := range testcases {
		t.Run(test.Desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expected, getValue(test.Input, BinaryPartition{
				Min:   0,
				Max:   7,
				Lower: 'L',
				Upper: 'R',
			}))
		})
	}
}
