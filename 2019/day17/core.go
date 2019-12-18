package day17

import (
	"fmt"
	"strings"

	"github.com/seanhagen/advent-of-code/2019/lib2019"
	"github.com/seanhagen/advent-of-code/lib"
	"github.com/seanhagen/advent-of-code/lib/facing"
)

type tile string

const (
	tileEmpty   tile = ""
	tileScaf    tile = "#"
	tileSpace   tile = "."
	tileDead    tile = "x"
	tileVisited tile = "C"
)

var validTiles = []tile{tileScaf, tileSpace, tileDead, tileVisited}

// valid ...
func (t tile) valid() bool {
	for _, vt := range validTiles {
		if t == vt {
			return true
		}
	}
	return false
}

// Vacuum ...
type Vacuum struct {
	code  string
	brain *lib2019.Program
	mover *facing.Mover

	start map[int]map[int]tile
	sx    int
	sy    int
}

// CreateVacuum ...
func CreateVacuum(input string) (*Vacuum, error) {
	v := &Vacuum{
		code: input,
	}
	var err error

	err = v.getStartState()
	if err != nil {
		return nil, err
	}

	return v, nil
}

// CompressPath ...
func (v *Vacuum) CompressPath() string {
	p := v.solvePath()
	return v.compressPath(p)
}

// getStartState ...
func (v *Vacuum) getStartState() error {
	f := lib.LoadInput(v.code)
	p, err := lib2019.ReadProgram(f)
	if err != nil {
		return err
	}

	data := map[int]map[int]tile{}
	x, y := 0, 0
	maxX := 0

	sx, sy := 0, 0
	var fa facing.Direction

	p.SetOutputFn(func(i int) bool {
		tmp := fmt.Sprintf("%v", string(i))
		v := tile(tmp)
		if v == "\n" {
			y++
			x = 0
			return false
		}

		if facing.IsFacing(tmp) {
			sx, sy = x, y
			fa = facing.DirectionFromString(tmp)
			v = tileVisited
		}

		row, ok := data[y]
		if !ok {
			row = map[int]tile{}
		}
		row[x] = v
		data[y] = row

		if y == 0 && x > maxX {
			maxX = x
		}
		x++
		return false
	})

	err = p.Run()
	if err != nil {
		return err
	}

	var no tile
	mcnf := facing.Config{
		StartX:    sx,
		StartY:    sy,
		Facing:    fa,
		NewObj:    func() interface{} { return no },
		Type:      no,
		NumMovers: 1,
	}
	mvr, err := facing.NewMover(&mcnf)
	if err != nil {
		return err
	}

	for j := 0; j <= y; j++ {
		row := data[j]
		for i := 0; i <= maxX; i++ {
			v := row[i]
			mvr.SetAt(i, j, v)
		}
	}

	v.mover = mvr
	v.start = data
	return nil
}

// getSpecial ...
func (v Vacuum) getSpecial() map[int]map[int]int {
	checks := []tile{"#", "^", "v", "<", ">", "C"}
	fn := func(i tile) bool {
		for _, v := range checks {
			if v == i {
				return true
			}
		}
		return false
	}

	minY, maxY, minX, maxX := 100, -100, 100, -100
	for y, row := range v.start {
		if y == 0 {
			for x := range row {
				if x > maxX {
					maxX = x
				}
				if x < minY {
					minX = x
				}
			}
		}

		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	endpoints := map[int]map[int]int{}
	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			d := v.start[j][i]
			if d == "." {
				continue
			}

			count := 0
			// can check north if j > 0
			if j > 0 {
				n := v.start[j-1][i]
				if fn(n) {
					count++
				}
			}

			// can check south if j < len(v.start)-1
			if j < len(v.start)-1 {
				s := v.start[j+1][i]
				if fn(s) {
					count++
				}
			}

			// can check east if i < len(row) -1
			row := v.start[j]
			if i < len(row)-1 {
				e := v.start[j][i+1]
				if fn(e) {
					count++
				}
			}

			// can check west if i > 0
			if i > 0 {
				w := v.start[j][i-1]
				if fn(w) {
					count++
				}
			}

			// count == 1 is an endpoint ( start or end )
			// count == 3 is probably impossible for this problem
			// count == 4 is an intersection
			// count >= 5 means something is horribly broken
			if count != 2 {
				tmp, ok := endpoints[j]
				if !ok {
					tmp = map[int]int{}
				}
				tmp[i] = count
				endpoints[j] = tmp
			}
		}
	}

	return endpoints
}

