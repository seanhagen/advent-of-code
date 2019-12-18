package lib2019

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/seanhagen/advent-of-code/lib"
)

// ErrHalt is returned when the Program reaches the Halt instruction
var ErrHalt = fmt.Errorf("halt instruction reached")

const (
	// OpAdd is the intcode for the add instruction
	OpAdd = 1
	// OpMul is the intcode for the multiply instruction
	OpMul = 2
	// OpSav is the intcode for the save/input instruction ( read from input and save to memory )
	OpSav = 3
	// OpOut is the intcode for the output instruction
	OpOut = 4
	// OpJit is the intcode for the "jump-if-true" instruction
	OpJit = 5
	// OpJif is the intcode for the "jump-if-false" instruction
	OpJif = 6
	// OpLt is the less-than instruction
	OpLt = 7
	// OpEq is the equal instruction
	OpEq = 8
	// OpAdj is the adjust relative base instruction
	OpAdj = 9

	// OpHalt is the halt instruction
	OpHalt = 99
)

var opName = map[int]string{
	OpAdd:  "ADD",
	OpMul:  "MUL",
	OpSav:  "SAV",
	OpOut:  "OUT",
	OpJit:  "JIT",
	OpJif:  "JIF",
	OpLt:   "LT ",
	OpEq:   "EQ ",
	OpAdj:  "ADJ",
	OpHalt: "FIN",
}

var opIncr = map[int]int{
	OpAdd:  4,
	OpMul:  4,
	OpSav:  2,
	OpOut:  2,
	OpJit:  3,
	OpJif:  3,
	OpLt:   4,
	OpEq:   4,
	OpAdj:  2,
	OpHalt: 1,
}

// InputFn is a function that outputs an integer to be consumed as input by the program
type InputFn func() (int, bool)

// OutputFn is a function that recieves the output of the program, and returns true if the program should halt
type OutputFn func(int) bool

// Program is an intcode program
type Program struct {
	code     string
	data     []int
	mode     int
	relBase  int
	position int

	inPtr   int
	inputs  []int
	outputs []int

	pauseOnOutput bool
	halted        bool

	hasInputFn bool
	inFn       InputFn

	hasOutputFn bool
	outFn       OutputFn
}

// FromString reads a program code from the string input
func FromString(in string) (*Program, error) {
	in = strings.Replace(in, "\n", "", -1)
	data := strings.Split(in, ",")
	prog := Program{code: in, data: []int{}, inPtr: 0, inputs: []int{}, outputs: []int{}}
	for x, v := range data {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("\ninstruction %v: '%v'\n", x, v)
			return nil, err
		}
		prog.data = append(prog.data, i)
	}
	return &prog, nil
}

// ReadProgram reads the program code from the file
func ReadProgram(f *os.File) (*Program, error) {
	d, err := lib.ReadLine(f)
	if err != nil {
		return nil, err
	}
	return FromString(string(d))
}

// SetInputFunc ...
func (p *Program) SetInputFunc(fn InputFn) {
	p.hasInputFn = true
	p.inFn = fn
}

// SetOutputFn ...
func (p *Program) SetOutputFn(fn OutputFn) {
	p.hasOutputFn = true
	p.outFn = fn
}

// Unhalt ...
func (p *Program) Unhalt() {
	p.halted = false
}

// AddInput ...
func (p *Program) AddInput(in int) {
	p.inputs = append(p.inputs, in)
}

// GetOutputs ...
func (p Program) GetOutputs() []int {
	return p.outputs
}

// SetPauseOnOutput ...
func (p *Program) SetPauseOnOutput(b bool) {
	p.pauseOnOutput = b
}

// CheckMemory ...
func (p *Program) checkMemory(pos, offset int, mode string) {
	// fmt.Printf("\n ----- check memory called! mode %v, ", mode)
	switch mode {
	case "0":
		// fmt.Printf("%v > %v = %v\n", offset, len(p.data), offset > len(p.data))
		if offset+1 > len(p.data) {
			// fmt.Printf("expanding memory!")
			for {
				p.data = append(p.data, 0)
				if offset+1 < len(p.data) {
					goto done
				}
			}
		}
	case "2":
		// fmt.Printf("%v > %v = %v\n", (pos+1)+offset+p.relBase, len(p.data), (pos+1)+offset+p.relBase > len(p.data))
		if (pos+1)+offset+p.relBase > len(p.data) {
			// fmt.Printf("expanding memory!")
			for {
				p.data = append(p.data, 0)
				if (pos+1)+offset+p.relBase < len(p.data) {
					goto done
				}
			}
		}
	}
done:
	// fmt.Printf("\n\n\tlength of data now: %v\n\n", len(p.data))
}

