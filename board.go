package golife

import (
	"fmt"
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

// Board represents the game board
type Board struct {
	Game *Game
	Grid [][]CellState
	Dimension
}

// newBoard returns a new Board instance
func newBoard(x, y, w, h int) *Board {
	return &Board{
		Dimension: Dimension{
			x,
			y,
			w,
			h,
		},
	}
}

// draw paints the Board grid to the screen
func (b *Board) draw() {
	s := b.Game.Screen

	for i := b.y; i < len(b.Grid); i++ {
		for j := b.x; j < len(b.Grid[0]); j++ {
			value := b.Grid[i][j]
			if value == 1 {
				cellStyle := tcell.StyleDefault.Foreground(b.Game.CellColor).Background(b.Game.Background)
				s.SetContent(j*2, i, tcell.RuneBlock, nil, cellStyle)
				s.SetContent(j*2+1, i, tcell.RuneBlock, nil, cellStyle)
			}
		}
	}
}

// generateGrid generates the grid for the game
func (b *Board) generateGrid() {
	grid := make([][]CellState, b.h)
	for i := 0; i < b.h; i++ {
		grid[i] = make([]CellState, b.w)
	}

	b.Grid = grid
}

// random populates the grid with a random
// arrangement of cells
func (b *Board) random() {
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			b.Grid[i][j] = CellState(rand.Int() % 2)
		}
	}
}

// preset populates the grid with a preset
// arrangement of cells
func (b *Board) preset() error {
	p := b.Game.PresetPattern
	ph := len(p.Cells)
	if ph == 0 {
		return fmt.Errorf("Something went wrong. Please check the pattern file or URL")
	}
	pw := len(p.Cells[0])

	if pw > b.w || ph > b.h {
		return fmt.Errorf("The current board is not big enough for this preset")
	}

	x := (b.w*2/2 - pw) / 2
	y := (b.h - ph) / 2

	for i := 0; i < ph; i++ {
		for j := 0; j < pw; j++ {
			if y+i < b.h && x+j < b.w*2 {
				b.Grid[y+i][x+j] = CellState(p.Cells[i][j])
			}
		}
	}

	return nil
}

// countNeighbours determines the number of live
// neighbours a cell has depending on whether
// wrapping is enabled or not
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

			if b.Game.EdgeWrap {
				row := (x + i + numberOfRows) % numberOfRows
				col := (y + j + numberOfCols) % numberOfCols

				sum += int(b.Grid[row][col])
			} else {
				row := x + i
				col := y + j

				if row >= 0 && col >= 0 && row < numberOfRows && col < numberOfCols {
					sum += int(b.Grid[row][col])
				}
			}
		}
	}

	return sum
}

// nextGeneration calculates the next generation
// based on the current one and returns the
// number of living cells
func (b *Board) nextGeneration() int {
	nextGrid := make([][]CellState, len(b.Grid))
	for i := 0; i < len(b.Grid); i++ {
		nextGrid[i] = make([]CellState, len(b.Grid[0]))
	}

	var livingCells int
	for i := 0; i < len(b.Grid); i++ {
		for j := 0; j < len(b.Grid[0]); j++ {
			value := b.Grid[i][j]
			neighbours := b.countNeighbours(i, j)
			if value == dead && containsInt(b.Game.Rule.Born, neighbours) {
				nextGrid[i][j] = alive
			} else if value == 1 && !containsInt(b.Game.Rule.Survives, neighbours) {
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
