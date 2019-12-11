package facing

// Location ...
func (m Mover) Location() (int, int) {
	return m.LocationOf(m.curMvr)
}

// LocationOf ...
func (m Mover) LocationOf(i int) (int, int) {
	if i >= m.numMvrs {
		return 0, 0
	}
	mvr := m.mvrs[i]
	return mvr.x, mvr.y
}
