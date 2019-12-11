package facing

type mvr struct {
	x    int
	xdir int
	y    int
	ydir int

	facing Direction
}

// moveForward ...
func (m *mvr) moveForward() {
	f := m.facing
	xdir, ydir := Vectors[f][0], Vectors[f][1]
	m.xdir = xdir
	m.ydir = ydir
	m.x += xdir
	m.y += ydir
}
