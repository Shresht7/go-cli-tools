package format

import "strings"

//	=======
//	PADDING
//	=======

//	Apply padding around the string
func Pad(s string, count int) string {
	space := strings.Repeat(" ", count)
	s = space + s + space
	return s
}

//	Apply padding to the left of the string
func PadLeft(s string, count int) string {
	return strings.Repeat(" ", count) + s
}

//	Apply padding to the right of the string
func PadRight(s string, count int) string {
	return s + strings.Repeat(" ", count)
}
