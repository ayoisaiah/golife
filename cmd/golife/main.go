package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// newScreen returns a screen with mouse
// support enabled
func NewScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	screen.EnableMouse()
	return screen
}

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen := NewScreen()

	defer screen.Fini()

	g := NewGame(screen)

	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					close(quit)
					return
				case tcell.KeyCtrlL:
					screen.Sync()
				}
			case *tcell.EventResize:
				screen.Sync()
			case *tcell.EventMouse:
				x, y := ev.Position()
				button := ev.Buttons()
				switch button {
				case tcell.Button1:
					// Make sure click position is within the limits of
					// the grid
					if x < len(g.Board.Grid) && y < len(g.Board.Grid[0]) {
						g.Board.Grid[x][y] = 1
					}
				}
			}
		}
	}()

	g.Init()
	g.Draw(screen)
	screen.Show()
	t := time.NewTicker(time.Second / time.Duration(g.FPS))
	for {
		select {
		case <-quit:
			return
		case <-t.C:
			screen.Clear()
			g.NextGen()
			g.Draw(screen)
			screen.Show()
		}
	}
}
