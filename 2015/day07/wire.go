package day07

// Wire ...
type Wire struct {
	name    string
	input   Operator
	outputs []Operator
	value   *uint16
}

// Value ...
func (w *Wire) Value() *uint16 {
	if w.value != nil {
		return w.value
	}

	v := w.input.OutputValue()
	if v != nil {
		w.value = v
	}

	return w.value
}
