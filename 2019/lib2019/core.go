package lib2019

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
	OP_SAV = 3
	OP_OUT = 4
	OP_JIT = 5
	OP_JIF = 6
	OP_LT  = 7
	OP_EQ  = 8
	OP_FIN = 99
)

var opIncr = map[int]int{
	OP_ADD: 4,
	OP_MUL: 4,
	OP_SAV: 2,
	OP_OUT: 2,
	OP_JIT: 3,
	OP_JIF: 3,
	OP_LT:  4,
	OP_EQ:  4,
	OP_FIN: 1,
}

type Program struct {
	code string
	data []int
	mode int

	inPtr  int
	inputs []int

	outputs []int
}

func FromString(in string) (*Program, error) {
	data := strings.Split(in, ",")
	prog := Program{code: in, data: []int{}, inPtr: 0, inputs: []int{}, outputs: []int{}}
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

// AddInput ...
func (p *Program) AddInput(in int) {
	p.inputs = append(p.inputs, in)
}

// GetOutput ...
func (p Program) GetOutputs() []int {
	return p.outputs
}

// Run ...
func (p *Program) Run() error {
	pos := 0
	max := len(p.data) - 1

	for {
		op := p.data[pos]
		opc := strconv.Itoa(op)
		bits := strings.Split(opc, "")

		pDE, pC, pB, pA := "0", "0", "0", "0"
		if len(bits) == 1 {
			pDE = bits[0]
		} else if len(bits) >= 2 {
			pDE = bits[len(bits)-2] + bits[len(bits)-1]
		}

		if len(bits) >= 3 {
			pC = bits[len(bits)-3]
		}

		if len(bits) >= 4 {
			pB = bits[len(bits)-4]
		}

		if len(bits) >= 5 {
			pA = bits[len(bits)-5]
		}

		op, err := strconv.Atoi(pDE)
		if err != nil {
			return fmt.Errorf("unable to parse opcode '%v', reason: %v", pDE, err)
		}

		incr := opIncr[op]
		if max < (pos + incr - 1) {
			return fmt.Errorf("not enough data to continue (pos '%v', max '%v')", pos, max)
		}

		switch op {
		case OP_ADD:
			var x, y int

			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			}

			z := p.data[pos+3]
			p.data[z] = x + y

		case OP_MUL:
			var x, y int
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			}

			z := p.data[pos+3]
			p.data[l] = x * y

		case OP_SAV:
			a := p.data[pos+1]
			p.data[a] = p.inputs[p.inPtr]
			p.inPtr++

		case OP_OUT:
			switch pC {
			case "0":
				a := p.data[pos+1]
				p.outputs = append(p.outputs, p.data[a])
			case "1":
				p.outputs = append(p.outputs, p.data[pos+1])
			}

		case OP_JIT:
			x, y := 0, 0
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			}

			if x != 0 {
				switch pB {
				case "0":
					b := p.data[pos+2]
					y = p.data[b]
				case "1":
					y = p.data[pos+2]
				}
				pos = y
				continue
			}

		case OP_JIF:
			x, y := 0, 0
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			}

			if x == 0 {
				switch pB {
				case "0":
					b := p.data[pos+2]
					y = p.data[b]
				case "1":
					y = p.data[pos+2]
				}
				pos = y
				continue
			}

		case OP_LT:
			var x, y, z int
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			}

			z = p.data[pos+3]
			if x < y {
				p.data[z] = 1
			} else {
				p.data[z] = 0
			}

		case OP_EQ:
			var x, y, z int
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			}

			z = p.data[pos+3]
			if x == y {
				p.data[z] = 1
			} else {
				p.data[z] = 0
			}

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
