package day07

import (
	"fmt"
	"strconv"
)

var _ Operator = Or{}

const scOrSS = "%s OR %s -> %s"
const scOrDS = "%d OR %s -> %s"
const scOrSD = "%s OR %d -> %s"
const scOrDD = "%d OR %d -> %s"

// Or ...
type Or struct {
	aWire *Wire
	aVal  *uint16

	bWire *Wire
	bVal  *uint16

	output *Wire
}

// OutputValue ...
func (o Or) OutputValue() *uint16 {
	var a, b *uint16
	if o.aWire != nil {
		a = o.aWire.Value()
	}
	if o.aVal != nil {
		a = o.aVal
	}

	if o.bWire != nil {
		b = o.bWire.Value()
	}
	if o.bVal != nil {
		b = o.bVal
	}

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

// addOr ...
func (b *Board) addOr(in string) error {
	var aI, bI *int
	var aS, bS, dest string

	_, err := fmt.Sscanf(in, scOrSS, &aS, &bS, &dest)
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

	op := Or{
		aWire: aWire,
		bWire: bWire,

		aVal: ua,
		bVal: ub,

		output: out,
	}

	out.input = op

	if aWire != nil {
		aWire.outputs = append(aWire.outputs, op)
		// _ = aWire.Value()
	}

	if bWire != nil {
		bWire.outputs = append(bWire.outputs, op)
		// _ = bWire.Value()
	}

	return nil
}
