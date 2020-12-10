package day5

import (
	"strings"
)

const (
	Front = "F"
	Back  = "B"
	Left  = "L"
	Right = "R"
)

var valid = []string{Front, Back, Left, Right}

var rows, seats []int

func init() {
	for i := 0; i <= 127; i++ {
		rows = append(rows, i)
	}
	for i := 0; i <= 7; i++ {
		seats = append(seats, i)
	}
}

func GetSeatID(row, seat int) int {
	return (row * 8) + seat
}

func GetRowSeat(path string) (row, seat int) {
	path = strings.TrimSpace(path)
	if path == "" {
		return
	}

	b := strings.Split(path, "")
	for _, v := range b {
		if !inar(v, valid) {
			return
		}
	}

	rowpath := strings.Join(b[:7], "")
	seatpath := strings.Join(b[7:], "")
	// fmt.Printf("row path: %v\nseat path: %v\n\n", rowpath, seatpath)

	r := choose(rows, rowpath)
	if len(r) >= 1 {
		row = r[0]
	}

	s := choose(seats, seatpath)
	if len(s) >= 1 {
		seat = s[0]
	}

	return
}

func inar(v string, ar []string) bool {
	f := false
	for _, x := range ar {
		if x == v {
			f = true
		}
	}
	return f
}

func choose(in []int, path string) []int {
	if len(path) <= 0 {
		return in
	}

	if len(in) == 1 {
		return in
	}

	x := strings.Split(path, "")
	choice := x[0]

	l := len(in)
	p := l / 2

	a, b := in[:p], in[p:]

	if len(x) == 1 {
		switch choice {
		case Left:
			fallthrough
		case Front:
			return a

		case Right:
			fallthrough
		case Back:
			return b
		default:
			return []int{}
		}
	}

	var half []int
	if choice == Front || choice == Left {
		half = a
	} else {
		half = b
	}

	rest := strings.Join(x[1:], "")
	if rest == "" {
		return half
	}

	// fmt.Printf("input: %v\na: %v\nb: %v\n", in, a, b)
	// fmt.Printf("path: %v\n", path)
	// fmt.Printf("choice: %v\nrest: %v\n-------\n", choice, rest)

	//spew.Dump(l, p, in, a, b, rest)

	return choose(half, rest)
}
