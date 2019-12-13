package day07

import "fmt"

var _ Operator = Not{}

const scNotS = "NOT %s -> %s"
const scNotD = "NOT %d -> %s"

// Not ...
type Not struct {
	input *Wire
	ival  *uint16

	output *Wire
}

// OutputValue ...
func (n Not) OutputValue() *uint16 {
	if n.input == nil && n.ival == nil {
		return nil
	}

	var v *uint16
	if n.input != nil {
		v = n.input.Value()
	}
	if n.ival != nil {
		v = n.ival
	}

	if v == nil {
		return nil
	}

	o := ^(*v)
	return &o
}

// OutputWire ...
func (n Not) OutputWire() *Wire {
	return n.output
}

// Valid ...
func (n Not) Valid() bool {
	return (n.input != nil || n.ival != nil) && n.output != nil
}

// Type ...
func (n Not) Type() OpType {
	return OpNot
}

// addNot ...
func (b *Board) addNot(in string) error {
	var aI *int
	var aS, dest string

	_, err := fmt.Sscanf(in, scNotD, &aI, &dest)
	if err != nil {
		_, err = fmt.Sscanf(in, scNotS, &aS, &dest)
		if err != nil {
			return err
		}
	}

	out := b.findOrCreateWire(dest)
	if out.input.Valid() {
		return fmt.Errorf("output wire '%v' already has an input operator", dest)
	}

	var w *Wire
	if aS != "" {
		w = b.findOrCreateWire(aS)
	}

	var v *uint16
	if aI != nil {
		x := uint16(*aI)
		v = &x
	}

	op := Not{
		input:  w,
		ival:   v,
		output: out,
	}

	out.input = op

	if w != nil {
		// _ = w.Value()
		w.outputs = append(w.outputs, op)
	}

	return nil
}
