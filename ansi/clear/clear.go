package clear

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

//	Clears the screen
func Screen() string {
	return codes.Escape("[J")
}

//	Clears cursor and below
func ClearCursorAndBelow() string {
	return codes.Escape("[0J")
}

//	Clears cursor and above
func ClearCursorAndAbove() string {
	return codes.Escape("[1J")
}

//	Clears the entire screen
func EntireScreen() string {
	return codes.Escape("[2J")
}

//	Clear line
func Line() string {
	return codes.Escape("[K")
}

//	Clear line from cursor
func ClearLineFromCursor() string {
	return codes.Escape("[0K")
}

//	Clear line to cursor
func ClearLineToCursor() string {
	return codes.Escape("[1K")
}

//	Clear entire line
func ClearEntireLine() string {
	return codes.Escape("[2K")
}
