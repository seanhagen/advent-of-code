package day07

import (
	"fmt"
	"regexp"
	"strings"
)

var reWire = regexp.MustCompile(`^[a-z]{1,2}\s?->\s?[a-z]{1,2}`)
var reInput = regexp.MustCompile(`^[\d]+ -> [a-z]{1,2}`)

// Board ...
type Board struct {
	wires map[string]*Wire
}

// NewBoard ...
func NewBoard() (*Board, error) {
	return &Board{wires: map[string]*Wire{}}, nil
}

// findOrCreateWire ...
func (b *Board) findOrCreateWire(n string) *Wire {
	w, ok := b.wires[n]
	if ok {
		return w
	}

	w = &Wire{
		name:    n,
		input:   Null{},
		outputs: []Operator{},
		value:   nil,
	}
	b.wires[n] = w
	return w
}

// AddWire ...
func (b *Board) AddWire(in string) error {
	in = strings.TrimSpace(in)
	if strings.Contains(in, string(OpAnd)) {
		return b.addAnd(in)
	}

	if strings.Contains(in, string(OpOr)) {
		return b.addOr(in)
	}

	if strings.Contains(in, string(OpNot)) {
		return b.addNot(in)
	}

	if strings.Contains(in, string(OpLShift)) {
		return b.addLShift(in)
	}

	if strings.Contains(in, string(OpRShift)) {
		return b.addRShift(in)
	}

	// lastly, check if it's an input instruction
	if reInput.MatchString(in) || reWire.MatchString(in) {
		return b.addInput(in)
	}

	return fmt.Errorf("input matched no known instructions")
}

// GetWireValues ...
func (b Board) GetWireValues() map[string]uint16 {
	out := map[string]uint16{}

	for n, w := range b.wires {
		v := w.Value()
		if v != nil {
			out[n] = *v
		}
	}

	return out
}

// Reset ...
func (b *Board) Reset() {
	for n, w := range b.wires {
		w.value = nil
		b.wires[n] = w
	}
}

// SetWireValue ...
func (b *Board) SetWireValue(n string, v int) error {
	w, ok := b.wires[n]
	if !ok {
		return fmt.Errorf("no wire named '%v'")
	}

	uv := uint16(v)
	w.value = &uv
	b.wires[n] = w

	return nil
}

// GetWire ...
func (b Board) GetWire(n string) *Wire {
	if w, ok := b.wires[n]; ok {
		return w
	}
	return nil
}
