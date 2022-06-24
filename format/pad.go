package format

import "strings"

//	=======
//	PADDING
//	=======

//	Apply padding around the string
func Pad(s, char string, count int) string {
	space := strings.Repeat(char, count)
	return space + s + space
}

//	Apply padding to the left of the string
func PadLeft(s, char string, count int) string {
	return strings.Repeat(char, count) + s
}

//	Apply padding to the right of the string
func PadRight(s, char string, count int) string {
	return s + strings.Repeat(char, count)
}
