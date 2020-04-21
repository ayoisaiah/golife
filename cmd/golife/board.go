package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

// CellState represents the state of a cell
// 0 is dead and 1 is alive
type CellState int

const (
	dead CellState = iota
	alive
)

type Board struct {
	Style tcell.Style
	Grid  [][]CellState
	Size
}

func NewBoard(style tcell.Style, x, y, w, h int) *Board {
	return &Board{
		Style: style,
		Size: Size{
			x,
			y,
			w,
			h,
		},
	}
}

func (b *Board) Draw(s tcell.Screen) {
	// Draw borders
	drawBox(s, b.x, b.y, b.w, b.h, b.Style, ' ')

	// Draw cells
	// Adding 1 to the x and y positions so that
	// the cells do not start at the borders
	for i := b.x + 1; i < len(b.Grid); i++ {
		for j := b.y + 1; j < len(b.Grid[0]); j++ {
			value := b.Grid[i][j]
			if value == 1 {
				s.SetContent(i, j, tcell.RuneBlock, nil, tcell.StyleDefault)
			}
		}
	}
}

// GenerateGrid generates the a random
// grid for the game
func (b *Board) GenerateGrid(w, h int) {
	grid := make([][]CellState, w)
	for i := 0; i < w; i++ {
		grid[i] = make([]CellState, h)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			grid[i][j] = CellState(rand.Int() % 2)
		}
	}

	b.Grid = grid
}

func (b *Board) countNeighbours(x int, y int) int {
	var sum int
	numberOfRows := len(b.Grid)
	numberOfCols := len(b.Grid[0])

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// skip over the current cell
			if i == 0 && j == 0 {
				continue
			}

			row := (x + i + numberOfRows) % numberOfRows
			col := (y + j + numberOfCols) % numberOfCols

			sum += int(b.Grid[row][col])
		}
	}

	return sum
}

// NextGeneration calculates the next generation
// based on the current one and returns the
// number of living cells
func (b *Board) NextGeneration() int {
	nextGrid := make([][]CellState, len(b.Grid))
	for i := 0; i < len(b.Grid); i++ {
		nextGrid[i] = make([]CellState, len(b.Grid[0]))
	}

	var livingCells int
	for i := 0; i < len(b.Grid); i++ {
		for j := 0; j < len(b.Grid[0]); j++ {
			value := b.Grid[i][j]
			neighbours := b.countNeighbours(i, j)
			if value == dead && neighbours == 3 {
				nextGrid[i][j] = alive
			} else if value == 1 && (neighbours < 2 || neighbours > 3) {
				nextGrid[i][j] = dead
			} else {
				nextGrid[i][j] = value
			}

			if nextGrid[i][j] == alive {
				livingCells++
			}
		}
	}

	b.Grid = nextGrid
	return livingCells
}
