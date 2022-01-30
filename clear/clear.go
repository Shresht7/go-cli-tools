package clear

import (
	"github.com/Shresht7/go-cli-tools/helpers"
)

//	Clears the screen
func Screen() string {
	return helpers.Escape("[J")
}

//	Clears cursor and below
func ClearCursorAndBelow() string {
	return helpers.Escape("[0J")
}

//	Clears cursor and above
func ClearCursorAndAbove() string {
	return helpers.Escape("[1J")
}

//	Clears the entire screen
func EntireScreen() string {
	return helpers.Escape("[2J")
}

//	Clear line
func Line() string {
	return helpers.Escape("[K")
}

//	Clear line from cursor
func ClearLineFromCursor() string {
	return helpers.Escape("[0K")
}

//	Clear line to cursor
func ClearLineToCursor() string {
	return helpers.Escape("[1K")
}

//	Clear entire line
func ClearEntireLine() string {
	return helpers.Escape("[2K")
}
