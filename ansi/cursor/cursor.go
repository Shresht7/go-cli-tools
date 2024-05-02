package cursor

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

// Moves the cursor back to home position (0, 0)
const ToHome = codes.CSI + "H"

// Moves the cursor back to the given row and column
func ToPos(row, column int) string {
	return codes.CSI + fmt.Sprintf("%d;%dH", row, column)
}

// Moves the cursor up by n number of lines
func Up(n int) string {
	return codes.CSI + fmt.Sprintf("%dA", n)
}

// Moves the cursor down by n number of lines
func Down(n int) string {
	return codes.CSI + fmt.Sprintf("%dB", n)
}

// Moves the cursor right by n number of lines
func Right(n int) string {
	return codes.CSI + fmt.Sprintf("%dC", n)
}

// Moves the cursor left by n number of lines
func Left(n int) string {
	return codes.CSI + fmt.Sprintf("%dD", n)
}

// Moves the cursor to the nth next line
func ToNextLine(n int) string {
	return codes.CSI + fmt.Sprintf("%dE", n)
}

// Moves the cursor to the nth prev line
func ToPrevLine(n int) string {
	return codes.CSI + fmt.Sprintf("%dF", n)
}

// Moves the cursor to a given column position
func ToColumn(n int) string {
	return codes.CSI + fmt.Sprintf("%dG", n)
}

// Returns the current cursor position
const RequestPosition = codes.CSI + "%6n"

// Show the cursor
const Show = codes.CSI + "?25h"

// Hide the cursor
const Hide = codes.CSI + "?25l"

// Save saves the current cursor position.
// The following commands have not been standardized and may have different functionality across different terminal applications. DEC Sequences are recommended.
func Save(mode string) string {
	if mode == "DEC" {
		return codes.ESC + "7"
	}
	return codes.CSI + "s"
}

// Restore restores the current cursor position.
func Restore(mode string) string {
	if mode == "DEC" {
		return codes.ESC + "8"
	}
	return codes.CSI + "u"
}
