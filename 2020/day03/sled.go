package day3

type Sled struct {
	px, py, vx, vy int
	g              *Grid
}

func NewSled(vx, vy int, gr string) (*Sled, error) {
	g, err := NewGrid(gr)
	if err != nil {
		return nil, err
	}

	return &Sled{vx: vx, vy: vy, g: g}, nil
}
