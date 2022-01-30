package helpers

import "fmt"

//	Escape Code
const ESC = "\u001b"

//	Reset Code
const RESET = "\u001b[0m"

//	Returns the escape sequence
func Escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", ESC, fmt.Sprintf(format, args...))
}

//	Wrap the string in escape codes
func Wrap(str string, codes [2]string) string {
	return Escape("[%sm", codes[0]) + str + Escape("[%sm", codes[1])
}
