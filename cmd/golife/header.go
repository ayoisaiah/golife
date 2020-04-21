package main

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type Header struct {
	Style            tcell.Style
	Instructions     []string
	GenerationCount  int
	LivingCellsCount int
	Size
}

type Size struct {
	x, y, w, h int
}

func (h *Header) Draw(s tcell.Screen) {
	drawBox(s, h.x, h.y, h.w, h.h, h.Style, ' ')
	for i, v := range h.Instructions {
		emitStr(s, h.x+1, h.y+i+1, h.Style.Foreground(tcell.ColorWhite), v)
	}
	emitStr(s, h.x+1, h.y+3, h.Style.Foreground(tcell.ColorWhite), fmt.Sprintf("Generation: %d | Living Cells: %d", h.GenerationCount, h.LivingCellsCount))
}

func NewHeader(style tcell.Style, x, y, w, h int) *Header {
	return &Header{
		Style: style,
		Instructions: []string{
			"Welcome to John Conway's Game of Life",
			"Instructions: Esc - quit",
		},
		Size: Size{
			x,
			y,
			w,
			h,
		},
	}
}
