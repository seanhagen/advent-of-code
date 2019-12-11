package day11

import (
	"fmt"

	"github.com/seanhagen/advent-of-code/2019/lib2019"
	"github.com/seanhagen/advent-of-code/lib/facing"
)

// Robot ...
type Robot struct {
	brain *lib2019.Program
	mover *facing.Mover
}

// NewRobot ...
func NewRobot(input string) (*Robot, error) {
	p, err := lib2019.FromString(input)
	if err != nil {
		return nil, err
	}

	newPanelFn := func() interface{} {
		return Panel{color: Black}
	}

	cnf := facing.Config{
		X:      0,
		Y:      0,
		Facing: facing.North,
		NewObj: newPanelFn,
		Type:   Panel{},
	}

	mover, err := facing.NewMover(&cnf)
	if err != nil {
		return nil, err
	}

	return &Robot{
		brain: p,
		mover: mover,
	}, nil
}

// turn ...
func (r *Robot) turn(i int) {
	r.mover.Turn(facing.Turn(i))
}

// paint ...
func (r *Robot) paint(c Color) {
	r.mover.ModifyCurrent(func(i interface{}) interface{} {
		p := i.(Panel)
		p.color = c
		p.timesPainted++
		return p
	})
}

// Run ...
func (r *Robot) Run() error {
	r.brain.SetPauseOnOutput(true)

	for {
		p := r.mover.GetCurrent().(Panel)

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

		r.paint(Color(paint))
		r.turn(turn)
	}

	return nil
}

// PanelsPainted ...
func (r *Robot) PanelsPainted() int {
	t := 0
	r.mover.Iterate(func(x, y int, tmp interface{}) {
		p := tmp.(Panel)
		if p.timesPainted > 0 {
			t++
		}
	})
	return t
}

// ColorPanel ...
func (r *Robot) ColorPanel(x, y int, c Color) {
	r.mover.ModifyAt(x, y, func(t interface{}) interface{} {
		p := t.(Panel)
		p.color = c
		return p
	})
}

// Print ...
func (r Robot) Print() {
	r.mover.Print(func(i interface{}) {
		p := i.(Panel)
		if p.color == Black {
			fmt.Printf("■")
		} else {
			fmt.Printf(" ")
		}
	})
}
