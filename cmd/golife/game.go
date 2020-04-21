package main

import "github.com/gdamore/tcell"

type Game struct {
	Header *Header
	Board  *Board
	FPS    int
}

// NewGame creates and returns a new game instance
func NewGame(screen tcell.Screen) *Game {
	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))

	w, h := screen.Size()

	g := &Game{}
	headerStyle := tcell.StyleDefault.
		Foreground(tcell.ColorTeal).Background(tcell.ColorBlack)
	g.Header = NewHeader(headerStyle, 0, 0, w-1, 4)

	boardStyle := tcell.StyleDefault.
		Foreground(tcell.ColorWhiteSmoke).Background(tcell.ColorBlack)
	g.Board = NewBoard(boardStyle, 0, g.Header.h+1, w-1, h-g.Header.h)
	g.FPS = 10

	return g
}

// Draw paints the components of the Game on the screen
func (g *Game) Draw(screen tcell.Screen) {
	g.Header.Draw(screen)
	g.Board.Draw(screen)
}

// Init sets the initial pattern of cells on the grid
func (g *Game) Init() {
	g.Board.GenerateGrid(g.Board.w, g.Board.h)
}

// NextGen computes the next generation of cells in the grid
func (g *Game) NextGen() {
	livingCells := g.Board.NextGeneration()
	g.Header.GenerationCount += 1
	g.Header.LivingCellsCount = livingCells
}
