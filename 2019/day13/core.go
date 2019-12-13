package day13

import (
	"fmt"
)

// Tile is a tile in the game
type Tile string

// EmptyTile ...
const EmptyTile Tile = " "

// WallTile ...
const WallTile Tile = "w"

// BlockTile ...
const BlockTile Tile = "b"

// PaddleTile ...
const PaddleTile Tile = "-"

// BallTile ...
const BallTile Tile = "o"

var idToTile = map[int]Tile{
	0: EmptyTile,
	1: WallTile,
	2: BlockTile,
	3: PaddleTile,
	4: BallTile,
}

// Game ...
type Game struct {
	screen map[int]map[int]Tile

	minX int
	maxX int

	minY int
	maxY int

	score int

	bX int
	pX int
}

// NewGame ...
func NewGame(input []int) (*Game, error) {
	if len(input)%3 != 0 {
		return nil, fmt.Errorf("length of input not divisible by 3")
	}

	g := &Game{screen: map[int]map[int]Tile{}}

	for i := range input {
		if (i+1)%3 == 0 {
			x := input[i-2]
			y := input[i-1]
			v := input[i]
			g.SetTile(x, y, v)
		}
	}

	return g, nil
}

// SetTile ...
func (g *Game) SetTile(x, y, t int) {
	if x == -1 && y == 0 {
		g.score = t
		return
	}

	row, ok := g.screen[y]
	if !ok {
		row = map[int]Tile{}
	}

	tile, ok := idToTile[t]
	if !ok {
		tile = EmptyTile
	}

	row[x] = tile
	g.screen[y] = row

	if tile == PaddleTile {
		g.pX = x
	}

	if tile == BallTile {
		g.bX = x
	}

	if x > g.maxX {
		g.maxX = x
	}
	if x < g.minX {
		g.minX = x
	}
	if y > g.maxY {
		g.maxY = y
	}
	if y < g.minY {
		g.minY = y
	}
}

// CountTileType ...
func (g Game) CountTileType(t Tile) int {
	sum := 0
	for _, row := range g.screen {
		for _, v := range row {
			if v == t {
				sum++
			}
		}
	}
	return sum
}

// Print ...
func (g Game) Print() {
	fmt.Printf("Score: %v\n", g.score)

	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			t, ok := g.screen[j][i]
			if !ok {
				t = EmptyTile
			}
			fmt.Printf("%v", t)
		}
		fmt.Printf("\n")
	}
}

// GetBallPosition ...
func (g Game) GetBallPosition() int {
	return g.bX
}

// GetPaddlePosition ...
func (g Game) GetPaddlePosition() int {
	return g.pX
}

// GetScore ...
func (g Game) GetScore() int {
	return g.score
}
