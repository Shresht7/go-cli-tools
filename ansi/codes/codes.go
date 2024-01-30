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

//	-------
//	HELPERS
//	-------

type ANSICode struct {
	open  string
	close string
}

// Helper function to format the given ANSI codes.
// The last code is the closing code.
func Code(codes ...string) *ANSICode {

	//	Check if the codes are valid
	if len(codes) < 2 {
		panic("Invalid ANSI codes")
	}

	//	Opening codes: "\u001b[__;__;__m"
	openCodes := codes[:len(codes)-1]
	open := CSI + strings.Join(openCodes, ";") + "m"

	//	The closing code: "\u001b[__m"
	closeCode := codes[len(codes)-1]
	close := CSI + closeCode + "m"

	return &ANSICode{open, close}

}

// Wrap ANSI Codes around string
func Wrap(str string, codes []string, enabled ...bool) string {
	if enabled[0] {
		ansiCode := Code(codes...)
		str = ansiCode.open + str + ansiCode.close
	}
	return str
}
