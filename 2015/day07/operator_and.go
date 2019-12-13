package day07

import (
	"fmt"
	"strconv"
)

var _ Operator = And{}

const scAndSS = "%s AND %s -> %s"
const scAndDS = "%d AND %s -> %s"
const scAndSD = "%s AND %d -> %s"
const scAndDD = "%d AND %d -> %s"

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
		// fmt.Printf("AND: wire A (%v) not nil, getting value\n", a.aWire.name)
		x = a.aWire.Value()
		// fmt.Printf("AND: wire A (%v) got value: %v\n", a.aWire.name, spew.Sdump(x))
	}
	if a.aVal != nil {
		x = a.aVal
		// fmt.Printf("AND: val A not nil, got value: %v\n", x)
	}
	if x == nil {
		// fmt.Printf("AND: val A is nil, returning\n")
		return nil
	}
	// fmt.Printf("AND: value A: %v\n", *x)

	if a.bWire != nil {
		// fmt.Printf("AND: wire B (%v) not nil, getting value\n", a.bWire.name)
		y = a.bWire.Value()
		// fmt.Printf("AND: wire B (%v) got value: %v\n", a.bWire.name, spew.Sdump(y))
	}
	if a.bVal != nil {
		y = a.bVal
		// fmt.Printf("AND: val B not nil, got value: %v\n", y)
	}
	if y == nil {
		return nil
	}
	// fmt.Printf("AND: value B: %v\n", *y)

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

// addInput ...
func (b *Board) addAnd(in string) error {
	var aI, bI *int
	var aS, bS, dest string

	_, err := fmt.Sscanf(in, scAndSS, &aS, &bS, &dest)
	if err != nil {
		return err
	}

	if v, err := strconv.Atoi(aS); err == nil {
		aI = &v
		aS = ""
	}

	if v, err := strconv.Atoi(bS); err == nil {
		bI = &v
		bS = ""
	}

	out := b.findOrCreateWire(dest)
	if out.input.Valid() {
		return fmt.Errorf("output wire '%v' already has an input operator", dest)
	}

	var aWire, bWire *Wire
	var ua, ub *uint16

	if aS != "" {
		aWire = b.findOrCreateWire(aS)
	}
	if aI != nil {
		u := uint16(*aI)
		ua = &u
	}

	if bS != "" {
		bWire = b.findOrCreateWire(bS)
	}
	if bI != nil {
		u := uint16(*bI)
		ub = &u
	}

	op := And{
		aWire: aWire,
		bWire: bWire,

		aVal: ua,
		bVal: ub,

		output: out,
	}

	if aWire != nil {
		aWire.outputs = append(aWire.outputs, op)
	}
	if bWire != nil {
		bWire.outputs = append(bWire.outputs, op)
	}

	out.input = op

	// _ = aWire.Value()
	// _ = bWire.Value()

	return nil
}
