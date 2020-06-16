package golife

import (
	"fmt"
	"time"

	"github.com/ayoisaiah/life"
	"github.com/gdamore/tcell"
	"github.com/urfave/cli/v2"
)

// Dimension represents the start coordinates
// as well as width and height of an entity
type Dimension struct {
	x, y, w, h int
}

// Stats is the current stats of the game
type Stats struct {
	GenerationCount int
	Population      int
}

// Game represents a game of life simulation
type Game struct {
	Screen        tcell.Screen
	PresetPattern *life.Pattern
	Header        *Header
	Board         *Board
	Paused        bool
	EdgeWrap      bool
	FPS           int
	Rule          Rule
	Theme
	Stats
}

// headerHeight is the height of the header
const headerHeight = 5

// NewGame creates and returns a new game instance
func NewGame(c *cli.Context) (*Game, error) {
	g := &Game{}

	screen, err := newScreen()
	if err != nil {
		return nil, err
	}

	g.SelectTheme(c.String("theme"))
	g.SelectRule(c.String("rule"))

	screen.SetStyle(tcell.StyleDefault.
		Foreground(g.Theme.Foreground).
		Background(g.Theme.Background))

	g.Screen = screen
	w, h := g.Screen.Size()

	g.Header = newHeader(0, h-headerHeight, w, headerHeight)
	g.Header.Game = g

	g.Board = newBoard(0, 0, w/2, h-headerHeight)
	g.Board.Game = g

	g.FPS = c.Int("refresh-rate")
	g.EdgeWrap = c.Bool("wrap")
	g.Paused = false

	filePath := c.String("file")
	presetURL := c.String("url")

	if filePath != "" && presetURL != "" {
		return nil, fmt.Errorf("Cannot load preset from file and url at the same time")
	}

	if filePath != "" {
		p, err := life.PresetFromFile(filePath)
		if err != nil {
			return nil, err
		}

		g.PresetPattern = p
	}

	if presetURL != "" {
		inputFormat := c.String("input-format")
		p, err := life.PresetFromURL(presetURL, inputFormat)
		if err != nil {
			return nil, err
		}

		g.PresetPattern = p
	}

	return g, nil
}

// Draw paints the components of the Game on the screen
func (g *Game) Draw() {
	g.Header.draw()
	g.Board.draw()
}

// SelectTheme sets the theme colours
func (g *Game) SelectTheme(name string) {
	var themeIndex int
	for i, v := range themes {
		if v.Name == name {
			themeIndex = i
			break
		}
	}

	g.Theme = themes[themeIndex]
}

// SelectRule selects a rule from the rulelist
// The default is Conway's Game of Life
func (g *Game) SelectRule(name string) {
	var ruleIndex int
	for i, v := range rules {
		if v.Name == name {
			ruleIndex = i
			break
		}
	}

	g.Rule = rules[ruleIndex]
}

// Init sets the initial pattern of cells on the grid
// and resets the generation count
func (g *Game) Init() error {
	g.Board.generateGrid()
	g.GenerationCount = 0
	if g.PresetPattern != nil {
		return g.Board.preset()
	} else {
		g.Board.random()
	}

	return nil
}

// TogglePause pauses or plays the game
func (g *Game) TogglePause() {
	g.Paused = !g.Paused
}

// NextGen computes the next generation of cells in the grid
func (g *Game) NextGen() {
	livingCells := g.Board.nextGeneration()
	g.GenerationCount += 1
	g.Population = livingCells
}

// Step pauses the game and displays the next
// generation
func (g *Game) Step() {
	g.Paused = true
	g.NextGen()
}

// Start begins the game loop
func (g *Game) Start() error {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	defer g.Screen.Fini()

	quit := make(chan struct{})
	go func() {
		for {
			ev := g.Screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					close(quit)
					return
				}
				switch ev.Rune() {
				case 'q':
					close(quit)
				case 'r':
					g.Init()
				case 'p':
					g.TogglePause()
				case 's':
					g.Step()
					g.Draw()
				}
			case *tcell.EventResize:
				g.Screen.Sync()
			case *tcell.EventMouse:
				x, y := ev.Position()
				x = x / 2
				button := ev.Buttons()
				switch button {
				case tcell.Button1:
					// Make sure cursor position is within the limits of
					// the grid
					if y < len(g.Board.Grid) && x < len(g.Board.Grid[0]) {
						g.Board.Grid[y][x] = 1
						g.Draw()
						g.Screen.Show()
					}
				}
			}
		}
	}()

	err := g.Init()
	if err != nil {
		return err
	}

	t := time.NewTicker(time.Second / time.Duration(g.FPS))
	for {
		select {
		case <-quit:
			return nil
		case <-t.C:
			g.Screen.Clear()
			if !g.Paused {
				g.NextGen()
			}
			g.Draw()
			g.Screen.Show()
		}
	}
}
