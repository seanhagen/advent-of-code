package day07

import "fmt"

var _ Operator = RShift{}

const scRSh = "%s RSHIFT %d -> %s"

// RShift ...
type RShift struct {
	input *Wire
	arg   uint16

	output *Wire
}

// OutputValue ...
func (rs RShift) OutputValue() *uint16 {
	if rs.input == nil {
		return nil
	}

	v := rs.input.Value()
	if v == nil {
		return nil
	}
	x := *v >> rs.arg
	return &x
}

// OutputWire ...
func (rs RShift) OutputWire() *Wire {
	return rs.output
}

// Valid ...
func (rs RShift) Valid() bool {
	return rs.input != nil && rs.output != nil
}

// Type ...
func (rs RShift) Type() OpType {
	return OpRShift
}

// addLshift ...
func (b *Board) addRShift(in string) error {
	var aI int
	var aS, dest string

	_, err := fmt.Sscanf(in, scRSh, &aS, &aI, &dest)
	if err != nil {
		return err
	}

	out := b.findOrCreateWire(dest)
	if out.input.Valid() {
		return fmt.Errorf("output wire '%v' already has an input operator", dest)
	}

	inw := b.findOrCreateWire(aS)
	v := uint16(aI)

	op := RShift{
		input:  inw,
		arg:    v,
		output: out,
	}

	out.input = op

	inw.outputs = append(inw.outputs, op)
	// _ = inw.Value()

	return nil
}
