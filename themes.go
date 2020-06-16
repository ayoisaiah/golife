package golife

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type Theme struct {
	Name       string
	Background tcell.Color
	Foreground tcell.Color
	CellColor  tcell.Color
}

var themes = []Theme{
	ThemeWhiteOnBlack, // This is the default
	ThemeBlackOnWhite,
}

var ThemeBlackOnWhite = Theme{
	Name:       "BlackOnWhite",
	Background: tcell.NewRGBColor(255, 255, 255),
	Foreground: tcell.NewRGBColor(51, 51, 51),
	CellColor:  tcell.NewRGBColor(34, 34, 34),
}

var ThemeWhiteOnBlack = Theme{
	Name:       "WhiteOnBlack",
	Background: tcell.NewRGBColor(51, 51, 51),
	Foreground: tcell.NewRGBColor(255, 255, 255),
	CellColor:  tcell.NewRGBColor(250, 250, 250),
}

func ListThemes() {
	for _, v := range themes {
		fmt.Println(v.Name)
	}
}
