package helpers

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

// Calculate the visual width of a string (the number of columns required to display it)
// considering that ansi escape codes do not take up any columns.
func StringWidth(s string) int {

	// Strip ansi escape codes
	s = codes.Strip(s)

	// ? This does not handle characters that take up more than one column -
	// ? For example, emojis, or characters from other languages. My goal for this
	// ? project is to have zero dependencies, so the implementation for this is
	// ? fairly simple.

	width := 0
	for range s {
		width++
	}

	// Return the length of the string
	return width
}
