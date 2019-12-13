package day07

// OpType ...
type OpType string

const (
	// OpIn is for the input instruction. Unique in that it's not from the example.
	// Handles input from instructions like "123 -> x", "456 -> y", or "x ->y"
	OpIn OpType = "IN"
	// OpAnd does a bitwise AND of its two inputs, and provides a single output.
	// Handles input from instructions like "1 AND x -> y" or "x AND y -> z"
	OpAnd OpType = "AND"
	// OpOr does a bitwise OR of its two inputs, and provides a single output.
	// Handles input from instructions like "1 AND x -> y" or "x OR y -> z"
	OpOr OpType = "OR"
	// OpNot is a bitwise compliment of the only input, and provides a single output.
	// Handles input from instructions like "NOT x -> z"
	OpNot OpType = "NOT"
	// OpLShift takes one input, and one argument and left-shifts the input by the argument
	// and provides a single output.
	// Handles input from instructions like "x LSHIFT 2 -> z"
	OpLShift OpType = "LSHIFT"
	// OpRShift is the same as LShift, but does a right-shift instead.
	// Handles input from instructions like "x RSHIFT 2 -> z"
	OpRShift OpType = "RSHIFT"
	// OpNull is for when a wire doesn't have an input yet
	OpNull OpType = "NULL"
)

// Operator ...
type Operator interface {
	OutputValue() *uint16
	OutputWire() *Wire
	Valid() bool
	Type() OpType
}
