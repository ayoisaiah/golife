package golife

import (
	"fmt"

	"github.com/gdamore/tcell"
)

// Header represents the section where
// instructions and game statistics are
// displayed
type Header struct {
	Game         *Game
	Instructions []string
	Dimension
}

// draw paints the header section on the screen
func (h *Header) draw() {
	s := h.Game.Screen
	headerStyle := tcell.StyleDefault.Foreground(h.Game.Theme.Foreground).Background(h.Game.Theme.Background)

	_, sh := s.Size()
	drawBox(s, h.x, h.y, h.w-1, sh-1, headerStyle, ' ')
	for i, v := range h.Instructions {
		emitStr(s, h.x+1, h.y+i+1, headerStyle, v)
	}
	emitStr(s, h.x+1, h.y+3, headerStyle, fmt.Sprintf("Generation: %d | Living Cells: %d", h.Game.GenerationCount, h.Game.Population))
}

// newHeader returns a new Header instance
func newHeader(x, y, w, h int) *Header {
	return &Header{
		Instructions: []string{
			"Welcome to John Conway's Game of Life",
			"Instructions: q - quit | r - restart | p - play/pause | s - step",
		},
		Dimension: Dimension{
			x,
			y,
			w,
			h,
		},
	}
}
