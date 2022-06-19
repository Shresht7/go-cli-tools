package codes

import (
	"fmt"
	"strings"
)

//	=================
//	ANSI ESCAPE CODES
//	=================

//	Escape Code
const ESC = "\u001b"

//	Control Sequence Introducer
const CSI = ESC + "["

//	Reset Code
const RESET = CSI + "0m"

//	Miscellaneous
const CTRL = "^["
const OSC = "\u001B]"
const BEL = "\u0007"

//	-------
//	HELPERS
//	-------

type ANSICode struct {
	open  string
	close string
}

//  Helper function to format the given ANSI codes
func Code(codes ...string) *ANSICode {

	//	Check if the codes are valid
	if len(codes) < 2 {
		panic("Invalid ANSI codes")
	}

	//	The slice of opening codes
	openCodes := codes[:len(codes)-1]
	//	The closing code
	closeCode := codes[len(codes)-1]

	//	Format the codes
	open := CSI + strings.Join(openCodes, ";") + "m"
	close := CSI + closeCode + "m"

	return &ANSICode{
		open,
		close,
	}
}

//	Wrap ANSI Codes around string
func Wrap(str string, codes []string, enabled ...bool) string {
	if enabled[0] {
		ansiCode := Code(codes...)
		str = ansiCode.open + str + ansiCode.close
	}
	return str
}

//	Returns the escape sequence
func Escape(format string, args ...any) string {
	return fmt.Sprintf("%s%s", ESC, fmt.Sprintf(format, args...))
}
