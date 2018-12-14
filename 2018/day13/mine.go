package day13

import (
	"fmt"
	"io"
	"strings"

	"github.com/seanhagen/advent-of-code/lib"
)

type row map[int]string

// Mine ...
type Mine struct {
	carts []*Cart

	cartPos map[int]map[int][]*Cart

	tracks map[int]row

	maxx int
	maxy int
}

// Print ...
func (m Mine) Print() string {
	out := ""

	for y := 0; y <= m.maxy; y++ {
		for x := 0; x <= m.maxx; x++ {
			crts := m.cartPos[y][x]

			cnt := len(crts)
			if cnt > 0 {
				if cnt > 1 {
					out = fmt.Sprintf("%vX", out)
				} else {
					out = fmt.Sprintf("%v%v", out, crts[0].facing)
				}
			} else {
				out = fmt.Sprintf("%v%v", out, m.tracks[y][x])
			}
		}
		out = fmt.Sprintf("%v\n", out)
	}

	return out
}

// PrintTrack ...
func (m Mine) PrintTrack() string {
	out := ""

	for y := 0; y <= m.maxy; y++ {
		for x := 0; x <= m.maxx; x++ {
			t := []*Cart{}
			for _, c := range m.carts {
				if c.x == x && c.y == y {
					t = append(t, c)
				}
			}

			out = fmt.Sprintf("%v%v", out, m.tracks[y][x])
		}
		out = fmt.Sprintf("%v\n", out)
	}

	return out
}

// Step ...
func (m *Mine) Step() {
	cp := map[int]map[int][]*Cart{}
	for i, c := range m.carts {
		nx, ny := c.NextPos()
		n := m.tracks[ny][nx]
		c.Process(n)

		tmpy, ok := cp[c.y]
		if !ok {
			tmpy = map[int][]*Cart{}
		}

		tmpx, ok := tmpy[c.x]
		if !ok {
			tmpx = []*Cart{}
		}
		tmpx = append(tmpx, c)
		tmpy[c.x] = tmpx
		cp[c.y] = tmpy

		m.carts[i] = c
	}
	m.cartPos = cp
}

// CheckCollision ...
func (m Mine) CheckCollision() (bool, int, int) {
	for y, row := range m.cartPos {
		for x, carts := range row {
			if len(carts) > 1 {
				return true, x, y
			}
		}
	}
	return false, 0, 0
}

// RemoveCollided ...
func (m *Mine) RemoveCollided() {
	newCarts := []*Cart{}
	for _, row := range m.cartPos {
		for _, carts := range row {
			if len(carts) > 1 {
				for _, c := range m.carts {
					found := false
					for _, cn := range carts {
						if c == cn {
							found = true
						}
					}
					if !found {
						newCarts = append(newCarts, c)
					}
				}
				// fmt.Printf("need to remove carts at <%v, %v>\n", x, y)
				// return
			}
		}
	}
	m.carts = newCarts

	newPos := map[int]map[int][]*Cart{}
	for _, c := range newCarts {
		tmpy, ok := newPos[c.y]
		if !ok {
			tmpy = map[int][]*Cart{}
		}

		tmpx, ok := tmpy[c.x]
		if !ok {
			tmpx = []*Cart{}
		}
		tmpx = append(tmpx, c)
		tmpy[c.x] = tmpx
		newPos[c.y] = tmpy
	}
	m.cartPos = newPos
}

// StepUntilCollision ...
func (m *Mine) StepUntilCollision() (int, int) {
	x := 0
	y := 0

	for {
		m.Step()
		c, cx, cy := m.CheckCollision()
		if c {
			x = cx
			y = cy
			break
		}
	}

	return x, y
}

// StepUntilOneCart ...
func (m Mine) StepUntilOneCart() (int, int) {
	x := 0
	y := 0
	for {
		m.Step()

		o := m.Print()
		fmt.Printf("mine:\n%v\n", o)

		m.RemoveCollided()
		if len(m.carts) == 1 {
			c := m.carts[0]
			x = c.x
			y = c.y
			break
		}
	}
	return x, y
}

// SetupMine ...
func SetupMine(path string) *Mine {
	f := lib.LoadInput(path)

	m := &Mine{
		carts:   []*Cart{},
		tracks:  map[int]row{},
		cartPos: map[int]map[int][]*Cart{},
	}

	y := 0
	err := lib.LoopOverLines(f, func(line []byte) error {
		l := string(line)
		bits := strings.Split(l, "")
		m.tracks[y] = row{}

		for x, v := range bits {
			c := CreateCart(v, x, y)
			if c != nil {
				m.carts = append(m.carts, c)

				tmpy, ok := m.cartPos[y]
				if !ok {
					tmpy = map[int][]*Cart{}
				}

				tmpx, ok := tmpy[x]
				if !ok {
					tmpx = []*Cart{}
				}
				tmpx = append(tmpx, c)
				tmpy[x] = tmpx
				m.cartPos[y] = tmpy
				m.tracks[y][x] = c.Under()
			} else {
				m.tracks[y][x] = v
			}
			if x > m.maxx {
				m.maxx = x
			}
		}
		if y > m.maxy {
			m.maxy = y
		}
		y++
		return nil
	})

	if err != io.EOF {
		panic(err)
	}

	return m
}
