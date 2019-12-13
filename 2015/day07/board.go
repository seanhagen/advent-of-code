package day07

import (
	"fmt"
	"regexp"
	"strings"
)

var reWire = regexp.MustCompile(`^[a-z]{1,2}\s?->\s?[a-z]{1,2}`)
var scWire = "%s -> %s"

var reInput = regexp.MustCompile(`^[\d]+ -> [a-z]{1,2}`)
var scInput = "%d -> %s"

// Board ...
type Board struct {
	wires map[string]*Wire
}

// NewBoard ...
func NewBoard() (*Board, error) {
	return &Board{wires: map[string]*Wire{}}, nil
}

// findOrCreateWire ...
func (b *Board) findOrCreateWire(n string) *Wire {
	w, ok := b.wires[n]
	if ok {
		return w
	}

	w = &Wire{
		name:    n,
		input:   Null{},
		outputs: []Operator{},
		value:   nil,
	}
	b.wires[n] = w

	return w
}

// AddWire ...
func (b *Board) AddWire(in string) error {
	in = strings.TrimSpace(in)

	if strings.Contains(in, string(OpAnd)) {
		fmt.Printf("AND\n")
		return nil
	}

	if strings.Contains(in, string(OpOr)) {
		fmt.Printf("OR\n")
		return nil
	}

	if strings.Contains(in, string(OpNot)) {
		fmt.Printf("NOT\n")
		return nil
	}

	if strings.Contains(in, string(OpLShift)) {
		fmt.Printf("LSHIFT\n")
		return nil
	}

	if strings.Contains(in, string(OpRShift)) {
		fmt.Printf("RSHIFT\n")
		return nil
	}

	if reInput.MatchString(in) || reWire.MatchString(in) {
		fmt.Printf("\n\n'%v' matched input regexes\n\n", in)
		return b.addInput(in)
	}

	return fmt.Errorf("input matched no known instructions")
}

// addInput ...
func (b *Board) addInput(in string) error {
	var val uint16
	dest := ""
	inw := ""
	_, err := fmt.Sscanf(in, scInput, &val, &dest)
	if err != nil {
		_, err = fmt.Sscanf(in, scWire, &inw, &dest)
		if err != nil {
			return err
		}
	}

	op := Input{}
	if inw == "" {
		op.argVal = &val
	} else {
		opw := b.findOrCreateWire(inw)
		op.argWire = opw
	}

	dw := b.findOrCreateWire(dest)
	if dit := dw.input.Type(); dit != OpNull {
		return fmt.Errorf("output wire '%v' already has an input (type: '%v')", dest, dit)
	}

	dw.input = op
	op.output = dw

	// try getting the value for the destination wire, just in case it's completable now
	_ = dw.Value()

	return nil
	// var x uint16 = 3
	// fmt.Printf("%b\n", x)
	// fmt.Printf("%b\n", ^x)

	// spew.Dump(val, dest, inw)
	// return fmt.Errorf("invalid input")
}

// GetWireValues ...
func (b Board) GetWireValues() map[string]int16 {
	out := map[string]int16{}

	return out
}
