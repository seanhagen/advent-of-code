package day07

var _ Operator = LShift{}

// LShift ...
type LShift struct {
	input *Wire
	arg   uint16

	output *Wire
}

// OutputValue ...
func (ls LShift) OutputValue() *uint16 {
	return nil
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
