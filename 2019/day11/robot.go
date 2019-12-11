package day11

import (
	"fmt"

	"github.com/seanhagen/advent-of-code/2019/lib2019"
)

type Robot struct {
	panels map[int]map[int]Panel
	brain  *lib2019.Program

	facing Facing
	x      int
	xdir   int
	y      int
	ydir   int

	paintedWhite int
}

func NewRobot(input string) (*Robot, error) {
	p, err := lib2019.FromString(input)
	if err != nil {
		return nil, err
	}

	panels := map[int]map[int]Panel{
		0: map[int]Panel{
			0: Panel{color: Black},
		},
	}

	return &Robot{
		panels: panels,
		brain:  p,
		facing: Up,
		x:      0,
		xdir:   dir[Up][0],
		y:      0,
		ydir:   dir[Up][1],
	}, nil
}

// turn ...
func (r *Robot) turn(i int) {
	switch i {
	case left:
		// fmt.Printf("turning left!\n")
		r.turnLeft()

	case right:
		// fmt.Printf("turning right!\n")
		r.turnRight()
	}
	r.move()
	// fmt.Printf("now facing: %v, at position %v,%v\n", string(r.facing), r.x, r.y)
}

// move ...
func (r *Robot) move() {
	r.x += r.xdir
	r.y += r.ydir

	// fmt.Printf("moved, now at %v,%v\n", r.x, r.y)

	xp, ok := r.panels[r.x]
	if !ok {
		xp = map[int]Panel{}
	}

	p, ok := xp[r.y]
	if !ok {
		p = Panel{color: Black}
		xp[r.y] = p
		r.panels[r.x] = xp
	}
}

// turnRight ...
func (r *Robot) turnRight() {
	switch r.facing {
	case Up:
		r.facing = Right
	case Right:
		r.facing = Down
	case Down:
		r.facing = Left
	case Left:
		r.facing = Up
	}
	r.xdir, r.ydir = dir[r.facing][0], dir[r.facing][1]
}

// turnLeft ...
func (r *Robot) turnLeft() {
	switch r.facing {
	case Up:
		r.facing = Left
	case Left:
		r.facing = Down
	case Down:
		r.facing = Right
	case Right:
		r.facing = Up
	}

	r.xdir, r.ydir = dir[r.facing][0], dir[r.facing][1]
}

// paint ...
func (r *Robot) paint(c Color) {
	xp, ok := r.panels[r.x]
	if !ok {
		xp = map[int]Panel{}
	}

	p, ok := xp[r.y]
	if !ok {
		p = Panel{color: Black, timesPainted: 0}
	}

	p.timesPainted++
	p.color = c

	xp[r.y] = p
	r.panels[r.x] = xp
}

// Run ...
func (r *Robot) Run() error {
	r.brain.SetPauseOnOutput(true)

	// fmt.Printf("now facing: %v, at position %v,%v\n", string(r.facing), r.x, r.y)

	for {
		p, ok := r.panels[r.x][r.y]
		if !ok {
			return fmt.Errorf("something went horribly wrong, no panel at %v,%v", r.x, r.y)
		}

		r.brain.AddInput(p.color.ToInt())
		err := r.brain.Run()
		if err == lib2019.ErrHalt {
			break
		}
		if err != nil {
			return fmt.Errorf("error running brain: %v", err)
		}

		err = r.brain.Run()
		if err == lib2019.ErrHalt {
			break
		}
		if err != nil {
			return fmt.Errorf("error running brain: %v", err)
		}

		outputs := r.brain.GetOutputs()
		l := len(outputs)

		paint, turn := outputs[l-2], outputs[l-1]
		if paint != 1 && paint != 0 {
			return fmt.Errorf("invalid paint number: %v", paint)
		}
		// fmt.Printf("paint: %v\tturn: %v\n", paint, turn)
		r.paint(Color(paint))
		r.turn(turn)
	}

	return nil
}

// PanelsPainted ...
func (r *Robot) PanelsPainted() int {
	t := 0
	for _, row := range r.panels {
		for _, p := range row {
			if p.timesPainted > 0 {
				t++
			}
		}
	}
	return t
}

// ColorPanel ...
func (r *Robot) ColorPanel(x, y int, c Color) {
	xp, ok := r.panels[x]
	if !ok {
		xp = map[int]Panel{}
	}

	p, ok := xp[y]
	if !ok {
		p = Panel{color: Black, timesPainted: 0}
	}

	p.color = c

	xp[y] = p
	r.panels[x] = xp
}

// Print ...
func (r Robot) Print() {
	minX := 10
	maxX := -10

	minY := 10
	maxY := -10

	for x, col := range r.panels {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		for y := range col {
			if y < minY {
				minY = y
			}
			if y > minY {
				maxY = y
			}
		}
	}

	fmt.Printf("min: %v,%v -> max: %v,%v\n\n", minX, minY, maxX, maxY)

	for j := 1; j >= -6; j-- {
		for i := -1; i <= 43; i++ {
			row := r.panels[i]
			p, ok := row[j]
			if !ok {
				fmt.Printf(" ")
			} else {
				if p.color == Black {
					fmt.Printf("■")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf(" -- row %v\n", j)
	}
	// jzpjragj
	// JZPJRAGJ
}
