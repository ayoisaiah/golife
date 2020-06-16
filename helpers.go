package golife

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, r rune) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Draw horizontal lines
	for row := x1; row <= x2; row++ {
		s.SetContent(row, y1, tcell.RuneHLine, nil, style)
		s.SetContent(row, y2, tcell.RuneHLine, nil, style)
	}

	// Draw vertical lines
	for col := y1 + 1; col < y2; col++ {
		s.SetContent(x1, col, tcell.RuneVLine, nil, style)
		s.SetContent(x2, col, tcell.RuneVLine, nil, style)
	}

	if y1 != y2 && x1 != x2 {
		// Only add corners if we need to
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	for col := y1 + 1; col < y2; col++ {
		for row := x1 + 1; row < x2; row++ {
			s.SetContent(row, col, ' ', nil, style)
		}
	}
}

// newScreen returns a screen with mouse support enabled
func newScreen() (tcell.Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	screen.EnableMouse()
	return screen, nil
}

// containsInt checks if a integer is present in
// a slice
func containsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// containsStr checks if a integer is present in
// a slice
func containsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// fetch accepts a URL and returns the response
// body as a byte slice
func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to read from URL")
	}

	return ioutil.ReadAll(resp.Body)
}

// isValidURL verifies if a string is a
// valid url
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
