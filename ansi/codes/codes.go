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

// A struct to represent an ANSI code
type ANSICode struct {
	open  string
	close string
}

// Create a new ANSI code with the given codes.
// The last code is used as the closing code.
func New(codes []string) *ANSICode {
	// Check if the codes are valid
	if len(codes) < 2 {
		codes = append(codes, "0")
	}

	// Opening codes: "\u001b[__;__;__m"
	openCodes := codes[:len(codes)-1]
	open := CSI + strings.Join(openCodes, ";") + "m"

	// The closing code: "\u001b[__m"
	closeCode := codes[len(codes)-1]
	close := CSI + closeCode + "m"

	return &ANSICode{open, close}
}

// Wrap ANSI codes around a strings
func (code *ANSICode) Wrap(str string) string {
	if isEnabled {
		return code.open + str + code.close
	}
	return str
}
