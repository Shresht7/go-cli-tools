package clear

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

const (
	//	Clears the screen
	Screen = codes.CSI + "J"

	//	Clears cursor and below
	CursorAndBelow = codes.CSI + "0J"

	//	Clears cursor and above
	CursorAndAbove = codes.CSI + "1J"

	//	Clears the entire screen
	EntireScreen = codes.CSI + "2J"

	//	Clears saved lines
	SavedLines = codes.CSI + "3J"

	//	Clear line
	Line = codes.CSI + "K"

	//	Clear line from cursor
	LineFromCursor = codes.CSI + "0K"

	//	Clear line to cursor
	LineToCursor = codes.CSI + "1K"

	//	Clear entire line
	EntireLine = codes.CSI + "2K"
)
