package facing

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
