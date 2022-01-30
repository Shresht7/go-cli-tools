package helpers

import "fmt"

//	Escape Code
const ESC = "\u001b"

//	Reset Code
const RESET = "\u001b[0m"

//	Bell
const BEL = "\u0007"
const OSC = "\u001B]"

//	Returns the escape sequence
func Escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", ESC, fmt.Sprintf(format, args...))
}

//	Wrap the string in escape codes
func Wrap(str string, codes [2]string) string {
	return Escape("[%sm", codes[0]) + str + Escape("[%sm", codes[1])
}

//	MISCELLANEOUS
//	-------------

//	Bell notification
func Bell() string {
	return BEL
}

//	Link
func Link(text, url string) string {
	return OSC + "8;;" + url + BEL + text + OSC + "8;;" + BEL
}

//	Save Screen
func SaveScreen() string {
	return Escape("[?47h")
}

//	Load Screen
func LoadScreen() string {
	return Escape("[?47l")
}

//	Enable Alt Buffer
func EnableAltBuffer() string {
	return Escape("[?1049h")
}

//	Disable Alt Buffer
func DisableAltBuffer() string {
	return Escape("[?1049l")
}
