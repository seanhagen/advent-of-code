package day07

var _ Operator = Or{}

// Or ...
type Or struct {
	aWire *Wire
	bWire *Wire

	output *Wire
}

// OutputValue ...
func (o Or) OutputValue() *uint16 {
	if o.aWire == nil || o.bWire == nil {
		return nil
	}

	a := o.aWire.Value()
	b := o.bWire.Value()

	if a == nil || b == nil {
		return nil
	}

	c := *a | *b
	return &c
}

// OutputWire ...
func (o Or) OutputWire() *Wire {
	return o.output
}

// Valid ...
func (o Or) Valid() bool {
	return o.aWire != nil && o.bWire != nil && o.output != nil
}

// Type ...
func (o Or) Type() OpType {
	return OpOr
}
