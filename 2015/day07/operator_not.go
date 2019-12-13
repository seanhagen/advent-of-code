package day07

var _ Operator = Not{}

// Not ...
type Not struct {
	input  *Wire
	output *Wire
}

// OutputValue ...
func (n Not) OutputValue() *uint16 {
	if n.input == nil {
		return nil
	}

	v := n.input.Value()
	if v == nil {
		return nil
	}
	o := 1 ^ *v
	return &o
}

// OutputWire ...
func (n Not) OutputWire() *Wire {
	return n.output
}

// Valid ...
func (n Not) Valid() bool {
	return n.input != nil && n.output != nil
}

// Type ...
func (n Not) Type() OpType {
	return OpNot
}
