package cursor

import (
	"github.com/Shresht7/go-cli-tools/helpers"
)

//	Moves the cursor back to home position (0, 0)
func ToHome() string {
	return helpers.Escape("[H")
}

//	Moves the cursor back to the given row and column
func ToPos(row, column int) string {
	return helpers.Escape("[%d;%dH", row, column)
}

//	Moves the cursor up by n number of lines
func Up(n int) string {
	return helpers.Escape("[%dA", n)
}

//	Moves the cursor down by n number of lines
func Down(n int) string {
	return helpers.Escape("[%dB", n)
}

//	Moves the cursor right by n number of lines
func Right(n int) string {
	return helpers.Escape("[%dC", n)
}

//	Moves the cursor left by n number of lines
func Left(n int) string {
	return helpers.Escape("[%dD", n)
}

//	Moves the cursor to the nth next line
func ToNextLine(n int) string {
	return helpers.Escape("[%dE", n)
}

//	Moves the cursor to the nth prev line
func ToPrevLine(n int) string {
	return helpers.Escape("[%dF", n)
}

//	Moves the cursor to a given column position
func ToColumn(n int) string {
	return helpers.Escape("[%dG", n)
}

//	Returns the current cursor position
func RequestPosition() string {
	return helpers.Escape("[%6n")
}

//	Show the cursor
func Show() string {
	return helpers.Escape("[?25h")
}

//	Hide the cursor
func Hide() string {
	return helpers.Escape("[?25l")
}

//	? Save and Restore