// Run ...
func (p *Program) Run() error {
	if p.halted {
		return ErrHalt
	}
	pos := p.position
	max := len(p.data) - 1
	for {
		op := p.data[pos]
		opc := strconv.Itoa(op)
		if len(opc) < 5 {
			for {
				opc = fmt.Sprintf("0%v", opc)
				if len(opc) == 5 {
					break
				}
			}
		}

		bits := strings.Split(opc, "")
		// fmt.Printf("\nop: %v", op)
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

		// if op == 1 {
		// 	fmt.Printf("\n%v -> %v => mode 1st: %v, mode 2nd: %v mode 3rd: %v -- ", opc, opName[op], pC, pB, pA)
		// }

		incr := opIncr[op]
		if max < (pos + incr - 1) {
			return fmt.Errorf("not enough data to continue (pos '%v', max '%v')", pos, max)
		}

		switch op {
		case OpAdd:
			var x, y int
			switch pC {
			case "0":
				a := p.data[pos+1]
				// fmt.Printf("add, 1st param mode 0 -- ")
				p.checkMemory(pos, a, pC)
				// fmt.Printf("a: %v", a)
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			case "2":
				a := p.data[pos+1]
				// fmt.Printf("add, 1st param mode 2 -- ")
				p.checkMemory(pos, a, pC)
				x = p.data[p.relBase+a]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				// fmt.Printf("add, 2nd param mode 0 -- ")
				p.checkMemory(pos, b, pB)
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			case "2":
				b := p.data[pos+2]
				// fmt.Printf("add, 2nd param mode 2! -- ")
				p.checkMemory(pos, b, pB)
				y = p.data[p.relBase+b]
			}

			// fmt.Printf("pos: %v -- ", pos)
			var z int
			switch pA {
			case "0":
				z = p.data[pos+3]
				// fmt.Printf("%v + %v stored in (%v)", x, y, z)
			case "2":
				z = p.relBase + p.data[pos+3]
				// fmt.Printf("%v + %v stored in (%v+%v -> %v)", x, y, v, p.relBase, z)
			}
			p.checkMemory(pos, z, "0")
			p.data[z] = x + y

		case OpMul:
			var x, y int
			switch pC {
			case "0":
				a := p.data[pos+1]
				p.checkMemory(pos, a, pC)
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			case "2":
				a := p.data[pos+1]
				p.checkMemory(pos, a, pC)
				x = p.data[p.relBase+a]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				p.checkMemory(pos, b, pB)
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			case "2":
				b := p.data[pos+2]
				p.checkMemory(pos, b, pB)
				y = p.data[p.relBase+b]
			}

			var z int
			switch pA {
			case "0":
				z = p.data[pos+3]
			case "2":
				z = p.relBase + p.data[pos+3]
			}

			// z := p.data[pos+3]
			p.checkMemory(pos, z, "0")
			// fmt.Printf(" -- %v * %v stored in %v", x, y, z)
			p.data[z] = x * y

		case OpSav:
			var in int
			var pause bool
			if p.hasInputFn {
				in, pause = p.inFn()
			} else {
				// fmt.Printf("trying to read from %v, inputs: %#v\n", p.inPtr, p.inputs)
				if len(p.inputs)-1 < p.inPtr {
					// fmt.Printf("not enough input, returning for now\n")
					goto done
				}
				in = p.inputs[p.inPtr]
				p.inPtr++
			}

			var a int
			switch pC {
			case "0":
				a = p.data[pos+1]
			case "2":
				// fmt.Printf(" -- op save, mode rel")
				b := p.data[pos+1]
				a = p.relBase + b
			}
			p.checkMemory(pos, a, pC)
			p.data[a] = in
			// fmt.Printf(" -- data at %v => %v", a, p.data[a])
			if pause {
				pos += incr
				p.position = pos
				goto done
			}

		case OpOut:
			switch pC {
			case "0":
				a := p.data[pos+1]
				// fmt.Printf(" -- %v to be stored in output", a)
				if p.hasOutputFn {
					p.pauseOnOutput = p.outFn(p.data[a])
				} else {
					p.outputs = append(p.outputs, p.data[a])
				}

			case "1":
				// fmt.Printf(" -- %v to be stored in output", p.data[pos+1])
				if p.hasOutputFn {
					p.pauseOnOutput = p.outFn(p.data[pos+1])
				} else {
					p.outputs = append(p.outputs, p.data[pos+1])
				}

			case "2":
				a := p.data[pos+1]
				p.checkMemory(pos, a, pC)

				// fmt.Printf(" -- %v to be stored in output", p.data[p.relBase+a])
				if p.hasOutputFn {
					p.pauseOnOutput = p.outFn(p.data[p.relBase+a])
				} else {
					p.outputs = append(p.outputs, p.data[p.relBase+a])
				}
			}

			if p.pauseOnOutput {
				pos += incr
				p.position = pos
				return nil
			}

		case OpJit:
			x, y := 0, 0
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			case "2":
				a := p.data[pos+1]
				x = p.data[p.relBase+a]
			}

			// fmt.Printf(" -- %v != 0? ", x)
			if x != 0 {
				switch pB {
				case "0":
					b := p.data[pos+2]
					y = p.data[b]
				case "1":
					y = p.data[pos+2]
				case "2":
					b := p.data[pos+2]
					y = p.data[p.relBase+b]
				}
				// fmt.Printf("pos now: %v", y)
				pos = y
				continue
			}

		case OpJif:
			x, y := 0, 0
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			case "2":
				a := p.data[pos+1]
				x = p.data[p.relBase+a]
			}

			// fmt.Printf(" -- %v == 0? ", x)
			if x == 0 {
				switch pB {
				case "0":
					b := p.data[pos+2]
					y = p.data[b]
				case "1":
					y = p.data[pos+2]
				case "2":
					b := p.data[pos+2]
					y = p.data[p.relBase+b]
				}
				// fmt.Printf("pos now: %v", y)
				pos = y
				continue
			}

		case OpLt:
			var x, y int
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
			case "1":
				x = p.data[pos+1]
			case "2":
				a := p.data[pos+1]
				x = p.data[p.relBase+a]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			case "2":
				b := p.data[pos+2]
				y = p.data[p.relBase+b]
			}

			var z int
			switch pA {
			case "0":
				z = p.data[pos+3]
			case "2":
				z = p.relBase + p.data[pos+3]
			}

			// z := p.data[pos+3]
			// fmt.Printf("\nx < y -> %v < %v, z: %v\n", x, y, z)
			p.checkMemory(pos, z, "0")
			if x < y {
				p.data[z] = 1
			} else {
				p.data[z] = 0
			}

		case OpEq:
			var x, y int
			switch pC {
			case "0":
				a := p.data[pos+1]
				x = p.data[a]
				// fmt.Printf("arg 1 (pos: %v): %v", a, x)
			case "1":
				x = p.data[pos+1]
				// fmt.Printf("arg 1: %v", x)
			case "2":
				a := p.data[pos+1]
				// fmt.Printf("data at %v+%v", p.relBase, a)
				x = p.data[p.relBase+a]
			}

			switch pB {
			case "0":
				b := p.data[pos+2]
				y = p.data[b]
			case "1":
				y = p.data[pos+2]
			case "2":
				b := p.data[pos+2]
				// fmt.Printf("data at %v", b)
				y = p.data[p.relBase+b]
			}

			var z int
			switch pA {
			case "0":
				// 	v := p.data[pos+3]
				// 	z = p.data[v]
				// case "1":
				z = p.data[pos+3]
			case "2":
				z = p.relBase + p.data[pos+3]
			}
			// z := p.data[pos+3]
			p.checkMemory(pos, z, "0")
			// fmt.Printf(" -- %v == %v? ", x, y)
			if x == y {
				// fmt.Printf("yes! 1 to %v", z)
				p.data[z] = 1
			} else {
				// fmt.Printf("no!  0 to %v", z)
				p.data[z] = 0
			}

		case OpAdj:
			var a int
			switch pC {
			case "0":
				v := p.data[pos+1]
				a = p.data[v]
			case "1":
				a = p.data[pos+1]
			case "2":
				v := p.data[pos+1]
				a = p.data[p.relBase+v]
			}
			// a := p.data[pos+1]
			p.relBase += a
			// fmt.Printf(" -- rel base adjusted by %v, now %v", a, p.relBase)

		case OpHalt:
			// fmt.Printf("\n\n\n")
			p.halted = true
			pos += incr
			p.position = pos
			goto done

		default:
			return fmt.Errorf("unknown opcode encountered (op: '%v', pos: '%v')", op, pos)
		}
		pos += incr
		p.position = pos
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
