package lib2019

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
		x := tt
		t.Run(fmt.Sprintf("input %v", x.input), func(t *testing.T) {
			t.Parallel()
			p, err := FromString(x.input)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", x.input, err)
			}

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			if !equal(x.output, p.data) {
				t.Fatalf("Invalid data, \n\texpected '%#v'\n\t     got '%#v", x.output, p.data)
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

func TestCodeMode(t *testing.T) {
	tests := []struct {
		input  string
		output []int
	}{
		{"1002,4,3,4,33", []int{1002, 4, 3, 4, 99}},
		{"1101,100,-1,4,0", []int{1101, 100, -1, 4, 99}},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("input %v", x.input), func(t *testing.T) {
			t.Parallel()
			p, err := FromString(x.input)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", x.input, err)
			}

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			if !equal(x.output, p.data) {
				t.Fatalf("Invalid data, \n\texpected '%#v'\n\t     got '%#v", x.output, p.data)
			}
		})
	}
}

func TestInputEqual(t *testing.T) {
	tests := []struct {
		code   string
		input  int
		output int
	}{
		{"3,9,8,9,10,9,4,9,99,-1,8", 1, 0},
		{"3,9,8,9,10,9,4,9,99,-1,8", 8, 1},
		{"3,3,1108,-1,8,3,4,3,99", 1, 0},
		{"3,3,1108,-1,8,3,4,3,99", 8, 1},
	}
	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("input %v", x.input), func(t *testing.T) {
			t.Parallel()
			p, err := FromString(x.code)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", x.input, err)
			}
			p.AddInput(x.input)

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			tmp := p.GetOutputs()
			out := tmp[0]
			if x.output != out {
				t.Errorf("wrong output, expected '%v', got '%v'", x.output, out)
			}
		})
	}
}

func TestInputLessThan(t *testing.T) {
	tests := []struct {
		code   string
		input  int
		output int
	}{
		{"3,9,7,9,10,9,4,9,99,-1,8", 1, 1},
		{"3,9,7,9,10,9,4,9,99,-1,8", 8, 0},
		{"3,9,7,9,10,9,4,9,99,-1,8", 9, 0},
		{"3,3,1107,-1,8,3,4,3,99", 1, 1},
		{"3,3,1107,-1,8,3,4,3,99", 8, 0},
		{"3,3,1107,-1,8,3,4,3,99", 9, 0},
	}
	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("input code %v input %v", x.code, x.input), func(t *testing.T) {
			// t.Parallel()
			p, err := FromString(x.code)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", x.input, err)
			}
			p.AddInput(x.input)

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			tmp := p.GetOutputs()
			out := tmp[0]
			if x.output != out {
				t.Errorf("wrong output, expected '%v', got '%v'", x.output, out)
			}
		})
	}
}

func TestJumpTests(t *testing.T) {
	tests := []struct {
		code   string
		input  int
		output int
	}{
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0, 0},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 1, 1},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", -1, 1},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 9, 1},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, 0},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 1, 1},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", -1, 1},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 9, 1},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("input_code%v_input%v", x.code, x.input), func(t *testing.T) {
			// 	t.Parallel()
			p, err := FromString(x.code)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", x.input, err)
			}
			p.AddInput(x.input)

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			tmp := p.GetOutputs()
			out := tmp[0]
			if out != x.output {
				t.Errorf("wrong output, expected '%v', got '%v'", x.output, out)
			}
		})
	}
}

func TestJumpLarge(t *testing.T) {
	program := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

	tests := []struct {
		code   string
		input  int
		output int
	}{
		{program, 1, 999},
		{program, 7, 999},
		{program, 8, 1000},
		{program, 9, 1001},
	}
	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("input %v", x.input), func(t *testing.T) {
			t.Parallel()
			p, err := FromString(x.code)
			if err != nil {
				t.Fatalf("unable to create program with input '%v', error: %v", x.input, err)
			}
			p.AddInput(x.input)

			err = p.Run()
			if err != nil {
				t.Fatalf("failure while running program: %v", err)
			}

			tmp := p.GetOutputs()
			out := tmp[0]
			if x.output != out {
				t.Errorf("wrong output, expected '%v', got '%v'", x.output, out)
			}
		})
	}
}
