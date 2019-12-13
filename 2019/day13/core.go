package day13

import "fmt"

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
