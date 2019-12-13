package day07

var _ Operator = Null{}

// Null is for when a wire doesn't have an input yet
type Null struct {
	w *Wire
}

// OutputValue ...
func (n Null) OutputValue() *uint16 {
	return nil
}

// OutputWire ...
func (n Null) OutputWire() *Wire {
	return nil
}

// Valid ...
func (n Null) Valid() bool {
	return false
}

// Type ...
func (n Null) Type() OpType {
	return OpNull
}
