package day02

import (
	"fmt"
	"testing"

	"github.com/seanhagen/advent-of-code/lib"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestReadProgram(t *testing.T) {
	f := lib.LoadInput("./test_input.txt")
	prog, err := ReadProgram(f)
	if err != nil {
		t.Errorf("unable to load program: %v\n", err)
	}

	if prog == nil {
		t.Fatalf("no error, but program nil?")
	}

	data := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	if !equal(data, prog.data) {
		t.Errorf("Invalid data, \n\texpected '%#v'\n\t     got '%#v", data, prog.data)
	}

	err = prog.Run()
	if err != nil {
		t.Fatalf("unexpected error running program: %v\n", err)
	}

	data = []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	if !equal(data, prog.data) {
		t.Errorf("Invalid data, \n\texpected '%#v'\n\t     got '%#v", data, prog.data)
	}
}

func TestPrograms(t *testing.T) {
	tests := []struct {
		input  string
		output []int
	}{
		{"1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		{"1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()
			p, err := FromString(tt.input)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", tt.input, err)
			}

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			if !equal(tt.output, p.data) {
				t.Fatalf("Invalid data, \n\texpected '%#v'\n\t     got '%#v", tt.output, p.data)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	start := 1
	expect := 2
	pos := 0
	p := Program{data: []int{start}}
	p.Replace(pos, expect)
	if p.data[pos] != expect {
		t.Errorf("invalid, p.data[%v] has value '%v', expected '%v'", pos, p.data[pos], expect)
	}
}
