package facing

// Turn takes a direction to turn, then moves forward in the new direction
func (m *Mover) Turn(i Turn) {
	mvr := m.cur()

	f := TurnTo(mvr.facing, i)
	m.Move(f)
}

// Move takes a direction to face, then moves forward
func (m *Mover) Move(f Direction) {
	mvr := m.cur()

	mvr.facing = f
	mvr.moveForward()

	row, ok := m.things[mvr.x]
	if !ok {
		row = map[int]interface{}{}
	}

	t, ok := row[mvr.y]
	if !ok {
		t = m.newObj()
		row[mvr.y] = t
		m.things[mvr.x] = row
	}

	m.setLast()
	m.incr()
}
