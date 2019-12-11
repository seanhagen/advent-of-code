package facing

import "fmt"

// GetAt will return the thing stored at X,Y. It will create a new thing
// by calling the configured NewObj function if nothing exists at that spot.
func (m Mover) GetAt(x, y int) interface{} {
	row, ok := m.things[x]
	if !ok {
		row = map[int]interface{}{}
	}

	t, ok := row[y]
	if !ok {
		t = m.newObj()
	}

	return t
}

// GetCurrent ...
func (m Mover) GetCurrent() interface{} {
	return m.GetAt(m.lastX, m.lastY)
}

// SetAt sets the thing at x,y to n. Returns an error if the type of n
// isn't the same as the type passed in in Config.Type
func (m *Mover) SetAt(x, y int, n interface{}) error {
	if !typeEqual(m.inType, n) {
		return fmt.Errorf("value %#v type %T isn't same type as %T", n, n, m.inType)
	}

	row, ok := m.things[x]
	if !ok {
		row = map[int]interface{}{}
	}

	row[y] = n
	m.things[x] = row
	return nil
}

// SetCurent ...
func (m *Mover) SetCurent(n interface{}) error {
	return m.SetAt(m.lastX, m.lastY, n)
}

// ModifyAt takes an x,y coordinate and a function. The function will be passed
// the current thing at x,y, and the return value will become the new thing at x,y.
// Will return an error if the value returned from fn isn't the same type as the
// value passed in.
func (m *Mover) ModifyAt(x, y int, fn func(interface{}) interface{}) error {
	t := m.GetAt(x, y)

	newT := fn(t)
	if !typeEqual(t, newT) {
		return fmt.Errorf("modify func returns different type")
	}

	return m.SetAt(x, y, newT)
}

// ModifyCurrent ...
func (m *Mover) ModifyCurrent(fn func(interface{}) interface{}) error {
	return m.ModifyAt(m.lastX, m.lastY, fn)
}
