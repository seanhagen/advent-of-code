package day07

var _ Operator = RShift{}

// RShift ...
type RShift struct {
	input  *Wire
	output *Wire

	arg uint16
}

// OutputValue ...
func (rs RShift) OutputValue() *uint16 {
	return nil
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
