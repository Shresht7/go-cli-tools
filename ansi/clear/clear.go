// Provides ANSI escape codes for clearing the terminal screen
package clear

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

const (
	// Clears the screen
	Screen = codes.CSI + "J"

	// Clears cursor and below
	CursorAndBelow = codes.CSI + "0J"

	// Clears cursor and above
	CursorAndAbove = codes.CSI + "1J"

	// Clears the entire screen
	EntireScreen = codes.CSI + "2J"

	// Clears saved lines
	SavedLines = codes.CSI + "3J"

	// Clears the current line
	Line = codes.CSI + "K"

	// Clears the line, starting from the cursor to the end
	LineFromCursor = codes.CSI + "0K"

	// Clears the line, starting from the beginning and ending on the cursor
	LineToCursor = codes.CSI + "1K"

	//	Clears the entire line
	EntireLine = codes.CSI + "2K"
)
