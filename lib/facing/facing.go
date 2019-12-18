package facing

import "fmt"

// Facing ...
func (m Mover) Facing() Direction {
	return m.FacingOf(m.curMvr)
}

// FacingOf ...
func (m Mover) FacingOf(i int) Direction {
	if i >= m.numMvrs {
		return North
	}
	mvr := m.mvrs[i]
	return mvr.facing
}

// IsFacing ...
func IsFacing(i string) bool {
	d := DirectionFromString(i)
	t := fmt.Sprintf("%v", string(d))
	if t != i {
		return false
	}

	return d == North ||
		d == South ||
		d == West ||
		d == East
}
