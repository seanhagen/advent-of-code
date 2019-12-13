package day07

var _ Operator = And{}

// And handles the bitwise AND of it's two inputs
type And struct {
	aWire *Wire
	aVal  *uint16

	bWire *Wire
	bVal  *uint16

	output *Wire
}

// OutputValue ...
func (a And) OutputValue() *uint16 {
	var x, y *uint16

	if a.aWire != nil {
		x = a.aWire.Value()
	}
	if a.aVal != nil {
		x = a.aVal
	}
	if x == nil {
		return nil
	}

	if a.bWire != nil {
		y = a.bWire.Value()
	}
	if a.bVal != nil {
		y = a.bVal
	}
	if y == nil {
		return nil
	}

	o := *x & *y
	return &o
}

// OutputWire ...
func (a And) OutputWire() *Wire {
	return a.output
}

// Valid ...
func (a And) Valid() bool {
	return true
}

// Type ...
func (a And) Type() OpType {
	return OpAnd
}
