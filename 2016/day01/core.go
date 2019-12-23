package day01

import (
	"strconv"
	"strings"

	"github.com/seanhagen/advent-of-code/lib"
	"github.com/seanhagen/advent-of-code/lib/facing"
)

type Answer struct {
	mvr  *facing.Mover
	dirs []string
}

func answerFromDirString(input string) (*Answer, error) {
	input = strings.Replace(input, "\n", "", -1)
	bits := strings.Split(input, ", ")

	ans := &Answer{dirs: bits}

	cnf := &facing.Config{
		NewObj: ans.newTile,
		Type:   0,
	}

	mvr, err := facing.NewMover(cnf)
	if err != nil {
		return nil, err
	}

	ans.mvr = mvr
	return ans, nil
}

// NewAnswer ...
func NewAnswer(path string) (*Answer, error) {
	input, err := lib.GetString(path)
	if err != nil {
		return nil, err
	}
	return answerFromDirString(input)
}

// newTile ...
func (a Answer) newTile() interface{} {
	return 0
}

// FirstTwice ...
func (a Answer) FirstTwice() (int, error) {
	ans := 0
	for _, d := range a.dirs {
		bits := strings.Split(d, "")

		dir := bits[0]
		rest, err := strconv.Atoi(strings.Join(bits[1:], ""))
		if err != nil {
			return ans, err
		}

		var t facing.Turn
		switch dir {
		case "R":
			t = facing.Right
		default:
			t = facing.Left
		}
		a.mvr.Turn(t)

		for i := 0; i < rest; i++ {
			a.mvr.MoveForward()
			cur := a.mvr.GetCurrent()
			i := cur.(int)
			i++

			if i > 1 {
				goto done
			}
			a.mvr.SetCurent(i)
		}
	}

done:

	x, y := a.mvr.Location()
	ans = lib.Abs(x) + lib.Abs(y)

	return ans, nil
}

// Process ...
func (a Answer) Process() (int, error) {
	ans := 0
	for _, d := range a.dirs {
		bits := strings.Split(d, "")

		dir := bits[0]
		rest, err := strconv.Atoi(strings.Join(bits[1:], ""))
		if err != nil {
			return ans, err
		}

		var t facing.Turn
		switch dir {
		case "R":
			t = facing.Right
		default:
			t = facing.Left
		}
		a.mvr.Turn(t)

		for i := 0; i < rest; i++ {
			a.mvr.MoveForward()
		}
	}

	x, y := a.mvr.Location()
	ans = lib.Abs(x) + lib.Abs(y)

	return ans, nil
}