// getEndpoints ...
func (v Vacuum) getEndpoints() [][]int {
	endpoints := v.getSpecial()

	out := [][]int{}
	for y, t := range endpoints {
		for x, d := range t {
			if d == 1 {
				out = append(out, []int{x, y})
			}
		}
	}

	return out
}

// Run ...
func (v Vacuum) Run(path string) error {
	f := lib.LoadInput(v.code)
	p, err := lib2019.ReadProgram(f)
	if err != nil {
		return err
	}

	path = fmt.Sprintf("%v\nn\n", path)
	prog := []rune{}
	for _, v := range strings.Split(path, "") {
		x := []rune(v)
		prog = append(prog, x[0])
	}

	idx := -1
	inputFn := func() (int, bool) {
		idx++
		ip := int(prog[idx])
		return ip, false
	}
	p.SetInputFunc(inputFn)

	var last int

	p.SetOutputFn(func(i int) bool {
		v := fmt.Sprintf("%v", string(i))
		fmt.Printf("%v", v)
		last = i
		return false
	})
	p.Replace(0, 2)
	err = p.Run()
	if err != nil {
		return err
	}

	fmt.Printf("output: %v\n", last)

	return nil
}

// solvePath ...
func (v Vacuum) solvePath() string {
	mvr := v.mover
	tmpe := v.getEndpoints()
	mx, my := mvr.Location()
	var endX, endY int
	for _, ep := range tmpe {
		if ep[0] != mx || ep[1] != my {
			endX, endY = ep[0], ep[1]
		}
	}

	path := []string{}
	mvCnt := 0
	let := map[string]string{
		"Left":  "L",
		"Right": "R",
	}

	for {
		nearby := mvr.NearCurrent()
		mvf := mvr.Facing()
		x, y := mvr.Location()

		if x == endX && y == endY {
			path = append(path, fmt.Sprintf("%v", mvCnt))
			break
		}

		// check what's in front of the robot
		ahead, ok := nearby[mvf]
		moveForward := false
		if ok {
			// ahead is a tile, let's take a look
			at := ahead.(tile)
			if at != tileSpace && at != tileEmpty {
				moveForward = true
			}
		} else {
			// ahead is empty, need to turn for sure
		}

		if !moveForward {
			// turning
			if mvCnt > 0 {
				// we've been moving, so append that to the list of directions and reset move counter
				path = append(path, fmt.Sprintf("%v", mvCnt))
				mvCnt = 0
			}

			var turn facing.Direction
			for dir, tmp := range nearby {
				t := tmp.(tile)
				if t == tileScaf {
					turn = dir
				}
			}

			turnDir := facing.TurnTowards(mvf, turn)
			tn := facing.TurnNames[turnDir]
			path = append(path, let[tn])
			mvr.TurnTowards(turn)
		} else {
			// moving forward
			// there's scaffolding ahead, move forward!
			x, y := mvr.Location()
			mvr.SetAt(x, y, tileVisited)
			mvr.MoveForward()
			mvCnt++
		}
	}

	return strings.Join(path, ",")
}

// compressPath ...
func (v Vacuum) compressPath(in string) string {
	parts := map[string]string{}
	have := []string{}
	partName := "A"
start:
	for _, v := range strings.Split(in, ",") {
		if v == "B" || v == "C" {
			for name, matcher := range parts {
				if name == v {
					continue
				}

				for n2, m2 := range parts {
					if strings.Contains(matcher, m2) && name != n2 {
						in = strings.Replace(in, name, matcher, -1)
						matcher = strings.Replace(matcher, ","+m2, "", -1)
						parts[name] = matcher

						in = strings.Replace(in, matcher, name, -1)
						in = strings.Replace(in, m2, n2, -1)
					}
				}
			}
		}

		if v == "A" || v == "B" || v == "C" {
			if len(have) < 1 {
				continue
			}

			remove := strings.Join(have, ",")
			parts[partName] = strings.Join(have, ",")
			in = strings.Replace(in, remove, partName, -1)
			partName = incrPartName(partName)
			have = []string{}
			goto start
		}

		have = append(have, v)
		if len(have) == 1 {
			continue
		}

		test := strings.Join(have, ",")
		times := strings.Count(in, test)

		if times == 1 {
			r := have[:len(have)-1]
			parts[partName] = strings.Join(r, ",")
			remove := strings.Join(r, ",")
			in = strings.Replace(in, remove, partName, -1)

			partName = incrPartName(partName)
			have = []string{}
			goto start
		}
	}

	out := fmt.Sprintf("%v\n%v\n%v\n%v", in, parts["A"], parts["B"], parts["C"])
	return out
}

func incrPartName(in string) string {
	switch in {
	case "A":
		in = "B"
	case "B":
		in = "C"
	}
	return in
}
