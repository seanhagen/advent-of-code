package day15

import (
	"fmt"

	"github.com/seanhagen/advent-of-code/2019/lib2019"
)

type statusCode int

const (
	statusHitWall  statusCode = 0
	statusOneStep  statusCode = 1
	statusAtOxySys statusCode = 2
)

type tileType string

const (
	tileEmpty tileType = "."
	tileWall  tileType = "#"
	tileOxy   tileType = "o"
)

var outputToTileType = map[int]tileType{
	0: tileWall,
	1: tileEmpty,
	2: tileOxy,
}

type move int

const (
	north move = 1
	south move = 2
	west  move = 3
	east  move = 4
)

var reverseMove = map[move]move{
	north: south,
	south: north,
	west:  east,
	east:  west,
}

var moveNames = map[move]string{
	north: "north",
	south: "south",
	west:  "west",
	east:  "east",
}

var moveDir = map[move][]int{
	north: []int{0, 1},
	south: []int{0, -1},
	west:  []int{1, 0},
	east:  []int{-1, 0},
}

type tile struct {
	x int
	y int

	t tileType

	idx int

	neighbours map[move]*tile
}

// RepairDroid ...
type RepairDroid struct {
	brain *lib2019.Program

	//         y       x   seen
	tiles map[int]map[int]*tile

	xpos int
	ypos int
	xdir int
	ydir int
	minX int
	maxX int
	minY int
	maxY int

	nextDir move
	idx     int

	backtrackMode bool
	btr           bool // back-track mode running

	shouldExit bool
}

// CreateRepairDroid ...
func CreateRepairDroid(p *lib2019.Program) (*RepairDroid, error) {
	robo := &RepairDroid{
		tiles: map[int]map[int]*tile{
			0: map[int]*tile{
				0: &tile{0, 0, tileEmpty, 0, map[move]*tile{}},
			},
		},
		nextDir: north,
		xdir:    0,
		ydir:    1,
	}

	p.SetInputFunc(robo.inputFn)
	p.SetOutputFn(robo.outputFn)
	robo.brain = p

	return robo, nil
}

