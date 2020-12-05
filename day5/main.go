package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Seat struct {
	Row    int
	Column int
}

func (s *Seat) GetID() int {
	return s.Row*8 + s.Column
}

func (s *Seat) String() string {
	return fmt.Sprintf("Row: %d, Col: %d, ID: %d", s.Row, s.Column, s.GetID())
}

type BinaryPartition struct {
	Min   int
	Lower rune
	Max   int
	Upper rune
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("ERROR: Cannot read file: %v", err)
		return
	}
	passes := strings.Split(string(data), "\n")

	seats := getSeats(passes)
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].GetID() > seats[j].GetID()
	})

	fmt.Printf("Part One result: %s\n", seats[0].String())

	for i := 0; i < len(seats)-2; i++ {
		if seats[i+1].GetID() != (seats[i].GetID() - 1) {
			fmt.Printf("Bingo, Part Two, seat between:\n %s and\n %s\n", seats[i].String(), seats[i+1].String())
			return
		}
	}
}

func getSeats(passes []string) []Seat {
	var seats []Seat
	for _, pass := range passes {
		if len(pass) != 10 {
			continue
		}
		seats = append(seats, getSeat(pass))
	}

	return seats
}

func getSeat(pass string) Seat {
	return Seat{
		Row: getValue(pass[:7], BinaryPartition{
			Min:   0,
			Max:   127,
			Lower: 'F',
			Upper: 'B',
		}),
		Column: getValue(pass[:7], BinaryPartition{
			Min:   0,
			Max:   127,
			Lower: 'F',
			Upper: 'B',
		}),
	}
}

func getValue(rowCode string, partition BinaryPartition) int {
	min := partition.Min
	max := partition.Max
	for _, c := range rowCode {
		switch c {
		case partition.Lower:
			max = min + ((max - min) / 2)
		case partition.Upper:
			min = min + (((max - min) / 2) + 1)
		default:
			fmt.Printf("ERROR: Unknown code %c\n", c)
		}
	}

	return max
}
