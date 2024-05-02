package codes

import (
	"strings"
)

//	=================
//	ANSI ESCAPE CODES
//	=================

// Escape Code
const ESC = "\u001b"

// Control Sequence Introducer
const CSI = ESC + "["

// Operating System Command
const OSC = ESC + "]"

// Reset Code
const RESET = CSI + "0m"

// Miscellaneous
const CTRL = "^["
const BEL = "\u0007"

// Global variable to control whether the ANSI codes are enabled or not
var IS_ENABLED = true

//	-------
//	HELPERS
//	-------

// Helper function to format the given ANSI codes.
// The last code is the closing code.
func Code(codes ...string) (open, close string) {

	//	Check if the codes are valid
	if len(codes) < 2 {
		panic("Invalid ANSI codes")
	}

	//	Opening codes: "\u001b[__;__;__m"
	openCodes := codes[:len(codes)-1]
	open = CSI + strings.Join(openCodes, ";") + "m"

	//	The closing code: "\u001b[__m"
	closeCode := codes[len(codes)-1]
	close = CSI + closeCode + "m"

	return open, close

}

// Wrap ANSI Codes around string
func Wrap(str string, codes []string) string {
	if IS_ENABLED {
		open, close := Code(codes...)
		str = open + str + close
	}
	return str
}
