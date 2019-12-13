package day07

import (
	"fmt"
)

var _ Operator = LShift{}

const scLSh = "%s LSHIFT %d -> %s"

// LShift ...
type LShift struct {
	input *Wire
	arg   uint16

	output *Wire
}

// OutputValue ...
func (ls LShift) OutputValue() *uint16 {
	if ls.input == nil {
		return nil
	}

	v := ls.input.Value()
	if v == nil {
		return nil
	}

	x := *v << ls.arg
	return &x
}

// OutputWire ...
func (ls LShift) OutputWire() *Wire {
	return ls.output
}

// Valid ...
func (ls LShift) Valid() bool {
	return ls.input != nil
}

// Type ...
func (ls LShift) Type() OpType {
	return OpLShift
}

// addLshift ...
func (b *Board) addLShift(in string) error {
	var aI int
	var aS, dest string

	_, err := fmt.Sscanf(in, scLSh, &aS, &aI, &dest)
	if err != nil {
		return err
	}

	out := b.findOrCreateWire(dest)
	if out.input.Valid() {
		return fmt.Errorf("output wire '%v' already has an input operator", dest)
	}

	inw := b.findOrCreateWire(aS)
	v := uint16(aI)

	op := LShift{
		input:  inw,
		arg:    v,
		output: out,
	}

	out.input = op
	b.wires[dest] = out

	// _ = inw.Value()
	inw.outputs = append(inw.outputs, op)
	b.wires[aS] = inw

	return nil
}
