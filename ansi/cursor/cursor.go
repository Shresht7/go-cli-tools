package cursor

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

//	Moves the cursor back to home position (0, 0)
func ToHome() string {
	return codes.Escape("[H")
}

//	Moves the cursor back to the given row and column
func ToPos(row, column int) string {
	return codes.Escape("[%d;%dH", row, column)
}

//	Moves the cursor up by n number of lines
func Up(n int) string {
	return codes.Escape("[%dA", n)
}

//	Moves the cursor down by n number of lines
func Down(n int) string {
	return codes.Escape("[%dB", n)
}

//	Moves the cursor right by n number of lines
func Right(n int) string {
	return codes.Escape("[%dC", n)
}

//	Moves the cursor left by n number of lines
func Left(n int) string {
	return codes.Escape("[%dD", n)
}

//	Moves the cursor to the nth next line
func ToNextLine(n int) string {
	return codes.Escape("[%dE", n)
}

//	Moves the cursor to the nth prev line
func ToPrevLine(n int) string {
	return codes.Escape("[%dF", n)
}

//	Moves the cursor to a given column position
func ToColumn(n int) string {
	return codes.Escape("[%dG", n)
}

//	Returns the current cursor position
func RequestPosition() string {
	return codes.Escape("[%6n")
}

//	Show the cursor
func Show() string {
	return codes.Escape("[?25h")
}

//	Hide the cursor
func Hide() string {
	return codes.Escape("[?25l")
}

//	? Save and Restore
