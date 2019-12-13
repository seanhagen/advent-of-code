package day07

import "fmt"

// Wire ...
type Wire struct {
	name    string
	input   Operator
	outputs []Operator
	value   *uint16
}

// Value ...
func (w *Wire) Value() *uint16 {
	if w == nil {
		// fmt.Printf("wire is nil, returning from Value()\n")
		return nil
	}

	if w.value != nil {
		// fmt.Printf("WIRE(%v) - wire value not nill, returning\n", w.name)
		return w.value
	}

	if !w.input.Valid() {
		// fmt.Printf("WIRE(%v) - wire input (operation %v) is not valid?\n", w.name, w.input.Type())
		return nil
	}

	// fmt.Printf("WIRE(%v) - getting value from input operation ( %v )\n", w.name, w.input.Type())
	w.value = w.input.OutputValue()
	// fmt.Printf("WIRE(%v) - value now: %v\n", w.name, w.value)
	return w.value
}

// String ...
func (w Wire) String() string {
	if w.value != nil {
		return fmt.Sprintf("Wire<Name: %v, Input: %v, Current: %v>", w.name, w.input.Type(), *w.value)
	}
	return fmt.Sprintf("Wire<Name: %v, Input: %v, Current: %v>", w.name, w.input.Type(), w.value)
}

// GoString ...
func (w Wire) GoString() string {
	if w.value != nil {
		return fmt.Sprintf("Wire<Name: %v, Input: %v, Current: %v>", w.name, w.input.Type(), *w.value)
	}
	return fmt.Sprintf("Wire<Name: %v, Input: %v, Current: %v>", w.name, w.input.Type(), w.value)
}
