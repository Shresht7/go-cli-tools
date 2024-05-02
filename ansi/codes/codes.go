package codes

import (
	"strings"
)

//	=================
//	ANSI ESCAPE CODES
//	=================

const (
	ESC   = "\u001b"   // Escape Code
	CSI   = ESC + "["  // Control Sequence Introducer
	OSC   = ESC + "]"  // Operating System Command
	RESET = CSI + "0m" // Reset Code
	CTRL  = "^["       // Control character prefix
	BEL   = "\u0007"   // Bell character
)

// ----------
// IS ENABLED
// ----------

// Global variable to control whether the ANSI codes are enabled or not
var isEnabled = true

// Enable ANSI codes
func Enable() {
	isEnabled = true
}

// Disable ANSI codes
func Disable() {
	isEnabled = false
}

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
	if isEnabled {
		open, close := Code(codes...)
		str = open + str + close
	}
	return str
}
