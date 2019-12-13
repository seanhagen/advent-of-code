package day07

var _ Operator = Input{}

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
	return i.output != nil
}

// Wire ...
func (i Input) Wire() *Wire {
	return i.output
}

// Type ...
func (i Input) Type() OpType {
	return OpIn
}
