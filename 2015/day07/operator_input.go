package day07

import "fmt"

var _ Operator = Input{}

const scWire = "%s -> %s"
const scInput = "%d -> %s"

// Input is the operator for instructions like "123 -> x". The argument is stored, and
// provided to the wire upon calling Output()
type Input struct {
	argWire *Wire
	argVal  *uint16

	output *Wire
}

// OutputValue returns the only argument
func (i Input) OutputValue() *uint16 {
	if i.argWire != nil {
		return i.argWire.Value()
	}

	return i.argVal
}

// OutputWire ...
func (i Input) OutputWire() *Wire {
	return i.output
}

// Valid returns true if there is an attached wire
func (i Input) Valid() bool {
	return i.argWire != nil || i.argVal != nil
}

// Wire ...
func (i Input) Wire() *Wire {
	return i.output
}

// Type ...
func (i Input) Type() OpType {
	return OpIn
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

	// try getting input from input wire if it exists
	// if op.argWire != nil {
	// 	_ = op.argWire.Value()
	// }

	return nil
}
