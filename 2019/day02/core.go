package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/seanhagen/advent-of-code/lib"
)

const (
	OP_ADD = 1
	OP_MUL = 2
	OP_FIN = 99
)

var opIncr = map[int]int{
	OP_ADD: 4,
	OP_MUL: 4,
	OP_FIN: 1,
}

type Program struct {
	data []int
}

func FromString(in string) (*Program, error) {
	data := strings.Split(in, ",")
	prog := Program{data: []int{}}
	for _, v := range data {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		prog.data = append(prog.data, i)
	}
	return &prog, nil
}

func ReadProgram(f *os.File) (*Program, error) {
	d, err := lib.ReadLine(f)
	if err != nil {
		return nil, err
	}
	return FromString(string(d))
}

// Run ...
func (p Program) Run() error {
	pos := 0
	max := len(p.data) - 1

	for {
		op := p.data[pos]
		incr := opIncr[op]
		if max < (incr - 1) {
			return fmt.Errorf("not enough data to continue (pos '%v', max '%v')", pos, max)
		}

		switch op {
		case OP_ADD:
			a := p.data[pos+1]
			b := p.data[pos+2]

			x := p.data[a]
			y := p.data[b]

			l := p.data[pos+3]
			p.data[l] = x + y

		case OP_MUL:
			a := p.data[pos+1]
			b := p.data[pos+2]

			x := p.data[a]
			y := p.data[b]

			l := p.data[pos+3]

			p.data[l] = x * y

		case OP_FIN:
			goto done

		default:
			return fmt.Errorf("unknown opcode encountered (op: '%v', pos: '%v')", op, pos)
		}
		pos += incr
	}

done:
	return nil
}

// Replace ...
func (p Program) Replace(x, v int) error {
	if x > len(p.data)-1 {
		return fmt.Errorf("pos %v invalid (len %v)", x, len(p.data))
	}

	p.data[x] = v
	return nil
}

// Value ...
func (p Program) Value(x int) (int, error) {
	if x > len(p.data)-1 {
		return 0, fmt.Errorf("pos %v invalid (len %v)", x, len(p.data))
	}

	return p.data[x], nil
}