// Print ...
func (r RepairDroid) Print() {
	for j := r.maxY + 1; j >= r.minY-1; j-- {
		for i := r.minX - 1; i <= r.maxX+1; i++ {
			if i == 0 && j == 0 {
				fmt.Printf("s")
			} else if i == r.xpos && j == r.ypos {
				fmt.Printf("d")
			} else {
				t, ok := r.tiles[j][i]
				if !ok {
					fmt.Printf(" ")
				} else {
					fmt.Printf("%v", t.t)
				}

			}

		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// inputFn ...
func (r *RepairDroid) inputFn() int {
	switch r.nextDir {
	case north:
		return 1
	case south:
		return 2
	case west:
		return 3
	case east:
		return 4
	}
	// default to going north
	return 1
}

// outputFn ...
func (r *RepairDroid) outputFn(in int) bool {
	nt := outputToTileType[in]
	r.addNewTile(nt)

	switch outputToTileType[in] {
	case tileEmpty:
		// moved and found an empty tile
		r.move()
		r.turn()

	case tileOxy:
		// moved and found oxygen system
		r.foundOxygen()

	case tileWall:
		// didn't move, wall in the way
		r.turn()
	}
	return true
}

// addNewTile ...
func (r *RepairDroid) addNewTile(nt tileType) {

	current := r.tiles[r.ypos][r.xpos]
	nx := r.xpos + r.xdir
	ny := r.ypos + r.ydir

	tcol, ok := r.tiles[ny]
	if !ok {
		tcol = map[int]*tile{}
	}

	t, ok := tcol[nx]
	if !ok {
		re := reverseMove[r.nextDir]
		t = &tile{
			x:   nx,
			y:   ny,
			t:   nt,
			idx: r.idx,
			neighbours: map[move]*tile{
				re: current,
			},
		}
		tcol[nx] = t
		r.tiles[ny] = tcol
		r.idx++
		r.btr = false
		r.backtrackMode = false
	}
	current.neighbours[r.nextDir] = t
}

// foundEmptyTile ...
func (r *RepairDroid) move() {
	// don't change direction, keep going

	nx := r.xpos + r.xdir
	ny := r.ypos + r.ydir

	r.xpos = nx
	r.ypos = ny

	if nx > r.maxX {
		r.maxX = nx
	}
	if nx < r.minX {
		r.minX = nx
	}
	if ny > r.maxY {
		r.maxY = ny
	}
	if ny < r.minY {
		r.minY = ny
	}
}

// foundWall ...
func (r *RepairDroid) turn() {
	// need new direction
	// prioritize directions we haven't gone yet
	//   ie, if we went north, hit a dead end and then went south, don't want to end up going back and forth
	//   in the same hallway when we hit the bottom -- should go left or right instead
	//
	//   this should also help us turn left if we get to the top of a north-south hallway and the only direction
	//   left to go is east
	//
	// when turning, turn clockwise to choose new direction

	if r.backtrackMode {
		r.moveBacktrack()
		return
	}
	r.moveNormal()
}

// moveBacktrack ...
func (r *RepairDroid) moveBacktrack() {

	current := r.tiles[r.ypos][r.xpos]

	nn, nok := current.neighbours[north]
	sn, sok := current.neighbours[south]
	wn, wok := current.neighbours[west]
	en, eok := current.neighbours[east]

	if r.btr {
		if !nok {
			// haven't gone north from here yet, give it a go
			r.nextDir = north
			goto setdir
		} else if nok && !wok {
			// north has been expored, but west hasn't, go that way
			r.nextDir = west
			goto setdir
		} else if nok && wok && !sok {
			// north and west have been expored, but south hasn't
			r.nextDir = south
			goto setdir
		} else if nok && sok && wok && !eok {
			// last try, go east
			r.nextDir = east
			goto setdir
		}

		var nextDir move
		for dir, ni := range current.neighbours {
			if ni.idx < current.idx && ni.t != tileWall {
				nextDir = dir
			}
		}

		r.nextDir = nextDir
		goto setdir
	}

	r.btr = true
	if nok && nn.t != tileWall {
		r.nextDir = north
		goto setdir
	}

	if sok && sn.t != tileWall {
		r.nextDir = south
		goto setdir
	}

	if wok && wn.t != tileWall {
		r.nextDir = west
		goto setdir
	}

	if eok && en.t != tileWall {
		r.nextDir = east
		goto setdir
	}

	r.shouldExit = true
	return

setdir:
	dirs := moveDir[r.nextDir]
	r.xdir = dirs[0]
	r.ydir = dirs[1]

}

func (r *RepairDroid) moveNormal() {
	current := r.tiles[r.ypos][r.xpos]

	_, nok := current.neighbours[north]
	_, sok := current.neighbours[south]
	_, wok := current.neighbours[west]
	_, eok := current.neighbours[east]

	var turnTo move

	if !nok {
		// haven't gone north from here yet, give it a go
		turnTo = north
	} else if nok && !wok {
		// north has been expored, but west hasn't, go that way
		turnTo = west
	} else if nok && wok && !sok {
		// north and west have been expored, but south hasn't
		turnTo = south
	} else if nok && sok && wok && !eok {
		// last try, go east
		turnTo = east
	} else {
		r.backtrackMode = true
		return
	}

	r.nextDir = turnTo
	dirs := moveDir[turnTo]
	r.xdir = dirs[0]
	r.ydir = dirs[1]
}

// foundOxygen ...
func (r *RepairDroid) foundOxygen() {
}

// FindOxygenSystem ...
func (r *RepairDroid) FindOxygenSystem() error {
	current := r.tiles[r.ypos][r.xpos]
	for {
		if current.t == tileOxy {
			break
		}

		err := r.brain.Run()
		if err != nil {
			return err
		}

		if r.shouldExit {
			break
		}
	}

	// spew.Dump(r)
	r.Print()

	return nil
}
